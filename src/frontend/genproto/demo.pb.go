// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package ymotongpoo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ParseRequest struct {
	TargetStr            string   `protobuf:"bytes,1,opt,name=target_str,json=targetStr,proto3" json:"target_str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseRequest) Reset()         { *m = ParseRequest{} }
func (m *ParseRequest) String() string { return proto.CompactTextString(m) }
func (*ParseRequest) ProtoMessage()    {}
func (*ParseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_6398af5df3e25236, []int{0}
}
func (m *ParseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseRequest.Unmarshal(m, b)
}
func (m *ParseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseRequest.Marshal(b, m, deterministic)
}
func (dst *ParseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseRequest.Merge(dst, src)
}
func (m *ParseRequest) XXX_Size() int {
	return xxx_messageInfo_ParseRequest.Size(m)
}
func (m *ParseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParseRequest proto.InternalMessageInfo

func (m *ParseRequest) GetTargetStr() string {
	if m != nil {
		return m.TargetStr
	}
	return ""
}

type ParsedArray struct {
	Number               []int64  `protobuf:"varint,1,rep,packed,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParsedArray) Reset()         { *m = ParsedArray{} }
func (m *ParsedArray) String() string { return proto.CompactTextString(m) }
func (*ParsedArray) ProtoMessage()    {}
func (*ParsedArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_6398af5df3e25236, []int{1}
}
func (m *ParsedArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParsedArray.Unmarshal(m, b)
}
func (m *ParsedArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParsedArray.Marshal(b, m, deterministic)
}
func (dst *ParsedArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParsedArray.Merge(dst, src)
}
func (m *ParsedArray) XXX_Size() int {
	return xxx_messageInfo_ParsedArray.Size(m)
}
func (m *ParsedArray) XXX_DiscardUnknown() {
	xxx_messageInfo_ParsedArray.DiscardUnknown(m)
}

var xxx_messageInfo_ParsedArray proto.InternalMessageInfo

func (m *ParsedArray) GetNumber() []int64 {
	if m != nil {
		return m.Number
	}
	return nil
}

type AddRequest struct {
	Number               []int64  `protobuf:"varint,1,rep,packed,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_6398af5df3e25236, []int{2}
}
func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (dst *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(dst, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetNumber() []int64 {
	if m != nil {
		return m.Number
	}
	return nil
}

type AddResult struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResult) Reset()         { *m = AddResult{} }
func (m *AddResult) String() string { return proto.CompactTextString(m) }
func (*AddResult) ProtoMessage()    {}
func (*AddResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_6398af5df3e25236, []int{3}
}
func (m *AddResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResult.Unmarshal(m, b)
}
func (m *AddResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResult.Marshal(b, m, deterministic)
}
func (dst *AddResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResult.Merge(dst, src)
}
func (m *AddResult) XXX_Size() int {
	return xxx_messageInfo_AddResult.Size(m)
}
func (m *AddResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResult.DiscardUnknown(m)
}

var xxx_messageInfo_AddResult proto.InternalMessageInfo

func (m *AddResult) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*ParseRequest)(nil), "ymotongpoo.ParseRequest")
	proto.RegisterType((*ParsedArray)(nil), "ymotongpoo.ParsedArray")
	proto.RegisterType((*AddRequest)(nil), "ymotongpoo.AddRequest")
	proto.RegisterType((*AddResult)(nil), "ymotongpoo.AddResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArrayParseServiceClient is the client API for ArrayParseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArrayParseServiceClient interface {
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParsedArray, error)
}

type arrayParseServiceClient struct {
	cc *grpc.ClientConn
}

func NewArrayParseServiceClient(cc *grpc.ClientConn) ArrayParseServiceClient {
	return &arrayParseServiceClient{cc}
}

func (c *arrayParseServiceClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParsedArray, error) {
	out := new(ParsedArray)
	err := c.cc.Invoke(ctx, "/ymotongpoo.ArrayParseService/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArrayParseServiceServer is the server API for ArrayParseService service.
type ArrayParseServiceServer interface {
	Parse(context.Context, *ParseRequest) (*ParsedArray, error)
}

func RegisterArrayParseServiceServer(s *grpc.Server, srv ArrayParseServiceServer) {
	s.RegisterService(&_ArrayParseService_serviceDesc, srv)
}

func _ArrayParseService_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArrayParseServiceServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ymotongpoo.ArrayParseService/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArrayParseServiceServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArrayParseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ymotongpoo.ArrayParseService",
	HandlerType: (*ArrayParseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _ArrayParseService_Parse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// AddNumberServiceClient is the client API for AddNumberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddNumberServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResult, error)
}

type addNumberServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddNumberServiceClient(cc *grpc.ClientConn) AddNumberServiceClient {
	return &addNumberServiceClient{cc}
}

func (c *addNumberServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResult, error) {
	out := new(AddResult)
	err := c.cc.Invoke(ctx, "/ymotongpoo.AddNumberService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddNumberServiceServer is the server API for AddNumberService service.
type AddNumberServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResult, error)
}

func RegisterAddNumberServiceServer(s *grpc.Server, srv AddNumberServiceServer) {
	s.RegisterService(&_AddNumberService_serviceDesc, srv)
}

func _AddNumberService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddNumberServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ymotongpoo.AddNumberService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddNumberServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddNumberService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ymotongpoo.AddNumberService",
	HandlerType: (*AddNumberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AddNumberService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_demo_6398af5df3e25236) }

var fileDescriptor_demo_6398af5df3e25236 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0x87, 0x40,
	0x10, 0xc7, 0x7f, 0x22, 0x09, 0x4e, 0x1d, 0x6a, 0x21, 0x13, 0x21, 0x88, 0xad, 0xc0, 0x4b, 0x1e,
	0x0c, 0x3a, 0x75, 0xd9, 0x6b, 0x87, 0x28, 0x7d, 0x80, 0xd0, 0x66, 0x90, 0x20, 0x5d, 0x9b, 0x5d,
	0x03, 0xdf, 0x3e, 0x5c, 0x13, 0x0d, 0xe9, 0x38, 0xb3, 0x9f, 0xfd, 0xfe, 0x19, 0x00, 0xa4, 0x56,
	0x67, 0x3d, 0x6b, 0xab, 0x05, 0x8c, 0xad, 0xb6, 0xba, 0x6b, 0x7a, 0xad, 0xe5, 0x1d, 0x9c, 0xbc,
	0x54, 0x6c, 0xa8, 0xa0, 0xaf, 0x81, 0x8c, 0x15, 0x97, 0x00, 0xb6, 0xe2, 0x86, 0xec, 0x9b, 0xb1,
	0x1c, 0x7b, 0x57, 0x5e, 0x1a, 0x16, 0xe1, 0xbc, 0x29, 0x2d, 0xcb, 0x5b, 0x38, 0x76, 0x38, 0x2a,
	0xe6, 0x6a, 0x14, 0x11, 0x04, 0xdd, 0xd0, 0xd6, 0x34, 0x91, 0x7e, 0xea, 0x17, 0xbf, 0x93, 0xbc,
	0x01, 0x50, 0x88, 0x8b, 0xe6, 0x7f, 0xd4, 0x35, 0x84, 0x8e, 0x32, 0xc3, 0xe7, 0x5f, 0xc8, 0x5b,
	0xa1, 0xfc, 0x15, 0xce, 0x9c, 0x97, 0xb3, 0x2d, 0x89, 0xbf, 0x3f, 0xde, 0x49, 0x3c, 0xc2, 0x91,
	0x9b, 0x45, 0x9c, 0xad, 0x5d, 0xb2, 0x6d, 0x91, 0xe4, 0x62, 0xf7, 0x32, 0x67, 0x96, 0x87, 0xfc,
	0x09, 0x4e, 0x15, 0xe2, 0xb3, 0xd3, 0x5f, 0x14, 0x1f, 0xc0, 0x57, 0x88, 0x22, 0xda, 0xfe, 0x5a,
	0x2b, 0x24, 0xe7, 0xbb, 0xfd, 0x14, 0x5a, 0x1e, 0xea, 0xc0, 0x9d, 0xf4, 0xfe, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x07, 0x9d, 0xc3, 0x57, 0x60, 0x01, 0x00, 0x00,
}
