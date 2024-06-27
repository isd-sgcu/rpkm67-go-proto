// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpkm67/auth/user/v1/user.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Nickname    string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Title       string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Firstname   string `protobuf:"bytes,5,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname    string `protobuf:"bytes,6,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Year        int32  `protobuf:"varint,7,opt,name=year,proto3" json:"year,omitempty"`
	Faculty     string `protobuf:"bytes,8,opt,name=faculty,proto3" json:"faculty,omitempty"`
	Tel         string `protobuf:"bytes,9,opt,name=tel,proto3" json:"tel,omitempty"`
	ParentTel   string `protobuf:"bytes,10,opt,name=parentTel,proto3" json:"parentTel,omitempty"`
	Parent      string `protobuf:"bytes,11,opt,name=parent,proto3" json:"parent,omitempty"`
	FoodAllergy string `protobuf:"bytes,12,opt,name=foodAllergy,proto3" json:"foodAllergy,omitempty"`
	DrugAllergy string `protobuf:"bytes,13,opt,name=drugAllergy,proto3" json:"drugAllergy,omitempty"`
	Illness     string `protobuf:"bytes,14,opt,name=illness,proto3" json:"illness,omitempty"`
	Role        string `protobuf:"bytes,15,opt,name=role,proto3" json:"role,omitempty"`
	PhotoKey    string `protobuf:"bytes,16,opt,name=photoKey,proto3" json:"photoKey,omitempty"`
	PhotoUrl    string `protobuf:"bytes,17,opt,name=photoUrl,proto3" json:"photoUrl,omitempty"`
	Baan        string `protobuf:"bytes,18,opt,name=baan,proto3" json:"baan,omitempty"`
	GroupId     string `protobuf:"bytes,19,opt,name=groupId,proto3" json:"groupId,omitempty"`
	ReceiveGift int32  `protobuf:"varint,20,opt,name=receiveGift,proto3" json:"receiveGift,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_rpkm67_auth_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *User) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *User) GetFaculty() string {
	if x != nil {
		return x.Faculty
	}
	return ""
}

func (x *User) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *User) GetParentTel() string {
	if x != nil {
		return x.ParentTel
	}
	return ""
}

func (x *User) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *User) GetFoodAllergy() string {
	if x != nil {
		return x.FoodAllergy
	}
	return ""
}

func (x *User) GetDrugAllergy() string {
	if x != nil {
		return x.DrugAllergy
	}
	return ""
}

func (x *User) GetIllness() string {
	if x != nil {
		return x.Illness
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetPhotoKey() string {
	if x != nil {
		return x.PhotoKey
	}
	return ""
}

func (x *User) GetPhotoUrl() string {
	if x != nil {
		return x.PhotoUrl
	}
	return ""
}

func (x *User) GetBaan() string {
	if x != nil {
		return x.Baan
	}
	return ""
}

func (x *User) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *User) GetReceiveGift() int32 {
	if x != nil {
		return x.ReceiveGift
	}
	return 0
}

// Create
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Role  string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_auth_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_auth_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// FindOne
type FindOneUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneUserRequest) Reset() {
	*x = FindOneUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneUserRequest) ProtoMessage() {}

func (x *FindOneUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneUserRequest.ProtoReflect.Descriptor instead.
func (*FindOneUserRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_auth_user_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *FindOneUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindOneUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FindOneUserResponse) Reset() {
	*x = FindOneUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneUserResponse) ProtoMessage() {}

func (x *FindOneUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneUserResponse.ProtoReflect.Descriptor instead.
func (*FindOneUserResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_auth_user_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// FindByEmail
type FindByEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FindByEmailRequest) Reset() {
	*x = FindByEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByEmailRequest) ProtoMessage() {}

func (x *FindByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByEmailRequest.ProtoReflect.Descriptor instead.
func (*FindByEmailRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_auth_user_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *FindByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type FindByEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FindByEmailResponse) Reset() {
	*x = FindByEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByEmailResponse) ProtoMessage() {}

func (x *FindByEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByEmailResponse.ProtoReflect.Descriptor instead.
func (*FindByEmailResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_auth_user_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *FindByEmailResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// Update
type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Nickname    string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Title       string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Firstname   string `protobuf:"bytes,5,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname    string `protobuf:"bytes,6,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Year        int32  `protobuf:"varint,7,opt,name=year,proto3" json:"year,omitempty"`
	Faculty     string `protobuf:"bytes,8,opt,name=faculty,proto3" json:"faculty,omitempty"`
	Tel         string `protobuf:"bytes,9,opt,name=tel,proto3" json:"tel,omitempty"`
	ParentTel   string `protobuf:"bytes,10,opt,name=parentTel,proto3" json:"parentTel,omitempty"`
	Parent      string `protobuf:"bytes,11,opt,name=parent,proto3" json:"parent,omitempty"`
	FoodAllergy string `protobuf:"bytes,12,opt,name=foodAllergy,proto3" json:"foodAllergy,omitempty"`
	DrugAllergy string `protobuf:"bytes,13,opt,name=drugAllergy,proto3" json:"drugAllergy,omitempty"`
	Illness     string `protobuf:"bytes,14,opt,name=illness,proto3" json:"illness,omitempty"`
	Role        string `protobuf:"bytes,15,opt,name=role,proto3" json:"role,omitempty"`
	PhotoKey    string `protobuf:"bytes,16,opt,name=photoKey,proto3" json:"photoKey,omitempty"`
	PhotoUrl    string `protobuf:"bytes,17,opt,name=photoUrl,proto3" json:"photoUrl,omitempty"`
	Baan        string `protobuf:"bytes,18,opt,name=baan,proto3" json:"baan,omitempty"`
	GroupId     string `protobuf:"bytes,19,opt,name=groupId,proto3" json:"groupId,omitempty"`
	ReceiveGift int32  `protobuf:"varint,20,opt,name=receiveGift,proto3" json:"receiveGift,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_rpkm67_auth_user_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UpdateUserRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateUserRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UpdateUserRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UpdateUserRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *UpdateUserRequest) GetFaculty() string {
	if x != nil {
		return x.Faculty
	}
	return ""
}

func (x *UpdateUserRequest) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *UpdateUserRequest) GetParentTel() string {
	if x != nil {
		return x.ParentTel
	}
	return ""
}

func (x *UpdateUserRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *UpdateUserRequest) GetFoodAllergy() string {
	if x != nil {
		return x.FoodAllergy
	}
	return ""
}

func (x *UpdateUserRequest) GetDrugAllergy() string {
	if x != nil {
		return x.DrugAllergy
	}
	return ""
}

func (x *UpdateUserRequest) GetIllness() string {
	if x != nil {
		return x.Illness
	}
	return ""
}

func (x *UpdateUserRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UpdateUserRequest) GetPhotoKey() string {
	if x != nil {
		return x.PhotoKey
	}
	return ""
}

func (x *UpdateUserRequest) GetPhotoUrl() string {
	if x != nil {
		return x.PhotoUrl
	}
	return ""
}

func (x *UpdateUserRequest) GetBaan() string {
	if x != nil {
		return x.Baan
	}
	return ""
}

func (x *UpdateUserRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *UpdateUserRequest) GetReceiveGift() int32 {
	if x != nil {
		return x.ReceiveGift
	}
	return 0
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm67_auth_user_v1_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_rpkm67_auth_user_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_rpkm67_auth_user_v1_user_proto protoreflect.FileDescriptor

var file_rpkm67_auth_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x88, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6c, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x6f, 0x6f, 0x64, 0x41, 0x6c,
	0x6c, 0x65, 0x72, 0x67, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6f, 0x6f,
	0x64, 0x41, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x72, 0x75, 0x67,
	0x41, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x72, 0x75, 0x67, 0x41, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6c,
	0x6c, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6c, 0x6c,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x4b, 0x65, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x61, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x61, 0x61, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x47, 0x69, 0x66, 0x74, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x47, 0x69, 0x66, 0x74,
	0x22, 0x3d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22,
	0x43, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x13, 0x46, 0x69,
	0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x2a, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x44, 0x0a, 0x13,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x95, 0x04, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x6f, 0x6f, 0x64, 0x41, 0x6c, 0x6c, 0x65, 0x72, 0x67, 0x79, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x6f, 0x6f, 0x64, 0x41, 0x6c, 0x6c, 0x65, 0x72,
	0x67, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x72, 0x75, 0x67, 0x41, 0x6c, 0x6c, 0x65, 0x72, 0x67,
	0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x72, 0x75, 0x67, 0x41, 0x6c, 0x6c,
	0x65, 0x72, 0x67, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6c, 0x6c, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61,
	0x61, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x61, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x47, 0x69, 0x66, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x47, 0x69, 0x66, 0x74, 0x22, 0x43, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32,
	0x8b, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x72, 0x70, 0x6b, 0x6d,
	0x36, 0x37, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x07,
	0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x27, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0b,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x27, 0x2e, 0x72, 0x70,
	0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x72, 0x70, 0x6b,
	0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a,
	0x13, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x37, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpkm67_auth_user_v1_user_proto_rawDescOnce sync.Once
	file_rpkm67_auth_user_v1_user_proto_rawDescData = file_rpkm67_auth_user_v1_user_proto_rawDesc
)

func file_rpkm67_auth_user_v1_user_proto_rawDescGZIP() []byte {
	file_rpkm67_auth_user_v1_user_proto_rawDescOnce.Do(func() {
		file_rpkm67_auth_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpkm67_auth_user_v1_user_proto_rawDescData)
	})
	return file_rpkm67_auth_user_v1_user_proto_rawDescData
}

var file_rpkm67_auth_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rpkm67_auth_user_v1_user_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: rpkm67.auth.user.v1.User
	(*CreateUserRequest)(nil),   // 1: rpkm67.auth.user.v1.CreateUserRequest
	(*CreateUserResponse)(nil),  // 2: rpkm67.auth.user.v1.CreateUserResponse
	(*FindOneUserRequest)(nil),  // 3: rpkm67.auth.user.v1.FindOneUserRequest
	(*FindOneUserResponse)(nil), // 4: rpkm67.auth.user.v1.FindOneUserResponse
	(*FindByEmailRequest)(nil),  // 5: rpkm67.auth.user.v1.FindByEmailRequest
	(*FindByEmailResponse)(nil), // 6: rpkm67.auth.user.v1.FindByEmailResponse
	(*UpdateUserRequest)(nil),   // 7: rpkm67.auth.user.v1.UpdateUserRequest
	(*UpdateUserResponse)(nil),  // 8: rpkm67.auth.user.v1.UpdateUserResponse
}
var file_rpkm67_auth_user_v1_user_proto_depIdxs = []int32{
	0, // 0: rpkm67.auth.user.v1.CreateUserResponse.user:type_name -> rpkm67.auth.user.v1.User
	0, // 1: rpkm67.auth.user.v1.FindOneUserResponse.user:type_name -> rpkm67.auth.user.v1.User
	0, // 2: rpkm67.auth.user.v1.FindByEmailResponse.user:type_name -> rpkm67.auth.user.v1.User
	0, // 3: rpkm67.auth.user.v1.UpdateUserResponse.user:type_name -> rpkm67.auth.user.v1.User
	1, // 4: rpkm67.auth.user.v1.UserService.Create:input_type -> rpkm67.auth.user.v1.CreateUserRequest
	3, // 5: rpkm67.auth.user.v1.UserService.FindOne:input_type -> rpkm67.auth.user.v1.FindOneUserRequest
	5, // 6: rpkm67.auth.user.v1.UserService.FindByEmail:input_type -> rpkm67.auth.user.v1.FindByEmailRequest
	7, // 7: rpkm67.auth.user.v1.UserService.Update:input_type -> rpkm67.auth.user.v1.UpdateUserRequest
	2, // 8: rpkm67.auth.user.v1.UserService.Create:output_type -> rpkm67.auth.user.v1.CreateUserResponse
	4, // 9: rpkm67.auth.user.v1.UserService.FindOne:output_type -> rpkm67.auth.user.v1.FindOneUserResponse
	6, // 10: rpkm67.auth.user.v1.UserService.FindByEmail:output_type -> rpkm67.auth.user.v1.FindByEmailResponse
	8, // 11: rpkm67.auth.user.v1.UserService.Update:output_type -> rpkm67.auth.user.v1.UpdateUserResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpkm67_auth_user_v1_user_proto_init() }
func file_rpkm67_auth_user_v1_user_proto_init() {
	if File_rpkm67_auth_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpkm67_auth_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_rpkm67_auth_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_rpkm67_auth_user_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_rpkm67_auth_user_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneUserRequest); i {
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
		file_rpkm67_auth_user_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneUserResponse); i {
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
		file_rpkm67_auth_user_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByEmailRequest); i {
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
		file_rpkm67_auth_user_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByEmailResponse); i {
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
		file_rpkm67_auth_user_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_rpkm67_auth_user_v1_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserResponse); i {
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
			RawDescriptor: file_rpkm67_auth_user_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpkm67_auth_user_v1_user_proto_goTypes,
		DependencyIndexes: file_rpkm67_auth_user_v1_user_proto_depIdxs,
		MessageInfos:      file_rpkm67_auth_user_v1_user_proto_msgTypes,
	}.Build()
	File_rpkm67_auth_user_v1_user_proto = out.File
	file_rpkm67_auth_user_v1_user_proto_rawDesc = nil
	file_rpkm67_auth_user_v1_user_proto_goTypes = nil
	file_rpkm67_auth_user_v1_user_proto_depIdxs = nil
}
