// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: processor/data.proto

package processor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Action int32

const (
	Action_CREDIT Action = 0
	Action_DEBIT  Action = 1
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "CREDIT",
		1: "DEBIT",
	}
	Action_value = map[string]int32{
		"CREDIT": 0,
		"DEBIT":  1,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_processor_data_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_processor_data_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{0}
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           int64                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	mi := &file_processor_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ResetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResetRequest) Reset() {
	*x = ResetRequest{}
	mi := &file_processor_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetRequest) ProtoMessage() {}

func (x *ResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetRequest.ProtoReflect.Descriptor instead.
func (*ResetRequest) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{1}
}

type CreditRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid           int64                  `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Amount        int32                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreditRequest) Reset() {
	*x = CreditRequest{}
	mi := &file_processor_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditRequest) ProtoMessage() {}

func (x *CreditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditRequest.ProtoReflect.Descriptor instead.
func (*CreditRequest) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{2}
}

func (x *CreditRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreditRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CreditRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreditResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Balance       uint32                 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreditResponse) Reset() {
	*x = CreditResponse{}
	mi := &file_processor_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditResponse) ProtoMessage() {}

func (x *CreditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditResponse.ProtoReflect.Descriptor instead.
func (*CreditResponse) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{3}
}

func (x *CreditResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreditResponse) GetBalance() uint32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type DebitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid           int64                  `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Amount        int32                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DebitRequest) Reset() {
	*x = DebitRequest{}
	mi := &file_processor_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DebitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebitRequest) ProtoMessage() {}

func (x *DebitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebitRequest.ProtoReflect.Descriptor instead.
func (*DebitRequest) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{4}
}

func (x *DebitRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DebitRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DebitRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type DebitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Balance       uint32                 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DebitResponse) Reset() {
	*x = DebitResponse{}
	mi := &file_processor_data_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DebitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebitResponse) ProtoMessage() {}

func (x *DebitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebitResponse.ProtoReflect.Descriptor instead.
func (*DebitResponse) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{5}
}

func (x *DebitResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DebitResponse) GetBalance() uint32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type GetUserBalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           int64                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserBalanceRequest) Reset() {
	*x = GetUserBalanceRequest{}
	mi := &file_processor_data_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBalanceRequest) ProtoMessage() {}

func (x *GetUserBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetUserBalanceRequest) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserBalanceRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetUserBalanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Balance       uint32                 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserBalanceResponse) Reset() {
	*x = GetUserBalanceResponse{}
	mi := &file_processor_data_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBalanceResponse) ProtoMessage() {}

func (x *GetUserBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetUserBalanceResponse) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserBalanceResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetUserBalanceResponse) GetBalance() uint32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type ListUserTransactionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           int64                  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Limit         int64                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page          int64                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserTransactionsRequest) Reset() {
	*x = ListUserTransactionsRequest{}
	mi := &file_processor_data_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTransactionsRequest) ProtoMessage() {}

func (x *ListUserTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserTransactionsRequest.ProtoReflect.Descriptor instead.
func (*ListUserTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{8}
}

func (x *ListUserTransactionsRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ListUserTransactionsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListUserTransactionsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ListUserTransactionsResponse struct {
	state         protoimpl.MessageState               `protogen:"open.v1"`
	Msg           string                               `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Data          []*ListUserTransactionsResponse_Data `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Total         int32                                `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserTransactionsResponse) Reset() {
	*x = ListUserTransactionsResponse{}
	mi := &file_processor_data_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTransactionsResponse) ProtoMessage() {}

func (x *ListUserTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserTransactionsResponse.ProtoReflect.Descriptor instead.
func (*ListUserTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{9}
}

func (x *ListUserTransactionsResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ListUserTransactionsResponse) GetData() []*ListUserTransactionsResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListUserTransactionsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type StreamTransactionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid           int64                  `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Amo           int32                  `protobuf:"varint,3,opt,name=amo,proto3" json:"amo,omitempty"`
	Act           Action                 `protobuf:"varint,4,opt,name=act,proto3,enum=backend.processor.api.Action" json:"act,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamTransactionRequest) Reset() {
	*x = StreamTransactionRequest{}
	mi := &file_processor_data_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamTransactionRequest) ProtoMessage() {}

func (x *StreamTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamTransactionRequest.ProtoReflect.Descriptor instead.
func (*StreamTransactionRequest) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{10}
}

func (x *StreamTransactionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StreamTransactionRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *StreamTransactionRequest) GetAmo() int32 {
	if x != nil {
		return x.Amo
	}
	return 0
}

func (x *StreamTransactionRequest) GetAct() Action {
	if x != nil {
		return x.Act
	}
	return Action_CREDIT
}

type StreamTransactionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	B             uint32                 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	T             int64                  `protobuf:"varint,3,opt,name=t,proto3" json:"t,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamTransactionResponse) Reset() {
	*x = StreamTransactionResponse{}
	mi := &file_processor_data_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamTransactionResponse) ProtoMessage() {}

func (x *StreamTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamTransactionResponse.ProtoReflect.Descriptor instead.
func (*StreamTransactionResponse) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{11}
}

func (x *StreamTransactionResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *StreamTransactionResponse) GetB() uint32 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *StreamTransactionResponse) GetT() int64 {
	if x != nil {
		return x.T
	}
	return 0
}

type ListUserTransactionsResponse_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount        int32                  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Action        string                 `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserTransactionsResponse_Data) Reset() {
	*x = ListUserTransactionsResponse_Data{}
	mi := &file_processor_data_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserTransactionsResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTransactionsResponse_Data) ProtoMessage() {}

func (x *ListUserTransactionsResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_processor_data_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserTransactionsResponse_Data.ProtoReflect.Descriptor instead.
func (*ListUserTransactionsResponse_Data) Descriptor() ([]byte, []int) {
	return file_processor_data_proto_rawDescGZIP(), []int{9, 0}
}

func (x *ListUserTransactionsResponse_Data) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListUserTransactionsResponse_Data) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ListUserTransactionsResponse_Data) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

var File_processor_data_proto protoreflect.FileDescriptor

var file_processor_data_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x28, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x48, 0x0a, 0x0c, 0x44, 0x65, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x0d, 0x44, 0x65,
	0x62, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x29, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x22, 0x44, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x59, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x22, 0xdc, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x4c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x1a, 0x46, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x6d, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61,
	0x6d, 0x6f, 0x12, 0x2f, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03,
	0x61, 0x63, 0x74, 0x22, 0x49, 0x0a, 0x19, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x62,
	0x12, 0x0c, 0x0a, 0x01, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x74, 0x2a, 0x1f,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x44,
	0x49, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x49, 0x54, 0x10, 0x01, 0x42,
	0xc5, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x6e, 0x61, 0x72, 0x64, 0x38, 0x34, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0xa2, 0x02, 0x03, 0x42, 0x50, 0x41, 0xaa, 0x02, 0x15, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x41, 0x70,
	0x69, 0xca, 0x02, 0x15, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5c, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x5c, 0x41, 0x70, 0x69, 0xe2, 0x02, 0x21, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x5c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5c, 0x41, 0x70,
	0x69, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_processor_data_proto_rawDescOnce sync.Once
	file_processor_data_proto_rawDescData []byte
)

func file_processor_data_proto_rawDescGZIP() []byte {
	file_processor_data_proto_rawDescOnce.Do(func() {
		file_processor_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_processor_data_proto_rawDesc), len(file_processor_data_proto_rawDesc)))
	})
	return file_processor_data_proto_rawDescData
}

var file_processor_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_processor_data_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_processor_data_proto_goTypes = []any{
	(Action)(0),                               // 0: backend.processor.api.Action
	(*CreateAccountRequest)(nil),              // 1: backend.processor.api.CreateAccountRequest
	(*ResetRequest)(nil),                      // 2: backend.processor.api.ResetRequest
	(*CreditRequest)(nil),                     // 3: backend.processor.api.CreditRequest
	(*CreditResponse)(nil),                    // 4: backend.processor.api.CreditResponse
	(*DebitRequest)(nil),                      // 5: backend.processor.api.DebitRequest
	(*DebitResponse)(nil),                     // 6: backend.processor.api.DebitResponse
	(*GetUserBalanceRequest)(nil),             // 7: backend.processor.api.GetUserBalanceRequest
	(*GetUserBalanceResponse)(nil),            // 8: backend.processor.api.GetUserBalanceResponse
	(*ListUserTransactionsRequest)(nil),       // 9: backend.processor.api.ListUserTransactionsRequest
	(*ListUserTransactionsResponse)(nil),      // 10: backend.processor.api.ListUserTransactionsResponse
	(*StreamTransactionRequest)(nil),          // 11: backend.processor.api.StreamTransactionRequest
	(*StreamTransactionResponse)(nil),         // 12: backend.processor.api.StreamTransactionResponse
	(*ListUserTransactionsResponse_Data)(nil), // 13: backend.processor.api.ListUserTransactionsResponse.Data
}
var file_processor_data_proto_depIdxs = []int32{
	13, // 0: backend.processor.api.ListUserTransactionsResponse.data:type_name -> backend.processor.api.ListUserTransactionsResponse.Data
	0,  // 1: backend.processor.api.StreamTransactionRequest.act:type_name -> backend.processor.api.Action
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_processor_data_proto_init() }
func file_processor_data_proto_init() {
	if File_processor_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_processor_data_proto_rawDesc), len(file_processor_data_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_processor_data_proto_goTypes,
		DependencyIndexes: file_processor_data_proto_depIdxs,
		EnumInfos:         file_processor_data_proto_enumTypes,
		MessageInfos:      file_processor_data_proto_msgTypes,
	}.Build()
	File_processor_data_proto = out.File
	file_processor_data_proto_goTypes = nil
	file_processor_data_proto_depIdxs = nil
}
