// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpkm67/backend/count/count.proto

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

type Count struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Count) Reset() {
	*x = Count{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_backend_count_count_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Count) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Count) ProtoMessage() {}

func (x *Count) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_backend_count_count_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Count.ProtoReflect.Descriptor instead.
func (*Count) Descriptor() ([]byte, []int) {
	return file_rpkm67_backend_count_count_proto_rawDescGZIP(), []int{0}
}

func (x *Count) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// FindAll
type FindAllCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllCountRequest) Reset() {
	*x = FindAllCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_backend_count_count_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllCountRequest) ProtoMessage() {}

func (x *FindAllCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_backend_count_count_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllCountRequest.ProtoReflect.Descriptor instead.
func (*FindAllCountRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_backend_count_count_proto_rawDescGZIP(), []int{1}
}

type FindAllCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counts []*Count `protobuf:"bytes,1,rep,name=counts,proto3" json:"counts,omitempty"`
}

func (x *FindAllCountResponse) Reset() {
	*x = FindAllCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_backend_count_count_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllCountResponse) ProtoMessage() {}

func (x *FindAllCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_backend_count_count_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllCountResponse.ProtoReflect.Descriptor instead.
func (*FindAllCountResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_backend_count_count_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllCountResponse) GetCounts() []*Count {
	if x != nil {
		return x.Counts
	}
	return nil
}

// Create
type CreateCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCountRequest) Reset() {
	*x = CreateCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_backend_count_count_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountRequest) ProtoMessage() {}

func (x *CreateCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_backend_count_count_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountRequest.ProtoReflect.Descriptor instead.
func (*CreateCountRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_backend_count_count_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count *Count `protobuf:"bytes,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CreateCountResponse) Reset() {
	*x = CreateCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_backend_count_count_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountResponse) ProtoMessage() {}

func (x *CreateCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_backend_count_count_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountResponse.ProtoReflect.Descriptor instead.
func (*CreateCountResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_backend_count_count_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCountResponse) GetCount() *Count {
	if x != nil {
		return x.Count
	}
	return nil
}

var File_rpkm67_backend_count_count_proto protoreflect.FileDescriptor

var file_rpkm67_backend_count_count_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x1b, 0x0a, 0x05, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4e, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22,
	0x28, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xdf, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x12, 0x2c, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x65, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x72, 0x70,
	0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36,
	0x37, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x72, 0x70, 0x6b, 0x6d,
	0x36, 0x37, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpkm67_backend_count_count_proto_rawDescOnce sync.Once
	file_rpkm67_backend_count_count_proto_rawDescData = file_rpkm67_backend_count_count_proto_rawDesc
)

func file_rpkm67_backend_count_count_proto_rawDescGZIP() []byte {
	file_rpkm67_backend_count_count_proto_rawDescOnce.Do(func() {
		file_rpkm67_backend_count_count_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpkm67_backend_count_count_proto_rawDescData)
	})
	return file_rpkm67_backend_count_count_proto_rawDescData
}

var file_rpkm67_backend_count_count_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpkm67_backend_count_count_proto_goTypes = []interface{}{
	(*Count)(nil),                // 0: rpkm67.backend.count.v1.Count
	(*FindAllCountRequest)(nil),  // 1: rpkm67.backend.count.v1.FindAllCountRequest
	(*FindAllCountResponse)(nil), // 2: rpkm67.backend.count.v1.FindAllCountResponse
	(*CreateCountRequest)(nil),   // 3: rpkm67.backend.count.v1.CreateCountRequest
	(*CreateCountResponse)(nil),  // 4: rpkm67.backend.count.v1.CreateCountResponse
}
var file_rpkm67_backend_count_count_proto_depIdxs = []int32{
	0, // 0: rpkm67.backend.count.v1.FindAllCountResponse.counts:type_name -> rpkm67.backend.count.v1.Count
	0, // 1: rpkm67.backend.count.v1.CreateCountResponse.count:type_name -> rpkm67.backend.count.v1.Count
	1, // 2: rpkm67.backend.count.v1.CountService.FindAll:input_type -> rpkm67.backend.count.v1.FindAllCountRequest
	3, // 3: rpkm67.backend.count.v1.CountService.Create:input_type -> rpkm67.backend.count.v1.CreateCountRequest
	2, // 4: rpkm67.backend.count.v1.CountService.FindAll:output_type -> rpkm67.backend.count.v1.FindAllCountResponse
	4, // 5: rpkm67.backend.count.v1.CountService.Create:output_type -> rpkm67.backend.count.v1.CreateCountResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpkm67_backend_count_count_proto_init() }
func file_rpkm67_backend_count_count_proto_init() {
	if File_rpkm67_backend_count_count_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpkm67_backend_count_count_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Count); i {
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
		file_rpkm67_backend_count_count_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllCountRequest); i {
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
		file_rpkm67_backend_count_count_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllCountResponse); i {
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
		file_rpkm67_backend_count_count_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountRequest); i {
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
		file_rpkm67_backend_count_count_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountResponse); i {
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
			RawDescriptor: file_rpkm67_backend_count_count_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpkm67_backend_count_count_proto_goTypes,
		DependencyIndexes: file_rpkm67_backend_count_count_proto_depIdxs,
		MessageInfos:      file_rpkm67_backend_count_count_proto_msgTypes,
	}.Build()
	File_rpkm67_backend_count_count_proto = out.File
	file_rpkm67_backend_count_count_proto_rawDesc = nil
	file_rpkm67_backend_count_count_proto_goTypes = nil
	file_rpkm67_backend_count_count_proto_depIdxs = nil
}
