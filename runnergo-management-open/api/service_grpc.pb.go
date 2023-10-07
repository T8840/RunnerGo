// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: api/service.proto

package services

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

// KpControllerClient is the client API for KpController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KpControllerClient interface {
	// 注册机器
	RegisterMachine(ctx context.Context, in *RegisterMachineReq, opts ...grpc.CallOption) (*RegisterMachineResp, error)
	// 发起压测任务
	RunStress(ctx context.Context, in *RunStressReq, opts ...grpc.CallOption) (*RunStressResp, error)
	// 通知压测任务停止
	NotifyStopStress(ctx context.Context, in *NotifyStopStressReq, opts ...grpc.CallOption) (*NotifyStopStressResp, error)
}

type kpControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewKpControllerClient(cc grpc.ClientConnInterface) KpControllerClient {
	return &kpControllerClient{cc}
}

func (c *kpControllerClient) RegisterMachine(ctx context.Context, in *RegisterMachineReq, opts ...grpc.CallOption) (*RegisterMachineResp, error) {
	out := new(RegisterMachineResp)
	err := c.cc.Invoke(ctx, "/service.KpController/RegisterMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kpControllerClient) RunStress(ctx context.Context, in *RunStressReq, opts ...grpc.CallOption) (*RunStressResp, error) {
	out := new(RunStressResp)
	err := c.cc.Invoke(ctx, "/service.KpController/RunStress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kpControllerClient) NotifyStopStress(ctx context.Context, in *NotifyStopStressReq, opts ...grpc.CallOption) (*NotifyStopStressResp, error) {
	out := new(NotifyStopStressResp)
	err := c.cc.Invoke(ctx, "/service.KpController/NotifyStopStress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KpControllerServer is the server API for KpController service.
// All implementations must embed UnimplementedKpControllerServer
// for forward compatibility
type KpControllerServer interface {
	// 注册机器
	RegisterMachine(context.Context, *RegisterMachineReq) (*RegisterMachineResp, error)
	// 发起压测任务
	RunStress(context.Context, *RunStressReq) (*RunStressResp, error)
	// 通知压测任务停止
	NotifyStopStress(context.Context, *NotifyStopStressReq) (*NotifyStopStressResp, error)
	mustEmbedUnimplementedKpControllerServer()
}

// UnimplementedKpControllerServer must be embedded to have forward compatible implementations.
type UnimplementedKpControllerServer struct {
}

func (UnimplementedKpControllerServer) RegisterMachine(context.Context, *RegisterMachineReq) (*RegisterMachineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMachine not implemented")
}
func (UnimplementedKpControllerServer) RunStress(context.Context, *RunStressReq) (*RunStressResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunStress not implemented")
}
func (UnimplementedKpControllerServer) NotifyStopStress(context.Context, *NotifyStopStressReq) (*NotifyStopStressResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyStopStress not implemented")
}
func (UnimplementedKpControllerServer) mustEmbedUnimplementedKpControllerServer() {}

// UnsafeKpControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KpControllerServer will
// result in compilation errors.
type UnsafeKpControllerServer interface {
	mustEmbedUnimplementedKpControllerServer()
}

func RegisterKpControllerServer(s grpc.ServiceRegistrar, srv KpControllerServer) {
	s.RegisterService(&KpController_ServiceDesc, srv)
}

func _KpController_RegisterMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterMachineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KpControllerServer).RegisterMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.KpController/RegisterMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KpControllerServer).RegisterMachine(ctx, req.(*RegisterMachineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KpController_RunStress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunStressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KpControllerServer).RunStress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.KpController/RunStress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KpControllerServer).RunStress(ctx, req.(*RunStressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KpController_NotifyStopStress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyStopStressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KpControllerServer).NotifyStopStress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.KpController/NotifyStopStress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KpControllerServer).NotifyStopStress(ctx, req.(*NotifyStopStressReq))
	}
	return interceptor(ctx, in, info, handler)
}

// KpController_ServiceDesc is the grpc.ServiceDesc for KpController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KpController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.KpController",
	HandlerType: (*KpControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterMachine",
			Handler:    _KpController_RegisterMachine_Handler,
		},
		{
			MethodName: "RunStress",
			Handler:    _KpController_RunStress_Handler,
		},
		{
			MethodName: "NotifyStopStress",
			Handler:    _KpController_NotifyStopStress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}
