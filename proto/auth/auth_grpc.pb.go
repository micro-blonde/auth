// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v5.28.0
// source: auth.proto

package auth

import (
	context "context"
	account "github.com/micro-blonde/auth/proto/auth/account"
	profile "github.com/micro-blonde/auth/proto/auth/account/profile"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	ListAccounts(ctx context.Context, in *Request, opts ...grpc.CallOption) (*account.Accounts, error)
	GetAccount(ctx context.Context, in *Request, opts ...grpc.CallOption) (*account.Account, error)
	ListProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*profile.Profiles, error)
	GetProfile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*profile.Profile, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) ListAccounts(ctx context.Context, in *Request, opts ...grpc.CallOption) (*account.Accounts, error) {
	out := new(account.Accounts)
	err := c.cc.Invoke(ctx, "/auth.Auth/ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAccount(ctx context.Context, in *Request, opts ...grpc.CallOption) (*account.Account, error) {
	out := new(account.Account)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ListProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*profile.Profiles, error) {
	out := new(profile.Profiles)
	err := c.cc.Invoke(ctx, "/auth.Auth/ListProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetProfile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*profile.Profile, error) {
	out := new(profile.Profile)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	ListAccounts(context.Context, *Request) (*account.Accounts, error)
	GetAccount(context.Context, *Request) (*account.Account, error)
	ListProfiles(context.Context, *Request) (*profile.Profiles, error)
	GetProfile(context.Context, *Request) (*profile.Profile, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) ListAccounts(context.Context, *Request) (*account.Accounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedAuthServer) GetAccount(context.Context, *Request) (*account.Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAuthServer) ListProfiles(context.Context, *Request) (*profile.Profiles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProfiles not implemented")
}
func (UnimplementedAuthServer) GetProfile(context.Context, *Request) (*profile.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ListAccounts(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAccount(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ListProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ListProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ListProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ListProfiles(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetProfile(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAccounts",
			Handler:    _Auth_ListAccounts_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Auth_GetAccount_Handler,
		},
		{
			MethodName: "ListProfiles",
			Handler:    _Auth_ListProfiles_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _Auth_GetProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
