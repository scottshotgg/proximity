// Code generated by protoc-gen-go. DO NOT EDIT.
// source: listener.proto

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

type Message struct {
	Route                string   `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Contents             string   `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f75aade3a9f7de9c, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *Message) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "proto.Message")
}

func init() {
	proto.RegisterFile("listener.proto", fileDescriptor_f75aade3a9f7de9c)
}

var fileDescriptor_f75aade3a9f7de9c = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xc9, 0x2c, 0x2e,
	0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xd6,
	0x5c, 0xec, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x22, 0x5c, 0xac, 0x45, 0xf9, 0xa5,
	0x25, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x14, 0x17, 0x47, 0x72,
	0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x49, 0xb1, 0x04, 0x13, 0x58, 0x02, 0xce, 0x4f, 0x62, 0x03, 0x9b,
	0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x27, 0x16, 0x27, 0x83, 0x5c, 0x00, 0x00, 0x00,
}
