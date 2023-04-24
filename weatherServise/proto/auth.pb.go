// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserInfo struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

type AuthResp struct {
	Reply                bool     `protobuf:"varint,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResp) Reset()         { *m = AuthResp{} }
func (m *AuthResp) String() string { return proto.CompactTextString(m) }
func (*AuthResp) ProtoMessage()    {}
func (*AuthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AuthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResp.Unmarshal(m, b)
}
func (m *AuthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResp.Marshal(b, m, deterministic)
}
func (m *AuthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResp.Merge(m, src)
}
func (m *AuthResp) XXX_Size() int {
	return xxx_messageInfo_AuthResp.Size(m)
}
func (m *AuthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResp.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResp proto.InternalMessageInfo

func (m *AuthResp) GetReply() bool {
	if m != nil {
		return m.Reply
	}
	return false
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "UserInfo")
	proto.RegisterType((*AuthResp)(nil), "AuthResp")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe0, 0xe2, 0x08, 0x2d, 0x4e, 0x2d, 0xf2, 0xcc,
	0x4b, 0xcb, 0x17, 0x12, 0xe1, 0x62, 0xcd, 0xc9, 0x4f, 0xcf, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0x70, 0x40, 0x2a, 0x1c, 0x4b, 0x4b, 0x32, 0x82, 0x52, 0x8b, 0x0b, 0x40, 0x2a,
	0x8a, 0x52, 0x0b, 0x72, 0x2a, 0xc1, 0x2a, 0x38, 0x82, 0x20, 0x1c, 0x23, 0x35, 0x2e, 0x16, 0x90,
	0x0a, 0x21, 0x39, 0x2e, 0x36, 0xcf, 0x62, 0x30, 0x8b, 0x53, 0x0f, 0x66, 0xa8, 0x14, 0xa7, 0x1e,
	0x4c, 0xb7, 0x13, 0x7f, 0x14, 0x6f, 0x79, 0x6a, 0x62, 0x49, 0x46, 0x6a, 0x91, 0x3e, 0xd8, 0xf2,
	0x24, 0x36, 0x30, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x79, 0xf0, 0x04, 0xe0, 0x91, 0x00,
	0x00, 0x00,
}
