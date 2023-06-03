// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: ProtoInner.proto

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

type InnerProtoCode int32

const (
	InnerProtoCode_INNER_INVALID               InnerProtoCode = 0
	InnerProtoCode_INNER_SERVER_HAND_SHAKE     InnerProtoCode = -1
	InnerProtoCode_INNER_HEART_BEAT_REQ        InnerProtoCode = -2
	InnerProtoCode_INNER_HEART_BEAT_RES        InnerProtoCode = -3
	InnerProtoCode_INNER_LOGIN_REQ             InnerProtoCode = -4
	InnerProtoCode_INNER_LOGIN_RES             InnerProtoCode = -5
	InnerProtoCode_INNER_PLAYER_DISCONNECT_REQ InnerProtoCode = -6
	InnerProtoCode_INNER_PLAYER_DISCONNECT_RES InnerProtoCode = -7
)

// Enum value maps for InnerProtoCode.
var (
	InnerProtoCode_name = map[int32]string{
		0:  "INNER_INVALID",
		-1: "INNER_SERVER_HAND_SHAKE",
		-2: "INNER_HEART_BEAT_REQ",
		-3: "INNER_HEART_BEAT_RES",
		-4: "INNER_LOGIN_REQ",
		-5: "INNER_LOGIN_RES",
		-6: "INNER_PLAYER_DISCONNECT_REQ",
		-7: "INNER_PLAYER_DISCONNECT_RES",
	}
	InnerProtoCode_value = map[string]int32{
		"INNER_INVALID":               0,
		"INNER_SERVER_HAND_SHAKE":     -1,
		"INNER_HEART_BEAT_REQ":        -2,
		"INNER_HEART_BEAT_RES":        -3,
		"INNER_LOGIN_REQ":             -4,
		"INNER_LOGIN_RES":             -5,
		"INNER_PLAYER_DISCONNECT_REQ": -6,
		"INNER_PLAYER_DISCONNECT_RES": -7,
	}
)

func (x InnerProtoCode) Enum() *InnerProtoCode {
	p := new(InnerProtoCode)
	*p = x
	return p
}

func (x InnerProtoCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InnerProtoCode) Descriptor() protoreflect.EnumDescriptor {
	return file_ProtoInner_proto_enumTypes[0].Descriptor()
}

func (InnerProtoCode) Type() protoreflect.EnumType {
	return &file_ProtoInner_proto_enumTypes[0]
}

func (x InnerProtoCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InnerProtoCode.Descriptor instead.
func (InnerProtoCode) EnumDescriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{0}
}

type InnerHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SendType   int32 `protobuf:"varint,2,opt,name=sendType,proto3" json:"sendType,omitempty"`
	ProtoCode  int32 `protobuf:"varint,3,opt,name=protoCode,proto3" json:"protoCode,omitempty"`
	CallbackId int64 `protobuf:"varint,4,opt,name=callbackId,proto3" json:"callbackId,omitempty"`
}

func (x *InnerHead) Reset() {
	*x = InnerHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerHead) ProtoMessage() {}

func (x *InnerHead) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerHead.ProtoReflect.Descriptor instead.
func (*InnerHead) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{0}
}

func (x *InnerHead) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InnerHead) GetSendType() int32 {
	if x != nil {
		return x.SendType
	}
	return 0
}

func (x *InnerHead) GetProtoCode() int32 {
	if x != nil {
		return x.ProtoCode
	}
	return 0
}

func (x *InnerHead) GetCallbackId() int64 {
	if x != nil {
		return x.CallbackId
	}
	return 0
}

type InnerHeartBeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InnerHeartBeatRequest) Reset() {
	*x = InnerHeartBeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerHeartBeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerHeartBeatRequest) ProtoMessage() {}

func (x *InnerHeartBeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerHeartBeatRequest.ProtoReflect.Descriptor instead.
func (*InnerHeartBeatRequest) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{1}
}

type InnerHeartBeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InnerHeartBeatResponse) Reset() {
	*x = InnerHeartBeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerHeartBeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerHeartBeatResponse) ProtoMessage() {}

func (x *InnerHeartBeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerHeartBeatResponse.ProtoReflect.Descriptor instead.
func (*InnerHeartBeatResponse) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{2}
}

type InnerLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid    int64 `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"` //sessionId
	RoleId int64 `protobuf:"varint,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
}

func (x *InnerLoginRequest) Reset() {
	*x = InnerLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerLoginRequest) ProtoMessage() {}

func (x *InnerLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerLoginRequest.ProtoReflect.Descriptor instead.
func (*InnerLoginRequest) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{3}
}

func (x *InnerLoginRequest) GetSid() int64 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *InnerLoginRequest) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type InnerLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid    int64 `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"` //sessionId
	RoleId int64 `protobuf:"varint,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
}

func (x *InnerLoginResponse) Reset() {
	*x = InnerLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerLoginResponse) ProtoMessage() {}

func (x *InnerLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerLoginResponse.ProtoReflect.Descriptor instead.
func (*InnerLoginResponse) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{4}
}

func (x *InnerLoginResponse) GetSid() int64 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *InnerLoginResponse) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type InnerPlayerDisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid    int64 `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`
	RoleId int64 `protobuf:"varint,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
}

func (x *InnerPlayerDisconnectRequest) Reset() {
	*x = InnerPlayerDisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerPlayerDisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerPlayerDisconnectRequest) ProtoMessage() {}

func (x *InnerPlayerDisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerPlayerDisconnectRequest.ProtoReflect.Descriptor instead.
func (*InnerPlayerDisconnectRequest) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{5}
}

func (x *InnerPlayerDisconnectRequest) GetSid() int64 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *InnerPlayerDisconnectRequest) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type InnerPlayerDisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid    int64 `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`
	RoleId int64 `protobuf:"varint,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
}

func (x *InnerPlayerDisconnectResponse) Reset() {
	*x = InnerPlayerDisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerPlayerDisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerPlayerDisconnectResponse) ProtoMessage() {}

func (x *InnerPlayerDisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerPlayerDisconnectResponse.ProtoReflect.Descriptor instead.
func (*InnerPlayerDisconnectResponse) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{6}
}

func (x *InnerPlayerDisconnectResponse) GetSid() int64 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *InnerPlayerDisconnectResponse) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type InnerLoginWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId           int64  `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	GatewayServerUid int64  `protobuf:"varint,2,opt,name=gatewayServerUid,proto3" json:"gatewayServerUid,omitempty"`
	GameServerUid    int64  `protobuf:"varint,3,opt,name=gameServerUid,proto3" json:"gameServerUid,omitempty"`
	Name             string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	SessionId        int64  `protobuf:"varint,5,opt,name=sessionId,proto3" json:"sessionId,omitempty"` // 网关id
}

func (x *InnerLoginWorldRequest) Reset() {
	*x = InnerLoginWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerLoginWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerLoginWorldRequest) ProtoMessage() {}

func (x *InnerLoginWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerLoginWorldRequest.ProtoReflect.Descriptor instead.
func (*InnerLoginWorldRequest) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{7}
}

func (x *InnerLoginWorldRequest) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *InnerLoginWorldRequest) GetGatewayServerUid() int64 {
	if x != nil {
		return x.GatewayServerUid
	}
	return 0
}

func (x *InnerLoginWorldRequest) GetGameServerUid() int64 {
	if x != nil {
		return x.GameServerUid
	}
	return 0
}

func (x *InnerLoginWorldRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InnerLoginWorldRequest) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

type InnerLoginWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode int32 `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	UnitId    int64 `protobuf:"varint,2,opt,name=unitId,proto3" json:"unitId,omitempty"`
}

func (x *InnerLoginWorldResponse) Reset() {
	*x = InnerLoginWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerLoginWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerLoginWorldResponse) ProtoMessage() {}

func (x *InnerLoginWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerLoginWorldResponse.ProtoReflect.Descriptor instead.
func (*InnerLoginWorldResponse) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{8}
}

func (x *InnerLoginWorldResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *InnerLoginWorldResponse) GetUnitId() int64 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

type InnerLogoutNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId    int64 `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	SessionId int64 `protobuf:"varint,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *InnerLogoutNotify) Reset() {
	*x = InnerLogoutNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerLogoutNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerLogoutNotify) ProtoMessage() {}

func (x *InnerLogoutNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerLogoutNotify.ProtoReflect.Descriptor instead.
func (*InnerLogoutNotify) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{9}
}

func (x *InnerLogoutNotify) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *InnerLogoutNotify) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

type InnerLoginInitNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId int64 `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
}

func (x *InnerLoginInitNotify) Reset() {
	*x = InnerLoginInitNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerLoginInitNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerLoginInitNotify) ProtoMessage() {}

func (x *InnerLoginInitNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerLoginInitNotify.ProtoReflect.Descriptor instead.
func (*InnerLoginInitNotify) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{10}
}

func (x *InnerLoginInitNotify) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

type InnerServerHandShake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromServerId   int64 `protobuf:"varint,1,opt,name=fromServerId,proto3" json:"fromServerId,omitempty"`
	FromServerType int32 `protobuf:"varint,2,opt,name=fromServerType,proto3" json:"fromServerType,omitempty"`
}

func (x *InnerServerHandShake) Reset() {
	*x = InnerServerHandShake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoInner_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerServerHandShake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerServerHandShake) ProtoMessage() {}

func (x *InnerServerHandShake) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoInner_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerServerHandShake.ProtoReflect.Descriptor instead.
func (*InnerServerHandShake) Descriptor() ([]byte, []int) {
	return file_ProtoInner_proto_rawDescGZIP(), []int{11}
}

func (x *InnerServerHandShake) GetFromServerId() int64 {
	if x != nil {
		return x.FromServerId
	}
	return 0
}

func (x *InnerServerHandShake) GetFromServerType() int32 {
	if x != nil {
		return x.FromServerType
	}
	return 0
}

var File_ProtoInner_proto protoreflect.FileDescriptor

var file_ProtoInner_proto_rawDesc = []byte{
	0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x75, 0x0a, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x65, 0x61, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x11,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x73, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x12, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x73, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x1c, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x1d, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0xb4, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x17, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x11, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x14, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x49, 0x6e, 0x69, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x14, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66,
	0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x9f, 0x02, 0x0a, 0x0e, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e,
	0x4e, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a,
	0x17, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x48, 0x41,
	0x4e, 0x44, 0x5f, 0x53, 0x48, 0x41, 0x4b, 0x45, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x12, 0x21, 0x0a, 0x14, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x48, 0x45, 0x41,
	0x52, 0x54, 0x5f, 0x42, 0x45, 0x41, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xfe, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x21, 0x0a, 0x14, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x5f,
	0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x42, 0x45, 0x41, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x10, 0xfd,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1c, 0x0a, 0x0f, 0x49, 0x4e, 0x4e,
	0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xfc, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1c, 0x0a, 0x0f, 0x49, 0x4e, 0x4e, 0x45, 0x52,
	0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x10, 0xfb, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x28, 0x0a, 0x1b, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x50,
	0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x5f, 0x52, 0x45, 0x51, 0x10, 0xfa, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x28, 0x0a, 0x1b, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f,
	0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x10, 0xf9,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x42, 0x1b, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x09, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x47, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ProtoInner_proto_rawDescOnce sync.Once
	file_ProtoInner_proto_rawDescData = file_ProtoInner_proto_rawDesc
)

func file_ProtoInner_proto_rawDescGZIP() []byte {
	file_ProtoInner_proto_rawDescOnce.Do(func() {
		file_ProtoInner_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProtoInner_proto_rawDescData)
	})
	return file_ProtoInner_proto_rawDescData
}

var file_ProtoInner_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ProtoInner_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_ProtoInner_proto_goTypes = []interface{}{
	(InnerProtoCode)(0),                   // 0: InnerProtoCode
	(*InnerHead)(nil),                     // 1: InnerHead
	(*InnerHeartBeatRequest)(nil),         // 2: InnerHeartBeatRequest
	(*InnerHeartBeatResponse)(nil),        // 3: InnerHeartBeatResponse
	(*InnerLoginRequest)(nil),             // 4: InnerLoginRequest
	(*InnerLoginResponse)(nil),            // 5: InnerLoginResponse
	(*InnerPlayerDisconnectRequest)(nil),  // 6: InnerPlayerDisconnectRequest
	(*InnerPlayerDisconnectResponse)(nil), // 7: InnerPlayerDisconnectResponse
	(*InnerLoginWorldRequest)(nil),        // 8: InnerLoginWorldRequest
	(*InnerLoginWorldResponse)(nil),       // 9: InnerLoginWorldResponse
	(*InnerLogoutNotify)(nil),             // 10: InnerLogoutNotify
	(*InnerLoginInitNotify)(nil),          // 11: InnerLoginInitNotify
	(*InnerServerHandShake)(nil),          // 12: InnerServerHandShake
}
var file_ProtoInner_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ProtoInner_proto_init() }
func file_ProtoInner_proto_init() {
	if File_ProtoInner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ProtoInner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerHead); i {
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
		file_ProtoInner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerHeartBeatRequest); i {
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
		file_ProtoInner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerHeartBeatResponse); i {
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
		file_ProtoInner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerLoginRequest); i {
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
		file_ProtoInner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerLoginResponse); i {
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
		file_ProtoInner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerPlayerDisconnectRequest); i {
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
		file_ProtoInner_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerPlayerDisconnectResponse); i {
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
		file_ProtoInner_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerLoginWorldRequest); i {
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
		file_ProtoInner_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerLoginWorldResponse); i {
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
		file_ProtoInner_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerLogoutNotify); i {
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
		file_ProtoInner_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerLoginInitNotify); i {
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
		file_ProtoInner_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerServerHandShake); i {
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
			RawDescriptor: file_ProtoInner_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProtoInner_proto_goTypes,
		DependencyIndexes: file_ProtoInner_proto_depIdxs,
		EnumInfos:         file_ProtoInner_proto_enumTypes,
		MessageInfos:      file_ProtoInner_proto_msgTypes,
	}.Build()
	File_ProtoInner_proto = out.File
	file_ProtoInner_proto_rawDesc = nil
	file_ProtoInner_proto_goTypes = nil
	file_ProtoInner_proto_depIdxs = nil
}
