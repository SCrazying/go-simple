// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package myexample

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	Buffer               string   `protobuf:"bytes,1,opt,name=buffer,proto3" json:"buffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetBuffer() string {
	if m != nil {
		return m.Buffer
	}
	return ""
}

type RouteRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=A,json=a,proto3" json:"A,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=B,json=b,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteRequest) Reset()         { *m = RouteRequest{} }
func (m *RouteRequest) String() string { return proto.CompactTextString(m) }
func (*RouteRequest) ProtoMessage()    {}
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}

func (m *RouteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteRequest.Unmarshal(m, b)
}
func (m *RouteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteRequest.Marshal(b, m, deterministic)
}
func (m *RouteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteRequest.Merge(m, src)
}
func (m *RouteRequest) XXX_Size() int {
	return xxx_messageInfo_RouteRequest.Size(m)
}
func (m *RouteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteRequest proto.InternalMessageInfo

func (m *RouteRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *RouteRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type RouteReply struct {
	C                    int32    `protobuf:"varint,1,opt,name=C,json=c,proto3" json:"C,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteReply) Reset()         { *m = RouteReply{} }
func (m *RouteReply) String() string { return proto.CompactTextString(m) }
func (*RouteReply) ProtoMessage()    {}
func (*RouteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{3}
}

func (m *RouteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteReply.Unmarshal(m, b)
}
func (m *RouteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteReply.Marshal(b, m, deterministic)
}
func (m *RouteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteReply.Merge(m, src)
}
func (m *RouteReply) XXX_Size() int {
	return xxx_messageInfo_RouteReply.Size(m)
}
func (m *RouteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteReply.DiscardUnknown(m)
}

var xxx_messageInfo_RouteReply proto.InternalMessageInfo

func (m *RouteReply) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

type RouterStream struct {
	//这个是客户端来的数据
	Recv string `protobuf:"bytes,1,opt,name=recv,proto3" json:"recv,omitempty"`
	//这个用于发生到客户端
	Send                 string   `protobuf:"bytes,2,opt,name=send,proto3" json:"send,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouterStream) Reset()         { *m = RouterStream{} }
func (m *RouterStream) String() string { return proto.CompactTextString(m) }
func (*RouterStream) ProtoMessage()    {}
func (*RouterStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{4}
}

func (m *RouterStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouterStream.Unmarshal(m, b)
}
func (m *RouterStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouterStream.Marshal(b, m, deterministic)
}
func (m *RouterStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouterStream.Merge(m, src)
}
func (m *RouterStream) XXX_Size() int {
	return xxx_messageInfo_RouterStream.Size(m)
}
func (m *RouterStream) XXX_DiscardUnknown() {
	xxx_messageInfo_RouterStream.DiscardUnknown(m)
}

var xxx_messageInfo_RouterStream proto.InternalMessageInfo

func (m *RouterStream) GetRecv() string {
	if m != nil {
		return m.Recv
	}
	return ""
}

func (m *RouterStream) GetSend() string {
	if m != nil {
		return m.Send
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "myexample.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "myexample.LoginReply")
	proto.RegisterType((*RouteRequest)(nil), "myexample.RouteRequest")
	proto.RegisterType((*RouteReply)(nil), "myexample.RouteReply")
	proto.RegisterType((*RouterStream)(nil), "myexample.RouterStream")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xbb, 0xd5, 0x46, 0x33, 0x44, 0x84, 0x05, 0xb5, 0xe4, 0x24, 0x8b, 0x87, 0xe2, 0x21,
	0x94, 0x08, 0x7a, 0xf0, 0x20, 0x6a, 0x11, 0x04, 0x4f, 0xe9, 0x13, 0x6c, 0xe2, 0x54, 0x84, 0x24,
	0xbb, 0xee, 0x26, 0x6a, 0x5e, 0xc2, 0x67, 0x96, 0xfd, 0x93, 0xd8, 0x22, 0x3d, 0xf4, 0x36, 0xdf,
	0x7e, 0xb3, 0xbf, 0xf9, 0x76, 0x58, 0x38, 0xc2, 0x6f, 0x5e, 0xc9, 0x12, 0x13, 0xa9, 0x44, 0x23,
	0x68, 0x58, 0x75, 0xfe, 0x80, 0x3d, 0x41, 0xf4, 0x22, 0xde, 0xde, 0xeb, 0x0c, 0x3f, 0x5a, 0xd4,
	0x0d, 0x8d, 0xe1, 0xb0, 0xd5, 0xa8, 0x6a, 0x5e, 0xe1, 0x94, 0x9c, 0x93, 0x59, 0x98, 0x0d, 0xda,
	0x78, 0x92, 0x6b, 0xfd, 0x25, 0xd4, 0xeb, 0x74, 0xec, 0xbc, 0x5e, 0xb3, 0x0b, 0x00, 0xcf, 0x91,
	0x65, 0x47, 0x4f, 0x21, 0xc8, 0xdb, 0xd5, 0x0a, 0x95, 0x67, 0x78, 0xc5, 0x2e, 0x21, 0xca, 0x44,
	0xdb, 0x60, 0x3f, 0x2d, 0x02, 0x72, 0x6f, 0x5b, 0x26, 0x19, 0xe1, 0x46, 0x3d, 0x58, 0xf0, 0x24,
	0x23, 0x39, 0x8b, 0x01, 0x7c, 0xaf, 0x21, 0x46, 0x40, 0x1e, 0xfb, 0xce, 0x82, 0x5d, 0x7b, 0x8e,
	0x5a, 0x36, 0x0a, 0x79, 0x45, 0x29, 0xec, 0x2b, 0x2c, 0x3e, 0xfd, 0x34, 0x5b, 0x9b, 0x33, 0x8d,
	0x75, 0x9f, 0xd4, 0xd6, 0xe9, 0x02, 0x26, 0x36, 0x25, 0xbd, 0x85, 0x83, 0x85, 0x70, 0xe5, 0x59,
	0x32, 0x6c, 0x23, 0x59, 0x5f, 0x45, 0x7c, 0xf2, 0xdf, 0x90, 0x65, 0xc7, 0x46, 0xe9, 0xcf, 0x18,
	0x02, 0x37, 0x9e, 0xde, 0xc0, 0x5e, 0x91, 0xea, 0x0d, 0xc6, 0xfa, 0x03, 0x37, 0x18, 0x7f, 0xaf,
	0x61, 0x23, 0x7a, 0x07, 0x61, 0x91, 0x6a, 0xed, 0xe2, 0xef, 0x7c, 0x7d, 0x4e, 0x2c, 0xc0, 0x5d,
	0x4f, 0x97, 0xbb, 0x03, 0x66, 0x84, 0x3e, 0xc3, 0xf1, 0x00, 0xd8, 0x96, 0xc3, 0xef, 0x37, 0xde,
	0x66, 0x18, 0xd0, 0x9c, 0xe4, 0x81, 0xfd, 0x56, 0x57, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x53,
	0x1a, 0x61, 0x2f, 0x67, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginClient interface {
	DoLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type loginClient struct {
	cc *grpc.ClientConn
}

func NewLoginClient(cc *grpc.ClientConn) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) DoLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/myexample.Login/DoLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
type LoginServer interface {
	DoLogin(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_DoLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).DoLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myexample.Login/DoLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).DoLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myexample.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoLogin",
			Handler:    _Login_DoLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	C2S(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteReply, error)
	C2Sstream(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (Router_C2SstreamClient, error)
	Cstream2S(ctx context.Context, opts ...grpc.CallOption) (Router_Cstream2SClient, error)
	Cstream2Sstream(ctx context.Context, opts ...grpc.CallOption) (Router_Cstream2SstreamClient, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) C2S(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteReply, error) {
	out := new(RouteReply)
	err := c.cc.Invoke(ctx, "/myexample.Router/c2s", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) C2Sstream(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (Router_C2SstreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Router_serviceDesc.Streams[0], "/myexample.Router/c2sstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &routerC2SstreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Router_C2SstreamClient interface {
	Recv() (*RouteReply, error)
	grpc.ClientStream
}

type routerC2SstreamClient struct {
	grpc.ClientStream
}

func (x *routerC2SstreamClient) Recv() (*RouteReply, error) {
	m := new(RouteReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routerClient) Cstream2S(ctx context.Context, opts ...grpc.CallOption) (Router_Cstream2SClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Router_serviceDesc.Streams[1], "/myexample.Router/cstream2S", opts...)
	if err != nil {
		return nil, err
	}
	x := &routerCstream2SClient{stream}
	return x, nil
}

type Router_Cstream2SClient interface {
	Send(*RouteRequest) error
	CloseAndRecv() (*RouteReply, error)
	grpc.ClientStream
}

type routerCstream2SClient struct {
	grpc.ClientStream
}

func (x *routerCstream2SClient) Send(m *RouteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routerCstream2SClient) CloseAndRecv() (*RouteReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routerClient) Cstream2Sstream(ctx context.Context, opts ...grpc.CallOption) (Router_Cstream2SstreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Router_serviceDesc.Streams[2], "/myexample.Router/cstream2Sstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &routerCstream2SstreamClient{stream}
	return x, nil
}

type Router_Cstream2SstreamClient interface {
	Send(*RouterStream) error
	Recv() (*RouterStream, error)
	grpc.ClientStream
}

type routerCstream2SstreamClient struct {
	grpc.ClientStream
}

func (x *routerCstream2SstreamClient) Send(m *RouterStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routerCstream2SstreamClient) Recv() (*RouterStream, error) {
	m := new(RouterStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	C2S(context.Context, *RouteRequest) (*RouteReply, error)
	C2Sstream(*RouteRequest, Router_C2SstreamServer) error
	Cstream2S(Router_Cstream2SServer) error
	Cstream2Sstream(Router_Cstream2SstreamServer) error
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_C2S_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).C2S(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myexample.Router/C2S",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).C2S(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_C2Sstream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RouteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouterServer).C2Sstream(m, &routerC2SstreamServer{stream})
}

type Router_C2SstreamServer interface {
	Send(*RouteReply) error
	grpc.ServerStream
}

type routerC2SstreamServer struct {
	grpc.ServerStream
}

func (x *routerC2SstreamServer) Send(m *RouteReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Router_Cstream2S_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouterServer).Cstream2S(&routerCstream2SServer{stream})
}

type Router_Cstream2SServer interface {
	SendAndClose(*RouteReply) error
	Recv() (*RouteRequest, error)
	grpc.ServerStream
}

type routerCstream2SServer struct {
	grpc.ServerStream
}

func (x *routerCstream2SServer) SendAndClose(m *RouteReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routerCstream2SServer) Recv() (*RouteRequest, error) {
	m := new(RouteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Router_Cstream2Sstream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouterServer).Cstream2Sstream(&routerCstream2SstreamServer{stream})
}

type Router_Cstream2SstreamServer interface {
	Send(*RouterStream) error
	Recv() (*RouterStream, error)
	grpc.ServerStream
}

type routerCstream2SstreamServer struct {
	grpc.ServerStream
}

func (x *routerCstream2SstreamServer) Send(m *RouterStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routerCstream2SstreamServer) Recv() (*RouterStream, error) {
	m := new(RouterStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myexample.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "c2s",
			Handler:    _Router_C2S_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "c2sstream",
			Handler:       _Router_C2Sstream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "cstream2S",
			Handler:       _Router_Cstream2S_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "cstream2Sstream",
			Handler:       _Router_Cstream2Sstream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "example.proto",
}
