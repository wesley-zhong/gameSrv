// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.26.0
// source: ProtoTask.proto

package protoGen

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

type IntIntProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntValue1 int32 `protobuf:"varint,1,opt,name=intValue1,proto3" json:"intValue1,omitempty"`
	IntValue2 int32 `protobuf:"varint,2,opt,name=intValue2,proto3" json:"intValue2,omitempty"`
}

func (x *IntIntProto) Reset() {
	*x = IntIntProto{}
	mi := &file_ProtoTask_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntIntProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntIntProto) ProtoMessage() {}

func (x *IntIntProto) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoTask_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntIntProto.ProtoReflect.Descriptor instead.
func (*IntIntProto) Descriptor() ([]byte, []int) {
	return file_ProtoTask_proto_rawDescGZIP(), []int{0}
}

func (x *IntIntProto) GetIntValue1() int32 {
	if x != nil {
		return x.IntValue1
	}
	return 0
}

func (x *IntIntProto) GetIntValue2() int32 {
	if x != nil {
		return x.IntValue2
	}
	return 0
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_ProtoTask_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoTask_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_ProtoTask_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ItemProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId int32 `protobuf:"varint,1,opt,name=itemId,proto3" json:"itemId,omitempty"`
	Count  int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ItemProto) Reset() {
	*x = ItemProto{}
	mi := &file_ProtoTask_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemProto) ProtoMessage() {}

func (x *ItemProto) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoTask_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemProto.ProtoReflect.Descriptor instead.
func (*ItemProto) Descriptor() ([]byte, []int) {
	return file_ProtoTask_proto_rawDescGZIP(), []int{2}
}

func (x *ItemProto) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ItemProto) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId  int64  `protobuf:"varint,1,opt,name=accountId,proto3" json:"accountId,omitempty"`   // 账号
	RoleId     int64  `protobuf:"varint,2,opt,name=roleId,proto3" json:"roleId,omitempty"`         // 角色
	LoginToken string `protobuf:"bytes,3,opt,name=loginToken,proto3" json:"loginToken,omitempty"`  // login token
	GameTicket int32  `protobuf:"varint,4,opt,name=gameTicket,proto3" json:"gameTicket,omitempty"` // game ticket
	ServerId   int32  `protobuf:"varint,5,opt,name=serverId,proto3" json:"serverId,omitempty"`     // 游戏服务器id
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_ProtoTask_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoTask_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_ProtoTask_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *LoginRequest) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *LoginRequest) GetLoginToken() string {
	if x != nil {
		return x.LoginToken
	}
	return ""
}

func (x *LoginRequest) GetGameTicket() int32 {
	if x != nil {
		return x.GameTicket
	}
	return 0
}

func (x *LoginRequest) GetServerId() int32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode  int32 `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`   // 错误码
	ServerTime int64 `protobuf:"varint,2,opt,name=serverTime,proto3" json:"serverTime,omitempty"` // 服务器当前时间
	RoleId     int64 `protobuf:"varint,3,opt,name=roleId,proto3" json:"roleId,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_ProtoTask_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoTask_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_ProtoTask_proto_rawDescGZIP(), []int{4}
}

func (x *LoginResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *LoginResponse) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *LoginResponse) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId int64 `protobuf:"varint,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_ProtoTask_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoTask_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_ProtoTask_proto_rawDescGZIP(), []int{5}
}

func (x *LogoutRequest) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode int32 `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_ProtoTask_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoTask_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_ProtoTask_proto_rawDescGZIP(), []int{6}
}

func (x *LogoutResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

type PerformanceTestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeId   int64  `protobuf:"varint,1,opt,name=someId,proto3" json:"someId,omitempty"`
	SomeBody string `protobuf:"bytes,2,opt,name=someBody,proto3" json:"someBody,omitempty"`
}

func (x *PerformanceTestReq) Reset() {
	*x = PerformanceTestReq{}
	mi := &file_ProtoTask_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PerformanceTestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerformanceTestReq) ProtoMessage() {}

func (x *PerformanceTestReq) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoTask_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerformanceTestReq.ProtoReflect.Descriptor instead.
func (*PerformanceTestReq) Descriptor() ([]byte, []int) {
	return file_ProtoTask_proto_rawDescGZIP(), []int{7}
}

func (x *PerformanceTestReq) GetSomeId() int64 {
	if x != nil {
		return x.SomeId
	}
	return 0
}

func (x *PerformanceTestReq) GetSomeBody() string {
	if x != nil {
		return x.SomeBody
	}
	return ""
}

type PerformanceTestRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeId    int64  `protobuf:"varint,1,opt,name=someId,proto3" json:"someId,omitempty"`
	ResBody   string `protobuf:"bytes,2,opt,name=resBody,proto3" json:"resBody,omitempty"`
	SomeIdAdd int64  `protobuf:"varint,3,opt,name=someIdAdd,proto3" json:"someIdAdd,omitempty"`
}

func (x *PerformanceTestRes) Reset() {
	*x = PerformanceTestRes{}
	mi := &file_ProtoTask_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PerformanceTestRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerformanceTestRes) ProtoMessage() {}

func (x *PerformanceTestRes) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoTask_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerformanceTestRes.ProtoReflect.Descriptor instead.
func (*PerformanceTestRes) Descriptor() ([]byte, []int) {
	return file_ProtoTask_proto_rawDescGZIP(), []int{8}
}

func (x *PerformanceTestRes) GetSomeId() int64 {
	if x != nil {
		return x.SomeId
	}
	return 0
}

func (x *PerformanceTestRes) GetResBody() string {
	if x != nil {
		return x.ResBody
	}
	return ""
}

func (x *PerformanceTestRes) GetSomeIdAdd() int64 {
	if x != nil {
		return x.SomeIdAdd
	}
	return 0
}

type EchoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestBody string `protobuf:"bytes,1,opt,name=requestBody,proto3" json:"requestBody,omitempty"`
	SomeId      int64  `protobuf:"varint,2,opt,name=someId,proto3" json:"someId,omitempty"`
}

func (x *EchoReq) Reset() {
	*x = EchoReq{}
	mi := &file_ProtoTask_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoReq) ProtoMessage() {}

func (x *EchoReq) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoTask_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoReq.ProtoReflect.Descriptor instead.
func (*EchoReq) Descriptor() ([]byte, []int) {
	return file_ProtoTask_proto_rawDescGZIP(), []int{9}
}

func (x *EchoReq) GetRequestBody() string {
	if x != nil {
		return x.RequestBody
	}
	return ""
}

func (x *EchoReq) GetSomeId() int64 {
	if x != nil {
		return x.SomeId
	}
	return 0
}

var File_ProtoTask_proto protoreflect.FileDescriptor

var file_ProtoTask_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x49, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x22, 0x2a, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0xa0, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x61,
	0x6d, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x67, 0x61, 0x6d, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x2d, 0x0a,
	0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0e,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x48, 0x0a, 0x12,
	0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f,
	0x6d, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f,
	0x6d, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x64, 0x0a, 0x12, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6f,
	0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x42, 0x6f, 0x64, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x41, 0x64, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x6f, 0x6d, 0x65, 0x49, 0x64, 0x41, 0x64, 0x64, 0x22, 0x43, 0x0a, 0x07,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x6d,
	0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6f, 0x6d, 0x65, 0x49,
	0x64, 0x42, 0x1b, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ProtoTask_proto_rawDescOnce sync.Once
	file_ProtoTask_proto_rawDescData = file_ProtoTask_proto_rawDesc
)

func file_ProtoTask_proto_rawDescGZIP() []byte {
	file_ProtoTask_proto_rawDescOnce.Do(func() {
		file_ProtoTask_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProtoTask_proto_rawDescData)
	})
	return file_ProtoTask_proto_rawDescData
}

var file_ProtoTask_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_ProtoTask_proto_goTypes = []any{
	(*IntIntProto)(nil),        // 0: IntIntProto
	(*Role)(nil),               // 1: Role
	(*ItemProto)(nil),          // 2: ItemProto
	(*LoginRequest)(nil),       // 3: LoginRequest
	(*LoginResponse)(nil),      // 4: LoginResponse
	(*LogoutRequest)(nil),      // 5: LogoutRequest
	(*LogoutResponse)(nil),     // 6: LogoutResponse
	(*PerformanceTestReq)(nil), // 7: PerformanceTestReq
	(*PerformanceTestRes)(nil), // 8: PerformanceTestRes
	(*EchoReq)(nil),            // 9: EchoReq
}
var file_ProtoTask_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ProtoTask_proto_init() }
func file_ProtoTask_proto_init() {
	if File_ProtoTask_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ProtoTask_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProtoTask_proto_goTypes,
		DependencyIndexes: file_ProtoTask_proto_depIdxs,
		MessageInfos:      file_ProtoTask_proto_msgTypes,
	}.Build()
	File_ProtoTask_proto = out.File
	file_ProtoTask_proto_rawDesc = nil
	file_ProtoTask_proto_goTypes = nil
	file_ProtoTask_proto_depIdxs = nil
}
