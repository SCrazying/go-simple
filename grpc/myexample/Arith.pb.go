// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Arith.proto

//定义包名

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

type ArithRequest struct {
	Width                int32    `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArithRequest) Reset()         { *m = ArithRequest{} }
func (m *ArithRequest) String() string { return proto.CompactTextString(m) }
func (*ArithRequest) ProtoMessage()    {}
func (*ArithRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89835e6438b55cb, []int{0}
}

func (m *ArithRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArithRequest.Unmarshal(m, b)
}
func (m *ArithRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArithRequest.Marshal(b, m, deterministic)
}
func (m *ArithRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithRequest.Merge(m, src)
}
func (m *ArithRequest) XXX_Size() int {
	return xxx_messageInfo_ArithRequest.Size(m)
}
func (m *ArithRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArithRequest proto.InternalMessageInfo

func (m *ArithRequest) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ArithRequest) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Arithreply struct {
	Circumference        int32    `protobuf:"varint,1,opt,name=circumference,proto3" json:"circumference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Arithreply) Reset()         { *m = Arithreply{} }
func (m *Arithreply) String() string { return proto.CompactTextString(m) }
func (*Arithreply) ProtoMessage()    {}
func (*Arithreply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89835e6438b55cb, []int{1}
}

func (m *Arithreply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Arithreply.Unmarshal(m, b)
}
func (m *Arithreply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Arithreply.Marshal(b, m, deterministic)
}
func (m *Arithreply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Arithreply.Merge(m, src)
}
func (m *Arithreply) XXX_Size() int {
	return xxx_messageInfo_Arithreply.Size(m)
}
func (m *Arithreply) XXX_DiscardUnknown() {
	xxx_messageInfo_Arithreply.DiscardUnknown(m)
}

var xxx_messageInfo_Arithreply proto.InternalMessageInfo

func (m *Arithreply) GetCircumference() int32 {
	if m != nil {
		return m.Circumference
	}
	return 0
}

type CountRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountRequest) Reset()         { *m = CountRequest{} }
func (m *CountRequest) String() string { return proto.CompactTextString(m) }
func (*CountRequest) ProtoMessage()    {}
func (*CountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89835e6438b55cb, []int{2}
}

func (m *CountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountRequest.Unmarshal(m, b)
}
func (m *CountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountRequest.Marshal(b, m, deterministic)
}
func (m *CountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountRequest.Merge(m, src)
}
func (m *CountRequest) XXX_Size() int {
	return xxx_messageInfo_CountRequest.Size(m)
}
func (m *CountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountRequest proto.InternalMessageInfo

func (m *CountRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type CountReply struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountReply) Reset()         { *m = CountReply{} }
func (m *CountReply) String() string { return proto.CompactTextString(m) }
func (*CountReply) ProtoMessage()    {}
func (*CountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89835e6438b55cb, []int{3}
}

func (m *CountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountReply.Unmarshal(m, b)
}
func (m *CountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountReply.Marshal(b, m, deterministic)
}
func (m *CountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountReply.Merge(m, src)
}
func (m *CountReply) XXX_Size() int {
	return xxx_messageInfo_CountReply.Size(m)
}
func (m *CountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CountReply.DiscardUnknown(m)
}

var xxx_messageInfo_CountReply proto.InternalMessageInfo

func (m *CountReply) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type FibonacciRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciRequest) Reset()         { *m = FibonacciRequest{} }
func (m *FibonacciRequest) String() string { return proto.CompactTextString(m) }
func (*FibonacciRequest) ProtoMessage()    {}
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89835e6438b55cb, []int{4}
}

func (m *FibonacciRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciRequest.Unmarshal(m, b)
}
func (m *FibonacciRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciRequest.Marshal(b, m, deterministic)
}
func (m *FibonacciRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciRequest.Merge(m, src)
}
func (m *FibonacciRequest) XXX_Size() int {
	return xxx_messageInfo_FibonacciRequest.Size(m)
}
func (m *FibonacciRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciRequest proto.InternalMessageInfo

func (m *FibonacciRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type FibonacciReply struct {
	Ans                  int32    `protobuf:"varint,1,opt,name=ans,proto3" json:"ans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibonacciReply) Reset()         { *m = FibonacciReply{} }
func (m *FibonacciReply) String() string { return proto.CompactTextString(m) }
func (*FibonacciReply) ProtoMessage()    {}
func (*FibonacciReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89835e6438b55cb, []int{5}
}

func (m *FibonacciReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibonacciReply.Unmarshal(m, b)
}
func (m *FibonacciReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibonacciReply.Marshal(b, m, deterministic)
}
func (m *FibonacciReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibonacciReply.Merge(m, src)
}
func (m *FibonacciReply) XXX_Size() int {
	return xxx_messageInfo_FibonacciReply.Size(m)
}
func (m *FibonacciReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FibonacciReply.DiscardUnknown(m)
}

var xxx_messageInfo_FibonacciReply proto.InternalMessageInfo

func (m *FibonacciReply) GetAns() int32 {
	if m != nil {
		return m.Ans
	}
	return 0
}

type ChatRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatRequest) Reset()         { *m = ChatRequest{} }
func (m *ChatRequest) String() string { return proto.CompactTextString(m) }
func (*ChatRequest) ProtoMessage()    {}
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89835e6438b55cb, []int{6}
}

func (m *ChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatRequest.Unmarshal(m, b)
}
func (m *ChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatRequest.Marshal(b, m, deterministic)
}
func (m *ChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatRequest.Merge(m, src)
}
func (m *ChatRequest) XXX_Size() int {
	return xxx_messageInfo_ChatRequest.Size(m)
}
func (m *ChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChatRequest proto.InternalMessageInfo

func (m *ChatRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ChatReply struct {
	Reply                string   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatReply) Reset()         { *m = ChatReply{} }
func (m *ChatReply) String() string { return proto.CompactTextString(m) }
func (*ChatReply) ProtoMessage()    {}
func (*ChatReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b89835e6438b55cb, []int{7}
}

func (m *ChatReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatReply.Unmarshal(m, b)
}
func (m *ChatReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatReply.Marshal(b, m, deterministic)
}
func (m *ChatReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatReply.Merge(m, src)
}
func (m *ChatReply) XXX_Size() int {
	return xxx_messageInfo_ChatReply.Size(m)
}
func (m *ChatReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChatReply proto.InternalMessageInfo

func (m *ChatReply) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*ArithRequest)(nil), "myexample.ArithRequest")
	proto.RegisterType((*Arithreply)(nil), "myexample.Arithreply")
	proto.RegisterType((*CountRequest)(nil), "myexample.CountRequest")
	proto.RegisterType((*CountReply)(nil), "myexample.CountReply")
	proto.RegisterType((*FibonacciRequest)(nil), "myexample.FibonacciRequest")
	proto.RegisterType((*FibonacciReply)(nil), "myexample.FibonacciReply")
	proto.RegisterType((*ChatRequest)(nil), "myexample.ChatRequest")
	proto.RegisterType((*ChatReply)(nil), "myexample.ChatReply")
}

func init() { proto.RegisterFile("Arith.proto", fileDescriptor_b89835e6438b55cb) }

var fileDescriptor_b89835e6438b55cb = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x6b, 0xea, 0x40,
	0x14, 0xc5, 0xcd, 0xd3, 0xc8, 0xcb, 0xd1, 0x96, 0x30, 0x58, 0x6b, 0xed, 0x46, 0x07, 0x17, 0xae,
	0x44, 0xd2, 0xad, 0x14, 0x4a, 0xa4, 0x85, 0x42, 0x29, 0x48, 0xa1, 0xeb, 0x31, 0x4e, 0x9b, 0x40,
	0x32, 0xb1, 0x71, 0x42, 0xeb, 0xb7, 0x2f, 0xb9, 0x8e, 0x9a, 0xd8, 0xd2, 0xdd, 0xfd, 0x73, 0xe6,
	0xe4, 0xdc, 0x1f, 0x41, 0xeb, 0x2e, 0x8b, 0x74, 0x38, 0x59, 0x67, 0xa9, 0x4e, 0x99, 0x93, 0x6c,
	0xe5, 0x97, 0x48, 0xd6, 0xb1, 0xe4, 0x33, 0xb4, 0x69, 0xb3, 0x90, 0x1f, 0xb9, 0xdc, 0x68, 0xd6,
	0x81, 0xfd, 0x19, 0xad, 0x74, 0xd8, 0xb3, 0x06, 0xd6, 0xd8, 0x5e, 0xec, 0x1a, 0xd6, 0x45, 0x33,
	0x94, 0xd1, 0x7b, 0xa8, 0x7b, 0xff, 0x68, 0x6c, 0x3a, 0xee, 0x01, 0xf4, 0x3a, 0x93, 0xeb, 0x78,
	0xcb, 0x46, 0x38, 0x0b, 0xa2, 0x2c, 0xc8, 0x93, 0x37, 0x99, 0x49, 0x15, 0x48, 0xe3, 0x51, 0x1d,
	0xf2, 0x01, 0xda, 0x7e, 0x9a, 0x2b, 0xbd, 0xff, 0xa2, 0x8b, 0xba, 0xca, 0x13, 0xa3, 0x2d, 0x4a,
	0xce, 0x01, 0xa3, 0x28, 0x5c, 0x3b, 0xb0, 0x75, 0xaa, 0x45, 0xbc, 0x4f, 0x44, 0x0d, 0x1f, 0xc1,
	0xbd, 0x8f, 0x96, 0xa9, 0x12, 0x41, 0x10, 0xfd, 0xe5, 0x74, 0x5e, 0x52, 0x15, 0x6e, 0x2e, 0xea,
	0x42, 0x6d, 0xf6, 0x1a, 0xa1, 0x36, 0x7c, 0x88, 0x96, 0x1f, 0x8a, 0x43, 0x1c, 0x86, 0xc6, 0x4a,
	0x68, 0x41, 0x0a, 0x67, 0x41, 0x35, 0x1f, 0xc2, 0xd9, 0x49, 0x4c, 0x1e, 0x3a, 0xd7, 0x28, 0x76,
	0x8d, 0xf7, 0x04, 0x9b, 0x48, 0xb0, 0x39, 0x5c, 0x5f, 0xc4, 0x7e, 0xf9, 0x64, 0x76, 0x39, 0x39,
	0x00, 0x9f, 0x94, 0x69, 0xf7, 0x2f, 0x4e, 0x17, 0x64, 0xc6, 0x6b, 0xde, 0x03, 0x6c, 0x42, 0xc0,
	0x6e, 0xf1, 0xdf, 0x17, 0xf1, 0x4b, 0x71, 0x73, 0xc5, 0xa6, 0x8c, 0xb0, 0x62, 0x73, 0x24, 0xc7,
	0x6b, 0x63, 0xcb, 0x7b, 0x85, 0x73, 0x20, 0xc0, 0x1e, 0xd1, 0xf6, 0x45, 0x7c, 0xec, 0xaf, 0x4b,
	0xef, 0x4e, 0x69, 0xf6, 0xaf, 0x7e, 0x5f, 0x92, 0xf1, 0xd4, 0xf2, 0xe6, 0x68, 0x14, 0x4c, 0xd8,
	0x0c, 0xcd, 0x67, 0x45, 0x55, 0xb7, 0x9c, 0xe2, 0x48, 0xb4, 0xdf, 0xf9, 0x31, 0x37, 0xe1, 0xa6,
	0xd6, 0xb2, 0x49, 0x3f, 0xe4, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x23, 0xb2, 0xe8, 0x65,
	0x9f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArithClient is the client API for Arith service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArithClient interface {
	//rpc
	CalCircumference(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*Arithreply, error)
}

type arithClient struct {
	cc *grpc.ClientConn
}

func NewArithClient(cc *grpc.ClientConn) ArithClient {
	return &arithClient{cc}
}

func (c *arithClient) CalCircumference(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*Arithreply, error) {
	out := new(Arithreply)
	err := c.cc.Invoke(ctx, "/myexample.Arith/CalCircumference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithServer is the server API for Arith service.
type ArithServer interface {
	//rpc
	CalCircumference(context.Context, *ArithRequest) (*Arithreply, error)
}

func RegisterArithServer(s *grpc.Server, srv ArithServer) {
	s.RegisterService(&_Arith_serviceDesc, srv)
}

func _Arith_CalCircumference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithServer).CalCircumference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myexample.Arith/CalCircumference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithServer).CalCircumference(ctx, req.(*ArithRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Arith_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myexample.Arith",
	HandlerType: (*ArithServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalCircumference",
			Handler:    _Arith_CalCircumference_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Arith.proto",
}

// CountClient is the client API for Count service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountClient interface {
	//rpc
	CalTotal(ctx context.Context, opts ...grpc.CallOption) (Count_CalTotalClient, error)
}

type countClient struct {
	cc *grpc.ClientConn
}

func NewCountClient(cc *grpc.ClientConn) CountClient {
	return &countClient{cc}
}

func (c *countClient) CalTotal(ctx context.Context, opts ...grpc.CallOption) (Count_CalTotalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Count_serviceDesc.Streams[0], "/myexample.Count/CalTotal", opts...)
	if err != nil {
		return nil, err
	}
	x := &countCalTotalClient{stream}
	return x, nil
}

type Count_CalTotalClient interface {
	Send(*CountRequest) error
	CloseAndRecv() (*CountReply, error)
	grpc.ClientStream
}

type countCalTotalClient struct {
	grpc.ClientStream
}

func (x *countCalTotalClient) Send(m *CountRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *countCalTotalClient) CloseAndRecv() (*CountReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CountReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CountServer is the server API for Count service.
type CountServer interface {
	//rpc
	CalTotal(Count_CalTotalServer) error
}

func RegisterCountServer(s *grpc.Server, srv CountServer) {
	s.RegisterService(&_Count_serviceDesc, srv)
}

func _Count_CalTotal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CountServer).CalTotal(&countCalTotalServer{stream})
}

type Count_CalTotalServer interface {
	SendAndClose(*CountReply) error
	Recv() (*CountRequest, error)
	grpc.ServerStream
}

type countCalTotalServer struct {
	grpc.ServerStream
}

func (x *countCalTotalServer) SendAndClose(m *CountReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *countCalTotalServer) Recv() (*CountRequest, error) {
	m := new(CountRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Count_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myexample.Count",
	HandlerType: (*CountServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalTotal",
			Handler:       _Count_CalTotal_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Arith.proto",
}

// FibonacciClient is the client API for Fibonacci service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FibonacciClient interface {
	//rpc
	CalFibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (Fibonacci_CalFibonacciClient, error)
}

type fibonacciClient struct {
	cc *grpc.ClientConn
}

func NewFibonacciClient(cc *grpc.ClientConn) FibonacciClient {
	return &fibonacciClient{cc}
}

func (c *fibonacciClient) CalFibonacci(ctx context.Context, in *FibonacciRequest, opts ...grpc.CallOption) (Fibonacci_CalFibonacciClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fibonacci_serviceDesc.Streams[0], "/myexample.Fibonacci/CalFibonacci", opts...)
	if err != nil {
		return nil, err
	}
	x := &fibonacciCalFibonacciClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fibonacci_CalFibonacciClient interface {
	Recv() (*FibonacciReply, error)
	grpc.ClientStream
}

type fibonacciCalFibonacciClient struct {
	grpc.ClientStream
}

func (x *fibonacciCalFibonacciClient) Recv() (*FibonacciReply, error) {
	m := new(FibonacciReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FibonacciServer is the server API for Fibonacci service.
type FibonacciServer interface {
	//rpc
	CalFibonacci(*FibonacciRequest, Fibonacci_CalFibonacciServer) error
}

func RegisterFibonacciServer(s *grpc.Server, srv FibonacciServer) {
	s.RegisterService(&_Fibonacci_serviceDesc, srv)
}

func _Fibonacci_CalFibonacci_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FibonacciRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FibonacciServer).CalFibonacci(m, &fibonacciCalFibonacciServer{stream})
}

type Fibonacci_CalFibonacciServer interface {
	Send(*FibonacciReply) error
	grpc.ServerStream
}

type fibonacciCalFibonacciServer struct {
	grpc.ServerStream
}

func (x *fibonacciCalFibonacciServer) Send(m *FibonacciReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Fibonacci_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myexample.Fibonacci",
	HandlerType: (*FibonacciServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalFibonacci",
			Handler:       _Fibonacci_CalFibonacci_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "Arith.proto",
}

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	OnChat(ctx context.Context, opts ...grpc.CallOption) (Chat_OnChatClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) OnChat(ctx context.Context, opts ...grpc.CallOption) (Chat_OnChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/myexample.Chat/OnChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatOnChatClient{stream}
	return x, nil
}

type Chat_OnChatClient interface {
	Send(*ChatRequest) error
	Recv() (*ChatReply, error)
	grpc.ClientStream
}

type chatOnChatClient struct {
	grpc.ClientStream
}

func (x *chatOnChatClient) Send(m *ChatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatOnChatClient) Recv() (*ChatReply, error) {
	m := new(ChatReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	OnChat(Chat_OnChatServer) error
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_OnChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).OnChat(&chatOnChatServer{stream})
}

type Chat_OnChatServer interface {
	Send(*ChatReply) error
	Recv() (*ChatRequest, error)
	grpc.ServerStream
}

type chatOnChatServer struct {
	grpc.ServerStream
}

func (x *chatOnChatServer) Send(m *ChatReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatOnChatServer) Recv() (*ChatRequest, error) {
	m := new(ChatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myexample.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnChat",
			Handler:       _Chat_OnChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Arith.proto",
}