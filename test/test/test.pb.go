// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

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

type TestRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_aab672adb641c8c6, []int{0}
}
func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (dst *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(dst, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TestReply struct {
	TestReply            string   `protobuf:"bytes,1,opt,name=testReply,proto3" json:"testReply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestReply) Reset()         { *m = TestReply{} }
func (m *TestReply) String() string { return proto.CompactTextString(m) }
func (*TestReply) ProtoMessage()    {}
func (*TestReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_aab672adb641c8c6, []int{1}
}
func (m *TestReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestReply.Unmarshal(m, b)
}
func (m *TestReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestReply.Marshal(b, m, deterministic)
}
func (dst *TestReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestReply.Merge(dst, src)
}
func (m *TestReply) XXX_Size() int {
	return xxx_messageInfo_TestReply.Size(m)
}
func (m *TestReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TestReply.DiscardUnknown(m)
}

var xxx_messageInfo_TestReply proto.InternalMessageInfo

func (m *TestReply) GetTestReply() string {
	if m != nil {
		return m.TestReply
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "test.TestRequest")
	proto.RegisterType((*TestReply)(nil), "test.TestReply")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_aab672adb641c8c6) }

var fileDescriptor_test_aab672adb641c8c6 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x14, 0xb9, 0xb8, 0x43, 0x52,
	0x8b, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x4d, 0x2e, 0x4e, 0x88, 0x92,
	0x82, 0x9c, 0x4a, 0x21, 0x19, 0x2e, 0xce, 0x12, 0x18, 0x07, 0xaa, 0x0a, 0x21, 0x60, 0x64, 0xc6,
	0x05, 0x36, 0x55, 0x48, 0x8f, 0x8b, 0xcd, 0x25, 0x1f, 0xa4, 0x49, 0x48, 0x50, 0x0f, 0x6c, 0x25,
	0x92, 0x1d, 0x52, 0xfc, 0xc8, 0x42, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x27, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x70, 0x15, 0x65, 0x30, 0xa0, 0x00, 0x00, 0x00,
}
