// Code generated by ncraft. DO NOT EDIT.
// Rerunning ncraft will overwrite this file.
// Version: 0.1.0
// Version Date:

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/ncraft-io/armory/go/pkg/armory/auth"

	// this service api
	pb "github.com/ncraft-io/armory/go/pkg/armory/auth/v1"
)

var (
	_ = auth.Account{}
)

// MakeGRPCServer makes a set of endpoints available as a gRPC AuthingServer.
func MakeGRPCServer(endpoints Endpoints, tracer stdopentracing.Tracer, logger log.Logger) pb.AuthingServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
		grpctransport.ServerErrorLogger(logger),
	}

	addTracerOption := func(methodName string) []grpctransport.ServerOption {
		if tracer != nil {
			return append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, methodName, logger)))
		}
		return serverOptions
	}

	return &grpcServer{
		// Authing

		createAccount: grpctransport.NewServer(
			endpoints.CreateAccountEndpoint,
			DecodeGRPCCreateAccountRequest,
			EncodeGRPCCreateAccountResponse,
			addTracerOption("create_account")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "create_account", logger)))...,
		),
	}
}

// grpcServer implements the AuthingServer interface
type grpcServer struct {
	pb.UnimplementedAuthingServer

	createAccount grpctransport.Handler
}

// Methods for grpcServer to implement AuthingServer interface

func (s *grpcServer) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*auth.Account, error) {
	_, rep, err := s.createAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*auth.Account), nil
}

// Server Decode

// DecodeGRPCCreateAccountRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateAccount request to a user-domain CreateAccount request. Primarily useful in a server.
func DecodeGRPCCreateAccountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateAccountRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCCreateAccountResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain CreateAccount response to a gRPC CreateAccount reply. Primarily useful in a server.
func EncodeGRPCCreateAccountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*auth.Account)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
