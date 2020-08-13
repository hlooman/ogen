// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: chain.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SubscribeValidatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey [][]byte `protobuf:"bytes,1,rep,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *SubscribeValidatorRequest) Reset() {
	*x = SubscribeValidatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeValidatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeValidatorRequest) ProtoMessage() {}

func (x *SubscribeValidatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeValidatorRequest.ProtoReflect.Descriptor instead.
func (*SubscribeValidatorRequest) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeValidatorRequest) GetPublicKey() [][]byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type ChainInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash   string          `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	BlockHeight uint64          `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Validators  *ValidatorsInfo `protobuf:"bytes,3,opt,name=validators,proto3" json:"validators,omitempty"`
}

func (x *ChainInfo) Reset() {
	*x = ChainInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainInfo) ProtoMessage() {}

func (x *ChainInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainInfo.ProtoReflect.Descriptor instead.
func (*ChainInfo) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{1}
}

func (x *ChainInfo) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *ChainInfo) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *ChainInfo) GetValidators() *ValidatorsInfo {
	if x != nil {
		return x.Validators
	}
	return nil
}

type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Balance *Balance `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce   uint64   `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Txs     []string `protobuf:"bytes,4,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_chain_proto_rawDescGZIP(), []int{2}
}

func (x *AccountInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AccountInfo) GetBalance() *Balance {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *AccountInfo) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *AccountInfo) GetTxs() []string {
	if x != nil {
		return x.Txs
	}
	return nil
}

var File_chain_proto protoreflect.FileDescriptor

var file_chain_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x19, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x7e, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x73, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x78, 0x73, 0x32, 0xb6, 0x05, 0x0a, 0x05, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x12, 0x3c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x06, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2f, 0x72, 0x61, 0x77, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x7b, 0x68, 0x61, 0x73, 0x68, 0x7d,
	0x12, 0x36, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x05, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x1a, 0x06, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x1b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x2f, 0x7b, 0x68, 0x61, 0x73, 0x68, 0x7d, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x07, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x1a, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x12, 0x19, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x68, 0x61,
	0x73, 0x68, 0x2f, 0x7b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x7d, 0x12, 0x4a, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x08, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x7b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x12, 0x36, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68,
	0x1a, 0x03, 0x2e, 0x54, 0x78, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x78, 0x2f, 0x7b, 0x68, 0x61, 0x73, 0x68, 0x7d, 0x12,
	0x35, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x08,
	0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x12, 0x12, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x7b, 0x68,
	0x61, 0x73, 0x68, 0x7d, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x08, 0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x30, 0x01, 0x12, 0x53,
	0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69,
	0x72, 0x73, 0x1a, 0x08, 0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x01,
	0x2a, 0x30, 0x01, 0x12, 0x5e, 0x0a, 0x1e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x73,
	0x1a, 0x08, 0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x3a, 0x01,
	0x2a, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chain_proto_rawDescOnce sync.Once
	file_chain_proto_rawDescData = file_chain_proto_rawDesc
)

func file_chain_proto_rawDescGZIP() []byte {
	file_chain_proto_rawDescOnce.Do(func() {
		file_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_chain_proto_rawDescData)
	})
	return file_chain_proto_rawDescData
}

var file_chain_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chain_proto_goTypes = []interface{}{
	(*SubscribeValidatorRequest)(nil), // 0: SubscribeValidatorRequest
	(*ChainInfo)(nil),                 // 1: ChainInfo
	(*AccountInfo)(nil),               // 2: AccountInfo
	(*ValidatorsInfo)(nil),            // 3: ValidatorsInfo
	(*Balance)(nil),                   // 4: Balance
	(*Empty)(nil),                     // 5: Empty
	(*Hash)(nil),                      // 6: Hash
	(*Number)(nil),                    // 7: Number
	(*Account)(nil),                   // 8: Account
	(*KeyPairs)(nil),                  // 9: KeyPairs
	(*Block)(nil),                     // 10: Block
	(*Tx)(nil),                        // 11: Tx
	(*RawData)(nil),                   // 12: RawData
}
var file_chain_proto_depIdxs = []int32{
	3,  // 0: ChainInfo.validators:type_name -> ValidatorsInfo
	4,  // 1: AccountInfo.balance:type_name -> Balance
	5,  // 2: Chain.GetChainInfo:input_type -> Empty
	6,  // 3: Chain.GetRawBlock:input_type -> Hash
	6,  // 4: Chain.GetBlock:input_type -> Hash
	7,  // 5: Chain.GetBlockHash:input_type -> Number
	8,  // 6: Chain.GetAccountInfo:input_type -> Account
	6,  // 7: Chain.GetTransaction:input_type -> Hash
	6,  // 8: Chain.Sync:input_type -> Hash
	5,  // 9: Chain.SubscribeBlocks:input_type -> Empty
	9,  // 10: Chain.SubscribeTransactions:input_type -> KeyPairs
	9,  // 11: Chain.SubscribeValidatorTransactions:input_type -> KeyPairs
	1,  // 12: Chain.GetChainInfo:output_type -> ChainInfo
	10, // 13: Chain.GetRawBlock:output_type -> Block
	10, // 14: Chain.GetBlock:output_type -> Block
	6,  // 15: Chain.GetBlockHash:output_type -> Hash
	2,  // 16: Chain.GetAccountInfo:output_type -> AccountInfo
	11, // 17: Chain.GetTransaction:output_type -> Tx
	12, // 18: Chain.Sync:output_type -> RawData
	12, // 19: Chain.SubscribeBlocks:output_type -> RawData
	12, // 20: Chain.SubscribeTransactions:output_type -> RawData
	12, // 21: Chain.SubscribeValidatorTransactions:output_type -> RawData
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_chain_proto_init() }
func file_chain_proto_init() {
	if File_chain_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeValidatorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chain_proto_goTypes,
		DependencyIndexes: file_chain_proto_depIdxs,
		MessageInfos:      file_chain_proto_msgTypes,
	}.Build()
	File_chain_proto = out.File
	file_chain_proto_rawDesc = nil
	file_chain_proto_goTypes = nil
	file_chain_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChainClient is the client API for Chain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChainClient interface {
	GetChainInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChainInfo, error)
	GetRawBlock(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Block, error)
	GetBlock(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Block, error)
	GetBlockHash(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Hash, error)
	GetAccountInfo(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountInfo, error)
	GetTransaction(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Tx, error)
	Sync(ctx context.Context, in *Hash, opts ...grpc.CallOption) (Chain_SyncClient, error)
	SubscribeBlocks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Chain_SubscribeBlocksClient, error)
	SubscribeTransactions(ctx context.Context, in *KeyPairs, opts ...grpc.CallOption) (Chain_SubscribeTransactionsClient, error)
	SubscribeValidatorTransactions(ctx context.Context, in *KeyPairs, opts ...grpc.CallOption) (Chain_SubscribeValidatorTransactionsClient, error)
}

type chainClient struct {
	cc grpc.ClientConnInterface
}

func NewChainClient(cc grpc.ClientConnInterface) ChainClient {
	return &chainClient{cc}
}

func (c *chainClient) GetChainInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ChainInfo, error) {
	out := new(ChainInfo)
	err := c.cc.Invoke(ctx, "/Chain/GetChainInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetRawBlock(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/Chain/GetRawBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetBlock(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/Chain/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetBlockHash(ctx context.Context, in *Number, opts ...grpc.CallOption) (*Hash, error) {
	out := new(Hash)
	err := c.cc.Invoke(ctx, "/Chain/GetBlockHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetAccountInfo(ctx context.Context, in *Account, opts ...grpc.CallOption) (*AccountInfo, error) {
	out := new(AccountInfo)
	err := c.cc.Invoke(ctx, "/Chain/GetAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) GetTransaction(ctx context.Context, in *Hash, opts ...grpc.CallOption) (*Tx, error) {
	out := new(Tx)
	err := c.cc.Invoke(ctx, "/Chain/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) Sync(ctx context.Context, in *Hash, opts ...grpc.CallOption) (Chain_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chain_serviceDesc.Streams[0], "/Chain/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chain_SyncClient interface {
	Recv() (*RawData, error)
	grpc.ClientStream
}

type chainSyncClient struct {
	grpc.ClientStream
}

func (x *chainSyncClient) Recv() (*RawData, error) {
	m := new(RawData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chainClient) SubscribeBlocks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Chain_SubscribeBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chain_serviceDesc.Streams[1], "/Chain/SubscribeBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainSubscribeBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chain_SubscribeBlocksClient interface {
	Recv() (*RawData, error)
	grpc.ClientStream
}

type chainSubscribeBlocksClient struct {
	grpc.ClientStream
}

func (x *chainSubscribeBlocksClient) Recv() (*RawData, error) {
	m := new(RawData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chainClient) SubscribeTransactions(ctx context.Context, in *KeyPairs, opts ...grpc.CallOption) (Chain_SubscribeTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chain_serviceDesc.Streams[2], "/Chain/SubscribeTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainSubscribeTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chain_SubscribeTransactionsClient interface {
	Recv() (*RawData, error)
	grpc.ClientStream
}

type chainSubscribeTransactionsClient struct {
	grpc.ClientStream
}

func (x *chainSubscribeTransactionsClient) Recv() (*RawData, error) {
	m := new(RawData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chainClient) SubscribeValidatorTransactions(ctx context.Context, in *KeyPairs, opts ...grpc.CallOption) (Chain_SubscribeValidatorTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chain_serviceDesc.Streams[3], "/Chain/SubscribeValidatorTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainSubscribeValidatorTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chain_SubscribeValidatorTransactionsClient interface {
	Recv() (*RawData, error)
	grpc.ClientStream
}

type chainSubscribeValidatorTransactionsClient struct {
	grpc.ClientStream
}

func (x *chainSubscribeValidatorTransactionsClient) Recv() (*RawData, error) {
	m := new(RawData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChainServer is the server API for Chain service.
type ChainServer interface {
	GetChainInfo(context.Context, *Empty) (*ChainInfo, error)
	GetRawBlock(context.Context, *Hash) (*Block, error)
	GetBlock(context.Context, *Hash) (*Block, error)
	GetBlockHash(context.Context, *Number) (*Hash, error)
	GetAccountInfo(context.Context, *Account) (*AccountInfo, error)
	GetTransaction(context.Context, *Hash) (*Tx, error)
	Sync(*Hash, Chain_SyncServer) error
	SubscribeBlocks(*Empty, Chain_SubscribeBlocksServer) error
	SubscribeTransactions(*KeyPairs, Chain_SubscribeTransactionsServer) error
	SubscribeValidatorTransactions(*KeyPairs, Chain_SubscribeValidatorTransactionsServer) error
}

// UnimplementedChainServer can be embedded to have forward compatible implementations.
type UnimplementedChainServer struct {
}

func (*UnimplementedChainServer) GetChainInfo(context.Context, *Empty) (*ChainInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChainInfo not implemented")
}
func (*UnimplementedChainServer) GetRawBlock(context.Context, *Hash) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRawBlock not implemented")
}
func (*UnimplementedChainServer) GetBlock(context.Context, *Hash) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (*UnimplementedChainServer) GetBlockHash(context.Context, *Number) (*Hash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHash not implemented")
}
func (*UnimplementedChainServer) GetAccountInfo(context.Context, *Account) (*AccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
func (*UnimplementedChainServer) GetTransaction(context.Context, *Hash) (*Tx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (*UnimplementedChainServer) Sync(*Hash, Chain_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (*UnimplementedChainServer) SubscribeBlocks(*Empty, Chain_SubscribeBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeBlocks not implemented")
}
func (*UnimplementedChainServer) SubscribeTransactions(*KeyPairs, Chain_SubscribeTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeTransactions not implemented")
}
func (*UnimplementedChainServer) SubscribeValidatorTransactions(*KeyPairs, Chain_SubscribeValidatorTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeValidatorTransactions not implemented")
}

func RegisterChainServer(s *grpc.Server, srv ChainServer) {
	s.RegisterService(&_Chain_serviceDesc, srv)
}

func _Chain_GetChainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetChainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetChainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetChainInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_GetRawBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetRawBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetRawBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetRawBlock(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetBlock(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_GetBlockHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Number)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetBlockHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetBlockHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetBlockHash(ctx, req.(*Number))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetAccountInfo(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).GetTransaction(ctx, req.(*Hash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Hash)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainServer).Sync(m, &chainSyncServer{stream})
}

type Chain_SyncServer interface {
	Send(*RawData) error
	grpc.ServerStream
}

type chainSyncServer struct {
	grpc.ServerStream
}

func (x *chainSyncServer) Send(m *RawData) error {
	return x.ServerStream.SendMsg(m)
}

func _Chain_SubscribeBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainServer).SubscribeBlocks(m, &chainSubscribeBlocksServer{stream})
}

type Chain_SubscribeBlocksServer interface {
	Send(*RawData) error
	grpc.ServerStream
}

type chainSubscribeBlocksServer struct {
	grpc.ServerStream
}

func (x *chainSubscribeBlocksServer) Send(m *RawData) error {
	return x.ServerStream.SendMsg(m)
}

func _Chain_SubscribeTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KeyPairs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainServer).SubscribeTransactions(m, &chainSubscribeTransactionsServer{stream})
}

type Chain_SubscribeTransactionsServer interface {
	Send(*RawData) error
	grpc.ServerStream
}

type chainSubscribeTransactionsServer struct {
	grpc.ServerStream
}

func (x *chainSubscribeTransactionsServer) Send(m *RawData) error {
	return x.ServerStream.SendMsg(m)
}

func _Chain_SubscribeValidatorTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KeyPairs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainServer).SubscribeValidatorTransactions(m, &chainSubscribeValidatorTransactionsServer{stream})
}

type Chain_SubscribeValidatorTransactionsServer interface {
	Send(*RawData) error
	grpc.ServerStream
}

type chainSubscribeValidatorTransactionsServer struct {
	grpc.ServerStream
}

func (x *chainSubscribeValidatorTransactionsServer) Send(m *RawData) error {
	return x.ServerStream.SendMsg(m)
}

var _Chain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Chain",
	HandlerType: (*ChainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChainInfo",
			Handler:    _Chain_GetChainInfo_Handler,
		},
		{
			MethodName: "GetRawBlock",
			Handler:    _Chain_GetRawBlock_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _Chain_GetBlock_Handler,
		},
		{
			MethodName: "GetBlockHash",
			Handler:    _Chain_GetBlockHash_Handler,
		},
		{
			MethodName: "GetAccountInfo",
			Handler:    _Chain_GetAccountInfo_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Chain_GetTransaction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _Chain_Sync_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeBlocks",
			Handler:       _Chain_SubscribeBlocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeTransactions",
			Handler:       _Chain_SubscribeTransactions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeValidatorTransactions",
			Handler:       _Chain_SubscribeValidatorTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chain.proto",
}