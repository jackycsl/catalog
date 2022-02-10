// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/bff/admin/v1/shop_admin.proto

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

// ShopAdminClient is the client API for ShopAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopAdminClient interface {
	ListGame(ctx context.Context, in *ListGameReq, opts ...grpc.CallOption) (*ListGameReply, error)
	CreateGame(ctx context.Context, in *CreateGameReq, opts ...grpc.CallOption) (*CreateGameReply, error)
	UpdateGame(ctx context.Context, in *UpdateGameReq, opts ...grpc.CallOption) (*UpdateGameReply, error)
	DeleteGame(ctx context.Context, in *DeleteGameReq, opts ...grpc.CallOption) (*DeleteGameReply, error)
	GetGame(ctx context.Context, in *GetGameReq, opts ...grpc.CallOption) (*GetGameReply, error)
}

type shopAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewShopAdminClient(cc grpc.ClientConnInterface) ShopAdminClient {
	return &shopAdminClient{cc}
}

func (c *shopAdminClient) ListGame(ctx context.Context, in *ListGameReq, opts ...grpc.CallOption) (*ListGameReply, error) {
	out := new(ListGameReply)
	err := c.cc.Invoke(ctx, "/shop.admin.v1.ShopAdmin/ListGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopAdminClient) CreateGame(ctx context.Context, in *CreateGameReq, opts ...grpc.CallOption) (*CreateGameReply, error) {
	out := new(CreateGameReply)
	err := c.cc.Invoke(ctx, "/shop.admin.v1.ShopAdmin/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopAdminClient) UpdateGame(ctx context.Context, in *UpdateGameReq, opts ...grpc.CallOption) (*UpdateGameReply, error) {
	out := new(UpdateGameReply)
	err := c.cc.Invoke(ctx, "/shop.admin.v1.ShopAdmin/UpdateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopAdminClient) DeleteGame(ctx context.Context, in *DeleteGameReq, opts ...grpc.CallOption) (*DeleteGameReply, error) {
	out := new(DeleteGameReply)
	err := c.cc.Invoke(ctx, "/shop.admin.v1.ShopAdmin/DeleteGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopAdminClient) GetGame(ctx context.Context, in *GetGameReq, opts ...grpc.CallOption) (*GetGameReply, error) {
	out := new(GetGameReply)
	err := c.cc.Invoke(ctx, "/shop.admin.v1.ShopAdmin/GetGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopAdminServer is the server API for ShopAdmin service.
// All implementations must embed UnimplementedShopAdminServer
// for forward compatibility
type ShopAdminServer interface {
	ListGame(context.Context, *ListGameReq) (*ListGameReply, error)
	CreateGame(context.Context, *CreateGameReq) (*CreateGameReply, error)
	UpdateGame(context.Context, *UpdateGameReq) (*UpdateGameReply, error)
	DeleteGame(context.Context, *DeleteGameReq) (*DeleteGameReply, error)
	GetGame(context.Context, *GetGameReq) (*GetGameReply, error)
	mustEmbedUnimplementedShopAdminServer()
}

// UnimplementedShopAdminServer must be embedded to have forward compatible implementations.
type UnimplementedShopAdminServer struct {
}

func (UnimplementedShopAdminServer) ListGame(context.Context, *ListGameReq) (*ListGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGame not implemented")
}
func (UnimplementedShopAdminServer) CreateGame(context.Context, *CreateGameReq) (*CreateGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedShopAdminServer) UpdateGame(context.Context, *UpdateGameReq) (*UpdateGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGame not implemented")
}
func (UnimplementedShopAdminServer) DeleteGame(context.Context, *DeleteGameReq) (*DeleteGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGame not implemented")
}
func (UnimplementedShopAdminServer) GetGame(context.Context, *GetGameReq) (*GetGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedShopAdminServer) mustEmbedUnimplementedShopAdminServer() {}

// UnsafeShopAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopAdminServer will
// result in compilation errors.
type UnsafeShopAdminServer interface {
	mustEmbedUnimplementedShopAdminServer()
}

func RegisterShopAdminServer(s grpc.ServiceRegistrar, srv ShopAdminServer) {
	s.RegisterService(&ShopAdmin_ServiceDesc, srv)
}

func _ShopAdmin_ListGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopAdminServer).ListGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.admin.v1.ShopAdmin/ListGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopAdminServer).ListGame(ctx, req.(*ListGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopAdmin_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopAdminServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.admin.v1.ShopAdmin/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopAdminServer).CreateGame(ctx, req.(*CreateGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopAdmin_UpdateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopAdminServer).UpdateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.admin.v1.ShopAdmin/UpdateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopAdminServer).UpdateGame(ctx, req.(*UpdateGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopAdmin_DeleteGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopAdminServer).DeleteGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.admin.v1.ShopAdmin/DeleteGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopAdminServer).DeleteGame(ctx, req.(*DeleteGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopAdmin_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopAdminServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.admin.v1.ShopAdmin/GetGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopAdminServer).GetGame(ctx, req.(*GetGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopAdmin_ServiceDesc is the grpc.ServiceDesc for ShopAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.admin.v1.ShopAdmin",
	HandlerType: (*ShopAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGame",
			Handler:    _ShopAdmin_ListGame_Handler,
		},
		{
			MethodName: "CreateGame",
			Handler:    _ShopAdmin_CreateGame_Handler,
		},
		{
			MethodName: "UpdateGame",
			Handler:    _ShopAdmin_UpdateGame_Handler,
		},
		{
			MethodName: "DeleteGame",
			Handler:    _ShopAdmin_DeleteGame_Handler,
		},
		{
			MethodName: "GetGame",
			Handler:    _ShopAdmin_GetGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/bff/admin/v1/shop_admin.proto",
}