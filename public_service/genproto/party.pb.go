// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: protos/party.proto

package public

import (
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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_party_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_protos_party_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_protos_party_proto_rawDescGZIP(), []int{0}
}

type Party struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,4,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   int64  `protobuf:"varint,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Party) Reset() {
	*x = Party{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_party_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Party) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Party) ProtoMessage() {}

func (x *Party) ProtoReflect() protoreflect.Message {
	mi := &file_protos_party_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Party.ProtoReflect.Descriptor instead.
func (*Party) Descriptor() ([]byte, []int) {
	return file_protos_party_proto_rawDescGZIP(), []int{1}
}

func (x *Party) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Party) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Party) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *Party) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *Party) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Party) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Party) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Party) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type PartyCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,4,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PartyCreate) Reset() {
	*x = PartyCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_party_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyCreate) ProtoMessage() {}

func (x *PartyCreate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_party_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyCreate.ProtoReflect.Descriptor instead.
func (*PartyCreate) Descriptor() ([]byte, []int) {
	return file_protos_party_proto_rawDescGZIP(), []int{2}
}

func (x *PartyCreate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PartyCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PartyCreate) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *PartyCreate) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *PartyCreate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PartyUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,4,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PartyUpdate) Reset() {
	*x = PartyUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_party_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyUpdate) ProtoMessage() {}

func (x *PartyUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_party_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyUpdate.ProtoReflect.Descriptor instead.
func (*PartyUpdate) Descriptor() ([]byte, []int) {
	return file_protos_party_proto_rawDescGZIP(), []int{3}
}

func (x *PartyUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PartyUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PartyUpdate) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *PartyUpdate) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *PartyUpdate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PartyDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PartyDelete) Reset() {
	*x = PartyDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_party_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyDelete) ProtoMessage() {}

func (x *PartyDelete) ProtoReflect() protoreflect.Message {
	mi := &file_protos_party_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyDelete.ProtoReflect.Descriptor instead.
func (*PartyDelete) Descriptor() ([]byte, []int) {
	return file_protos_party_proto_rawDescGZIP(), []int{4}
}

func (x *PartyDelete) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PartyById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PartyById) Reset() {
	*x = PartyById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_party_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyById) ProtoMessage() {}

func (x *PartyById) ProtoReflect() protoreflect.Message {
	mi := &file_protos_party_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyById.ProtoReflect.Descriptor instead.
func (*PartyById) Descriptor() ([]byte, []int) {
	return file_protos_party_proto_rawDescGZIP(), []int{5}
}

func (x *PartyById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenedDate  string `protobuf:"bytes,1,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetAllPartyRequest) Reset() {
	*x = GetAllPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_party_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPartyRequest) ProtoMessage() {}

func (x *GetAllPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_party_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPartyRequest.ProtoReflect.Descriptor instead.
func (*GetAllPartyRequest) Descriptor() ([]byte, []int) {
	return file_protos_party_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllPartyRequest) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *GetAllPartyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllPartyRequest) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *GetAllPartyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetAllPartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parties []*Party `protobuf:"bytes,1,rep,name=parties,proto3" json:"parties,omitempty"`
}

func (x *GetAllPartyResponse) Reset() {
	*x = GetAllPartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_party_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPartyResponse) ProtoMessage() {}

func (x *GetAllPartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_party_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPartyResponse.ProtoReflect.Descriptor instead.
func (*GetAllPartyResponse) Descriptor() ([]byte, []int) {
	return file_protos_party_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllPartyResponse) GetParties() []*Party {
	if x != nil {
		return x.Parties
	}
	return nil
}

var File_protos_party_proto protoreflect.FileDescriptor

var file_protos_party_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x06, 0x0a, 0x04,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0xe3, 0x01, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70,
	0x65, 0x6e, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65,
	0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6f,
	0x70, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x32, 0x8f, 0x02, 0x0a, 0x0c, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_party_proto_rawDescOnce sync.Once
	file_protos_party_proto_rawDescData = file_protos_party_proto_rawDesc
)

func file_protos_party_proto_rawDescGZIP() []byte {
	file_protos_party_proto_rawDescOnce.Do(func() {
		file_protos_party_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_party_proto_rawDescData)
	})
	return file_protos_party_proto_rawDescData
}

var file_protos_party_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_party_proto_goTypes = []interface{}{
	(*Void)(nil),                // 0: protos.Void
	(*Party)(nil),               // 1: protos.Party
	(*PartyCreate)(nil),         // 2: protos.PartyCreate
	(*PartyUpdate)(nil),         // 3: protos.PartyUpdate
	(*PartyDelete)(nil),         // 4: protos.PartyDelete
	(*PartyById)(nil),           // 5: protos.PartyById
	(*GetAllPartyRequest)(nil),  // 6: protos.GetAllPartyRequest
	(*GetAllPartyResponse)(nil), // 7: protos.GetAllPartyResponse
}
var file_protos_party_proto_depIdxs = []int32{
	1, // 0: protos.GetAllPartyResponse.parties:type_name -> protos.Party
	2, // 1: protos.PartyService.Create:input_type -> protos.PartyCreate
	3, // 2: protos.PartyService.Update:input_type -> protos.PartyUpdate
	4, // 3: protos.PartyService.Delete:input_type -> protos.PartyDelete
	5, // 4: protos.PartyService.GetById:input_type -> protos.PartyById
	6, // 5: protos.PartyService.GetAll:input_type -> protos.GetAllPartyRequest
	0, // 6: protos.PartyService.Create:output_type -> protos.Void
	0, // 7: protos.PartyService.Update:output_type -> protos.Void
	0, // 8: protos.PartyService.Delete:output_type -> protos.Void
	1, // 9: protos.PartyService.GetById:output_type -> protos.Party
	7, // 10: protos.PartyService.GetAll:output_type -> protos.GetAllPartyResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_party_proto_init() }
func file_protos_party_proto_init() {
	if File_protos_party_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_party_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_protos_party_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Party); i {
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
		file_protos_party_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyCreate); i {
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
		file_protos_party_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyUpdate); i {
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
		file_protos_party_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyDelete); i {
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
		file_protos_party_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyById); i {
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
		file_protos_party_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPartyRequest); i {
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
		file_protos_party_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPartyResponse); i {
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
			RawDescriptor: file_protos_party_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_party_proto_goTypes,
		DependencyIndexes: file_protos_party_proto_depIdxs,
		MessageInfos:      file_protos_party_proto_msgTypes,
	}.Build()
	File_protos_party_proto = out.File
	file_protos_party_proto_rawDesc = nil
	file_protos_party_proto_goTypes = nil
	file_protos_party_proto_depIdxs = nil
}
