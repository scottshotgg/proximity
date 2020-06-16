// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package buffs

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
	return fileDescriptor_0c843d59d2d938e7, []int{0}
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
	return fileDescriptor_0c843d59d2d938e7, []int{1}
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
	return fileDescriptor_0c843d59d2d938e7, []int{2}
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

type AttachReq struct {
	// id is a requested ID only; no gaurantee
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// route is a requested route; the server can kick it back
	Route                string   `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachReq) Reset()         { *m = AttachReq{} }
func (m *AttachReq) String() string { return proto.CompactTextString(m) }
func (*AttachReq) ProtoMessage()    {}
func (*AttachReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{3}
}

func (m *AttachReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachReq.Unmarshal(m, b)
}
func (m *AttachReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachReq.Marshal(b, m, deterministic)
}
func (m *AttachReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachReq.Merge(m, src)
}
func (m *AttachReq) XXX_Size() int {
	return xxx_messageInfo_AttachReq.Size(m)
}
func (m *AttachReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachReq.DiscardUnknown(m)
}

var xxx_messageInfo_AttachReq proto.InternalMessageInfo

func (m *AttachReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AttachReq) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

type AttachRes struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachRes) Reset()         { *m = AttachRes{} }
func (m *AttachRes) String() string { return proto.CompactTextString(m) }
func (*AttachRes) ProtoMessage()    {}
func (*AttachRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{4}
}

func (m *AttachRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachRes.Unmarshal(m, b)
}
func (m *AttachRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachRes.Marshal(b, m, deterministic)
}
func (m *AttachRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachRes.Merge(m, src)
}
func (m *AttachRes) XXX_Size() int {
	return xxx_messageInfo_AttachRes.Size(m)
}
func (m *AttachRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachRes.DiscardUnknown(m)
}

var xxx_messageInfo_AttachRes proto.InternalMessageInfo

func (m *AttachRes) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type DiscoverReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoverReq) Reset()         { *m = DiscoverReq{} }
func (m *DiscoverReq) String() string { return proto.CompactTextString(m) }
func (*DiscoverReq) ProtoMessage()    {}
func (*DiscoverReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{5}
}

func (m *DiscoverReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverReq.Unmarshal(m, b)
}
func (m *DiscoverReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverReq.Marshal(b, m, deterministic)
}
func (m *DiscoverReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverReq.Merge(m, src)
}
func (m *DiscoverReq) XXX_Size() int {
	return xxx_messageInfo_DiscoverReq.Size(m)
}
func (m *DiscoverReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverReq.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverReq proto.InternalMessageInfo

type DiscoverRes struct {
	// TODO: this may need to be changed to clients, listeners, and nodes or something; a map of network maps? idk if getting too complicated
	Nodes                []string `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoverRes) Reset()         { *m = DiscoverRes{} }
func (m *DiscoverRes) String() string { return proto.CompactTextString(m) }
func (*DiscoverRes) ProtoMessage()    {}
func (*DiscoverRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{6}
}

func (m *DiscoverRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverRes.Unmarshal(m, b)
}
func (m *DiscoverRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverRes.Marshal(b, m, deterministic)
}
func (m *DiscoverRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverRes.Merge(m, src)
}
func (m *DiscoverRes) XXX_Size() int {
	return xxx_messageInfo_DiscoverRes.Size(m)
}
func (m *DiscoverRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverRes.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverRes proto.InternalMessageInfo

func (m *DiscoverRes) GetNodes() []string {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "buffs.Message")
	proto.RegisterType((*SendReq)(nil), "buffs.SendReq")
	proto.RegisterType((*SendRes)(nil), "buffs.SendRes")
	proto.RegisterType((*AttachReq)(nil), "buffs.AttachReq")
	proto.RegisterType((*AttachRes)(nil), "buffs.AttachRes")
	proto.RegisterType((*DiscoverReq)(nil), "buffs.DiscoverReq")
	proto.RegisterType((*DiscoverRes)(nil), "buffs.DiscoverRes")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xe5, 0xf4, 0x4f, 0x9a, 0xab, 0xa8, 0xd0, 0x89, 0x21, 0xca, 0x14, 0x99, 0x25, 0x12,
	0x52, 0x04, 0x41, 0x4c, 0x4c, 0x48, 0xac, 0x30, 0x84, 0x4f, 0xd0, 0xc6, 0xd7, 0x92, 0xa1, 0x31,
	0xcd, 0xb9, 0x7c, 0x12, 0x3e, 0x30, 0xb2, 0x63, 0xb7, 0x01, 0xc4, 0xf8, 0xee, 0xe5, 0xde, 0xfb,
	0x5d, 0x0c, 0xd0, 0x69, 0x45, 0xe5, 0x47, 0xaf, 0x8d, 0xc6, 0xd9, 0xe6, 0xb8, 0xdd, 0xb2, 0x7c,
	0x84, 0xf8, 0x85, 0x98, 0xd7, 0x3b, 0xc2, 0x2b, 0x98, 0xf5, 0xfa, 0x68, 0x28, 0x15, 0xb9, 0x28,
	0x92, 0x7a, 0x10, 0x98, 0xc1, 0xa2, 0xd1, 0x9d, 0xa1, 0xce, 0x70, 0x1a, 0x39, 0xe3, 0xa4, 0xe5,
	0x0d, 0xc4, 0x6f, 0xd4, 0xa9, 0x9a, 0x0e, 0x98, 0xc3, 0x64, 0xcf, 0x3b, 0xb7, 0xba, 0xac, 0x56,
	0xa5, 0x0b, 0x2f, 0x7d, 0x72, 0x6d, 0x2d, 0x99, 0x84, 0x8f, 0x59, 0xde, 0x41, 0xf2, 0x64, 0xcc,
	0xba, 0x79, 0xb7, 0x9b, 0x2b, 0x88, 0x5a, 0xe5, 0x3b, 0xa3, 0x56, 0x9d, 0x31, 0xa2, 0x11, 0x86,
	0x7c, 0x38, 0xaf, 0x30, 0x16, 0x10, 0xef, 0x87, 0xe8, 0x7f, 0x0a, 0x83, 0x2d, 0x2f, 0x60, 0xf9,
	0xdc, 0x72, 0xa3, 0x3f, 0xa9, 0xaf, 0xe9, 0x20, 0xaf, 0xc7, 0x92, 0x6d, 0x95, 0xfd, 0x23, 0x9c,
	0x8a, 0x7c, 0x62, 0xab, 0x9c, 0xa8, 0xbe, 0x04, 0x4c, 0x5f, 0xb5, 0x22, 0x2c, 0x60, 0x6a, 0x89,
	0x31, 0xa4, 0xfb, 0x5b, 0xb3, 0x9f, 0x9a, 0x0b, 0x81, 0x25, 0xcc, 0x07, 0x3a, 0xbc, 0xf4, 0xde,
	0xe9, 0xbe, 0xec, 0xf7, 0x84, 0x6f, 0x05, 0x56, 0xb0, 0x08, 0x1c, 0x88, 0xde, 0x1f, 0x71, 0x66,
	0x7f, 0x67, 0xbc, 0x99, 0xbb, 0x77, 0xbb, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xf6, 0x0b,
	0xf8, 0xc5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	Send(ctx context.Context, opts ...grpc.CallOption) (Node_SendClient, error)
	Attach(ctx context.Context, in *AttachReq, opts ...grpc.CallOption) (Node_AttachClient, error)
	Discover(ctx context.Context, in *DiscoverReq, opts ...grpc.CallOption) (*DiscoverRes, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Send(ctx context.Context, opts ...grpc.CallOption) (Node_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Node_serviceDesc.Streams[0], "/buffs.Node/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeSendClient{stream}
	return x, nil
}

type Node_SendClient interface {
	Send(*SendReq) error
	CloseAndRecv() (*SendRes, error)
	grpc.ClientStream
}

type nodeSendClient struct {
	grpc.ClientStream
}

func (x *nodeSendClient) Send(m *SendReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeSendClient) CloseAndRecv() (*SendRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) Attach(ctx context.Context, in *AttachReq, opts ...grpc.CallOption) (Node_AttachClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Node_serviceDesc.Streams[1], "/buffs.Node/Attach", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAttachClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Node_AttachClient interface {
	Recv() (*AttachRes, error)
	grpc.ClientStream
}

type nodeAttachClient struct {
	grpc.ClientStream
}

func (x *nodeAttachClient) Recv() (*AttachRes, error) {
	m := new(AttachRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) Discover(ctx context.Context, in *DiscoverReq, opts ...grpc.CallOption) (*DiscoverRes, error) {
	out := new(DiscoverRes)
	err := c.cc.Invoke(ctx, "/buffs.Node/Discover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	Send(Node_SendServer) error
	Attach(*AttachReq, Node_AttachServer) error
	Discover(context.Context, *DiscoverReq) (*DiscoverRes, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) Send(srv Node_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedNodeServer) Attach(req *AttachReq, srv Node_AttachServer) error {
	return status.Errorf(codes.Unimplemented, "method Attach not implemented")
}
func (*UnimplementedNodeServer) Discover(ctx context.Context, req *DiscoverReq) (*DiscoverRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Discover not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).Send(&nodeSendServer{stream})
}

type Node_SendServer interface {
	SendAndClose(*SendRes) error
	Recv() (*SendReq, error)
	grpc.ServerStream
}

type nodeSendServer struct {
	grpc.ServerStream
}

func (x *nodeSendServer) SendAndClose(m *SendRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeSendServer) Recv() (*SendReq, error) {
	m := new(SendReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_Attach_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AttachReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServer).Attach(m, &nodeAttachServer{stream})
}

type Node_AttachServer interface {
	Send(*AttachRes) error
	grpc.ServerStream
}

type nodeAttachServer struct {
	grpc.ServerStream
}

func (x *nodeAttachServer) Send(m *AttachRes) error {
	return x.ServerStream.SendMsg(m)
}

func _Node_Discover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Discover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buffs.Node/Discover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Discover(ctx, req.(*DiscoverReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buffs.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Discover",
			Handler:    _Node_Discover_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _Node_Send_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Attach",
			Handler:       _Node_Attach_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node.proto",
}
