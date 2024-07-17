package handlers

import (
	"context"
	"github.com/ncraft-io/armory/go/pkg/armory/auth"

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
func (s authingServer) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*auth.Account, error) {
	resp := &auth.Account{}
	return resp, nil
}
