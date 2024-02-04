// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// GostTunelClient is the client API for GostTunel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GostTunelClient interface {
	Tunnel(ctx context.Context, opts ...grpc.CallOption) (GostTunel_TunnelClient, error)
}

type gostTunelClient struct {
	cc grpc.ClientConnInterface
}

func NewGostTunelClient(cc grpc.ClientConnInterface) GostTunelClient {
	return &gostTunelClient{cc}
}

func (c *gostTunelClient) Tunnel(ctx context.Context, opts ...grpc.CallOption) (GostTunel_TunnelClient, error) {
	stream, err := c.cc.NewStream(ctx, &GostTunel_ServiceDesc.Streams[0], "/GostTunel/Tunnel", opts...)
	if err != nil {
		return nil, err
	}
	x := &gostTunelTunnelClient{stream}
	return x, nil
}

type GostTunel_TunnelClient interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type gostTunelTunnelClient struct {
	grpc.ClientStream
}

func (x *gostTunelTunnelClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gostTunelTunnelClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GostTunelServer is the server API for GostTunel service.
// All implementations must embed UnimplementedGostTunelServer
// for forward compatibility
type GostTunelServer interface {
	Tunnel(GostTunel_TunnelServer) error
	mustEmbedUnimplementedGostTunelServer()
}

// UnimplementedGostTunelServer must be embedded to have forward compatible implementations.
type UnimplementedGostTunelServer struct {
}

func (UnimplementedGostTunelServer) Tunnel(GostTunel_TunnelServer) error {
	return status.Errorf(codes.Unimplemented, "method Tunnel not implemented")
}
func (UnimplementedGostTunelServer) mustEmbedUnimplementedGostTunelServer() {}

// UnsafeGostTunelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GostTunelServer will
// result in compilation errors.
type UnsafeGostTunelServer interface {
	mustEmbedUnimplementedGostTunelServer()
}

func RegisterGostTunelServer(s grpc.ServiceRegistrar, srv GostTunelServer) {
	s.RegisterService(&GostTunel_ServiceDesc, srv)
}

func _GostTunel_Tunnel_Handler(srv any, stream grpc.ServerStream) error {
	return srv.(GostTunelServer).Tunnel(&gostTunelTunnelServer{stream})
}

type GostTunel_TunnelServer interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type gostTunelTunnelServer struct {
	grpc.ServerStream
}

func (x *gostTunelTunnelServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gostTunelTunnelServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GostTunel_ServiceDesc is the grpc.ServiceDesc for GostTunel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GostTunel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GostTunel",
	HandlerType: (*GostTunelServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tunnel",
			Handler:       _GostTunel_Tunnel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gost.proto",
}
