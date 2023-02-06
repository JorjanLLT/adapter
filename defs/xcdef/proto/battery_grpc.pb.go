// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: battery.proto

package proto

import (
	context "context"
	entpb "github.com/auroraride/adapter/defs/xcdef/proto/entpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BatteryClient is the client API for Battery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BatteryClient interface {
	GetBatteryDetail(ctx context.Context, in *BatteryQueryRequest, opts ...grpc.CallOption) (*entpb.Battery, error)
	GetBatterySample(ctx context.Context, in *BatteryBatchQueryRequest, opts ...grpc.CallOption) (*BatterySampleResponse, error)
}

type batteryClient struct {
	cc grpc.ClientConnInterface
}

func NewBatteryClient(cc grpc.ClientConnInterface) BatteryClient {
	return &batteryClient{cc}
}

func (c *batteryClient) GetBatteryDetail(ctx context.Context, in *BatteryQueryRequest, opts ...grpc.CallOption) (*entpb.Battery, error) {
	out := new(entpb.Battery)
	err := c.cc.Invoke(ctx, "/proto.Battery/GetBatteryDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batteryClient) GetBatterySample(ctx context.Context, in *BatteryBatchQueryRequest, opts ...grpc.CallOption) (*BatterySampleResponse, error) {
	out := new(BatterySampleResponse)
	err := c.cc.Invoke(ctx, "/proto.Battery/GetBatterySample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BatteryServer is the server API for Battery service.
// All implementations must embed UnimplementedBatteryServer
// for forward compatibility
type BatteryServer interface {
	GetBatteryDetail(context.Context, *BatteryQueryRequest) (*entpb.Battery, error)
	GetBatterySample(context.Context, *BatteryBatchQueryRequest) (*BatterySampleResponse, error)
	mustEmbedUnimplementedBatteryServer()
}

// UnimplementedBatteryServer must be embedded to have forward compatible implementations.
type UnimplementedBatteryServer struct {
}

func (UnimplementedBatteryServer) GetBatteryDetail(context.Context, *BatteryQueryRequest) (*entpb.Battery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatteryDetail not implemented")
}
func (UnimplementedBatteryServer) GetBatterySample(context.Context, *BatteryBatchQueryRequest) (*BatterySampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatterySample not implemented")
}
func (UnimplementedBatteryServer) mustEmbedUnimplementedBatteryServer() {}

// UnsafeBatteryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BatteryServer will
// result in compilation errors.
type UnsafeBatteryServer interface {
	mustEmbedUnimplementedBatteryServer()
}

func RegisterBatteryServer(s grpc.ServiceRegistrar, srv BatteryServer) {
	s.RegisterService(&Battery_ServiceDesc, srv)
}

func _Battery_GetBatteryDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatteryQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatteryServer).GetBatteryDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Battery/GetBatteryDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatteryServer).GetBatteryDetail(ctx, req.(*BatteryQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Battery_GetBatterySample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatteryBatchQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatteryServer).GetBatterySample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Battery/GetBatterySample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatteryServer).GetBatterySample(ctx, req.(*BatteryBatchQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Battery_ServiceDesc is the grpc.ServiceDesc for Battery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Battery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Battery",
	HandlerType: (*BatteryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBatteryDetail",
			Handler:    _Battery_GetBatteryDetail_Handler,
		},
		{
			MethodName: "GetBatterySample",
			Handler:    _Battery_GetBatterySample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "battery.proto",
}
