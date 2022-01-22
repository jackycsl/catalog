// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/catalog/service/v1/catalog.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogClient interface {
	ListGame(ctx context.Context, in *ListGameReq, opts ...grpc.CallOption) (*ListGameReply, error)
	CreateGame(ctx context.Context, in *CreateGameReq, opts ...grpc.CallOption) (*CreateGameReply, error)
	GetGame(ctx context.Context, in *GetGameReq, opts ...grpc.CallOption) (*GetGameReply, error)
	UpdateGame(ctx context.Context, in *UpdateGameReq, opts ...grpc.CallOption) (*UpdateGameReply, error)
	DeleteGame(ctx context.Context, in *DeleteGameReq, opts ...grpc.CallOption) (*DeleteGameReply, error)
	HealthCheck(ctx context.Context, in *HealthReq, opts ...grpc.CallOption) (*HealthReply, error)
}

type catalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogClient(cc grpc.ClientConnInterface) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) ListGame(ctx context.Context, in *ListGameReq, opts ...grpc.CallOption) (*ListGameReply, error) {
	out := new(ListGameReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/ListGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) CreateGame(ctx context.Context, in *CreateGameReq, opts ...grpc.CallOption) (*CreateGameReply, error) {
	out := new(CreateGameReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) GetGame(ctx context.Context, in *GetGameReq, opts ...grpc.CallOption) (*GetGameReply, error) {
	out := new(GetGameReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/GetGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) UpdateGame(ctx context.Context, in *UpdateGameReq, opts ...grpc.CallOption) (*UpdateGameReply, error) {
	out := new(UpdateGameReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/UpdateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) DeleteGame(ctx context.Context, in *DeleteGameReq, opts ...grpc.CallOption) (*DeleteGameReply, error) {
	out := new(DeleteGameReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/DeleteGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) HealthCheck(ctx context.Context, in *HealthReq, opts ...grpc.CallOption) (*HealthReply, error) {
	out := new(HealthReply)
	err := c.cc.Invoke(ctx, "/catalog.service.v1.Catalog/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
// All implementations must embed UnimplementedCatalogServer
// for forward compatibility
type CatalogServer interface {
	ListGame(context.Context, *ListGameReq) (*ListGameReply, error)
	CreateGame(context.Context, *CreateGameReq) (*CreateGameReply, error)
	GetGame(context.Context, *GetGameReq) (*GetGameReply, error)
	UpdateGame(context.Context, *UpdateGameReq) (*UpdateGameReply, error)
	DeleteGame(context.Context, *DeleteGameReq) (*DeleteGameReply, error)
	HealthCheck(context.Context, *HealthReq) (*HealthReply, error)
	mustEmbedUnimplementedCatalogServer()
}

// UnimplementedCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServer struct {
}

func (UnimplementedCatalogServer) ListGame(context.Context, *ListGameReq) (*ListGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGame not implemented")
}
func (UnimplementedCatalogServer) CreateGame(context.Context, *CreateGameReq) (*CreateGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedCatalogServer) GetGame(context.Context, *GetGameReq) (*GetGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedCatalogServer) UpdateGame(context.Context, *UpdateGameReq) (*UpdateGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGame not implemented")
}
func (UnimplementedCatalogServer) DeleteGame(context.Context, *DeleteGameReq) (*DeleteGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGame not implemented")
}
func (UnimplementedCatalogServer) HealthCheck(context.Context, *HealthReq) (*HealthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedCatalogServer) mustEmbedUnimplementedCatalogServer() {}

// UnsafeCatalogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServer will
// result in compilation errors.
type UnsafeCatalogServer interface {
	mustEmbedUnimplementedCatalogServer()
}

func RegisterCatalogServer(s grpc.ServiceRegistrar, srv CatalogServer) {
	s.RegisterService(&Catalog_ServiceDesc, srv)
}

func _Catalog_ListGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).ListGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/ListGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).ListGame(ctx, req.(*ListGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).CreateGame(ctx, req.(*CreateGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/GetGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).GetGame(ctx, req.(*GetGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_UpdateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).UpdateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/UpdateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).UpdateGame(ctx, req.(*UpdateGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_DeleteGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).DeleteGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/DeleteGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).DeleteGame(ctx, req.(*DeleteGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/catalog.service.v1.Catalog/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).HealthCheck(ctx, req.(*HealthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Catalog_ServiceDesc is the grpc.ServiceDesc for Catalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.service.v1.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGame",
			Handler:    _Catalog_ListGame_Handler,
		},
		{
			MethodName: "CreateGame",
			Handler:    _Catalog_CreateGame_Handler,
		},
		{
			MethodName: "GetGame",
			Handler:    _Catalog_GetGame_Handler,
		},
		{
			MethodName: "UpdateGame",
			Handler:    _Catalog_UpdateGame_Handler,
		},
		{
			MethodName: "DeleteGame",
			Handler:    _Catalog_DeleteGame_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _Catalog_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/catalog/service/v1/catalog.proto",
}
