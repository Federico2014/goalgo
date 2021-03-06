// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_broker.proto

package goalgo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RobotExchangeInfoRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	RobotId              string   `protobuf:"bytes,2,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RobotExchangeInfoRequest) Reset()         { *m = RobotExchangeInfoRequest{} }
func (m *RobotExchangeInfoRequest) String() string { return proto.CompactTextString(m) }
func (*RobotExchangeInfoRequest) ProtoMessage()    {}
func (*RobotExchangeInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{0}
}

func (m *RobotExchangeInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RobotExchangeInfoRequest.Unmarshal(m, b)
}
func (m *RobotExchangeInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RobotExchangeInfoRequest.Marshal(b, m, deterministic)
}
func (m *RobotExchangeInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RobotExchangeInfoRequest.Merge(m, src)
}
func (m *RobotExchangeInfoRequest) XXX_Size() int {
	return xxx_messageInfo_RobotExchangeInfoRequest.Size(m)
}
func (m *RobotExchangeInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RobotExchangeInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RobotExchangeInfoRequest proto.InternalMessageInfo

func (m *RobotExchangeInfoRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RobotExchangeInfoRequest) GetRobotId() string {
	if m != nil {
		return m.RobotId
	}
	return ""
}

type RobotExchangeInfo struct {
	Label                string   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AccessKey            string   `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey            string   `protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RobotExchangeInfo) Reset()         { *m = RobotExchangeInfo{} }
func (m *RobotExchangeInfo) String() string { return proto.CompactTextString(m) }
func (*RobotExchangeInfo) ProtoMessage()    {}
func (*RobotExchangeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{1}
}

func (m *RobotExchangeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RobotExchangeInfo.Unmarshal(m, b)
}
func (m *RobotExchangeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RobotExchangeInfo.Marshal(b, m, deterministic)
}
func (m *RobotExchangeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RobotExchangeInfo.Merge(m, src)
}
func (m *RobotExchangeInfo) XXX_Size() int {
	return xxx_messageInfo_RobotExchangeInfo.Size(m)
}
func (m *RobotExchangeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RobotExchangeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RobotExchangeInfo proto.InternalMessageInfo

func (m *RobotExchangeInfo) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *RobotExchangeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RobotExchangeInfo) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *RobotExchangeInfo) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

type RobotExchangeInfoReply struct {
	Exchanges            []*RobotExchangeInfo `protobuf:"bytes,1,rep,name=exchanges,proto3" json:"exchanges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RobotExchangeInfoReply) Reset()         { *m = RobotExchangeInfoReply{} }
func (m *RobotExchangeInfoReply) String() string { return proto.CompactTextString(m) }
func (*RobotExchangeInfoReply) ProtoMessage()    {}
func (*RobotExchangeInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{2}
}

func (m *RobotExchangeInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RobotExchangeInfoReply.Unmarshal(m, b)
}
func (m *RobotExchangeInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RobotExchangeInfoReply.Marshal(b, m, deterministic)
}
func (m *RobotExchangeInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RobotExchangeInfoReply.Merge(m, src)
}
func (m *RobotExchangeInfoReply) XXX_Size() int {
	return xxx_messageInfo_RobotExchangeInfoReply.Size(m)
}
func (m *RobotExchangeInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RobotExchangeInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_RobotExchangeInfoReply proto.InternalMessageInfo

func (m *RobotExchangeInfoReply) GetExchanges() []*RobotExchangeInfo {
	if m != nil {
		return m.Exchanges
	}
	return nil
}

type GetValueRequest struct {
	RobotId              string   `protobuf:"bytes,1,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetValueRequest) Reset()         { *m = GetValueRequest{} }
func (m *GetValueRequest) String() string { return proto.CompactTextString(m) }
func (*GetValueRequest) ProtoMessage()    {}
func (*GetValueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{3}
}

func (m *GetValueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetValueRequest.Unmarshal(m, b)
}
func (m *GetValueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetValueRequest.Marshal(b, m, deterministic)
}
func (m *GetValueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetValueRequest.Merge(m, src)
}
func (m *GetValueRequest) XXX_Size() int {
	return xxx_messageInfo_GetValueRequest.Size(m)
}
func (m *GetValueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetValueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetValueRequest proto.InternalMessageInfo

func (m *GetValueRequest) GetRobotId() string {
	if m != nil {
		return m.RobotId
	}
	return ""
}

func (m *GetValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetValueReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetValueReply) Reset()         { *m = GetValueReply{} }
func (m *GetValueReply) String() string { return proto.CompactTextString(m) }
func (*GetValueReply) ProtoMessage()    {}
func (*GetValueReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{4}
}

func (m *GetValueReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetValueReply.Unmarshal(m, b)
}
func (m *GetValueReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetValueReply.Marshal(b, m, deterministic)
}
func (m *GetValueReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetValueReply.Merge(m, src)
}
func (m *GetValueReply) XXX_Size() int {
	return xxx_messageInfo_GetValueReply.Size(m)
}
func (m *GetValueReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetValueReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetValueReply proto.InternalMessageInfo

func (m *GetValueReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetValueReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetValueReply) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetValueRequest struct {
	RobotId              string   `protobuf:"bytes,1,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetValueRequest) Reset()         { *m = SetValueRequest{} }
func (m *SetValueRequest) String() string { return proto.CompactTextString(m) }
func (*SetValueRequest) ProtoMessage()    {}
func (*SetValueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{5}
}

func (m *SetValueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetValueRequest.Unmarshal(m, b)
}
func (m *SetValueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetValueRequest.Marshal(b, m, deterministic)
}
func (m *SetValueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetValueRequest.Merge(m, src)
}
func (m *SetValueRequest) XXX_Size() int {
	return xxx_messageInfo_SetValueRequest.Size(m)
}
func (m *SetValueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetValueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetValueRequest proto.InternalMessageInfo

func (m *SetValueRequest) GetRobotId() string {
	if m != nil {
		return m.RobotId
	}
	return ""
}

func (m *SetValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetValueRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetValueReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetValueReply) Reset()         { *m = SetValueReply{} }
func (m *SetValueReply) String() string { return proto.CompactTextString(m) }
func (*SetValueReply) ProtoMessage()    {}
func (*SetValueReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{6}
}

func (m *SetValueReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetValueReply.Unmarshal(m, b)
}
func (m *SetValueReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetValueReply.Marshal(b, m, deterministic)
}
func (m *SetValueReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetValueReply.Merge(m, src)
}
func (m *SetValueReply) XXX_Size() int {
	return xxx_messageInfo_SetValueReply.Size(m)
}
func (m *SetValueReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SetValueReply.DiscardUnknown(m)
}

var xxx_messageInfo_SetValueReply proto.InternalMessageInfo

func (m *SetValueReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SetValueReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LogRequest struct {
	Sid                  int32    `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Time                 int64    `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Level                int32    `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{7}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetSid() int32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

func (m *LogRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LogRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *LogRequest) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *LogRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LogReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogReply) Reset()         { *m = LogReply{} }
func (m *LogReply) String() string { return proto.CompactTextString(m) }
func (*LogReply) ProtoMessage()    {}
func (*LogReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{8}
}

func (m *LogReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogReply.Unmarshal(m, b)
}
func (m *LogReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogReply.Marshal(b, m, deterministic)
}
func (m *LogReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogReply.Merge(m, src)
}
func (m *LogReply) XXX_Size() int {
	return xxx_messageInfo_LogReply.Size(m)
}
func (m *LogReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LogReply.DiscardUnknown(m)
}

var xxx_messageInfo_LogReply proto.InternalMessageInfo

func (m *LogReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LogReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UpdateStatusRequest struct {
	RobotId              string   `protobuf:"bytes,1,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStatusRequest) Reset()         { *m = UpdateStatusRequest{} }
func (m *UpdateStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateStatusRequest) ProtoMessage()    {}
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{9}
}

func (m *UpdateStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStatusRequest.Unmarshal(m, b)
}
func (m *UpdateStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStatusRequest.Merge(m, src)
}
func (m *UpdateStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateStatusRequest.Size(m)
}
func (m *UpdateStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStatusRequest proto.InternalMessageInfo

func (m *UpdateStatusRequest) GetRobotId() string {
	if m != nil {
		return m.RobotId
	}
	return ""
}

func (m *UpdateStatusRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type UpdateStatusReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStatusReply) Reset()         { *m = UpdateStatusReply{} }
func (m *UpdateStatusReply) String() string { return proto.CompactTextString(m) }
func (*UpdateStatusReply) ProtoMessage()    {}
func (*UpdateStatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e9beed3ec3b28, []int{10}
}

func (m *UpdateStatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStatusReply.Unmarshal(m, b)
}
func (m *UpdateStatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStatusReply.Marshal(b, m, deterministic)
}
func (m *UpdateStatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStatusReply.Merge(m, src)
}
func (m *UpdateStatusReply) XXX_Size() int {
	return xxx_messageInfo_UpdateStatusReply.Size(m)
}
func (m *UpdateStatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStatusReply proto.InternalMessageInfo

func (m *UpdateStatusReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UpdateStatusReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RobotExchangeInfoRequest)(nil), "goalgo.RobotExchangeInfoRequest")
	proto.RegisterType((*RobotExchangeInfo)(nil), "goalgo.RobotExchangeInfo")
	proto.RegisterType((*RobotExchangeInfoReply)(nil), "goalgo.RobotExchangeInfoReply")
	proto.RegisterType((*GetValueRequest)(nil), "goalgo.GetValueRequest")
	proto.RegisterType((*GetValueReply)(nil), "goalgo.GetValueReply")
	proto.RegisterType((*SetValueRequest)(nil), "goalgo.SetValueRequest")
	proto.RegisterType((*SetValueReply)(nil), "goalgo.SetValueReply")
	proto.RegisterType((*LogRequest)(nil), "goalgo.LogRequest")
	proto.RegisterType((*LogReply)(nil), "goalgo.LogReply")
	proto.RegisterType((*UpdateStatusRequest)(nil), "goalgo.UpdateStatusRequest")
	proto.RegisterType((*UpdateStatusReply)(nil), "goalgo.UpdateStatusReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RobotCtlClient is the client API for RobotCtl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RobotCtlClient interface {
	// 获取交易所配置
	GetRobotExchangeInfo(ctx context.Context, in *RobotExchangeInfoRequest, opts ...grpc.CallOption) (*RobotExchangeInfoReply, error)
	// 获取配置值
	GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueReply, error)
	// 保存配置值
	SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueReply, error)
	// 写入日志
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error)
	// 更新状态
	UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*UpdateStatusReply, error)
}

type robotCtlClient struct {
	cc *grpc.ClientConn
}

func NewRobotCtlClient(cc *grpc.ClientConn) RobotCtlClient {
	return &robotCtlClient{cc}
}

func (c *robotCtlClient) GetRobotExchangeInfo(ctx context.Context, in *RobotExchangeInfoRequest, opts ...grpc.CallOption) (*RobotExchangeInfoReply, error) {
	out := new(RobotExchangeInfoReply)
	err := c.cc.Invoke(ctx, "/goalgo.RobotCtl/GetRobotExchangeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotCtlClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueReply, error) {
	out := new(GetValueReply)
	err := c.cc.Invoke(ctx, "/goalgo.RobotCtl/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotCtlClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueReply, error) {
	out := new(SetValueReply)
	err := c.cc.Invoke(ctx, "/goalgo.RobotCtl/SetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotCtlClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error) {
	out := new(LogReply)
	err := c.cc.Invoke(ctx, "/goalgo.RobotCtl/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotCtlClient) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*UpdateStatusReply, error) {
	out := new(UpdateStatusReply)
	err := c.cc.Invoke(ctx, "/goalgo.RobotCtl/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotCtlServer is the server API for RobotCtl service.
type RobotCtlServer interface {
	// 获取交易所配置
	GetRobotExchangeInfo(context.Context, *RobotExchangeInfoRequest) (*RobotExchangeInfoReply, error)
	// 获取配置值
	GetValue(context.Context, *GetValueRequest) (*GetValueReply, error)
	// 保存配置值
	SetValue(context.Context, *SetValueRequest) (*SetValueReply, error)
	// 写入日志
	Log(context.Context, *LogRequest) (*LogReply, error)
	// 更新状态
	UpdateStatus(context.Context, *UpdateStatusRequest) (*UpdateStatusReply, error)
}

func RegisterRobotCtlServer(s *grpc.Server, srv RobotCtlServer) {
	s.RegisterService(&_RobotCtl_serviceDesc, srv)
}

func _RobotCtl_GetRobotExchangeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RobotExchangeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotCtlServer).GetRobotExchangeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalgo.RobotCtl/GetRobotExchangeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotCtlServer).GetRobotExchangeInfo(ctx, req.(*RobotExchangeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotCtl_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotCtlServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalgo.RobotCtl/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotCtlServer).GetValue(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotCtl_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotCtlServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalgo.RobotCtl/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotCtlServer).SetValue(ctx, req.(*SetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotCtl_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotCtlServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalgo.RobotCtl/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotCtlServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotCtl_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotCtlServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalgo.RobotCtl/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotCtlServer).UpdateStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RobotCtl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goalgo.RobotCtl",
	HandlerType: (*RobotCtlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRobotExchangeInfo",
			Handler:    _RobotCtl_GetRobotExchangeInfo_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _RobotCtl_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _RobotCtl_SetValue_Handler,
		},
		{
			MethodName: "Log",
			Handler:    _RobotCtl_Log_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _RobotCtl_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_broker.proto",
}

func init() { proto.RegisterFile("grpc_broker.proto", fileDescriptor_802e9beed3ec3b28) }

var fileDescriptor_802e9beed3ec3b28 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x25, 0x71, 0x9c, 0x3a, 0x43, 0x4b, 0x9b, 0xa5, 0x14, 0x27, 0x08, 0x14, 0xed, 0xa9, 0x17,
	0x72, 0x28, 0x07, 0x2e, 0xa8, 0x97, 0x0a, 0xa5, 0x55, 0x7b, 0x61, 0x0d, 0x48, 0x70, 0x89, 0x36,
	0xf6, 0x60, 0xa2, 0x3a, 0x59, 0xe3, 0x5d, 0x57, 0x58, 0x7c, 0x2a, 0x3f, 0x83, 0x76, 0xd7, 0x26,
	0x71, 0x9d, 0x20, 0xd4, 0xdc, 0x66, 0xe6, 0xcd, 0xbc, 0x7d, 0x33, 0x7e, 0x09, 0xf4, 0xe3, 0x2c,
	0x0d, 0xa7, 0xb3, 0x4c, 0xdc, 0x62, 0x36, 0x4e, 0x33, 0xa1, 0x04, 0xe9, 0xc6, 0x82, 0x27, 0xb1,
	0xa0, 0x13, 0xf0, 0x99, 0x98, 0x09, 0xf5, 0xfe, 0x67, 0xf8, 0x9d, 0x2f, 0x63, 0xbc, 0x5a, 0x7e,
	0x13, 0x0c, 0x7f, 0xe4, 0x28, 0x15, 0x39, 0x02, 0x27, 0x9f, 0x47, 0x7e, 0x6b, 0xd4, 0x3a, 0xed,
	0x31, 0x1d, 0x92, 0x01, 0x78, 0x99, 0xee, 0x9e, 0xce, 0x23, 0xbf, 0x6d, 0xca, 0x7b, 0x26, 0xbf,
	0x8a, 0xe8, 0x2f, 0xe8, 0x37, 0x88, 0xc8, 0x31, 0xb8, 0x09, 0x9f, 0x61, 0x52, 0x72, 0xd8, 0x84,
	0x10, 0xe8, 0x2c, 0xf9, 0x02, 0x4b, 0x06, 0x13, 0x93, 0x97, 0x00, 0x3c, 0x0c, 0x51, 0xca, 0xe9,
	0x2d, 0x16, 0xbe, 0x63, 0x90, 0x9e, 0xad, 0x5c, 0x63, 0xa1, 0x61, 0x89, 0x61, 0x86, 0xca, 0xc0,
	0x1d, 0x0b, 0xdb, 0xca, 0x35, 0x16, 0xf4, 0x03, 0x9c, 0x6c, 0xd8, 0x22, 0x4d, 0x0a, 0xf2, 0x16,
	0x7a, 0x58, 0x16, 0xa5, 0xdf, 0x1a, 0x39, 0xa7, 0x8f, 0xcf, 0x06, 0x63, 0xbb, 0xfb, 0xb8, 0x39,
	0xb2, 0xea, 0xa5, 0xe7, 0x70, 0x38, 0x41, 0xf5, 0x99, 0x27, 0x39, 0x56, 0xf7, 0x58, 0xdf, 0xbe,
	0x55, 0xdb, 0x5e, 0x9f, 0x4a, 0x0b, 0xb3, 0x1b, 0xe9, 0x90, 0x7e, 0x81, 0x83, 0xd5, 0xbc, 0x56,
	0xe2, 0xc3, 0x9e, 0xcc, 0xcd, 0x42, 0x66, 0xd8, 0x63, 0x55, 0xaa, 0x91, 0x05, 0x4a, 0xc9, 0xe3,
	0xea, 0x24, 0x55, 0xaa, 0xef, 0x77, 0xa7, 0x19, 0xcc, 0x41, 0xf6, 0x99, 0x4d, 0xe8, 0x47, 0x38,
	0x0c, 0x76, 0x90, 0xb6, 0x85, 0xf5, 0x02, 0x0e, 0x82, 0x5d, 0x05, 0xd3, 0x0c, 0xe0, 0x46, 0xc4,
	0x6b, 0x06, 0x92, 0xa5, 0x20, 0x97, 0xe9, 0x90, 0x3c, 0x81, 0x76, 0x69, 0x9d, 0x0e, 0x6b, 0xcf,
	0x23, 0x6d, 0x05, 0x35, 0x5f, 0x58, 0x25, 0x0e, 0x33, 0xb1, 0x31, 0x0d, 0xde, 0x61, 0x62, 0x3e,
	0xb3, 0xcb, 0x6c, 0xb2, 0xfe, 0xa6, 0x5b, 0x7f, 0xf3, 0x1c, 0x3c, 0xf3, 0xe6, 0x43, 0x35, 0x5f,
	0xc2, 0xd3, 0x4f, 0x69, 0xc4, 0x15, 0x06, 0x8a, 0xab, 0x5c, 0xfe, 0xc7, 0x49, 0x4f, 0xa0, 0x2b,
	0x4d, 0xaf, 0xa1, 0x72, 0x59, 0x99, 0xd1, 0x09, 0xf4, 0xeb, 0x4c, 0x0f, 0x94, 0x74, 0xf6, 0xbb,
	0x0d, 0x9e, 0x71, 0xe7, 0x85, 0x4a, 0xc8, 0x57, 0x38, 0x9e, 0xa0, 0x6a, 0xfe, 0xb8, 0x46, 0xdb,
	0x7d, 0x6c, 0x57, 0x18, 0xbe, 0xfa, 0x47, 0x47, 0x9a, 0x14, 0xf4, 0x11, 0x79, 0x07, 0x5e, 0xe5,
	0x52, 0xf2, 0xbc, 0xea, 0xbe, 0xe7, 0xfb, 0xe1, 0xb3, 0x26, 0xf0, 0x77, 0x3a, 0x68, 0x4c, 0x07,
	0xdb, 0xa6, 0x83, 0x7b, 0xd3, 0xaf, 0xc1, 0xb9, 0x11, 0x31, 0x21, 0x15, 0xbe, 0x32, 0xce, 0xf0,
	0xa8, 0x56, 0xb3, 0xed, 0x97, 0xb0, 0xbf, 0x7e, 0x5c, 0xf2, 0xa2, 0xea, 0xd9, 0xf0, 0xf1, 0x86,
	0x83, 0xcd, 0xa0, 0x61, 0x9a, 0x75, 0xcd, 0x5f, 0xe0, 0x9b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf2, 0xe8, 0xc4, 0xc8, 0x17, 0x05, 0x00, 0x00,
}
