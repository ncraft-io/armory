package handlers

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/ncraft-io/ncraft/go/pkg/gokit/middleware"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/auth"

	"github.com/ncraft-io/armory/service-go/pkg/authing-service/svc"

	// this service api
	pb "github.com/ncraft-io/armory/go/pkg/armory/auth/v1"
)

var (
	_ = auth.User{}
	_ = core.Null{}
	_ = auth.LogonUser{}
	_ = core.Ordering{}
	_ = core.FieldMask{}
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

	{ // create_user
		if tracer != nil {
			in.CreateUserEndpoint = opentracing.TraceServer(tracer, "create_user")(in.CreateUserEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateUserEndpoint = middleware.Instrumenting(latency.With("method", "create_user"), count.With("method", "create_user"))(in.CreateUserEndpoint)
		}
		in.CreateUserEndpoint = middleware.NewJWT()(in.CreateUserEndpoint)
		//if validator != nil {
		//	in.CreateUserEndpoint = validator.Validate()(in.CreateUserEndpoint)
		//}
	}
	{ // update_user
		if tracer != nil {
			in.UpdateUserEndpoint = opentracing.TraceServer(tracer, "update_user")(in.UpdateUserEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateUserEndpoint = middleware.Instrumenting(latency.With("method", "update_user"), count.With("method", "update_user"))(in.UpdateUserEndpoint)
		}
		in.UpdateUserEndpoint = middleware.NewJWT()(in.UpdateUserEndpoint)
		//if validator != nil {
		//	in.UpdateUserEndpoint = validator.Validate()(in.UpdateUserEndpoint)
		//}
	}
	{ // active_user
		if tracer != nil {
			in.ActiveUserEndpoint = opentracing.TraceServer(tracer, "active_user")(in.ActiveUserEndpoint)
		}
		if count != nil && latency != nil {
			in.ActiveUserEndpoint = middleware.Instrumenting(latency.With("method", "active_user"), count.With("method", "active_user"))(in.ActiveUserEndpoint)
		}
		in.ActiveUserEndpoint = middleware.NewJWT()(in.ActiveUserEndpoint)
		//if validator != nil {
		//	in.ActiveUserEndpoint = validator.Validate()(in.ActiveUserEndpoint)
		//}
	}
	{ // get_user
		if tracer != nil {
			in.GetUserEndpoint = opentracing.TraceServer(tracer, "get_user")(in.GetUserEndpoint)
		}
		if count != nil && latency != nil {
			in.GetUserEndpoint = middleware.Instrumenting(latency.With("method", "get_user"), count.With("method", "get_user"))(in.GetUserEndpoint)
		}
		in.GetUserEndpoint = middleware.NewJWT()(in.GetUserEndpoint)
		//if validator != nil {
		//	in.GetUserEndpoint = validator.Validate()(in.GetUserEndpoint)
		//}
	}
	{ // list_user
		if tracer != nil {
			in.ListUserEndpoint = opentracing.TraceServer(tracer, "list_user")(in.ListUserEndpoint)
		}
		if count != nil && latency != nil {
			in.ListUserEndpoint = middleware.Instrumenting(latency.With("method", "list_user"), count.With("method", "list_user"))(in.ListUserEndpoint)
		}
		in.ListUserEndpoint = middleware.NewJWT()(in.ListUserEndpoint)
		//if validator != nil {
		//	in.ListUserEndpoint = validator.Validate()(in.ListUserEndpoint)
		//}
	}
	{ // delete_user
		if tracer != nil {
			in.DeleteUserEndpoint = opentracing.TraceServer(tracer, "delete_user")(in.DeleteUserEndpoint)
		}
		if count != nil && latency != nil {
			in.DeleteUserEndpoint = middleware.Instrumenting(latency.With("method", "delete_user"), count.With("method", "delete_user"))(in.DeleteUserEndpoint)
		}
		in.DeleteUserEndpoint = middleware.NewJWT()(in.DeleteUserEndpoint)
		//if validator != nil {
		//	in.DeleteUserEndpoint = validator.Validate()(in.DeleteUserEndpoint)
		//}
	}
	{ // update_password
		if tracer != nil {
			in.UpdatePasswordEndpoint = opentracing.TraceServer(tracer, "update_password")(in.UpdatePasswordEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdatePasswordEndpoint = middleware.Instrumenting(latency.With("method", "update_password"), count.With("method", "update_password"))(in.UpdatePasswordEndpoint)
		}
		in.UpdatePasswordEndpoint = middleware.NewJWT()(in.UpdatePasswordEndpoint)
		//if validator != nil {
		//	in.UpdatePasswordEndpoint = validator.Validate()(in.UpdatePasswordEndpoint)
		//}
	}
	{ // login
		if tracer != nil {
			in.LoginEndpoint = opentracing.TraceServer(tracer, "login")(in.LoginEndpoint)
		}
		if count != nil && latency != nil {
			in.LoginEndpoint = middleware.Instrumenting(latency.With("method", "login"), count.With("method", "login"))(in.LoginEndpoint)
		}
		in.LoginEndpoint = middleware.NewJWT()(in.LoginEndpoint)
		//if validator != nil {
		//	in.LoginEndpoint = validator.Validate()(in.LoginEndpoint)
		//}
	}
	{ // logout
		if tracer != nil {
			in.LogoutEndpoint = opentracing.TraceServer(tracer, "logout")(in.LogoutEndpoint)
		}
		if count != nil && latency != nil {
			in.LogoutEndpoint = middleware.Instrumenting(latency.With("method", "logout"), count.With("method", "logout"))(in.LogoutEndpoint)
		}
		in.LogoutEndpoint = middleware.NewJWT()(in.LogoutEndpoint)
		//if validator != nil {
		//	in.LogoutEndpoint = validator.Validate()(in.LogoutEndpoint)
		//}
	}

	return in
}

func WrapService(in pb.AuthingServer, options map[string]interface{}) pb.AuthingServer {
	_ = options
	return in
}
