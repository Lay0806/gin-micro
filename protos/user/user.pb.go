// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Gender int32

const (
	// 女
	Gender_Women Gender = 0
	// 男
	Gender_Man Gender = 1
)

var Gender_name = map[int32]string{
	0: "Women",
	1: "Man",
}
var Gender_value = map[string]int32{
	"Women": 0,
	"Man":   1,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}
func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_4874d0e054b024d7, []int{0}
}

// 用户登陆结构体
type LoginParams struct {
	// 用户名
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// 密码
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// 验证码
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginParams) Reset()         { *m = LoginParams{} }
func (m *LoginParams) String() string { return proto.CompactTextString(m) }
func (*LoginParams) ProtoMessage()    {}
func (*LoginParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_4874d0e054b024d7, []int{0}
}
func (m *LoginParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginParams.Unmarshal(m, b)
}
func (m *LoginParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginParams.Marshal(b, m, deterministic)
}
func (dst *LoginParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginParams.Merge(dst, src)
}
func (m *LoginParams) XXX_Size() int {
	return xxx_messageInfo_LoginParams.Size(m)
}
func (m *LoginParams) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginParams.DiscardUnknown(m)
}

var xxx_messageInfo_LoginParams proto.InternalMessageInfo

func (m *LoginParams) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginParams) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginParams) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type UserInfo struct {
	// id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 账号
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	// 昵称
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// 密码
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// 邮箱
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_4874d0e054b024d7, []int{1}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserInfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// 响应结构体
type Result struct {
	// 响应码
	Status int32 `protobuf:"zigzag32,1,opt,name=status,proto3" json:"status,omitempty"`
	// 响应参数
	Map                  map[string]string `protobuf:"bytes,2,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_4874d0e054b024d7, []int{2}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Result) GetMap() map[string]string {
	if m != nil {
		return m.Map
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginParams)(nil), "user.LoginParams")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
	proto.RegisterType((*Result)(nil), "user.Result")
	proto.RegisterMapType((map[string]string)(nil), "user.Result.MapEntry")
	proto.RegisterEnum("user.Gender", Gender_name, Gender_value)
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_4874d0e054b024d7) }

var fileDescriptor_user_4874d0e054b024d7 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4b, 0xeb, 0x40,
	0x10, 0xc6, 0x5f, 0x92, 0x26, 0x6d, 0xa7, 0x8f, 0xd2, 0x0e, 0x2a, 0xa1, 0x78, 0x28, 0x3d, 0x68,
	0xf1, 0xd0, 0x43, 0x05, 0x11, 0xef, 0x22, 0x05, 0x0b, 0x92, 0x22, 0x3d, 0xca, 0x9a, 0x8c, 0x65,
	0x69, 0xb3, 0x1b, 0x76, 0x37, 0x95, 0x1e, 0xf5, 0x2f, 0x97, 0x6c, 0x92, 0x92, 0xea, 0x6d, 0x7e,
	0xb3, 0xb3, 0x33, 0xdf, 0xcc, 0x07, 0x90, 0x6b, 0x52, 0xb3, 0x4c, 0x49, 0x23, 0xb1, 0x55, 0xc4,
	0x93, 0x35, 0xf4, 0x9e, 0xe5, 0x86, 0x8b, 0x17, 0xa6, 0x58, 0xaa, 0x31, 0x84, 0x36, 0x8b, 0x63,
	0x99, 0x0b, 0x13, 0x3a, 0x63, 0x67, 0xda, 0x8d, 0x6a, 0xc4, 0x11, 0x74, 0x32, 0xa6, 0xf5, 0xa7,
	0x54, 0x49, 0xe8, 0xda, 0xa7, 0x23, 0x23, 0x42, 0x2b, 0x96, 0x09, 0x85, 0x9e, 0xcd, 0xdb, 0x78,
	0xf2, 0xed, 0x40, 0xe7, 0x55, 0x93, 0x5a, 0x88, 0x0f, 0x89, 0x7d, 0x70, 0x79, 0x52, 0x75, 0x74,
	0x79, 0xd2, 0x1c, 0xe3, 0xfe, 0x19, 0x23, 0x78, 0xbc, 0x15, 0x2c, 0xad, 0xdb, 0x1d, 0xf9, 0x44,
	0x42, 0xeb, 0x97, 0x84, 0x33, 0xf0, 0x29, 0x65, 0x7c, 0x17, 0xfa, 0xf6, 0xa1, 0x84, 0xc9, 0x97,
	0x03, 0x41, 0x44, 0x3a, 0xdf, 0x19, 0xbc, 0x80, 0x40, 0x1b, 0x66, 0x72, 0x6d, 0x65, 0x0c, 0xa3,
	0x8a, 0xf0, 0x1a, 0xbc, 0x94, 0x65, 0xa1, 0x3b, 0xf6, 0xa6, 0xbd, 0xf9, 0xf9, 0xcc, 0x1e, 0xa8,
	0xfc, 0x32, 0x5b, 0xb2, 0xec, 0x51, 0x18, 0x75, 0x88, 0x8a, 0x8a, 0xd1, 0x1d, 0x74, 0xea, 0x04,
	0x0e, 0xc0, 0xdb, 0xd2, 0xa1, 0x5a, 0xa8, 0x08, 0x8b, 0xf9, 0x7b, 0xb6, 0xcb, 0xa9, 0xda, 0xa7,
	0x84, 0x07, 0xf7, 0xde, 0xb9, 0xb9, 0x84, 0xe0, 0x89, 0x44, 0x42, 0x0a, 0xbb, 0xe0, 0xaf, 0x65,
	0x4a, 0x62, 0xf0, 0x0f, 0xdb, 0xe0, 0x2d, 0x99, 0x18, 0x38, 0xf3, 0x37, 0xe8, 0x15, 0x57, 0x5a,
	0x91, 0xda, 0xf3, 0x98, 0x70, 0x0a, 0xbe, 0xb5, 0x03, 0x87, 0xa5, 0x92, 0x86, 0x37, 0xa3, 0xff,
	0x4d, 0x71, 0x78, 0x05, 0xc1, 0x8a, 0x6f, 0xc4, 0x42, 0x60, 0xbf, 0xcc, 0xd7, 0xc7, 0x3e, 0xad,
	0x7b, 0x0f, 0xac, 0xdb, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x65, 0x79, 0x5c, 0xfb,
	0x01, 0x00, 0x00,
}
