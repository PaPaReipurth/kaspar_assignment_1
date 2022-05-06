// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/link_station.proto

package link_station

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

// LinkStationServiceClient is the client API for LinkStationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkStationServiceClient interface {
	GetNearest(ctx context.Context, in *Point, opts ...grpc.CallOption) (*LinkStation, error)
}

type linkStationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkStationServiceClient(cc grpc.ClientConnInterface) LinkStationServiceClient {
	return &linkStationServiceClient{cc}
}

func (c *linkStationServiceClient) GetNearest(ctx context.Context, in *Point, opts ...grpc.CallOption) (*LinkStation, error) {
	out := new(LinkStation)
	err := c.cc.Invoke(ctx, "/LinkStationService/GetNearest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkStationServiceServer is the server API for LinkStationService service.
// All implementations must embed UnimplementedLinkStationServiceServer
// for forward compatibility
type LinkStationServiceServer interface {
	GetNearest(context.Context, *Point) (*LinkStation, error)
	mustEmbedUnimplementedLinkStationServiceServer()
}

// UnimplementedLinkStationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLinkStationServiceServer struct {
}

func (UnimplementedLinkStationServiceServer) GetNearest(context.Context, *Point) (*LinkStation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNearest not implemented")
}
func (UnimplementedLinkStationServiceServer) mustEmbedUnimplementedLinkStationServiceServer() {}

// UnsafeLinkStationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkStationServiceServer will
// result in compilation errors.
type UnsafeLinkStationServiceServer interface {
	mustEmbedUnimplementedLinkStationServiceServer()
}

func RegisterLinkStationServiceServer(s grpc.ServiceRegistrar, srv LinkStationServiceServer) {
	s.RegisterService(&LinkStationService_ServiceDesc, srv)
}

func _LinkStationService_GetNearest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkStationServiceServer).GetNearest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LinkStationService/GetNearest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkStationServiceServer).GetNearest(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkStationService_ServiceDesc is the grpc.ServiceDesc for LinkStationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkStationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LinkStationService",
	HandlerType: (*LinkStationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNearest",
			Handler:    _LinkStationService_GetNearest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/link_station.proto",
}