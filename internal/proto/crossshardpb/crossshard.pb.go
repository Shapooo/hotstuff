// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: internal/proto/crossshardpb/crossshard.proto

package crossshardpb

import (
	_ "github.com/relab/gorums"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UseMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// repeated uint64 rrbList  = 2;
	// repeated uint64 hprice   = 3;
	Content   []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	TxHash    []byte `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *UseMsg) Reset() {
	*x = UseMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseMsg) ProtoMessage() {}

func (x *UseMsg) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseMsg.ProtoReflect.Descriptor instead.
func (*UseMsg) Descriptor() ([]byte, []int) {
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP(), []int{0}
}

func (x *UseMsg) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UseMsg) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *UseMsg) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *UseMsg) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type OcuMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// repeated uint64 vrrbList = 2;
	// repeated uint64 rrbList  = 3;
	Content   []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	TxHash    []byte `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *OcuMsg) Reset() {
	*x = OcuMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OcuMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OcuMsg) ProtoMessage() {}

func (x *OcuMsg) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OcuMsg.ProtoReflect.Descriptor instead.
func (*OcuMsg) Descriptor() ([]byte, []int) {
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP(), []int{1}
}

func (x *OcuMsg) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *OcuMsg) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *OcuMsg) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *OcuMsg) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Ocu2Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TxHash    []byte `protobuf:"bytes,2,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Ocu2Msg) Reset() {
	*x = Ocu2Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ocu2Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ocu2Msg) ProtoMessage() {}

func (x *Ocu2Msg) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ocu2Msg.ProtoReflect.Descriptor instead.
func (*Ocu2Msg) Descriptor() ([]byte, []int) {
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP(), []int{2}
}

func (x *Ocu2Msg) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Ocu2Msg) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *Ocu2Msg) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ReplyMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TxHash    []byte `protobuf:"bytes,2,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ReplyMsg) Reset() {
	*x = ReplyMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMsg) ProtoMessage() {}

func (x *ReplyMsg) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMsg.ProtoReflect.Descriptor instead.
func (*ReplyMsg) Descriptor() ([]byte, []int) {
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyMsg) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ReplyMsg) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *ReplyMsg) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ECDSASignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer uint32 `protobuf:"varint,1,opt,name=Signer,proto3" json:"Signer,omitempty"`
	R      []byte `protobuf:"bytes,2,opt,name=R,proto3" json:"R,omitempty"`
	S      []byte `protobuf:"bytes,3,opt,name=S,proto3" json:"S,omitempty"`
}

func (x *ECDSASignature) Reset() {
	*x = ECDSASignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECDSASignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECDSASignature) ProtoMessage() {}

func (x *ECDSASignature) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECDSASignature.ProtoReflect.Descriptor instead.
func (*ECDSASignature) Descriptor() ([]byte, []int) {
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP(), []int{4}
}

func (x *ECDSASignature) GetSigner() uint32 {
	if x != nil {
		return x.Signer
	}
	return 0
}

func (x *ECDSASignature) GetR() []byte {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *ECDSASignature) GetS() []byte {
	if x != nil {
		return x.S
	}
	return nil
}

type BLS12Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sig []byte `protobuf:"bytes,1,opt,name=Sig,proto3" json:"Sig,omitempty"`
}

func (x *BLS12Signature) Reset() {
	*x = BLS12Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLS12Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLS12Signature) ProtoMessage() {}

func (x *BLS12Signature) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLS12Signature.ProtoReflect.Descriptor instead.
func (*BLS12Signature) Descriptor() ([]byte, []int) {
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP(), []int{5}
}

func (x *BLS12Signature) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Sig:
	//
	//	*Signature_ECDSASig
	//	*Signature_BLS12Sig
	Sig isSignature_Sig `protobuf_oneof:"Sig"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP(), []int{6}
}

func (m *Signature) GetSig() isSignature_Sig {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (x *Signature) GetECDSASig() *ECDSASignature {
	if x, ok := x.GetSig().(*Signature_ECDSASig); ok {
		return x.ECDSASig
	}
	return nil
}

func (x *Signature) GetBLS12Sig() *BLS12Signature {
	if x, ok := x.GetSig().(*Signature_BLS12Sig); ok {
		return x.BLS12Sig
	}
	return nil
}

type isSignature_Sig interface {
	isSignature_Sig()
}

type Signature_ECDSASig struct {
	ECDSASig *ECDSASignature `protobuf:"bytes,1,opt,name=ECDSASig,proto3,oneof"`
}

type Signature_BLS12Sig struct {
	BLS12Sig *BLS12Signature `protobuf:"bytes,2,opt,name=BLS12Sig,proto3,oneof"`
}

func (*Signature_ECDSASig) isSignature_Sig() {}

func (*Signature_BLS12Sig) isSignature_Sig() {}

type ECDSAMultiSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sigs []*ECDSASignature `protobuf:"bytes,1,rep,name=Sigs,proto3" json:"Sigs,omitempty"`
}

func (x *ECDSAMultiSignature) Reset() {
	*x = ECDSAMultiSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECDSAMultiSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECDSAMultiSignature) ProtoMessage() {}

func (x *ECDSAMultiSignature) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECDSAMultiSignature.ProtoReflect.Descriptor instead.
func (*ECDSAMultiSignature) Descriptor() ([]byte, []int) {
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP(), []int{7}
}

func (x *ECDSAMultiSignature) GetSigs() []*ECDSASignature {
	if x != nil {
		return x.Sigs
	}
	return nil
}

type BLS12AggregateSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sig          []byte `protobuf:"bytes,1,opt,name=Sig,proto3" json:"Sig,omitempty"`
	Participants []byte `protobuf:"bytes,2,opt,name=participants,proto3" json:"participants,omitempty"`
}

func (x *BLS12AggregateSignature) Reset() {
	*x = BLS12AggregateSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLS12AggregateSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLS12AggregateSignature) ProtoMessage() {}

func (x *BLS12AggregateSignature) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLS12AggregateSignature.ProtoReflect.Descriptor instead.
func (*BLS12AggregateSignature) Descriptor() ([]byte, []int) {
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP(), []int{8}
}

func (x *BLS12AggregateSignature) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *BLS12AggregateSignature) GetParticipants() []byte {
	if x != nil {
		return x.Participants
	}
	return nil
}

type QuorumSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Sig:
	//
	//	*QuorumSignature_ECDSASigs
	//	*QuorumSignature_BLS12Sig
	Sig isQuorumSignature_Sig `protobuf_oneof:"Sig"`
}

func (x *QuorumSignature) Reset() {
	*x = QuorumSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuorumSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuorumSignature) ProtoMessage() {}

func (x *QuorumSignature) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_crossshardpb_crossshard_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuorumSignature.ProtoReflect.Descriptor instead.
func (*QuorumSignature) Descriptor() ([]byte, []int) {
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP(), []int{9}
}

func (m *QuorumSignature) GetSig() isQuorumSignature_Sig {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (x *QuorumSignature) GetECDSASigs() *ECDSAMultiSignature {
	if x, ok := x.GetSig().(*QuorumSignature_ECDSASigs); ok {
		return x.ECDSASigs
	}
	return nil
}

func (x *QuorumSignature) GetBLS12Sig() *BLS12AggregateSignature {
	if x, ok := x.GetSig().(*QuorumSignature_BLS12Sig); ok {
		return x.BLS12Sig
	}
	return nil
}

type isQuorumSignature_Sig interface {
	isQuorumSignature_Sig()
}

type QuorumSignature_ECDSASigs struct {
	ECDSASigs *ECDSAMultiSignature `protobuf:"bytes,1,opt,name=ECDSASigs,proto3,oneof"`
}

type QuorumSignature_BLS12Sig struct {
	BLS12Sig *BLS12AggregateSignature `protobuf:"bytes,2,opt,name=BLS12Sig,proto3,oneof"`
}

func (*QuorumSignature_ECDSASigs) isQuorumSignature_Sig() {}

func (*QuorumSignature_BLS12Sig) isQuorumSignature_Sig() {}

var File_internal_proto_crossshardpb_crossshard_proto protoreflect.FileDescriptor

var file_internal_proto_crossshardpb_crossshard_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x73, 0x68, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2f, 0x63, 0x72,
	0x6f, 0x73, 0x73, 0x73, 0x68, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x73, 0x68, 0x61, 0x72, 0x64, 0x70, 0x62, 0x1a, 0x0c, 0x67, 0x6f,
	0x72, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x4d, 0x73,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x68, 0x0a, 0x06, 0x4f, 0x63, 0x75, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x4f, 0x0a, 0x07, 0x4f,
	0x63, 0x75, 0x32, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x50, 0x0a, 0x08,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x44,
	0x0a, 0x0e, 0x45, 0x43, 0x44, 0x53, 0x41, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x52, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x01, 0x52, 0x12, 0x0c, 0x0a, 0x01, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x01, 0x53, 0x22, 0x22, 0x0a, 0x0e, 0x42, 0x4c, 0x53, 0x31, 0x32, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x53, 0x69, 0x67, 0x22, 0x8a, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x45, 0x43, 0x44, 0x53, 0x41, 0x53,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x45, 0x43, 0x44, 0x53, 0x41, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x08, 0x45, 0x43, 0x44, 0x53, 0x41, 0x53,
	0x69, 0x67, 0x12, 0x3a, 0x0a, 0x08, 0x42, 0x4c, 0x53, 0x31, 0x32, 0x53, 0x69, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x70, 0x62, 0x2e, 0x42, 0x4c, 0x53, 0x31, 0x32, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x48, 0x00, 0x52, 0x08, 0x42, 0x4c, 0x53, 0x31, 0x32, 0x53, 0x69, 0x67, 0x42, 0x05,
	0x0a, 0x03, 0x53, 0x69, 0x67, 0x22, 0x47, 0x0a, 0x13, 0x45, 0x43, 0x44, 0x53, 0x41, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x04,
	0x53, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x72, 0x6f,
	0x73, 0x73, 0x73, 0x68, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x45, 0x43, 0x44, 0x53, 0x41, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x04, 0x53, 0x69, 0x67, 0x73, 0x22, 0x4f,
	0x0a, 0x17, 0x42, 0x4c, 0x53, 0x31, 0x32, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x53, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x22,
	0xa0, 0x01, 0x0a, 0x0f, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x45, 0x43, 0x44, 0x53, 0x41, 0x53, 0x69, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x45, 0x43, 0x44, 0x53, 0x41, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x09, 0x45, 0x43, 0x44,
	0x53, 0x41, 0x53, 0x69, 0x67, 0x73, 0x12, 0x43, 0x0a, 0x08, 0x42, 0x4c, 0x53, 0x31, 0x32, 0x53,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x42, 0x4c, 0x53, 0x31, 0x32, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48,
	0x00, 0x52, 0x08, 0x42, 0x4c, 0x53, 0x31, 0x32, 0x53, 0x69, 0x67, 0x42, 0x05, 0x0a, 0x03, 0x53,
	0x69, 0x67, 0x32, 0x84, 0x02, 0x0a, 0x0a, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x53, 0x68, 0x61, 0x72,
	0x64, 0x12, 0x3d, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x2e, 0x63, 0x72, 0x6f,
	0x73, 0x73, 0x73, 0x68, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d,
	0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0x90, 0xb5, 0x18, 0x01,
	0x12, 0x3b, 0x0a, 0x05, 0x54, 0x58, 0x55, 0x73, 0x65, 0x12, 0x14, 0x2e, 0x63, 0x72, 0x6f, 0x73,
	0x73, 0x73, 0x68, 0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x4d, 0x73, 0x67, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0x90, 0xb5, 0x18, 0x01, 0x12, 0x3b, 0x0a,
	0x05, 0x54, 0x58, 0x4f, 0x63, 0x75, 0x12, 0x14, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x70, 0x62, 0x2e, 0x4f, 0x63, 0x75, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0x90, 0xb5, 0x18, 0x01, 0x12, 0x3d, 0x0a, 0x06, 0x54, 0x58,
	0x4f, 0x63, 0x75, 0x32, 0x12, 0x15, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x70, 0x62, 0x2e, 0x4f, 0x63, 0x75, 0x32, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x04, 0x90, 0xb5, 0x18, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x62, 0x2f, 0x68, 0x6f,
	0x74, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_crossshardpb_crossshard_proto_rawDescOnce sync.Once
	file_internal_proto_crossshardpb_crossshard_proto_rawDescData = file_internal_proto_crossshardpb_crossshard_proto_rawDesc
)

func file_internal_proto_crossshardpb_crossshard_proto_rawDescGZIP() []byte {
	file_internal_proto_crossshardpb_crossshard_proto_rawDescOnce.Do(func() {
		file_internal_proto_crossshardpb_crossshard_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_crossshardpb_crossshard_proto_rawDescData)
	})
	return file_internal_proto_crossshardpb_crossshard_proto_rawDescData
}

var file_internal_proto_crossshardpb_crossshard_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_proto_crossshardpb_crossshard_proto_goTypes = []interface{}{
	(*UseMsg)(nil),                  // 0: crossshardpb.UseMsg
	(*OcuMsg)(nil),                  // 1: crossshardpb.OcuMsg
	(*Ocu2Msg)(nil),                 // 2: crossshardpb.Ocu2Msg
	(*ReplyMsg)(nil),                // 3: crossshardpb.ReplyMsg
	(*ECDSASignature)(nil),          // 4: crossshardpb.ECDSASignature
	(*BLS12Signature)(nil),          // 5: crossshardpb.BLS12Signature
	(*Signature)(nil),               // 6: crossshardpb.Signature
	(*ECDSAMultiSignature)(nil),     // 7: crossshardpb.ECDSAMultiSignature
	(*BLS12AggregateSignature)(nil), // 8: crossshardpb.BLS12AggregateSignature
	(*QuorumSignature)(nil),         // 9: crossshardpb.QuorumSignature
	(*emptypb.Empty)(nil),           // 10: google.protobuf.Empty
}
var file_internal_proto_crossshardpb_crossshard_proto_depIdxs = []int32{
	4,  // 0: crossshardpb.Signature.ECDSASig:type_name -> crossshardpb.ECDSASignature
	5,  // 1: crossshardpb.Signature.BLS12Sig:type_name -> crossshardpb.BLS12Signature
	4,  // 2: crossshardpb.ECDSAMultiSignature.Sigs:type_name -> crossshardpb.ECDSASignature
	7,  // 3: crossshardpb.QuorumSignature.ECDSASigs:type_name -> crossshardpb.ECDSAMultiSignature
	8,  // 4: crossshardpb.QuorumSignature.BLS12Sig:type_name -> crossshardpb.BLS12AggregateSignature
	3,  // 5: crossshardpb.CrossShard.Reply:input_type -> crossshardpb.ReplyMsg
	0,  // 6: crossshardpb.CrossShard.TXUse:input_type -> crossshardpb.UseMsg
	1,  // 7: crossshardpb.CrossShard.TXOcu:input_type -> crossshardpb.OcuMsg
	2,  // 8: crossshardpb.CrossShard.TXOcu2:input_type -> crossshardpb.Ocu2Msg
	10, // 9: crossshardpb.CrossShard.Reply:output_type -> google.protobuf.Empty
	10, // 10: crossshardpb.CrossShard.TXUse:output_type -> google.protobuf.Empty
	10, // 11: crossshardpb.CrossShard.TXOcu:output_type -> google.protobuf.Empty
	10, // 12: crossshardpb.CrossShard.TXOcu2:output_type -> google.protobuf.Empty
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_internal_proto_crossshardpb_crossshard_proto_init() }
func file_internal_proto_crossshardpb_crossshard_proto_init() {
	if File_internal_proto_crossshardpb_crossshard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_crossshardpb_crossshard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseMsg); i {
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
		file_internal_proto_crossshardpb_crossshard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OcuMsg); i {
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
		file_internal_proto_crossshardpb_crossshard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ocu2Msg); i {
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
		file_internal_proto_crossshardpb_crossshard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyMsg); i {
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
		file_internal_proto_crossshardpb_crossshard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECDSASignature); i {
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
		file_internal_proto_crossshardpb_crossshard_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLS12Signature); i {
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
		file_internal_proto_crossshardpb_crossshard_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_internal_proto_crossshardpb_crossshard_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECDSAMultiSignature); i {
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
		file_internal_proto_crossshardpb_crossshard_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLS12AggregateSignature); i {
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
		file_internal_proto_crossshardpb_crossshard_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuorumSignature); i {
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
	file_internal_proto_crossshardpb_crossshard_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Signature_ECDSASig)(nil),
		(*Signature_BLS12Sig)(nil),
	}
	file_internal_proto_crossshardpb_crossshard_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*QuorumSignature_ECDSASigs)(nil),
		(*QuorumSignature_BLS12Sig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_crossshardpb_crossshard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_crossshardpb_crossshard_proto_goTypes,
		DependencyIndexes: file_internal_proto_crossshardpb_crossshard_proto_depIdxs,
		MessageInfos:      file_internal_proto_crossshardpb_crossshard_proto_msgTypes,
	}.Build()
	File_internal_proto_crossshardpb_crossshard_proto = out.File
	file_internal_proto_crossshardpb_crossshard_proto_rawDesc = nil
	file_internal_proto_crossshardpb_crossshard_proto_goTypes = nil
	file_internal_proto_crossshardpb_crossshard_proto_depIdxs = nil
}
