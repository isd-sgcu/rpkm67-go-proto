// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpkm67/checkin/checkin/v1/checkin.proto

package v1

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

type CheckIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Event       string `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	Timestamp   string `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IsDuplicate bool   `protobuf:"varint,6,opt,name=isDuplicate,proto3" json:"isDuplicate,omitempty"`
}

func (x *CheckIn) Reset() {
	*x = CheckIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIn) ProtoMessage() {}

func (x *CheckIn) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIn.ProtoReflect.Descriptor instead.
func (*CheckIn) Descriptor() ([]byte, []int) {
	return file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescGZIP(), []int{0}
}

func (x *CheckIn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CheckIn) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CheckIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CheckIn) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *CheckIn) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *CheckIn) GetIsDuplicate() bool {
	if x != nil {
		return x.IsDuplicate
	}
	return false
}

// Create
type CreateCheckInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Event  string `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *CreateCheckInRequest) Reset() {
	*x = CreateCheckInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCheckInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCheckInRequest) ProtoMessage() {}

func (x *CreateCheckInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCheckInRequest.ProtoReflect.Descriptor instead.
func (*CreateCheckInRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCheckInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateCheckInRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateCheckInRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

type CreateCheckInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckIn *CheckIn `protobuf:"bytes,1,opt,name=checkIn,proto3" json:"checkIn,omitempty"`
}

func (x *CreateCheckInResponse) Reset() {
	*x = CreateCheckInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCheckInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCheckInResponse) ProtoMessage() {}

func (x *CreateCheckInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCheckInResponse.ProtoReflect.Descriptor instead.
func (*CreateCheckInResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCheckInResponse) GetCheckIn() *CheckIn {
	if x != nil {
		return x.CheckIn
	}
	return nil
}

// FindByUserId
type FindByUserIdCheckInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *FindByUserIdCheckInRequest) Reset() {
	*x = FindByUserIdCheckInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByUserIdCheckInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByUserIdCheckInRequest) ProtoMessage() {}

func (x *FindByUserIdCheckInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByUserIdCheckInRequest.ProtoReflect.Descriptor instead.
func (*FindByUserIdCheckInRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescGZIP(), []int{3}
}

func (x *FindByUserIdCheckInRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FindByUserIdCheckInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckIns []*CheckIn `protobuf:"bytes,1,rep,name=checkIns,proto3" json:"checkIns,omitempty"`
}

func (x *FindByUserIdCheckInResponse) Reset() {
	*x = FindByUserIdCheckInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByUserIdCheckInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByUserIdCheckInResponse) ProtoMessage() {}

func (x *FindByUserIdCheckInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByUserIdCheckInResponse.ProtoReflect.Descriptor instead.
func (*FindByUserIdCheckInResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescGZIP(), []int{4}
}

func (x *FindByUserIdCheckInResponse) GetCheckIns() []*CheckIn {
	if x != nil {
		return x.CheckIns
	}
	return nil
}

// FindByEmail
type FindByEmailCheckInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FindByEmailCheckInRequest) Reset() {
	*x = FindByEmailCheckInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByEmailCheckInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByEmailCheckInRequest) ProtoMessage() {}

func (x *FindByEmailCheckInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByEmailCheckInRequest.ProtoReflect.Descriptor instead.
func (*FindByEmailCheckInRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescGZIP(), []int{5}
}

func (x *FindByEmailCheckInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type FindByEmailCheckInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckIns []*CheckIn `protobuf:"bytes,1,rep,name=checkIns,proto3" json:"checkIns,omitempty"`
}

func (x *FindByEmailCheckInResponse) Reset() {
	*x = FindByEmailCheckInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByEmailCheckInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByEmailCheckInResponse) ProtoMessage() {}

func (x *FindByEmailCheckInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByEmailCheckInResponse.ProtoReflect.Descriptor instead.
func (*FindByEmailCheckInResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescGZIP(), []int{6}
}

func (x *FindByEmailCheckInResponse) GetCheckIns() []*CheckIn {
	if x != nil {
		return x.CheckIns
	}
	return nil
}

var File_rpkm67_checkin_checkin_v1_checkin_proto protoreflect.FileDescriptor

var file_rpkm67_checkin_checkin_v1_checkin_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x70, 0x6b, 0x6d, 0x36,
	0x37, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x22, 0x9d, 0x01, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x44, 0x75, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x22, 0x5a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x55, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x70, 0x6b,
	0x6d, 0x36, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x07,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x22, 0x34, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a,
	0x1b, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x08,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x73, 0x22, 0x31, 0x0a, 0x19,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x5c, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x6e, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x73, 0x32, 0xfe, 0x02,
	0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x72, 0x70, 0x6b,
	0x6d, 0x36, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x72, 0x70,
	0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x7f, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x35, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x7c, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x34, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b,
	0x5a, 0x19, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescOnce sync.Once
	file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescData = file_rpkm67_checkin_checkin_v1_checkin_proto_rawDesc
)

func file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescGZIP() []byte {
	file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescOnce.Do(func() {
		file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescData)
	})
	return file_rpkm67_checkin_checkin_v1_checkin_proto_rawDescData
}

var file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpkm67_checkin_checkin_v1_checkin_proto_goTypes = []interface{}{
	(*CheckIn)(nil),                     // 0: rpkm67.checkin.checkin.v1.CheckIn
	(*CreateCheckInRequest)(nil),        // 1: rpkm67.checkin.checkin.v1.CreateCheckInRequest
	(*CreateCheckInResponse)(nil),       // 2: rpkm67.checkin.checkin.v1.CreateCheckInResponse
	(*FindByUserIdCheckInRequest)(nil),  // 3: rpkm67.checkin.checkin.v1.FindByUserIdCheckInRequest
	(*FindByUserIdCheckInResponse)(nil), // 4: rpkm67.checkin.checkin.v1.FindByUserIdCheckInResponse
	(*FindByEmailCheckInRequest)(nil),   // 5: rpkm67.checkin.checkin.v1.FindByEmailCheckInRequest
	(*FindByEmailCheckInResponse)(nil),  // 6: rpkm67.checkin.checkin.v1.FindByEmailCheckInResponse
}
var file_rpkm67_checkin_checkin_v1_checkin_proto_depIdxs = []int32{
	0, // 0: rpkm67.checkin.checkin.v1.CreateCheckInResponse.checkIn:type_name -> rpkm67.checkin.checkin.v1.CheckIn
	0, // 1: rpkm67.checkin.checkin.v1.FindByUserIdCheckInResponse.checkIns:type_name -> rpkm67.checkin.checkin.v1.CheckIn
	0, // 2: rpkm67.checkin.checkin.v1.FindByEmailCheckInResponse.checkIns:type_name -> rpkm67.checkin.checkin.v1.CheckIn
	1, // 3: rpkm67.checkin.checkin.v1.CheckInService.Create:input_type -> rpkm67.checkin.checkin.v1.CreateCheckInRequest
	3, // 4: rpkm67.checkin.checkin.v1.CheckInService.FindByUserId:input_type -> rpkm67.checkin.checkin.v1.FindByUserIdCheckInRequest
	5, // 5: rpkm67.checkin.checkin.v1.CheckInService.FindByEmail:input_type -> rpkm67.checkin.checkin.v1.FindByEmailCheckInRequest
	2, // 6: rpkm67.checkin.checkin.v1.CheckInService.Create:output_type -> rpkm67.checkin.checkin.v1.CreateCheckInResponse
	4, // 7: rpkm67.checkin.checkin.v1.CheckInService.FindByUserId:output_type -> rpkm67.checkin.checkin.v1.FindByUserIdCheckInResponse
	6, // 8: rpkm67.checkin.checkin.v1.CheckInService.FindByEmail:output_type -> rpkm67.checkin.checkin.v1.FindByEmailCheckInResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpkm67_checkin_checkin_v1_checkin_proto_init() }
func file_rpkm67_checkin_checkin_v1_checkin_proto_init() {
	if File_rpkm67_checkin_checkin_v1_checkin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIn); i {
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
		file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCheckInRequest); i {
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
		file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCheckInResponse); i {
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
		file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByUserIdCheckInRequest); i {
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
		file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByUserIdCheckInResponse); i {
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
		file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByEmailCheckInRequest); i {
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
		file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByEmailCheckInResponse); i {
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
			RawDescriptor: file_rpkm67_checkin_checkin_v1_checkin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpkm67_checkin_checkin_v1_checkin_proto_goTypes,
		DependencyIndexes: file_rpkm67_checkin_checkin_v1_checkin_proto_depIdxs,
		MessageInfos:      file_rpkm67_checkin_checkin_v1_checkin_proto_msgTypes,
	}.Build()
	File_rpkm67_checkin_checkin_v1_checkin_proto = out.File
	file_rpkm67_checkin_checkin_v1_checkin_proto_rawDesc = nil
	file_rpkm67_checkin_checkin_v1_checkin_proto_goTypes = nil
	file_rpkm67_checkin_checkin_v1_checkin_proto_depIdxs = nil
}
