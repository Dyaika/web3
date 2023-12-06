// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: length/length.proto

package __

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
	CalcLength_CalcLength_FullMethodName = "/prime.CalcLength/CalcLength"
)

// CalcLengthClient is the client API for CalcLength service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcLengthClient interface {
	CalcLength(ctx context.Context, in *LengthRequest, opts ...grpc.CallOption) (*LengthResponse, error)
}

type calcLengthClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcLengthClient(cc grpc.ClientConnInterface) CalcLengthClient {
	return &calcLengthClient{cc}
}

func (c *calcLengthClient) CalcLength(ctx context.Context, in *LengthRequest, opts ...grpc.CallOption) (*LengthResponse, error) {
	out := new(LengthResponse)
	err := c.cc.Invoke(ctx, CalcLength_CalcLength_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcLengthServer is the server API for CalcLength service.
// All implementations must embed UnimplementedCalcLengthServer
// for forward compatibility
type CalcLengthServer interface {
	CalcLength(context.Context, *LengthRequest) (*LengthResponse, error)
	mustEmbedUnimplementedCalcLengthServer()
}

// UnimplementedCalcLengthServer must be embedded to have forward compatible implementations.
type UnimplementedCalcLengthServer struct {
}

func (UnimplementedCalcLengthServer) CalcLength(context.Context, *LengthRequest) (*LengthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalcLength not implemented")
}
func (UnimplementedCalcLengthServer) mustEmbedUnimplementedCalcLengthServer() {}

// UnsafeCalcLengthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcLengthServer will
// result in compilation errors.
type UnsafeCalcLengthServer interface {
	mustEmbedUnimplementedCalcLengthServer()
}

func RegisterCalcLengthServer(s grpc.ServiceRegistrar, srv CalcLengthServer) {
	s.RegisterService(&CalcLength_ServiceDesc, srv)
}

func _CalcLength_CalcLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LengthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcLengthServer).CalcLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalcLength_CalcLength_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcLengthServer).CalcLength(ctx, req.(*LengthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalcLength_ServiceDesc is the grpc.ServiceDesc for CalcLength service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcLength_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prime.CalcLength",
	HandlerType: (*CalcLengthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalcLength",
			Handler:    _CalcLength_CalcLength_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "length/length.proto",
}