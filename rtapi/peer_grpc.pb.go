// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: peer.proto

package rtapi

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

const (
	NakamaPeerApi_Call_FullMethodName   = "/nakama.peer.NakamaPeerApi/Call"
	NakamaPeerApi_Stream_FullMethodName = "/nakama.peer.NakamaPeerApi/Stream"
)

// NakamaPeerApiClient is the client API for NakamaPeerApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NakamaPeerApiClient interface {
	Call(ctx context.Context, in *NakamaPeer_Envelope, opts ...grpc.CallOption) (*NakamaPeer_Envelope, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (NakamaPeerApi_StreamClient, error)
}

type nakamaPeerApiClient struct {
	cc grpc.ClientConnInterface
}

func NewNakamaPeerApiClient(cc grpc.ClientConnInterface) NakamaPeerApiClient {
	return &nakamaPeerApiClient{cc}
}

func (c *nakamaPeerApiClient) Call(ctx context.Context, in *NakamaPeer_Envelope, opts ...grpc.CallOption) (*NakamaPeer_Envelope, error) {
	out := new(NakamaPeer_Envelope)
	err := c.cc.Invoke(ctx, NakamaPeerApi_Call_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nakamaPeerApiClient) Stream(ctx context.Context, opts ...grpc.CallOption) (NakamaPeerApi_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &NakamaPeerApi_ServiceDesc.Streams[0], NakamaPeerApi_Stream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &nakamaPeerApiStreamClient{stream}
	return x, nil
}

type NakamaPeerApi_StreamClient interface {
	Send(*NakamaPeer_Envelope) error
	Recv() (*NakamaPeer_Envelope, error)
	grpc.ClientStream
}

type nakamaPeerApiStreamClient struct {
	grpc.ClientStream
}

func (x *nakamaPeerApiStreamClient) Send(m *NakamaPeer_Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nakamaPeerApiStreamClient) Recv() (*NakamaPeer_Envelope, error) {
	m := new(NakamaPeer_Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NakamaPeerApiServer is the server API for NakamaPeerApi service.
// All implementations must embed UnimplementedNakamaPeerApiServer
// for forward compatibility
type NakamaPeerApiServer interface {
	Call(context.Context, *NakamaPeer_Envelope) (*NakamaPeer_Envelope, error)
	Stream(NakamaPeerApi_StreamServer) error
	mustEmbedUnimplementedNakamaPeerApiServer()
}

// UnimplementedNakamaPeerApiServer must be embedded to have forward compatible implementations.
type UnimplementedNakamaPeerApiServer struct {
}

func (UnimplementedNakamaPeerApiServer) Call(context.Context, *NakamaPeer_Envelope) (*NakamaPeer_Envelope, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedNakamaPeerApiServer) Stream(NakamaPeerApi_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedNakamaPeerApiServer) mustEmbedUnimplementedNakamaPeerApiServer() {}

// UnsafeNakamaPeerApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NakamaPeerApiServer will
// result in compilation errors.
type UnsafeNakamaPeerApiServer interface {
	mustEmbedUnimplementedNakamaPeerApiServer()
}

func RegisterNakamaPeerApiServer(s grpc.ServiceRegistrar, srv NakamaPeerApiServer) {
	s.RegisterService(&NakamaPeerApi_ServiceDesc, srv)
}

func _NakamaPeerApi_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NakamaPeer_Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NakamaPeerApiServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NakamaPeerApi_Call_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NakamaPeerApiServer).Call(ctx, req.(*NakamaPeer_Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _NakamaPeerApi_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NakamaPeerApiServer).Stream(&nakamaPeerApiStreamServer{stream})
}

type NakamaPeerApi_StreamServer interface {
	Send(*NakamaPeer_Envelope) error
	Recv() (*NakamaPeer_Envelope, error)
	grpc.ServerStream
}

type nakamaPeerApiStreamServer struct {
	grpc.ServerStream
}

func (x *nakamaPeerApiStreamServer) Send(m *NakamaPeer_Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nakamaPeerApiStreamServer) Recv() (*NakamaPeer_Envelope, error) {
	m := new(NakamaPeer_Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NakamaPeerApi_ServiceDesc is the grpc.ServiceDesc for NakamaPeerApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NakamaPeerApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nakama.peer.NakamaPeerApi",
	HandlerType: (*NakamaPeerApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _NakamaPeerApi_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _NakamaPeerApi_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer.proto",
}