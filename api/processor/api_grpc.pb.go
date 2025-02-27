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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProcessorService_CreateAccount_FullMethodName        = "/backend.processor.api.ProcessorService/CreateAccount"
	ProcessorService_Reset_FullMethodName                = "/backend.processor.api.ProcessorService/Reset"
	ProcessorService_Credit_FullMethodName               = "/backend.processor.api.ProcessorService/Credit"
	ProcessorService_Debit_FullMethodName                = "/backend.processor.api.ProcessorService/Debit"
	ProcessorService_GetUserBalance_FullMethodName       = "/backend.processor.api.ProcessorService/GetUserBalance"
	ProcessorService_ListUserTransactions_FullMethodName = "/backend.processor.api.ProcessorService/ListUserTransactions"
	ProcessorService_CheckTransaction_FullMethodName     = "/backend.processor.api.ProcessorService/CheckTransaction"
	ProcessorService_StreamTransaction_FullMethodName    = "/backend.processor.api.ProcessorService/StreamTransaction"
	ProcessorService_ListWrongBalance_FullMethodName     = "/backend.processor.api.ProcessorService/ListWrongBalance"
)

// ProcessorServiceClient is the client API for ProcessorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessorServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Credit(ctx context.Context, in *CreditRequest, opts ...grpc.CallOption) (*CreditResponse, error)
	Debit(ctx context.Context, in *DebitRequest, opts ...grpc.CallOption) (*DebitResponse, error)
	GetUserBalance(ctx context.Context, in *GetUserBalanceRequest, opts ...grpc.CallOption) (*GetUserBalanceResponse, error)
	ListUserTransactions(ctx context.Context, in *ListUserTransactionsRequest, opts ...grpc.CallOption) (*ListUserTransactionsResponse, error)
	CheckTransaction(ctx context.Context, in *CheckTransactionRequest, opts ...grpc.CallOption) (*CheckTransactionResponse, error)
	StreamTransaction(ctx context.Context, opts ...grpc.CallOption) (ProcessorService_StreamTransactionClient, error)
	ListWrongBalance(ctx context.Context, in *ListWrongBalanceRequest, opts ...grpc.CallOption) (*ListWrongBalanceResponse, error)
}

type processorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessorServiceClient(cc grpc.ClientConnInterface) ProcessorServiceClient {
	return &processorServiceClient{cc}
}

func (c *processorServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProcessorService_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processorServiceClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProcessorService_Reset_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processorServiceClient) Credit(ctx context.Context, in *CreditRequest, opts ...grpc.CallOption) (*CreditResponse, error) {
	out := new(CreditResponse)
	err := c.cc.Invoke(ctx, ProcessorService_Credit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processorServiceClient) Debit(ctx context.Context, in *DebitRequest, opts ...grpc.CallOption) (*DebitResponse, error) {
	out := new(DebitResponse)
	err := c.cc.Invoke(ctx, ProcessorService_Debit_FullMethodName, in, out, opts...)
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

func (c *processorServiceClient) CheckTransaction(ctx context.Context, in *CheckTransactionRequest, opts ...grpc.CallOption) (*CheckTransactionResponse, error) {
	out := new(CheckTransactionResponse)
	err := c.cc.Invoke(ctx, ProcessorService_CheckTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processorServiceClient) StreamTransaction(ctx context.Context, opts ...grpc.CallOption) (ProcessorService_StreamTransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProcessorService_ServiceDesc.Streams[0], ProcessorService_StreamTransaction_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &processorServiceStreamTransactionClient{stream}
	return x, nil
}

type ProcessorService_StreamTransactionClient interface {
	Send(*StreamTransactionRequest) error
	Recv() (*StreamTransactionResponse, error)
	grpc.ClientStream
}

type processorServiceStreamTransactionClient struct {
	grpc.ClientStream
}

func (x *processorServiceStreamTransactionClient) Send(m *StreamTransactionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *processorServiceStreamTransactionClient) Recv() (*StreamTransactionResponse, error) {
	m := new(StreamTransactionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processorServiceClient) ListWrongBalance(ctx context.Context, in *ListWrongBalanceRequest, opts ...grpc.CallOption) (*ListWrongBalanceResponse, error) {
	out := new(ListWrongBalanceResponse)
	err := c.cc.Invoke(ctx, ProcessorService_ListWrongBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorServiceServer is the server API for ProcessorService service.
// All implementations must embed UnimplementedProcessorServiceServer
// for forward compatibility
type ProcessorServiceServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*emptypb.Empty, error)
	Reset(context.Context, *ResetRequest) (*emptypb.Empty, error)
	Credit(context.Context, *CreditRequest) (*CreditResponse, error)
	Debit(context.Context, *DebitRequest) (*DebitResponse, error)
	GetUserBalance(context.Context, *GetUserBalanceRequest) (*GetUserBalanceResponse, error)
	ListUserTransactions(context.Context, *ListUserTransactionsRequest) (*ListUserTransactionsResponse, error)
	CheckTransaction(context.Context, *CheckTransactionRequest) (*CheckTransactionResponse, error)
	StreamTransaction(ProcessorService_StreamTransactionServer) error
	ListWrongBalance(context.Context, *ListWrongBalanceRequest) (*ListWrongBalanceResponse, error)
	mustEmbedUnimplementedProcessorServiceServer()
}

// UnimplementedProcessorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProcessorServiceServer struct {
}

func (UnimplementedProcessorServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedProcessorServiceServer) Reset(context.Context, *ResetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (UnimplementedProcessorServiceServer) Credit(context.Context, *CreditRequest) (*CreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Credit not implemented")
}
func (UnimplementedProcessorServiceServer) Debit(context.Context, *DebitRequest) (*DebitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Debit not implemented")
}
func (UnimplementedProcessorServiceServer) GetUserBalance(context.Context, *GetUserBalanceRequest) (*GetUserBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBalance not implemented")
}
func (UnimplementedProcessorServiceServer) ListUserTransactions(context.Context, *ListUserTransactionsRequest) (*ListUserTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserTransactions not implemented")
}
func (UnimplementedProcessorServiceServer) CheckTransaction(context.Context, *CheckTransactionRequest) (*CheckTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTransaction not implemented")
}
func (UnimplementedProcessorServiceServer) StreamTransaction(ProcessorService_StreamTransactionServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTransaction not implemented")
}
func (UnimplementedProcessorServiceServer) ListWrongBalance(context.Context, *ListWrongBalanceRequest) (*ListWrongBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWrongBalance not implemented")
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

func _ProcessorService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessorService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessorService_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServiceServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessorService_Reset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServiceServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessorService_Credit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServiceServer).Credit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessorService_Credit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServiceServer).Credit(ctx, req.(*CreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessorService_Debit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServiceServer).Debit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessorService_Debit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServiceServer).Debit(ctx, req.(*DebitRequest))
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

func _ProcessorService_CheckTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServiceServer).CheckTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessorService_CheckTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServiceServer).CheckTransaction(ctx, req.(*CheckTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessorService_StreamTransaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProcessorServiceServer).StreamTransaction(&processorServiceStreamTransactionServer{stream})
}

type ProcessorService_StreamTransactionServer interface {
	Send(*StreamTransactionResponse) error
	Recv() (*StreamTransactionRequest, error)
	grpc.ServerStream
}

type processorServiceStreamTransactionServer struct {
	grpc.ServerStream
}

func (x *processorServiceStreamTransactionServer) Send(m *StreamTransactionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *processorServiceStreamTransactionServer) Recv() (*StreamTransactionRequest, error) {
	m := new(StreamTransactionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProcessorService_ListWrongBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWrongBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServiceServer).ListWrongBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessorService_ListWrongBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServiceServer).ListWrongBalance(ctx, req.(*ListWrongBalanceRequest))
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
			MethodName: "CreateAccount",
			Handler:    _ProcessorService_CreateAccount_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _ProcessorService_Reset_Handler,
		},
		{
			MethodName: "Credit",
			Handler:    _ProcessorService_Credit_Handler,
		},
		{
			MethodName: "Debit",
			Handler:    _ProcessorService_Debit_Handler,
		},
		{
			MethodName: "GetUserBalance",
			Handler:    _ProcessorService_GetUserBalance_Handler,
		},
		{
			MethodName: "ListUserTransactions",
			Handler:    _ProcessorService_ListUserTransactions_Handler,
		},
		{
			MethodName: "CheckTransaction",
			Handler:    _ProcessorService_CheckTransaction_Handler,
		},
		{
			MethodName: "ListWrongBalance",
			Handler:    _ProcessorService_ListWrongBalance_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTransaction",
			Handler:       _ProcessorService_StreamTransaction_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "processor/api.proto",
}
