// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpkm67/backend/stamp/v1/stamp.proto

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

type Stamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	PointA int32  `protobuf:"varint,3,opt,name=pointA,proto3" json:"pointA,omitempty"`
	PointB int32  `protobuf:"varint,4,opt,name=pointB,proto3" json:"pointB,omitempty"`
	PointC int32  `protobuf:"varint,5,opt,name=pointC,proto3" json:"pointC,omitempty"`
	PointD int32  `protobuf:"varint,6,opt,name=pointD,proto3" json:"pointD,omitempty"`
	Stamp  string `protobuf:"bytes,7,opt,name=stamp,proto3" json:"stamp,omitempty"`
}

func (x *Stamp) Reset() {
	*x = Stamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stamp) ProtoMessage() {}

func (x *Stamp) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stamp.ProtoReflect.Descriptor instead.
func (*Stamp) Descriptor() ([]byte, []int) {
	return file_rpkm67_backend_stamp_v1_stamp_proto_rawDescGZIP(), []int{0}
}

func (x *Stamp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Stamp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Stamp) GetPointA() int32 {
	if x != nil {
		return x.PointA
	}
	return 0
}

func (x *Stamp) GetPointB() int32 {
	if x != nil {
		return x.PointB
	}
	return 0
}

func (x *Stamp) GetPointC() int32 {
	if x != nil {
		return x.PointC
	}
	return 0
}

func (x *Stamp) GetPointD() int32 {
	if x != nil {
		return x.PointD
	}
	return 0
}

func (x *Stamp) GetStamp() string {
	if x != nil {
		return x.Stamp
	}
	return ""
}

// FindByUserId
type FindByUserIdStampRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *FindByUserIdStampRequest) Reset() {
	*x = FindByUserIdStampRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByUserIdStampRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByUserIdStampRequest) ProtoMessage() {}

func (x *FindByUserIdStampRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByUserIdStampRequest.ProtoReflect.Descriptor instead.
func (*FindByUserIdStampRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_backend_stamp_v1_stamp_proto_rawDescGZIP(), []int{1}
}

func (x *FindByUserIdStampRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FindByUserIdStampResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stamp *Stamp `protobuf:"bytes,1,opt,name=stamp,proto3" json:"stamp,omitempty"`
}

func (x *FindByUserIdStampResponse) Reset() {
	*x = FindByUserIdStampResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByUserIdStampResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByUserIdStampResponse) ProtoMessage() {}

func (x *FindByUserIdStampResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByUserIdStampResponse.ProtoReflect.Descriptor instead.
func (*FindByUserIdStampResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_backend_stamp_v1_stamp_proto_rawDescGZIP(), []int{2}
}

func (x *FindByUserIdStampResponse) GetStamp() *Stamp {
	if x != nil {
		return x.Stamp
	}
	return nil
}

// StampByUserId
type StampByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ActivityId string `protobuf:"bytes,2,opt,name=activityId,proto3" json:"activityId,omitempty"`
}

func (x *StampByUserIdRequest) Reset() {
	*x = StampByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StampByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StampByUserIdRequest) ProtoMessage() {}

func (x *StampByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StampByUserIdRequest.ProtoReflect.Descriptor instead.
func (*StampByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_backend_stamp_v1_stamp_proto_rawDescGZIP(), []int{3}
}

func (x *StampByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *StampByUserIdRequest) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

type StampByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stamp *Stamp `protobuf:"bytes,1,opt,name=stamp,proto3" json:"stamp,omitempty"`
}

func (x *StampByUserIdResponse) Reset() {
	*x = StampByUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StampByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StampByUserIdResponse) ProtoMessage() {}

func (x *StampByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StampByUserIdResponse.ProtoReflect.Descriptor instead.
func (*StampByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_backend_stamp_v1_stamp_proto_rawDescGZIP(), []int{4}
}

func (x *StampByUserIdResponse) GetStamp() *Stamp {
	if x != nil {
		return x.Stamp
	}
	return nil
}

var File_rpkm67_backend_stamp_v1_stamp_proto protoreflect.FileDescriptor

var file_rpkm67_backend_stamp_v1_stamp_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x22, 0xa5,
	0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x32, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x19, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4e, 0x0a,
	0x14, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x4d, 0x0a,
	0x15, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0xf9, 0x01, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a,
	0x0c, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x2e,
	0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x72, 0x70, 0x6b, 0x6d,
	0x36, 0x37, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpkm67_backend_stamp_v1_stamp_proto_rawDescOnce sync.Once
	file_rpkm67_backend_stamp_v1_stamp_proto_rawDescData = file_rpkm67_backend_stamp_v1_stamp_proto_rawDesc
)

func file_rpkm67_backend_stamp_v1_stamp_proto_rawDescGZIP() []byte {
	file_rpkm67_backend_stamp_v1_stamp_proto_rawDescOnce.Do(func() {
		file_rpkm67_backend_stamp_v1_stamp_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpkm67_backend_stamp_v1_stamp_proto_rawDescData)
	})
	return file_rpkm67_backend_stamp_v1_stamp_proto_rawDescData
}

var file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpkm67_backend_stamp_v1_stamp_proto_goTypes = []interface{}{
	(*Stamp)(nil),                     // 0: rpkm67.backend.stamp.v1.Stamp
	(*FindByUserIdStampRequest)(nil),  // 1: rpkm67.backend.stamp.v1.FindByUserIdStampRequest
	(*FindByUserIdStampResponse)(nil), // 2: rpkm67.backend.stamp.v1.FindByUserIdStampResponse
	(*StampByUserIdRequest)(nil),      // 3: rpkm67.backend.stamp.v1.StampByUserIdRequest
	(*StampByUserIdResponse)(nil),     // 4: rpkm67.backend.stamp.v1.StampByUserIdResponse
}
var file_rpkm67_backend_stamp_v1_stamp_proto_depIdxs = []int32{
	0, // 0: rpkm67.backend.stamp.v1.FindByUserIdStampResponse.stamp:type_name -> rpkm67.backend.stamp.v1.Stamp
	0, // 1: rpkm67.backend.stamp.v1.StampByUserIdResponse.stamp:type_name -> rpkm67.backend.stamp.v1.Stamp
	1, // 2: rpkm67.backend.stamp.v1.StampService.FindByUserId:input_type -> rpkm67.backend.stamp.v1.FindByUserIdStampRequest
	3, // 3: rpkm67.backend.stamp.v1.StampService.StampByUserId:input_type -> rpkm67.backend.stamp.v1.StampByUserIdRequest
	2, // 4: rpkm67.backend.stamp.v1.StampService.FindByUserId:output_type -> rpkm67.backend.stamp.v1.FindByUserIdStampResponse
	4, // 5: rpkm67.backend.stamp.v1.StampService.StampByUserId:output_type -> rpkm67.backend.stamp.v1.StampByUserIdResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpkm67_backend_stamp_v1_stamp_proto_init() }
func file_rpkm67_backend_stamp_v1_stamp_proto_init() {
	if File_rpkm67_backend_stamp_v1_stamp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stamp); i {
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
		file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByUserIdStampRequest); i {
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
		file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByUserIdStampResponse); i {
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
		file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StampByUserIdRequest); i {
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
		file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StampByUserIdResponse); i {
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
			RawDescriptor: file_rpkm67_backend_stamp_v1_stamp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpkm67_backend_stamp_v1_stamp_proto_goTypes,
		DependencyIndexes: file_rpkm67_backend_stamp_v1_stamp_proto_depIdxs,
		MessageInfos:      file_rpkm67_backend_stamp_v1_stamp_proto_msgTypes,
	}.Build()
	File_rpkm67_backend_stamp_v1_stamp_proto = out.File
	file_rpkm67_backend_stamp_v1_stamp_proto_rawDesc = nil
	file_rpkm67_backend_stamp_v1_stamp_proto_goTypes = nil
	file_rpkm67_backend_stamp_v1_stamp_proto_depIdxs = nil
}
