package handlers

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/ncraft-io/armory/service-go/pkg/synchro"
)

// sqlToken keeps quoted text and comments intact. Only unquoted ? tokens bind
// parameters; optional clauses use [WHERE ...], [AND ...], etc.
type sqlToken struct {
	text            string
	word, parameter bool
}
type queryPart struct {
	tokens   []sqlToken
	optional bool
}
type queryTemplate struct {
	parts []queryPart
	count int
}

func queryTokens(sql string) ([]sqlToken, error) {
	var tokens []sqlToken
	for i := 0; i < len(sql); {
		start := i
		c := sql[i]
		word := false
		switch {
		case c == '\'' || c == '"' || c == '`':
			i++
			closed := false
			for i < len(sql) {
				if sql[i] == c {
					i++
					if i < len(sql) && sql[i] == c {
						i++
						continue
					}
					closed = true
					break
				}
				if sql[i] == '\\' {
					return nil, fmt.Errorf("backslash escapes are not supported; use SQL doubled quotes")
				}
				i++
			}
			if !closed {
				return nil, fmt.Errorf("unterminated SQL quote")
			}
		case c == '-' && i+1 < len(sql) && sql[i+1] == '-':
			for i < len(sql) && sql[i] != '\n' {
				i++
			}
		case c == '/' && i+1 < len(sql) && sql[i+1] == '*':
			if strings.HasPrefix(sql[i:], "/*!") || strings.HasPrefix(sql[i:], "/*M!") {
				return nil, fmt.Errorf("executable SQL comments are not supported")
			}
			i += 2
			depth := 1
			for i < len(sql) && depth > 0 {
				if i+1 < len(sql) && sql[i:i+2] == "/*" {
					depth++
					i += 2
				} else if i+1 < len(sql) && sql[i:i+2] == "*/" {
					depth--
					i += 2
				} else {
					i++
				}
			}
			if depth != 0 {
				return nil, fmt.Errorf("unterminated SQL comment")
			}
		case c == '$':
			j := i + 1
			for j < len(sql) && ((sql[j] >= 'a' && sql[j] <= 'z') || (sql[j] >= 'A' && sql[j] <= 'Z') || sql[j] == '_' || (j > i+1 && sql[j] >= '0' && sql[j] <= '9')) {
				j++
			}
			if j < len(sql) && sql[j] == '$' {
				tag := sql[i : j+1]
				end := strings.Index(sql[j+1:], tag)
				if end < 0 {
					return nil, fmt.Errorf("unterminated dollar quote")
				}
				i = j + 1 + end + len(tag)
			} else {
				return nil, fmt.Errorf("use ? for positional SQL parameters")
			}
		case unicode.IsLetter(rune(c)) || c == '_':
			word = true
			i++
			for i < len(sql) && (unicode.IsLetter(rune(sql[i])) || unicode.IsDigit(rune(sql[i])) || sql[i] == '_' || sql[i] == '$') {
				i++
			}
		default:
			i++
		}
		tokens = append(tokens, sqlToken{text: sql[start:i], word: word, parameter: c == '?'})
	}
	return tokens, nil
}

func parseQueryTemplate(sql string) (*queryTemplate, error) {
	tokens, err := queryTokens(sql)
	if err != nil {
		return nil, err
	}
	first := ""
	ended := false
	forbidden := map[string]bool{"INSERT": true, "UPDATE": true, "DELETE": true, "MERGE": true, "REPLACE": true, "DROP": true, "ALTER": true, "CREATE": true, "TRUNCATE": true, "GRANT": true, "REVOKE": true, "CALL": true, "EXEC": true, "EXECUTE": true, "COPY": true, "INTO": true, "ATTACH": true, "DETACH": true, "PRAGMA": true, "VACUUM": true, "LOCK": true}
	for i, t := range tokens {
		if strings.TrimSpace(t.text) == "" || strings.HasPrefix(t.text, "--") || strings.HasPrefix(t.text, "/*") {
			continue
		}
		if ended {
			return nil, fmt.Errorf("only one SQL statement is allowed")
		}
		if t.text == ";" {
			ended = true
			tokens[i].text = ""
			continue
		}
		if t.word {
			w := strings.ToUpper(t.text)
			if first == "" {
				first = w
			}
			if forbidden[w] {
				return nil, fmt.Errorf("query contains unsupported SQL keyword %s", w)
			}
		}
	}
	if first != "SELECT" && first != "WITH" {
		return nil, fmt.Errorf("query must be a SELECT or WITH SELECT statement")
	}
	result := &queryTemplate{}
	part := queryPart{}
	brackets := 0
	for i, t := range tokens {
		if t.text == "[" {
			next := ""
			for _, n := range tokens[i+1:] {
				if strings.TrimSpace(n.text) != "" {
					next = strings.ToUpper(n.text)
					break
				}
			}
			optional := next == "WHERE" || next == "AND" || next == "OR" || next == "HAVING" || next == "LIMIT" || next == "OFFSET" || next == "ORDER" || next == "GROUP" || next == "JOIN" || next == "LEFT" || next == "INNER"
			if optional {
				if part.optional {
					return nil, fmt.Errorf("nested optional SQL clauses are not supported")
				}
				result.parts = append(result.parts, part)
				part = queryPart{optional: true}
				continue
			}
			brackets++
		}
		if t.text == "]" {
			if brackets > 0 {
				brackets--
			} else if part.optional {
				result.parts = append(result.parts, part)
				part = queryPart{}
				continue
			} else {
				return nil, fmt.Errorf("unmatched SQL bracket")
			}
		}
		part.tokens = append(part.tokens, t)
		if t.parameter {
			result.count++
		}
	}
	if brackets != 0 || part.optional {
		return nil, fmt.Errorf("unclosed SQL bracket")
	}
	result.parts = append(result.parts, part)
	for _, p := range result.parts {
		if p.optional {
			found := false
			for _, t := range p.tokens {
				found = found || t.parameter
			}
			if !found {
				return nil, fmt.Errorf("optional SQL clause requires a parameter")
			}
		}
	}
	return result, nil
}

func parameterType(p *unitable.DbQuery_Parameter) string {
	switch strings.ToLower(p.Type) {
	case "", "string":
		return "string"
	case "int", "int64", "integer":
		return "integer"
	case "number", "double", "float", "float64":
		return "float"
	case "bool", "boolean":
		return "bool"
	default:
		return ""
	}
}
func queryScalar(value interface{}, typ string) (interface{}, error) {
	// HTTP values are strings, while JSON bodies retain native scalar types.
	if value == nil {
		return nil, fmt.Errorf("null parameter")
	}
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Map || rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array || rv.Kind() == reflect.Ptr {
		return nil, fmt.Errorf("expected a scalar parameter")
	}
	text := fmt.Sprint(value)
	switch typ {
	case "string":
		if v, ok := value.(string); ok {
			return v, nil
		}
		return nil, fmt.Errorf("expected a string")
	case "integer":
		return strconv.ParseInt(text, 10, 64)
	case "float":
		v, err := strconv.ParseFloat(text, 64)
		if err == nil && (math.IsNaN(v) || math.IsInf(v, 0)) {
			err = fmt.Errorf("expected a finite number")
		}
		return v, err
	case "bool":
		return strconv.ParseBool(text)
	}
	return nil, fmt.Errorf("unsupported parameter type")
}
func queryValue(p *unitable.DbQuery_Parameter, value interface{}, dialect string) (interface{}, error) {
	if !p.IsArray {
		return queryScalar(value, parameterType(p))
	}
	if str, ok := value.(string); ok {
		value = strings.Split(str, ",")
	}
	v := reflect.ValueOf(value)
	if !v.IsValid() || (v.Kind() != reflect.Slice && v.Kind() != reflect.Array) {
		return nil, fmt.Errorf("parameter %s must be an array", p.Name)
	}
	values := make([]interface{}, v.Len())
	for i := range values {
		var err error
		values[i], err = queryScalar(v.Index(i).Interface(), parameterType(p))
		if err != nil {
			return nil, err
		}
	}
	if p.PgArray {
		if dialect != "postgres" || parameterType(p) != "string" {
			return nil, fmt.Errorf("pg_array requires PostgreSQL and string elements")
		}
		array := make(synchro.StringArray, len(values))
		for i, v := range values {
			array[i] = v.(string)
		}
		return array, nil
	}
	return values, nil
}
func missingQueryValue(v interface{}) bool {
	if v == nil {
		return true
	}
	r := reflect.ValueOf(v)
	return (r.Kind() == reflect.Array || r.Kind() == reflect.Slice) && r.Len() == 0
}

// bind emits native database/sql placeholders so question marks in literals and
// comments cannot consume arguments. Values are always passed separately.
func (t *queryTemplate) bind(q *unitable.DbQuery, input map[string]interface{}, dialect string) (string, []interface{}, error) {
	if t.count != len(q.Parameters) {
		return "", nil, fmt.Errorf("SQL has %d placeholders but %d parameter definitions", t.count, len(q.Parameters))
	}
	var sql strings.Builder
	var args []interface{}
	index := 0
	placeholder := func(v interface{}) string {
		args = append(args, v)
		if dialect == "postgres" {
			return "$" + strconv.Itoa(len(args))
		}
		return "?"
	}
	for _, part := range t.parts {
		start := index
		present, missing := 0, 0
		for _, token := range part.tokens {
			if token.parameter {
				p := q.Parameters[index]
				if p == nil {
					return "", nil, fmt.Errorf("nil parameter")
				}
				if missingQueryValue(input[p.Name]) {
					missing++
				} else {
					present++
				}
				index++
			}
		}
		if part.optional && present == 0 {
			sql.WriteByte(' ')
			continue
		}
		if missing > 0 {
			return "", nil, fmt.Errorf("missing required parameters or partially supplied optional clause")
		}
		index = start
		for i, token := range part.tokens {
			if !token.parameter {
				sql.WriteString(token.text)
				continue
			}
			p := q.Parameters[index]
			index++
			v, err := queryValue(p, input[p.Name], dialect)
			if err != nil {
				return "", nil, fmt.Errorf("parameter %s: %w", p.Name, err)
			}
			if array, ok := v.([]interface{}); ok {
				// Support both IN ? (legacy configuration) and IN (?).
				prev, next := "", ""
				for j := i - 1; j >= 0; j-- {
					if strings.TrimSpace(part.tokens[j].text) != "" {
						prev = part.tokens[j].text
						break
					}
				}
				for j := i + 1; j < len(part.tokens); j++ {
					if strings.TrimSpace(part.tokens[j].text) != "" {
						next = part.tokens[j].text
						break
					}
				}
				wrap := prev != "(" || next != ")"
				if wrap {
					sql.WriteByte('(')
				}
				for n, item := range array {
					if n > 0 {
						sql.WriteByte(',')
					}
					sql.WriteString(placeholder(item))
				}
				if wrap {
					sql.WriteByte(')')
				}
			} else {
				sql.WriteString(placeholder(v))
			}
		}
	}
	return sql.String(), args, nil
}
