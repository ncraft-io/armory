// Code generated by ncraft. DO NOT EDIT.
// Rerunning ncraft will overwrite this file.
// Version: 0.1.0
// Version Date:

package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net"

	"github.com/rs/cors"
	"net/http"
	"net/http/pprof"
	"strings"
	"time"

	// 3d Party
	"github.com/etherlabsio/healthcheck/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	"github.com/ncraft-io/ncraft-gokit/pkg/utils/network"
	_ "go.uber.org/automaxprocs"

	nclient "github.com/ncraft-io/ncraft-gokit/pkg/client"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"

	"github.com/ncraft-io/ncraft-gokit/pkg/kit"
	"github.com/ncraft-io/ncraft-gokit/pkg/metrics"
	"github.com/ncraft-io/ncraft-gokit/pkg/sd"
	nserver "github.com/ncraft-io/ncraft-gokit/pkg/server"
	"github.com/ncraft-io/ncraft-gokit/pkg/tracing"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/ncraft-io/armory/service-go/pkg/unitable-service/handlers"
	"github.com/ncraft-io/armory/service-go/pkg/unitable-service/svc"

	// This Service api
	pb "github.com/ncraft-io/armory/go/pkg/armory/unitable/v1"
)

var _ nclient.Config

func NewEndpoints(options map[string]interface{}) svc.Endpoints {
	// Business domain.
	var service pb.UnitableServer
	{
		service = handlers.NewService()
		// Wrap Service with middlewares. See handlers/middlewares.go
		service = handlers.WrapService(service, options)
	}

	// Endpoint domain.
	var (
		createTableEndpoint        = svc.MakeCreateTableEndpoint(service)
		updateTableEndpoint        = svc.MakeUpdateTableEndpoint(service)
		getTableEndpoint           = svc.MakeGetTableEndpoint(service)
		listTablesEndpoint         = svc.MakeListTablesEndpoint(service)
		deleteTableEndpoint        = svc.MakeDeleteTableEndpoint(service)
		syncTableEndpoint          = svc.MakeSyncTableEndpoint(service)
		createColumnEndpoint       = svc.MakeCreateColumnEndpoint(service)
		updateColumnEndpoint       = svc.MakeUpdateColumnEndpoint(service)
		getColumnEndpoint          = svc.MakeGetColumnEndpoint(service)
		deleteColumnEndpoint       = svc.MakeDeleteColumnEndpoint(service)
		listColumnsEndpoint        = svc.MakeListColumnsEndpoint(service)
		batchCreateColumnsEndpoint = svc.MakeBatchCreateColumnsEndpoint(service)
		batchUpdateColumnEndpoint  = svc.MakeBatchUpdateColumnEndpoint(service)
		batchDeleteColumnEndpoint  = svc.MakeBatchDeleteColumnEndpoint(service)
		createRowEndpoint          = svc.MakeCreateRowEndpoint(service)
		updateRowEndpoint          = svc.MakeUpdateRowEndpoint(service)
		getRowEndpoint             = svc.MakeGetRowEndpoint(service)
		deleteRowEndpoint          = svc.MakeDeleteRowEndpoint(service)
		listRowEndpoint            = svc.MakeListRowEndpoint(service)
		exportRowEndpoint          = svc.MakeExportRowEndpoint(service)
		batchCreateRowsEndpoint    = svc.MakeBatchCreateRowsEndpoint(service)
		batchUpdateRowsEndpoint    = svc.MakeBatchUpdateRowsEndpoint(service)
		batchDeleteRowsEndpoint    = svc.MakeBatchDeleteRowsEndpoint(service)
	)

	endpoints := svc.Endpoints{
		CreateTableEndpoint:        createTableEndpoint,
		UpdateTableEndpoint:        updateTableEndpoint,
		GetTableEndpoint:           getTableEndpoint,
		ListTablesEndpoint:         listTablesEndpoint,
		DeleteTableEndpoint:        deleteTableEndpoint,
		SyncTableEndpoint:          syncTableEndpoint,
		CreateColumnEndpoint:       createColumnEndpoint,
		UpdateColumnEndpoint:       updateColumnEndpoint,
		GetColumnEndpoint:          getColumnEndpoint,
		DeleteColumnEndpoint:       deleteColumnEndpoint,
		ListColumnsEndpoint:        listColumnsEndpoint,
		BatchCreateColumnsEndpoint: batchCreateColumnsEndpoint,
		BatchUpdateColumnEndpoint:  batchUpdateColumnEndpoint,
		BatchDeleteColumnEndpoint:  batchDeleteColumnEndpoint,
		CreateRowEndpoint:          createRowEndpoint,
		UpdateRowEndpoint:          updateRowEndpoint,
		GetRowEndpoint:             getRowEndpoint,
		DeleteRowEndpoint:          deleteRowEndpoint,
		ListRowEndpoint:            listRowEndpoint,
		ExportRowEndpoint:          exportRowEndpoint,
		BatchCreateRowsEndpoint:    batchCreateRowsEndpoint,
		BatchUpdateRowsEndpoint:    batchUpdateRowsEndpoint,
		BatchDeleteRowsEndpoint:    batchDeleteRowsEndpoint,
	}

	// Wrap selected Endpoints with middlewares. See handlers/middlewares.go
	endpoints = handlers.WrapEndpoints(endpoints, options)

	return endpoints
}

func RegisterService(cfg nserver.Config, r *mux.Router, s *grpc.Server) svc.Endpoints {
	const FullServiceName = "unitable.Unitable"

	// tracing init
	tracer, c := tracing.New(FullServiceName)
	if c != nil {
		defer c.Close()
	}

	// Create a single logger, which we'll use and give to other components.
	logger := kit.Logger()

	options := map[string]interface{}{
		"tracer": tracer,
		"logger": logger,
	}

	metricsConfig := metrics.NewConfig("metrics")
	if metricsConfig.Enabled() {
		fieldKeys := []string{"method", "access_key", "error"}
		count := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: metricsConfig.Department,
			Subsystem: metricsConfig.Project,
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys)

		latency := kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: metricsConfig.Department,
			Subsystem: metricsConfig.Project,
			Name:      "request_latency_seconds",
			Help:      "Total duration of requests in seconds.",
		}, fieldKeys)

		options["count"] = count
		options["latency"] = latency
	}

	sdConfig := sd.NewConfig("sd")
	sdClient := sd.New(sdConfig, logger)

	if sdClient != nil {
		url := "etcd://" + network.GetHost() + ":" + getGrpcPort(cfg.GrpcAddr)
		err := sdClient.Register(url, FullServiceName, []string{})
		if err != nil {
			panic(err)
		}
		defer sdClient.Deregister()
	}

	// required service clients ...
	//xxClient := xx_client.New(nclient.NewConfig("xx"), sdClient.Instancer(FullServiceName), tracer, logger)
	//defer xxClient.Close()

	endpoints := NewEndpoints(options)

	svc.RegisterHttpHandler(r, endpoints, tracer, logger)
	pb.RegisterUnitableServer(s, svc.MakeGRPCServer(endpoints, tracer, logger))

	return endpoints
}

// Run starts a new http server, gRPC server, and a debug server with the
// passed config and logger
func Run(cfg nserver.Config) {
	// Mechanical domain.
	errc := make(chan error)

	// Interrupt handler.
	go handlers.InterruptHandler(errc)

	// Debug listener.
	go func() {
		logs.Infow("begin debug server", "transport", "debug", "address", cfg.DebugAddr)

		m := http.NewServeMux()
		m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

		m.Handle("/metrics", promhttp.Handler())

		m.Handle("/health", healthcheck.Handler(
			// WithTimeout allows you to set a max overall timeout.
			healthcheck.WithTimeout(5*time.Second),
			healthcheck.WithChecker("alive", healthcheck.CheckerFunc(func(ctx context.Context) error {
				conn, err := net.DialTimeout("tcp", cfg.HttpAddr, time.Second)
				if err != nil {
					return err
				}
				return conn.Close()
			})),
		))

		errc <- http.ListenAndServe(cfg.DebugAddr, m)
	}()

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryServerFilter))
	r := mux.NewRouter()
	endpoints := RegisterService(cfg, r, s)

	// HTTP transport.
	go func() {
		logs.Infow("begin http server", "transport", "HTTP", "address", cfg.HttpAddr)
		h := cors.AllowAll().Handler(r)
		errc <- http.ListenAndServe(cfg.HttpAddr, h)
	}()

	// gRPC transport.
	go func() {
		logs.Infow("begin grpc server", "transport", "gRPC", "address", cfg.GrpcAddr)
		ln, err := net.Listen("tcp", cfg.GrpcAddr)
		if err != nil {
			errc <- err
			return
		}
		errc <- s.Serve(ln)
	}()

	//if watchObj, err := config.WatchFunc(level.ChangeLogLevel, level.LevelPath); err == nil {
	//    defer func() { _ = watchObj.Close() }()
	//} else {
	//    panic(err.Error())
	//}
	_ = endpoints

	// Run!
	logs.Info("unitable.UnitableServer", " started.")
	logs.Info("unitable.UnitableServer", <-errc)

	logs.Info("unitable.UnitableServer", " closed.")
}

func getGrpcPort(addr string) string {
	host := strings.Split(addr, ":")
	if len(host) < 2 {
		panic("host name is invalid (" + addr + ")")
	}
	return host[1]
}

func unaryServerFilter(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	//if err := middleware.Validate(req); err != nil {
	//	logs.Errorf("validate request failed, err: %s", err)
	//	return nil, core.NewError(http.StatusBadRequest, err.Error())
	//}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	resp, err = handler(ctx, req)
	if err != nil {
		return resp, err
	}

	//var validatorCfg middleware.ValidatorConfig
	//_ = config.ScanKey("validator", &validatorCfg)
	//if !validatorCfg.CheckResponse {
	//	return
	//}
	//if err = middleware.Validate(resp); err != nil {
	//	logs.Errorf("validate response failed, err: %s", err)
	//	return nil, core.NewError(http.StatusInternalServerError, err.Error())
	//}
	return
}