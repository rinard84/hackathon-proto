// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: processor/api.proto

package processor

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
	ProcessorService_CreateTransaction_FullMethodName    = "/backend.processor.api.ProcessorService/CreateTransaction"
	ProcessorService_GetUserBalance_FullMethodName       = "/backend.processor.api.ProcessorService/GetUserBalance"
	ProcessorService_ListUserTransactions_FullMethodName = "/backend.processor.api.ProcessorService/ListUserTransactions"
)

// ProcessorServiceClient is the client API for ProcessorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessorServiceClient interface {
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	GetUserBalance(ctx context.Context, in *GetUserBalanceRequest, opts ...grpc.CallOption) (*GetUserBalanceResponse, error)
	ListUserTransactions(ctx context.Context, in *ListUserTransactionsRequest, opts ...grpc.CallOption) (*ListUserTransactionsResponse, error)
}

type processorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessorServiceClient(cc grpc.ClientConnInterface) ProcessorServiceClient {
	return &processorServiceClient{cc}
}

func (c *processorServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, ProcessorService_CreateTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processorServiceClient) GetUserBalance(ctx context.Context, in *GetUserBalanceRequest, opts ...grpc.CallOption) (*GetUserBalanceResponse, error) {
	out := new(GetUserBalanceResponse)
	err := c.cc.Invoke(ctx, ProcessorService_GetUserBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processorServiceClient) ListUserTransactions(ctx context.Context, in *ListUserTransactionsRequest, opts ...grpc.CallOption) (*ListUserTransactionsResponse, error) {
	out := new(ListUserTransactionsResponse)
	err := c.cc.Invoke(ctx, ProcessorService_ListUserTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorServiceServer is the server API for ProcessorService service.
// All implementations must embed UnimplementedProcessorServiceServer
// for forward compatibility
type ProcessorServiceServer interface {
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	GetUserBalance(context.Context, *GetUserBalanceRequest) (*GetUserBalanceResponse, error)
	ListUserTransactions(context.Context, *ListUserTransactionsRequest) (*ListUserTransactionsResponse, error)
	mustEmbedUnimplementedProcessorServiceServer()
}

// UnimplementedProcessorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProcessorServiceServer struct {
}

func (UnimplementedProcessorServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedProcessorServiceServer) GetUserBalance(context.Context, *GetUserBalanceRequest) (*GetUserBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBalance not implemented")
}
func (UnimplementedProcessorServiceServer) ListUserTransactions(context.Context, *ListUserTransactionsRequest) (*ListUserTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserTransactions not implemented")
}
func (UnimplementedProcessorServiceServer) mustEmbedUnimplementedProcessorServiceServer() {}

// UnsafeProcessorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessorServiceServer will
// result in compilation errors.
type UnsafeProcessorServiceServer interface {
	mustEmbedUnimplementedProcessorServiceServer()
}

func RegisterProcessorServiceServer(s grpc.ServiceRegistrar, srv ProcessorServiceServer) {
	s.RegisterService(&ProcessorService_ServiceDesc, srv)
}

func _ProcessorService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessorService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessorService_GetUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServiceServer).GetUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessorService_GetUserBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServiceServer).GetUserBalance(ctx, req.(*GetUserBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessorService_ListUserTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServiceServer).ListUserTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessorService_ListUserTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServiceServer).ListUserTransactions(ctx, req.(*ListUserTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProcessorService_ServiceDesc is the grpc.ServiceDesc for ProcessorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backend.processor.api.ProcessorService",
	HandlerType: (*ProcessorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransaction",
			Handler:    _ProcessorService_CreateTransaction_Handler,
		},
		{
			MethodName: "GetUserBalance",
			Handler:    _ProcessorService_GetUserBalance_Handler,
		},
		{
			MethodName: "ListUserTransactions",
			Handler:    _ProcessorService_ListUserTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "processor/api.proto",
}
