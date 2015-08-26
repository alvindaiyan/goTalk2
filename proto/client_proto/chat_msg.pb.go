// Code generated by protoc-gen-go.
// source: chat_msg.proto
// DO NOT EDIT!

/*
Package chat_msg is a generated protocol buffer package.

It is generated from these files:
	chat_msg.proto

It has these top-level messages:
	Msg
*/
package chat_msg

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// only support the string for now
type Msg struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Chat service

type ChatClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (Chat_ChatClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Chat_ChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Chat_serviceDesc.Streams[0], c.cc, "/chat_msg.Chat/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatChatClient{stream}
	return x, nil
}

type Chat_ChatClient interface {
	Send(*Msg) error
	Recv() (*Msg, error)
	grpc.ClientStream
}

type chatChatClient struct {
	grpc.ClientStream
}

func (x *chatChatClient) Send(m *Msg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatChatClient) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Chat service

type ChatServer interface {
	Chat(Chat_ChatServer) error
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Chat(&chatChatServer{stream})
}

type Chat_ChatServer interface {
	Send(*Msg) error
	Recv() (*Msg, error)
	grpc.ServerStream
}

type chatChatServer struct {
	grpc.ServerStream
}

func (x *chatChatServer) Send(m *Msg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatChatServer) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat_msg.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _Chat_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}
