package handlers

import (
	"context"
	"sort"
	"strings"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
	"gorm.io/gorm"
)

type statPlan struct {
	functions []string
	groups    []string
	buckets   []string
}

func (s unitableServer) GetRowStat(ctx context.Context, in *pb.GetRowStatRequest) (*core.Object, error) {
	if in == nil {
		return nil, core.NewInvalidArgumentError("nil request")
	}
	table, err := s.getTable(ctx, in.Database, in.Table)
	if err != nil {
		return nil, err
	}
	columns := table.ColumnIndex()
	plans := map[string]*statPlan{}
	expressions := in.Stats
	if len(expressions) == 0 {
		for _, col := range table.Columns {
			if !col.Statistical {
				continue
			}
			switch {
			case col.Type == "integer" || col.Type == "float":
				expressions = append(expressions, "count "+col.Name, "sum "+col.Name, "avg "+col.Name, "min "+col.Name, "max "+col.Name)
			case isTimeColumn(col):
				expressions = append(expressions, "range "+col.Name)
			default:
				expressions = append(expressions, "group "+col.Name)
			}
		}
	}
	if len(expressions) == 0 {
		return nil, core.NewInvalidArgumentError("no statistics expressions or statistical columns")
	}
	var globalGroups []string
	for _, expression := range strings.Split(strings.Join(expressions, "|"), "|") {
		parts := strings.Fields(expression)
		if len(parts) != 2 && !(len(parts) == 4 && strings.EqualFold(parts[2], "group")) {
			return nil, core.NewInvalidArgumentError("invalid statistics expression %q", expression)
		}
		fun, name := strings.ToLower(parts[0]), parts[1]
		col := columns[name]
		if col == nil {
			return nil, core.NewInvalidArgumentError("unknown statistics column %s", name)
		}
		if plans[name] == nil {
			plans[name] = &statPlan{}
		}
		plan := plans[name]
		if len(parts) == 4 {
			if columns[parts[3]] == nil {
				return nil, core.NewInvalidArgumentError("unknown group column %s", parts[3])
			}
			plan.groups = appendUnique(plan.groups, parts[3])
		}
		switch fun {
		case "group":
			if len(parts) != 2 {
				return nil, core.NewInvalidArgumentError("invalid group expression")
			}
			plan.functions = appendUnique(plan.functions, "count")
			// Explicit group expressions apply to all requested aggregates. Automatic
			// statistics group each textual column independently.
			if len(in.Stats) > 0 {
				globalGroups = appendUnique(globalGroups, name)
			} else {
				plan.groups = appendUnique(plan.groups, name)
			}
		case "count", "min", "max":
			plan.functions = appendUnique(plan.functions, fun)
		case "sum", "avg":
			if col.Type != "integer" && col.Type != "float" {
				return nil, core.NewInvalidArgumentError("%s requires a numeric column", fun)
			}
			plan.functions = appendUnique(plan.functions, fun)
		case "range":
			if !isTimeColumn(col) && col.Type != "integer" && col.Type != "float" {
				return nil, core.NewInvalidArgumentError("range requires a numeric or time column")
			}
			plan.functions = appendUnique(plan.functions, "min", "max")
		case "years", "months", "days", "hours":
			if !isTimeColumn(col) {
				return nil, core.NewInvalidArgumentError("%s requires a time column", fun)
			}
			plan.buckets = appendUnique(plan.buckets, fun)
		default:
			return nil, core.NewInvalidArgumentError("unsupported statistics function %s", fun)
		}
	}
	qry, err := ParseQuery(in)
	if err != nil {
		return nil, err
	}
	meta := s.Synchro().GetMetaTable(table.Id, table)
	d := synchro.GetDataDB(table.Database)
	if d == nil {
		return nil, core.NewNotFoundError("database not found")
	}
	names := make([]string, 0, len(plans))
	for name := range plans {
		names = append(names, name)
	}
	sort.Strings(names)
	response := core.NewObject()
	for _, name := range names {
		plan := plans[name]
		plan.groups = appendUnique(plan.groups, globalGroups...)
		stmt := &gorm.Statement{DB: d.DB}
		field := stmt.Quote(name)
		var records []*core.Value
		run := func(selects, groups []string) error {
			tx := qry.Apply(d.WithContext(ctx).Table(table.Name), meta.FieldsInfo).Select(strings.Join(selects, ", "))
			if len(groups) > 0 {
				tx = tx.Group(strings.Join(groups, ", ")).Order(strings.Join(groups, ", "))
			}
			rows, err := tx.Rows()
			if err != nil {
				return err
			}
			defer rows.Close()
			names, err := rows.Columns()
			if err != nil {
				return err
			}
			for rows.Next() {
				values := make([]interface{}, len(names))
				ptrs := make([]interface{}, len(names))
				for i := range values {
					ptrs[i] = &values[i]
				}
				if err := rows.Scan(ptrs...); err != nil {
					return err
				}
				object := core.NewObject()
				for i, key := range names {
					value := values[i]
					if data, ok := value.([]byte); ok {
						value = string(data)
					}
					v, err := core.NewValue(value)
					if err != nil {
						return err
					}
					object.SetValue(key, v)
				}
				records = append(records, core.NewObjectValue(object))
			}
			return rows.Err()
		}
		var groups, selects []string
		for _, group := range plan.groups {
			quoted := stmt.Quote(group)
			groups = append(groups, quoted)
			selects = append(selects, quoted)
		}
		if len(plan.functions) > 0 {
			aggregate := append([]string{}, selects...)
			for _, fun := range plan.functions {
				aggregate = append(aggregate, strings.ToUpper(fun)+"("+field+") AS "+stmt.Quote(fun))
			}
			if err := run(aggregate, groups); err != nil {
				return nil, err
			}
		}
		for _, bucket := range plan.buckets {
			expression, err := timeBucket(d.Dialector.Name(), bucket, field)
			if err != nil {
				return nil, err
			}
			aggregate := append([]string{}, selects...)
			aggregate = append(aggregate, expression+" AS "+stmt.Quote(bucket), "COUNT("+field+") AS "+stmt.Quote("count"))
			bucketGroups := append([]string{}, groups...)
			bucketGroups = append(bucketGroups, expression)
			if err := run(aggregate, bucketGroups); err != nil {
				return nil, err
			}
		}
		response.SetValue(name, core.NewArrayValue(records...))
	}
	return response, nil
}

func appendUnique(values []string, extra ...string) []string {
	for _, value := range extra {
		found := false
		for _, old := range values {
			if old == value {
				found = true
				break
			}
		}
		if !found {
			values = append(values, value)
		}
	}
	return values
}
func isTimeColumn(col *unitable.Column) bool {
	return col.Type == "string" && (col.Format == "time" || col.Format == "datetime" || col.Format == "timestamp")
}
func timeBucket(driver, bucket, field string) (string, error) {
	formats := map[string]string{"years": "%Y", "months": "%Y-%m", "days": "%Y-%m-%d", "hours": "%Y-%m-%d %H"}
	switch driver {
	case "sqlite":
		return "strftime('" + formats[bucket] + "', " + field + ")", nil
	case "mysql":
		return "DATE_FORMAT(" + field + ", '" + formats[bucket] + "')", nil
	case "postgres":
		formats := map[string]string{"years": "YYYY", "months": "YYYY-MM", "days": "YYYY-MM-DD", "hours": "YYYY-MM-DD HH24"}
		return "TO_CHAR(" + field + ", '" + formats[bucket] + "')", nil
	default:
		return "", core.NewInvalidArgumentError("time bucket statistics unsupported for database driver %s", driver)
	}
}
