// Code generated by protoc-gen-go. DO NOT EDIT.
// source: codenamecreator.proto

package codenamecreator

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

type NameRequest struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameRequest) Reset()         { *m = NameRequest{} }
func (m *NameRequest) String() string { return proto.CompactTextString(m) }
func (*NameRequest) ProtoMessage()    {}
func (*NameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b22556833824cba, []int{0}
}

func (m *NameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameRequest.Unmarshal(m, b)
}
func (m *NameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameRequest.Marshal(b, m, deterministic)
}
func (m *NameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRequest.Merge(m, src)
}
func (m *NameRequest) XXX_Size() int {
	return xxx_messageInfo_NameRequest.Size(m)
}
func (m *NameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NameRequest proto.InternalMessageInfo

func (m *NameRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type NameResult struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameResult) Reset()         { *m = NameResult{} }
func (m *NameResult) String() string { return proto.CompactTextString(m) }
func (*NameResult) ProtoMessage()    {}
func (*NameResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b22556833824cba, []int{1}
}

func (m *NameResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameResult.Unmarshal(m, b)
}
func (m *NameResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameResult.Marshal(b, m, deterministic)
}
func (m *NameResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameResult.Merge(m, src)
}
func (m *NameResult) XXX_Size() int {
	return xxx_messageInfo_NameResult.Size(m)
}
func (m *NameResult) XXX_DiscardUnknown() {
	xxx_messageInfo_NameResult.DiscardUnknown(m)
}

var xxx_messageInfo_NameResult proto.InternalMessageInfo

func (m *NameResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*NameRequest)(nil), "NameRequest")
	proto.RegisterType((*NameResult)(nil), "NameResult")
}

func init() { proto.RegisterFile("codenamecreator.proto", fileDescriptor_1b22556833824cba) }

var fileDescriptor_1b22556833824cba = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0x4f, 0x49,
	0xcd, 0x4b, 0xcc, 0x4d, 0x4d, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x57, 0xd2, 0xe4, 0xe2, 0xf6, 0x4b, 0xcc, 0x4d, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x92, 0xe2, 0xe2, 0x48, 0x4e, 0x2c, 0x49, 0x4d, 0xcf, 0x2f, 0xaa, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95, 0x14, 0xb8, 0xb8, 0x20, 0x4a, 0x8b, 0x4b, 0x73, 0x4a, 0x84,
	0x84, 0xb8, 0x58, 0x40, 0xa6, 0x41, 0x55, 0x81, 0xd9, 0x46, 0x65, 0x5c, 0xfc, 0xce, 0x50, 0x5b,
	0x9c, 0x21, 0xb6, 0x08, 0x69, 0x71, 0x71, 0xbb, 0xa7, 0x96, 0xc0, 0x44, 0x85, 0x78, 0xf4, 0x90,
	0x6c, 0x93, 0xe2, 0xd6, 0x43, 0x18, 0xa8, 0xc4, 0x20, 0x64, 0xce, 0x25, 0xe2, 0x9d, 0x9a, 0x5a,
	0xe0, 0x9e, 0x5a, 0x52, 0x92, 0x99, 0x97, 0x0e, 0xd3, 0x53, 0x8c, 0x57, 0x93, 0x06, 0xa3, 0x01,
	0xa3, 0x93, 0x60, 0x14, 0x3f, 0x9a, 0xef, 0x92, 0xd8, 0xc0, 0xde, 0x33, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xf8, 0x4a, 0xe0, 0x00, 0xf7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CodenameCreatorClient is the client API for CodenameCreator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CodenameCreatorClient interface {
	GetCodename(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResult, error)
	KeepGettingCodenames(ctx context.Context, opts ...grpc.CallOption) (CodenameCreator_KeepGettingCodenamesClient, error)
}

type codenameCreatorClient struct {
	cc *grpc.ClientConn
}

func NewCodenameCreatorClient(cc *grpc.ClientConn) CodenameCreatorClient {
	return &codenameCreatorClient{cc}
}

func (c *codenameCreatorClient) GetCodename(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResult, error) {
	out := new(NameResult)
	err := c.cc.Invoke(ctx, "/CodenameCreator/GetCodename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codenameCreatorClient) KeepGettingCodenames(ctx context.Context, opts ...grpc.CallOption) (CodenameCreator_KeepGettingCodenamesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CodenameCreator_serviceDesc.Streams[0], "/CodenameCreator/KeepGettingCodenames", opts...)
	if err != nil {
		return nil, err
	}
	x := &codenameCreatorKeepGettingCodenamesClient{stream}
	return x, nil
}

type CodenameCreator_KeepGettingCodenamesClient interface {
	Send(*NameRequest) error
	Recv() (*NameResult, error)
	grpc.ClientStream
}

type codenameCreatorKeepGettingCodenamesClient struct {
	grpc.ClientStream
}

func (x *codenameCreatorKeepGettingCodenamesClient) Send(m *NameRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *codenameCreatorKeepGettingCodenamesClient) Recv() (*NameResult, error) {
	m := new(NameResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CodenameCreatorServer is the server API for CodenameCreator service.
type CodenameCreatorServer interface {
	GetCodename(context.Context, *NameRequest) (*NameResult, error)
	KeepGettingCodenames(CodenameCreator_KeepGettingCodenamesServer) error
}

func RegisterCodenameCreatorServer(s *grpc.Server, srv CodenameCreatorServer) {
	s.RegisterService(&_CodenameCreator_serviceDesc, srv)
}

func _CodenameCreator_GetCodename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodenameCreatorServer).GetCodename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodenameCreator/GetCodename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodenameCreatorServer).GetCodename(ctx, req.(*NameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodenameCreator_KeepGettingCodenames_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CodenameCreatorServer).KeepGettingCodenames(&codenameCreatorKeepGettingCodenamesServer{stream})
}

type CodenameCreator_KeepGettingCodenamesServer interface {
	Send(*NameResult) error
	Recv() (*NameRequest, error)
	grpc.ServerStream
}

type codenameCreatorKeepGettingCodenamesServer struct {
	grpc.ServerStream
}

func (x *codenameCreatorKeepGettingCodenamesServer) Send(m *NameResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *codenameCreatorKeepGettingCodenamesServer) Recv() (*NameRequest, error) {
	m := new(NameRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CodenameCreator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CodenameCreator",
	HandlerType: (*CodenameCreatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCodename",
			Handler:    _CodenameCreator_GetCodename_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "KeepGettingCodenames",
			Handler:       _CodenameCreator_KeepGettingCodenames_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "codenamecreator.proto",
}