// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/catalog/job/v1/job.proto

package v1

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CatalogJobClient is the client API for CatalogJob service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogJobClient interface {
}

type catalogJobClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogJobClient(cc grpc.ClientConnInterface) CatalogJobClient {
	return &catalogJobClient{cc}
}

// CatalogJobServer is the server API for CatalogJob service.
// All implementations must embed UnimplementedCatalogJobServer
// for forward compatibility
type CatalogJobServer interface {
	mustEmbedUnimplementedCatalogJobServer()
}

// UnimplementedCatalogJobServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogJobServer struct {
}

func (UnimplementedCatalogJobServer) mustEmbedUnimplementedCatalogJobServer() {}

// UnsafeCatalogJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogJobServer will
// result in compilation errors.
type UnsafeCatalogJobServer interface {
	mustEmbedUnimplementedCatalogJobServer()
}

func RegisterCatalogJobServer(s grpc.ServiceRegistrar, srv CatalogJobServer) {
	s.RegisterService(&CatalogJob_ServiceDesc, srv)
}

// CatalogJob_ServiceDesc is the grpc.ServiceDesc for CatalogJob service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogJob_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.job.v1.CatalogJob",
	HandlerType: (*CatalogJobServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api/catalog/job/v1/job.proto",
}
