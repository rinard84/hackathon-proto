// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: store/api.proto

package store

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_store_api_proto protoreflect.FileDescriptor

var file_store_api_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xe9, 0x08, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6a, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22,
	0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x57,
	0x0a, 0x05, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x65, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x12, 0x20, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01,
	0x2a, 0x22, 0x0b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x61,
	0x0a, 0x05, 0x44, 0x65, 0x62, 0x69, 0x74, 0x12, 0x1f, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x62, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x62,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x62, 0x69,
	0x74, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75,
	0x69, 0x64, 0x7d, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x90, 0x01, 0x0a,
	0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x80, 0x01, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x28, 0x01, 0x12, 0x88, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x72, 0x6f, 0x6e, 0x67,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x72, 0x6f, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x72, 0x6f, 0x6e,
	0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77,
	0x72, 0x6f, 0x6e, 0x67, 0x2d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0xac, 0x01,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x69, 0x6e, 0x61, 0x72, 0x64, 0x38, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x42, 0x53, 0x41, 0xaa, 0x02,
	0x11, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41,
	0x70, 0x69, 0xca, 0x02, 0x11, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5c, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x5c, 0x41, 0x70, 0x69, 0xe2, 0x02, 0x1d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var file_store_api_proto_goTypes = []any{
	(*CreateAccountRequest)(nil),         // 0: backend.store.api.CreateAccountRequest
	(*ResetRequest)(nil),                 // 1: backend.store.api.ResetRequest
	(*CreditRequest)(nil),                // 2: backend.store.api.CreditRequest
	(*DebitRequest)(nil),                 // 3: backend.store.api.DebitRequest
	(*GetUserBalanceRequest)(nil),        // 4: backend.store.api.GetUserBalanceRequest
	(*ListUserTransactionsRequest)(nil),  // 5: backend.store.api.ListUserTransactionsRequest
	(*CheckTransactionRequest)(nil),      // 6: backend.store.api.CheckTransactionRequest
	(*StreamTransactionRequest)(nil),     // 7: backend.store.api.StreamTransactionRequest
	(*ListWrongBalanceRequest)(nil),      // 8: backend.store.api.ListWrongBalanceRequest
	(*emptypb.Empty)(nil),                // 9: google.protobuf.Empty
	(*CreditResponse)(nil),               // 10: backend.store.api.CreditResponse
	(*DebitResponse)(nil),                // 11: backend.store.api.DebitResponse
	(*GetUserBalanceResponse)(nil),       // 12: backend.store.api.GetUserBalanceResponse
	(*ListUserTransactionsResponse)(nil), // 13: backend.store.api.ListUserTransactionsResponse
	(*CheckTransactionResponse)(nil),     // 14: backend.store.api.CheckTransactionResponse
	(*ListWrongBalanceResponse)(nil),     // 15: backend.store.api.ListWrongBalanceResponse
}
var file_store_api_proto_depIdxs = []int32{
	0,  // 0: backend.store.api.StoreService.CreateAccount:input_type -> backend.store.api.CreateAccountRequest
	1,  // 1: backend.store.api.StoreService.Reset:input_type -> backend.store.api.ResetRequest
	2,  // 2: backend.store.api.StoreService.Credit:input_type -> backend.store.api.CreditRequest
	3,  // 3: backend.store.api.StoreService.Debit:input_type -> backend.store.api.DebitRequest
	4,  // 4: backend.store.api.StoreService.GetUserBalance:input_type -> backend.store.api.GetUserBalanceRequest
	5,  // 5: backend.store.api.StoreService.ListUserTransactions:input_type -> backend.store.api.ListUserTransactionsRequest
	6,  // 6: backend.store.api.StoreService.CheckTransaction:input_type -> backend.store.api.CheckTransactionRequest
	7,  // 7: backend.store.api.StoreService.StreamTransaction:input_type -> backend.store.api.StreamTransactionRequest
	8,  // 8: backend.store.api.StoreService.ListWrongBalance:input_type -> backend.store.api.ListWrongBalanceRequest
	9,  // 9: backend.store.api.StoreService.CreateAccount:output_type -> google.protobuf.Empty
	9,  // 10: backend.store.api.StoreService.Reset:output_type -> google.protobuf.Empty
	10, // 11: backend.store.api.StoreService.Credit:output_type -> backend.store.api.CreditResponse
	11, // 12: backend.store.api.StoreService.Debit:output_type -> backend.store.api.DebitResponse
	12, // 13: backend.store.api.StoreService.GetUserBalance:output_type -> backend.store.api.GetUserBalanceResponse
	13, // 14: backend.store.api.StoreService.ListUserTransactions:output_type -> backend.store.api.ListUserTransactionsResponse
	14, // 15: backend.store.api.StoreService.CheckTransaction:output_type -> backend.store.api.CheckTransactionResponse
	9,  // 16: backend.store.api.StoreService.StreamTransaction:output_type -> google.protobuf.Empty
	15, // 17: backend.store.api.StoreService.ListWrongBalance:output_type -> backend.store.api.ListWrongBalanceResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_store_api_proto_init() }
func file_store_api_proto_init() {
	if File_store_api_proto != nil {
		return
	}
	file_store_data_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_store_api_proto_rawDesc), len(file_store_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_store_api_proto_goTypes,
		DependencyIndexes: file_store_api_proto_depIdxs,
	}.Build()
	File_store_api_proto = out.File
	file_store_api_proto_goTypes = nil
	file_store_api_proto_depIdxs = nil
}
