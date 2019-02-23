// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package ping

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mattmoor/protoc-plugin/proto"
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

type Request struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "ping.Request")
	proto.RegisterType((*Response)(nil), "ping.Response")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xc8, 0xcc, 0x4b, 0x97, 0x32, 0x4a,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x4d, 0x2c, 0x29, 0xc9, 0xcd,
	0xcf, 0x2f, 0xd2, 0x07, 0x2b, 0x48, 0xd6, 0x2d, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x83, 0xf0, 0xf4,
	0xb3, 0x8b, 0x0a, 0x92, 0x21, 0x3a, 0x95, 0xa4, 0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40,
	0x4c, 0x25, 0x19, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x4c, 0x59, 0xa3,
	0xf9, 0x8c, 0x5c, 0xdc, 0x01, 0x99, 0x79, 0xe9, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42,
	0x4e, 0x5c, 0x2c, 0x20, 0xae, 0x10, 0xaf, 0x1e, 0xc8, 0x35, 0x7a, 0x50, 0x63, 0xa5, 0xf8, 0x60,
	0x5c, 0x88, 0x41, 0x4a, 0xd2, 0x5d, 0xdb, 0x54, 0xc4, 0xb9, 0xd8, 0x53, 0x52, 0xd3, 0x12, 0x4b,
	0x73, 0x4a, 0x04, 0x44, 0x94, 0xb8, 0xb8, 0x98, 0xdd, 0xfc, 0xfd, 0x85, 0x98, 0x9d, 0x1c, 0x83,
	0x84, 0x7c, 0xb8, 0xb8, 0xc0, 0x46, 0x96, 0x14, 0xa5, 0x26, 0xe6, 0x92, 0x68, 0x12, 0x13, 0x92,
	0x49, 0x51, 0x1a, 0x8c, 0x06, 0x8c, 0x49, 0x6c, 0x60, 0x3f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x69, 0x23, 0x2f, 0xfd, 0x2c, 0x01, 0x00, 0x00,
}
