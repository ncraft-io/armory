package handlers

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/ncraft-io/ncraft-gokit/pkg/middleware"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"

	"github.com/ncraft-io/armory/service-go/pkg/unitable-service/svc"

	// this service api
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
)

var (
	_ = unitable.Table{}
	_ = core.Null{}
	_ = unitable.Column{}
	_ = core.Object{}
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints, options map[string]interface{}) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/ncraft-io//_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	var tracer stdopentracing.Tracer
	if value, ok := options["tracer"]; ok && value != nil {
		tracer = value.(stdopentracing.Tracer)
	}
	var count *kitprometheus.Counter
	if value, ok := options["count"]; ok && value != nil {
		count = value.(*kitprometheus.Counter)
	}
	var latency *kitprometheus.Histogram
	if value, ok := options["latency"]; ok && value != nil {
		latency = value.(*kitprometheus.Histogram)
	}
	//var validator *middleware.Validator
	//if value, ok := options["validator"]; ok && value != nil {
	//	validator = value.(*middleware.Validator)
	//}

	{ // create_table
		if tracer != nil {
			in.CreateTableEndpoint = opentracing.TraceServer(tracer, "create_table")(in.CreateTableEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateTableEndpoint = middleware.Instrumenting(latency.With("method", "create_table"), count.With("method", "create_table"))(in.CreateTableEndpoint)
		}
		//if validator != nil {
		//	in.CreateTableEndpoint = validator.Validate()(in.CreateTableEndpoint)
		//}
	}
	{ // update_table
		if tracer != nil {
			in.UpdateTableEndpoint = opentracing.TraceServer(tracer, "update_table")(in.UpdateTableEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateTableEndpoint = middleware.Instrumenting(latency.With("method", "update_table"), count.With("method", "update_table"))(in.UpdateTableEndpoint)
		}
		//if validator != nil {
		//	in.UpdateTableEndpoint = validator.Validate()(in.UpdateTableEndpoint)
		//}
	}
	{ // get_table
		if tracer != nil {
			in.GetTableEndpoint = opentracing.TraceServer(tracer, "get_table")(in.GetTableEndpoint)
		}
		if count != nil && latency != nil {
			in.GetTableEndpoint = middleware.Instrumenting(latency.With("method", "get_table"), count.With("method", "get_table"))(in.GetTableEndpoint)
		}
		//if validator != nil {
		//	in.GetTableEndpoint = validator.Validate()(in.GetTableEndpoint)
		//}
	}
	{ // list_tables
		if tracer != nil {
			in.ListTablesEndpoint = opentracing.TraceServer(tracer, "list_tables")(in.ListTablesEndpoint)
		}
		if count != nil && latency != nil {
			in.ListTablesEndpoint = middleware.Instrumenting(latency.With("method", "list_tables"), count.With("method", "list_tables"))(in.ListTablesEndpoint)
		}
		//if validator != nil {
		//	in.ListTablesEndpoint = validator.Validate()(in.ListTablesEndpoint)
		//}
	}
	{ // delete_table
		if tracer != nil {
			in.DeleteTableEndpoint = opentracing.TraceServer(tracer, "delete_table")(in.DeleteTableEndpoint)
		}
		if count != nil && latency != nil {
			in.DeleteTableEndpoint = middleware.Instrumenting(latency.With("method", "delete_table"), count.With("method", "delete_table"))(in.DeleteTableEndpoint)
		}
		//if validator != nil {
		//	in.DeleteTableEndpoint = validator.Validate()(in.DeleteTableEndpoint)
		//}
	}
	{ // sync_table
		if tracer != nil {
			in.SyncTableEndpoint = opentracing.TraceServer(tracer, "sync_table")(in.SyncTableEndpoint)
		}
		if count != nil && latency != nil {
			in.SyncTableEndpoint = middleware.Instrumenting(latency.With("method", "sync_table"), count.With("method", "sync_table"))(in.SyncTableEndpoint)
		}
		//if validator != nil {
		//	in.SyncTableEndpoint = validator.Validate()(in.SyncTableEndpoint)
		//}
	}
	{ // create_column
		if tracer != nil {
			in.CreateColumnEndpoint = opentracing.TraceServer(tracer, "create_column")(in.CreateColumnEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateColumnEndpoint = middleware.Instrumenting(latency.With("method", "create_column"), count.With("method", "create_column"))(in.CreateColumnEndpoint)
		}
		//if validator != nil {
		//	in.CreateColumnEndpoint = validator.Validate()(in.CreateColumnEndpoint)
		//}
	}
	{ // update_column
		if tracer != nil {
			in.UpdateColumnEndpoint = opentracing.TraceServer(tracer, "update_column")(in.UpdateColumnEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateColumnEndpoint = middleware.Instrumenting(latency.With("method", "update_column"), count.With("method", "update_column"))(in.UpdateColumnEndpoint)
		}
		//if validator != nil {
		//	in.UpdateColumnEndpoint = validator.Validate()(in.UpdateColumnEndpoint)
		//}
	}
	{ // get_column
		if tracer != nil {
			in.GetColumnEndpoint = opentracing.TraceServer(tracer, "get_column")(in.GetColumnEndpoint)
		}
		if count != nil && latency != nil {
			in.GetColumnEndpoint = middleware.Instrumenting(latency.With("method", "get_column"), count.With("method", "get_column"))(in.GetColumnEndpoint)
		}
		//if validator != nil {
		//	in.GetColumnEndpoint = validator.Validate()(in.GetColumnEndpoint)
		//}
	}
	{ // delete_column
		if tracer != nil {
			in.DeleteColumnEndpoint = opentracing.TraceServer(tracer, "delete_column")(in.DeleteColumnEndpoint)
		}
		if count != nil && latency != nil {
			in.DeleteColumnEndpoint = middleware.Instrumenting(latency.With("method", "delete_column"), count.With("method", "delete_column"))(in.DeleteColumnEndpoint)
		}
		//if validator != nil {
		//	in.DeleteColumnEndpoint = validator.Validate()(in.DeleteColumnEndpoint)
		//}
	}
	{ // list_columns
		if tracer != nil {
			in.ListColumnsEndpoint = opentracing.TraceServer(tracer, "list_columns")(in.ListColumnsEndpoint)
		}
		if count != nil && latency != nil {
			in.ListColumnsEndpoint = middleware.Instrumenting(latency.With("method", "list_columns"), count.With("method", "list_columns"))(in.ListColumnsEndpoint)
		}
		//if validator != nil {
		//	in.ListColumnsEndpoint = validator.Validate()(in.ListColumnsEndpoint)
		//}
	}
	{ // batch_create_columns
		if tracer != nil {
			in.BatchCreateColumnsEndpoint = opentracing.TraceServer(tracer, "batch_create_columns")(in.BatchCreateColumnsEndpoint)
		}
		if count != nil && latency != nil {
			in.BatchCreateColumnsEndpoint = middleware.Instrumenting(latency.With("method", "batch_create_columns"), count.With("method", "batch_create_columns"))(in.BatchCreateColumnsEndpoint)
		}
		//if validator != nil {
		//	in.BatchCreateColumnsEndpoint = validator.Validate()(in.BatchCreateColumnsEndpoint)
		//}
	}
	{ // batch_update_column
		if tracer != nil {
			in.BatchUpdateColumnEndpoint = opentracing.TraceServer(tracer, "batch_update_column")(in.BatchUpdateColumnEndpoint)
		}
		if count != nil && latency != nil {
			in.BatchUpdateColumnEndpoint = middleware.Instrumenting(latency.With("method", "batch_update_column"), count.With("method", "batch_update_column"))(in.BatchUpdateColumnEndpoint)
		}
		//if validator != nil {
		//	in.BatchUpdateColumnEndpoint = validator.Validate()(in.BatchUpdateColumnEndpoint)
		//}
	}
	{ // batch_delete_column
		if tracer != nil {
			in.BatchDeleteColumnEndpoint = opentracing.TraceServer(tracer, "batch_delete_column")(in.BatchDeleteColumnEndpoint)
		}
		if count != nil && latency != nil {
			in.BatchDeleteColumnEndpoint = middleware.Instrumenting(latency.With("method", "batch_delete_column"), count.With("method", "batch_delete_column"))(in.BatchDeleteColumnEndpoint)
		}
		//if validator != nil {
		//	in.BatchDeleteColumnEndpoint = validator.Validate()(in.BatchDeleteColumnEndpoint)
		//}
	}
	{ // create_row
		if tracer != nil {
			in.CreateRowEndpoint = opentracing.TraceServer(tracer, "create_row")(in.CreateRowEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateRowEndpoint = middleware.Instrumenting(latency.With("method", "create_row"), count.With("method", "create_row"))(in.CreateRowEndpoint)
		}
		//if validator != nil {
		//	in.CreateRowEndpoint = validator.Validate()(in.CreateRowEndpoint)
		//}
	}
	{ // update_row
		if tracer != nil {
			in.UpdateRowEndpoint = opentracing.TraceServer(tracer, "update_row")(in.UpdateRowEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateRowEndpoint = middleware.Instrumenting(latency.With("method", "update_row"), count.With("method", "update_row"))(in.UpdateRowEndpoint)
		}
		//if validator != nil {
		//	in.UpdateRowEndpoint = validator.Validate()(in.UpdateRowEndpoint)
		//}
	}
	{ // get_row
		if tracer != nil {
			in.GetRowEndpoint = opentracing.TraceServer(tracer, "get_row")(in.GetRowEndpoint)
		}
		if count != nil && latency != nil {
			in.GetRowEndpoint = middleware.Instrumenting(latency.With("method", "get_row"), count.With("method", "get_row"))(in.GetRowEndpoint)
		}
		//if validator != nil {
		//	in.GetRowEndpoint = validator.Validate()(in.GetRowEndpoint)
		//}
	}
	{ // delete_row
		if tracer != nil {
			in.DeleteRowEndpoint = opentracing.TraceServer(tracer, "delete_row")(in.DeleteRowEndpoint)
		}
		if count != nil && latency != nil {
			in.DeleteRowEndpoint = middleware.Instrumenting(latency.With("method", "delete_row"), count.With("method", "delete_row"))(in.DeleteRowEndpoint)
		}
		//if validator != nil {
		//	in.DeleteRowEndpoint = validator.Validate()(in.DeleteRowEndpoint)
		//}
	}
	{ // list_row
		if tracer != nil {
			in.ListRowEndpoint = opentracing.TraceServer(tracer, "list_row")(in.ListRowEndpoint)
		}
		if count != nil && latency != nil {
			in.ListRowEndpoint = middleware.Instrumenting(latency.With("method", "list_row"), count.With("method", "list_row"))(in.ListRowEndpoint)
		}
		//if validator != nil {
		//	in.ListRowEndpoint = validator.Validate()(in.ListRowEndpoint)
		//}
	}
	{ // export_row
		if tracer != nil {
			in.ExportRowEndpoint = opentracing.TraceServer(tracer, "export_row")(in.ExportRowEndpoint)
		}
		if count != nil && latency != nil {
			in.ExportRowEndpoint = middleware.Instrumenting(latency.With("method", "export_row"), count.With("method", "export_row"))(in.ExportRowEndpoint)
		}
		//if validator != nil {
		//	in.ExportRowEndpoint = validator.Validate()(in.ExportRowEndpoint)
		//}
	}
	{ // batch_create_rows
		if tracer != nil {
			in.BatchCreateRowsEndpoint = opentracing.TraceServer(tracer, "batch_create_rows")(in.BatchCreateRowsEndpoint)
		}
		if count != nil && latency != nil {
			in.BatchCreateRowsEndpoint = middleware.Instrumenting(latency.With("method", "batch_create_rows"), count.With("method", "batch_create_rows"))(in.BatchCreateRowsEndpoint)
		}
		//if validator != nil {
		//	in.BatchCreateRowsEndpoint = validator.Validate()(in.BatchCreateRowsEndpoint)
		//}
	}
	{ // batch_update_rows
		if tracer != nil {
			in.BatchUpdateRowsEndpoint = opentracing.TraceServer(tracer, "batch_update_rows")(in.BatchUpdateRowsEndpoint)
		}
		if count != nil && latency != nil {
			in.BatchUpdateRowsEndpoint = middleware.Instrumenting(latency.With("method", "batch_update_rows"), count.With("method", "batch_update_rows"))(in.BatchUpdateRowsEndpoint)
		}
		//if validator != nil {
		//	in.BatchUpdateRowsEndpoint = validator.Validate()(in.BatchUpdateRowsEndpoint)
		//}
	}
	{ // batch_delete_rows
		if tracer != nil {
			in.BatchDeleteRowsEndpoint = opentracing.TraceServer(tracer, "batch_delete_rows")(in.BatchDeleteRowsEndpoint)
		}
		if count != nil && latency != nil {
			in.BatchDeleteRowsEndpoint = middleware.Instrumenting(latency.With("method", "batch_delete_rows"), count.With("method", "batch_delete_rows"))(in.BatchDeleteRowsEndpoint)
		}
		//if validator != nil {
		//	in.BatchDeleteRowsEndpoint = validator.Validate()(in.BatchDeleteRowsEndpoint)
		//}
	}

	return in
}

func WrapService(in pb.UnitableServer, options map[string]interface{}) pb.UnitableServer {
	return in
}
