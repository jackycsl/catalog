// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/bff/interface/v1/shop_interface.proto

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

// ShopInterfaceClient is the client API for ShopInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopInterfaceClient interface {
	ListGame(ctx context.Context, in *ListGameReq, opts ...grpc.CallOption) (*ListGameReply, error)
	GetGame(ctx context.Context, in *GetGameReq, opts ...grpc.CallOption) (*GetGameReply, error)
}

type shopInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopInterfaceClient(cc grpc.ClientConnInterface) ShopInterfaceClient {
	return &shopInterfaceClient{cc}
}

func (c *shopInterfaceClient) ListGame(ctx context.Context, in *ListGameReq, opts ...grpc.CallOption) (*ListGameReply, error) {
	out := new(ListGameReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/ListGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) GetGame(ctx context.Context, in *GetGameReq, opts ...grpc.CallOption) (*GetGameReply, error) {
	out := new(GetGameReply)
	err := c.cc.Invoke(ctx, "/shop.interface.v1.ShopInterface/GetGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopInterfaceServer is the server API for ShopInterface service.
// All implementations must embed UnimplementedShopInterfaceServer
// for forward compatibility
type ShopInterfaceServer interface {
	ListGame(context.Context, *ListGameReq) (*ListGameReply, error)
	GetGame(context.Context, *GetGameReq) (*GetGameReply, error)
	mustEmbedUnimplementedShopInterfaceServer()
}

// UnimplementedShopInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedShopInterfaceServer struct {
}

func (UnimplementedShopInterfaceServer) ListGame(context.Context, *ListGameReq) (*ListGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGame not implemented")
}
func (UnimplementedShopInterfaceServer) GetGame(context.Context, *GetGameReq) (*GetGameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedShopInterfaceServer) mustEmbedUnimplementedShopInterfaceServer() {}

// UnsafeShopInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopInterfaceServer will
// result in compilation errors.
type UnsafeShopInterfaceServer interface {
	mustEmbedUnimplementedShopInterfaceServer()
}

func RegisterShopInterfaceServer(s grpc.ServiceRegistrar, srv ShopInterfaceServer) {
	s.RegisterService(&ShopInterface_ServiceDesc, srv)
}

func _ShopInterface_ListGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).ListGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/ListGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).ListGame(ctx, req.(*ListGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shop.interface.v1.ShopInterface/GetGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).GetGame(ctx, req.(*GetGameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopInterface_ServiceDesc is the grpc.ServiceDesc for ShopInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shop.interface.v1.ShopInterface",
	HandlerType: (*ShopInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGame",
			Handler:    _ShopInterface_ListGame_Handler,
		},
		{
			MethodName: "GetGame",
			Handler:    _ShopInterface_GetGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/bff/interface/v1/shop_interface.proto",
}
