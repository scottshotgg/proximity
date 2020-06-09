// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sender.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SendReq struct {
	Msg                  *Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendReq) Reset()         { *m = SendReq{} }
func (m *SendReq) String() string { return proto.CompactTextString(m) }
func (*SendReq) ProtoMessage()    {}
func (*SendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec65ed00f2bd807, []int{0}
}

func (m *SendReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendReq.Unmarshal(m, b)
}
func (m *SendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendReq.Marshal(b, m, deterministic)
}
func (m *SendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendReq.Merge(m, src)
}
func (m *SendReq) XXX_Size() int {
	return xxx_messageInfo_SendReq.Size(m)
}
func (m *SendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendReq proto.InternalMessageInfo

func (m *SendReq) GetMsg() *Message {
	if m != nil {
		return m.Msg
	}
	return nil
}

type SendRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRes) Reset()         { *m = SendRes{} }
func (m *SendRes) String() string { return proto.CompactTextString(m) }
func (*SendRes) ProtoMessage()    {}
func (*SendRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec65ed00f2bd807, []int{1}
}

func (m *SendRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRes.Unmarshal(m, b)
}
func (m *SendRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRes.Marshal(b, m, deterministic)
}
func (m *SendRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRes.Merge(m, src)
}
func (m *SendRes) XXX_Size() int {
	return xxx_messageInfo_SendRes.Size(m)
}
func (m *SendRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRes.DiscardUnknown(m)
}

var xxx_messageInfo_SendRes proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SendReq)(nil), "proto.SendReq")
	proto.RegisterType((*SendRes)(nil), "proto.SendRes")
}

func init() {
	proto.RegisterFile("sender.proto", fileDescriptor_dec65ed00f2bd807)
}

var fileDescriptor_dec65ed00f2bd807 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xcd, 0x4b,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x7c, 0x39, 0x99,
	0xc5, 0x25, 0xa9, 0x79, 0x30, 0x61, 0x29, 0xae, 0xa2, 0xd4, 0xe4, 0x32, 0x08, 0x5b, 0x49, 0x9b,
	0x8b, 0x3d, 0x38, 0x35, 0x2f, 0x25, 0x28, 0xb5, 0x50, 0x48, 0x81, 0x8b, 0x39, 0xb7, 0x38, 0x5d,
	0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x0f, 0x22, 0xaf, 0xe7, 0x9b, 0x5a, 0x5c, 0x9c, 0x98,
	0x9e, 0x1a, 0x04, 0x92, 0x52, 0xe2, 0x84, 0x29, 0x2e, 0x36, 0xea, 0x64, 0xe4, 0x62, 0x0b, 0x06,
	0xdb, 0x25, 0xa4, 0xc1, 0xc5, 0xe2, 0x5f, 0x90, 0x9a, 0x27, 0x04, 0xd3, 0x02, 0xe2, 0x04, 0xa5,
	0x16, 0x4a, 0xa1, 0xf2, 0x8b, 0x95, 0x18, 0x84, 0xb4, 0xb9, 0x58, 0x9d, 0x73, 0xf2, 0x8b, 0x53,
	0x85, 0xf8, 0xa1, 0x52, 0x60, 0x1e, 0x48, 0x2d, 0x9a, 0x00, 0x48, 0xb1, 0x06, 0x17, 0x0b, 0xc8,
	0x02, 0xb8, 0xb1, 0x50, 0x67, 0x4a, 0xa1, 0xf2, 0x8b, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x02, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xa7, 0x04, 0xb9, 0xfd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SenderClient is the client API for Sender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SenderClient interface {
	Open(ctx context.Context, in *OpenReq, opts ...grpc.CallOption) (*OpenRes, error)
	Close(ctx context.Context, in *CloseReq, opts ...grpc.CallOption) (*CloseRes, error)
	Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendRes, error)
}

type senderClient struct {
	cc grpc.ClientConnInterface
}

func NewSenderClient(cc grpc.ClientConnInterface) SenderClient {
	return &senderClient{cc}
}

func (c *senderClient) Open(ctx context.Context, in *OpenReq, opts ...grpc.CallOption) (*OpenRes, error) {
	out := new(OpenRes)
	err := c.cc.Invoke(ctx, "/proto.Sender/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *senderClient) Close(ctx context.Context, in *CloseReq, opts ...grpc.CallOption) (*CloseRes, error) {
	out := new(CloseRes)
	err := c.cc.Invoke(ctx, "/proto.Sender/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *senderClient) Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendRes, error) {
	out := new(SendRes)
	err := c.cc.Invoke(ctx, "/proto.Sender/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SenderServer is the server API for Sender service.
type SenderServer interface {
	Open(context.Context, *OpenReq) (*OpenRes, error)
	Close(context.Context, *CloseReq) (*CloseRes, error)
	Send(context.Context, *SendReq) (*SendRes, error)
}

// UnimplementedSenderServer can be embedded to have forward compatible implementations.
type UnimplementedSenderServer struct {
}

func (*UnimplementedSenderServer) Open(ctx context.Context, req *OpenReq) (*OpenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (*UnimplementedSenderServer) Close(ctx context.Context, req *CloseReq) (*CloseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (*UnimplementedSenderServer) Send(ctx context.Context, req *SendReq) (*SendRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterSenderServer(s *grpc.Server, srv SenderServer) {
	s.RegisterService(&_Sender_serviceDesc, srv)
}

func _Sender_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sender/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).Open(ctx, req.(*OpenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sender_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sender/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).Close(ctx, req.(*CloseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sender_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sender/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).Send(ctx, req.(*SendReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sender_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Sender",
	HandlerType: (*SenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _Sender_Open_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Sender_Close_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Sender_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sender.proto",
}
