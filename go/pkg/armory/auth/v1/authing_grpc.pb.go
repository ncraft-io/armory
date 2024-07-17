// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: armory/auth/v1/authing.proto

package auth

import (
	context "context"
	auth "github.com/ncraft-io/armory/go/pkg/armory/auth"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Authing_CreateAccount_FullMethodName = "/armory.auth.v1.Authing/create_account"
)

// AuthingClient is the client API for Authing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthingClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*auth.Account, error)
}

type authingClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthingClient(cc grpc.ClientConnInterface) AuthingClient {
	return &authingClient{cc}
}

func (c *authingClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*auth.Account, error) {
	out := new(auth.Account)
	err := c.cc.Invoke(ctx, Authing_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthingServer is the server API for Authing service.
// All implementations must embed UnimplementedAuthingServer
// for forward compatibility
type AuthingServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*auth.Account, error)
	mustEmbedUnimplementedAuthingServer()
}

// UnimplementedAuthingServer must be embedded to have forward compatible implementations.
type UnimplementedAuthingServer struct {
}

func (UnimplementedAuthingServer) CreateAccount(context.Context, *CreateAccountRequest) (*auth.Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAuthingServer) mustEmbedUnimplementedAuthingServer() {}

// UnsafeAuthingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthingServer will
// result in compilation errors.
type UnsafeAuthingServer interface {
	mustEmbedUnimplementedAuthingServer()
}

func RegisterAuthingServer(s grpc.ServiceRegistrar, srv AuthingServer) {
	s.RegisterService(&Authing_ServiceDesc, srv)
}

func _Authing_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthingServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authing_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthingServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authing_ServiceDesc is the grpc.ServiceDesc for Authing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "armory.auth.v1.Authing",
	HandlerType: (*AuthingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create_account",
			Handler:    _Authing_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "armory/auth/v1/authing.proto",
}
