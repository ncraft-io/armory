package handlers

import (
	"context"

	// this service api
	pb "github.com/ncraft-io/armory/go/pkg/armory/auth/v1"
)

type authingServer struct {
	pb.UnimplementedAuthingServer
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.AuthingServer {
	return authingServer{}
}

// CreateAccount implements Interface.
func (s authingServer) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}
