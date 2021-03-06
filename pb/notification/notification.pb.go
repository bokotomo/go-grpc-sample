// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

package notification

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

type NotificationRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationRequest) Reset()         { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string { return proto.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()    {}
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

func (m *NotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationRequest.Unmarshal(m, b)
}
func (m *NotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationRequest.Marshal(b, m, deterministic)
}
func (m *NotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationRequest.Merge(m, src)
}
func (m *NotificationRequest) XXX_Size() int {
	return xxx_messageInfo_NotificationRequest.Size(m)
}
func (m *NotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationRequest proto.InternalMessageInfo

func (m *NotificationRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type NotificationReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationReply) Reset()         { *m = NotificationReply{} }
func (m *NotificationReply) String() string { return proto.CompactTextString(m) }
func (*NotificationReply) ProtoMessage()    {}
func (*NotificationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{1}
}

func (m *NotificationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationReply.Unmarshal(m, b)
}
func (m *NotificationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationReply.Marshal(b, m, deterministic)
}
func (m *NotificationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationReply.Merge(m, src)
}
func (m *NotificationReply) XXX_Size() int {
	return xxx_messageInfo_NotificationReply.Size(m)
}
func (m *NotificationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationReply.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationReply proto.InternalMessageInfo

func (m *NotificationReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*NotificationRequest)(nil), "notification.NotificationRequest")
	proto.RegisterType((*NotificationReply)(nil), "notification.NotificationReply")
}

func init() { proto.RegisterFile("notification.proto", fileDescriptor_736a457d4a5efa07) }

var fileDescriptor_736a457d4a5efa07 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcb, 0x2f, 0xc9,
	0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x53, 0x52, 0xe7, 0x12, 0xf6, 0x43, 0xe2, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x09, 0x70, 0x31, 0xe7, 0x95, 0xe6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0x98,
	0x4a, 0xba, 0x5c, 0x82, 0xa8, 0x0a, 0x0b, 0x72, 0x2a, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b,
	0x8b, 0x13, 0xd3, 0x53, 0xc1, 0x4a, 0x39, 0x83, 0x60, 0x5c, 0xa3, 0x34, 0x2e, 0x1e, 0x64, 0xe5,
	0x42, 0x61, 0x68, 0x7c, 0x45, 0x3d, 0x14, 0xa7, 0x61, 0x71, 0x83, 0x94, 0x3c, 0x3e, 0x25, 0x05,
	0x39, 0x95, 0x4a, 0x0c, 0x06, 0x8c, 0x4e, 0x26, 0x5c, 0xb2, 0x99, 0xf9, 0x7a, 0xe9, 0x45, 0x05,
	0xc9, 0x7a, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0xc5, 0x28, 0xda, 0x9c, 0x50, 0x5c, 0x1d,
	0x00, 0x0a, 0x81, 0x00, 0xc6, 0x24, 0x36, 0x70, 0x50, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0x89, 0x5c, 0x6f, 0x20, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationClient interface {
	Notification(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (Notification_NotificationClient, error)
}

type notificationClient struct {
	cc *grpc.ClientConn
}

func NewNotificationClient(cc *grpc.ClientConn) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) Notification(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (Notification_NotificationClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Notification_serviceDesc.Streams[0], "/notification.Notification/Notification", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationNotificationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Notification_NotificationClient interface {
	Recv() (*NotificationReply, error)
	grpc.ClientStream
}

type notificationNotificationClient struct {
	grpc.ClientStream
}

func (x *notificationNotificationClient) Recv() (*NotificationReply, error) {
	m := new(NotificationReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationServer is the server API for Notification service.
type NotificationServer interface {
	Notification(*NotificationRequest, Notification_NotificationServer) error
}

// UnimplementedNotificationServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (*UnimplementedNotificationServer) Notification(req *NotificationRequest, srv Notification_NotificationServer) error {
	return status.Errorf(codes.Unimplemented, "method Notification not implemented")
}

func RegisterNotificationServer(s *grpc.Server, srv NotificationServer) {
	s.RegisterService(&_Notification_serviceDesc, srv)
}

func _Notification_Notification_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServer).Notification(m, &notificationNotificationServer{stream})
}

type Notification_NotificationServer interface {
	Send(*NotificationReply) error
	grpc.ServerStream
}

type notificationNotificationServer struct {
	grpc.ServerStream
}

func (x *notificationNotificationServer) Send(m *NotificationReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Notification_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notification.Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Notification",
			Handler:       _Notification_Notification_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notification.proto",
}
