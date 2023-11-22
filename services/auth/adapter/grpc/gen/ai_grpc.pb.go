// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: ai.proto

package gen

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
	AiService_GetAiById_FullMethodName = "/AiService/GetAiById"
)

// AiServiceClient is the client API for AiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AiServiceClient interface {
	GetAiById(ctx context.Context, in *GetAiByIdMsg, opts ...grpc.CallOption) (*AI, error)
}

type aiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAiServiceClient(cc grpc.ClientConnInterface) AiServiceClient {
	return &aiServiceClient{cc}
}

func (c *aiServiceClient) GetAiById(ctx context.Context, in *GetAiByIdMsg, opts ...grpc.CallOption) (*AI, error) {
	out := new(AI)
	err := c.cc.Invoke(ctx, AiService_GetAiById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AiServiceServer is the server API for AiService service.
// All implementations must embed UnimplementedAiServiceServer
// for forward compatibility
type AiServiceServer interface {
	GetAiById(context.Context, *GetAiByIdMsg) (*AI, error)
	mustEmbedUnimplementedAiServiceServer()
}

// UnimplementedAiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAiServiceServer struct {
}

func (UnimplementedAiServiceServer) GetAiById(context.Context, *GetAiByIdMsg) (*AI, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAiById not implemented")
}
func (UnimplementedAiServiceServer) mustEmbedUnimplementedAiServiceServer() {}

// UnsafeAiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AiServiceServer will
// result in compilation errors.
type UnsafeAiServiceServer interface {
	mustEmbedUnimplementedAiServiceServer()
}

func RegisterAiServiceServer(s grpc.ServiceRegistrar, srv AiServiceServer) {
	s.RegisterService(&AiService_ServiceDesc, srv)
}

func _AiService_GetAiById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAiByIdMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiServiceServer).GetAiById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AiService_GetAiById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiServiceServer).GetAiById(ctx, req.(*GetAiByIdMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// AiService_ServiceDesc is the grpc.ServiceDesc for AiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AiService",
	HandlerType: (*AiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAiById",
			Handler:    _AiService_GetAiById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ai.proto",
}
