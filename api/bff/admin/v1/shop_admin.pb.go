// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api/bff/admin/v1/shop_admin.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListGameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int64 `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListGameReq) Reset() {
	*x = ListGameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGameReq) ProtoMessage() {}

func (x *ListGameReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGameReq.ProtoReflect.Descriptor instead.
func (*ListGameReq) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{0}
}

func (x *ListGameReq) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListGameReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListGameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ListGameReply_Game `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListGameReply) Reset() {
	*x = ListGameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGameReply) ProtoMessage() {}

func (x *ListGameReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGameReply.ProtoReflect.Descriptor instead.
func (*ListGameReply) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{1}
}

func (x *ListGameReply) GetResults() []*ListGameReply_Game {
	if x != nil {
		return x.Results
	}
	return nil
}

type GetGameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGameReq) Reset() {
	*x = GetGameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameReq) ProtoMessage() {}

func (x *GetGameReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameReq.ProtoReflect.Descriptor instead.
func (*GetGameReq) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{2}
}

func (x *GetGameReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Count       int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetGameReply) Reset() {
	*x = GetGameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameReply) ProtoMessage() {}

func (x *GetGameReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameReply.ProtoReflect.Descriptor instead.
func (*GetGameReply) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{3}
}

func (x *GetGameReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetGameReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetGameReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetGameReply) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateGameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Count       int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CreateGameReq) Reset() {
	*x = CreateGameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameReq) ProtoMessage() {}

func (x *CreateGameReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameReq.ProtoReflect.Descriptor instead.
func (*CreateGameReq) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{4}
}

func (x *CreateGameReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGameReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateGameReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateGameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateGameReply) Reset() {
	*x = CreateGameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameReply) ProtoMessage() {}

func (x *CreateGameReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameReply.ProtoReflect.Descriptor instead.
func (*CreateGameReply) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{5}
}

type UpdateGameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Count       int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *UpdateGameReq) Reset() {
	*x = UpdateGameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGameReq) ProtoMessage() {}

func (x *UpdateGameReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGameReq.ProtoReflect.Descriptor instead.
func (*UpdateGameReq) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateGameReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateGameReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGameReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateGameReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UpdateGameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateGameReply) Reset() {
	*x = UpdateGameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGameReply) ProtoMessage() {}

func (x *UpdateGameReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGameReply.ProtoReflect.Descriptor instead.
func (*UpdateGameReply) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{7}
}

type DeleteGameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGameReq) Reset() {
	*x = DeleteGameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGameReq) ProtoMessage() {}

func (x *DeleteGameReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGameReq.ProtoReflect.Descriptor instead.
func (*DeleteGameReq) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteGameReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteGameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGameReply) Reset() {
	*x = DeleteGameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGameReply) ProtoMessage() {}

func (x *DeleteGameReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGameReply.ProtoReflect.Descriptor instead.
func (*DeleteGameReply) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{9}
}

type ListGameReply_Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Count       int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListGameReply_Game) Reset() {
	*x = ListGameReply_Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGameReply_Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGameReply_Game) ProtoMessage() {}

func (x *ListGameReply_Game) ProtoReflect() protoreflect.Message {
	mi := &file_api_bff_admin_v1_shop_admin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGameReply_Game.ProtoReflect.Descriptor instead.
func (*ListGameReply_Game) Descriptor() ([]byte, []int) {
	return file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListGameReply_Game) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListGameReply_Game) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListGameReply_Game) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ListGameReply_Game) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_api_bff_admin_v1_shop_admin_proto protoreflect.FileDescriptor

var file_api_bff_admin_v1_shop_admin_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x66, 0x66, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x45, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x62, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x6b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xb1, 0x04, 0x0a, 0x09, 0x53, 0x68, 0x6f,
	0x70, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x65, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x6e, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x22, 0x17, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x73, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x1a, 0x1c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a,
	0x01, 0x2a, 0x12, 0x70, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1e,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x1c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12,
	0x1b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x16, 0x5a, 0x14,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_bff_admin_v1_shop_admin_proto_rawDescOnce sync.Once
	file_api_bff_admin_v1_shop_admin_proto_rawDescData = file_api_bff_admin_v1_shop_admin_proto_rawDesc
)

func file_api_bff_admin_v1_shop_admin_proto_rawDescGZIP() []byte {
	file_api_bff_admin_v1_shop_admin_proto_rawDescOnce.Do(func() {
		file_api_bff_admin_v1_shop_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_bff_admin_v1_shop_admin_proto_rawDescData)
	})
	return file_api_bff_admin_v1_shop_admin_proto_rawDescData
}

var file_api_bff_admin_v1_shop_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_bff_admin_v1_shop_admin_proto_goTypes = []interface{}{
	(*ListGameReq)(nil),        // 0: shop.admin.v1.ListGameReq
	(*ListGameReply)(nil),      // 1: shop.admin.v1.ListGameReply
	(*GetGameReq)(nil),         // 2: shop.admin.v1.GetGameReq
	(*GetGameReply)(nil),       // 3: shop.admin.v1.GetGameReply
	(*CreateGameReq)(nil),      // 4: shop.admin.v1.CreateGameReq
	(*CreateGameReply)(nil),    // 5: shop.admin.v1.CreateGameReply
	(*UpdateGameReq)(nil),      // 6: shop.admin.v1.UpdateGameReq
	(*UpdateGameReply)(nil),    // 7: shop.admin.v1.UpdateGameReply
	(*DeleteGameReq)(nil),      // 8: shop.admin.v1.DeleteGameReq
	(*DeleteGameReply)(nil),    // 9: shop.admin.v1.DeleteGameReply
	(*ListGameReply_Game)(nil), // 10: shop.admin.v1.ListGameReply.Game
}
var file_api_bff_admin_v1_shop_admin_proto_depIdxs = []int32{
	10, // 0: shop.admin.v1.ListGameReply.results:type_name -> shop.admin.v1.ListGameReply.Game
	0,  // 1: shop.admin.v1.ShopAdmin.ListGame:input_type -> shop.admin.v1.ListGameReq
	4,  // 2: shop.admin.v1.ShopAdmin.CreateGame:input_type -> shop.admin.v1.CreateGameReq
	6,  // 3: shop.admin.v1.ShopAdmin.UpdateGame:input_type -> shop.admin.v1.UpdateGameReq
	8,  // 4: shop.admin.v1.ShopAdmin.DeleteGame:input_type -> shop.admin.v1.DeleteGameReq
	2,  // 5: shop.admin.v1.ShopAdmin.GetGame:input_type -> shop.admin.v1.GetGameReq
	1,  // 6: shop.admin.v1.ShopAdmin.ListGame:output_type -> shop.admin.v1.ListGameReply
	5,  // 7: shop.admin.v1.ShopAdmin.CreateGame:output_type -> shop.admin.v1.CreateGameReply
	7,  // 8: shop.admin.v1.ShopAdmin.UpdateGame:output_type -> shop.admin.v1.UpdateGameReply
	9,  // 9: shop.admin.v1.ShopAdmin.DeleteGame:output_type -> shop.admin.v1.DeleteGameReply
	3,  // 10: shop.admin.v1.ShopAdmin.GetGame:output_type -> shop.admin.v1.GetGameReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_bff_admin_v1_shop_admin_proto_init() }
func file_api_bff_admin_v1_shop_admin_proto_init() {
	if File_api_bff_admin_v1_shop_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGameReq); i {
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
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGameReply); i {
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
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGameReq); i {
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
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGameReply); i {
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
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameReq); i {
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
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameReply); i {
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
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGameReq); i {
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
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGameReply); i {
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
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGameReq); i {
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
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGameReply); i {
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
		file_api_bff_admin_v1_shop_admin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGameReply_Game); i {
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
			RawDescriptor: file_api_bff_admin_v1_shop_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_bff_admin_v1_shop_admin_proto_goTypes,
		DependencyIndexes: file_api_bff_admin_v1_shop_admin_proto_depIdxs,
		MessageInfos:      file_api_bff_admin_v1_shop_admin_proto_msgTypes,
	}.Build()
	File_api_bff_admin_v1_shop_admin_proto = out.File
	file_api_bff_admin_v1_shop_admin_proto_rawDesc = nil
	file_api_bff_admin_v1_shop_admin_proto_goTypes = nil
	file_api_bff_admin_v1_shop_admin_proto_depIdxs = nil
}
