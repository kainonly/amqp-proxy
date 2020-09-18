// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package amqp_proxy

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

type Response struct {
	Error                uint32   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
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

func (m *Response) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PublishParameter struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Mandatory            bool     `protobuf:"varint,3,opt,name=mandatory,proto3" json:"mandatory,omitempty"`
	Immediate            bool     `protobuf:"varint,4,opt,name=immediate,proto3" json:"immediate,omitempty"`
	ContentType          string   `protobuf:"bytes,5,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Body                 []byte   `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishParameter) Reset()         { *m = PublishParameter{} }
func (m *PublishParameter) String() string { return proto.CompactTextString(m) }
func (*PublishParameter) ProtoMessage()    {}
func (*PublishParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *PublishParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishParameter.Unmarshal(m, b)
}
func (m *PublishParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishParameter.Marshal(b, m, deterministic)
}
func (m *PublishParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishParameter.Merge(m, src)
}
func (m *PublishParameter) XXX_Size() int {
	return xxx_messageInfo_PublishParameter.Size(m)
}
func (m *PublishParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishParameter.DiscardUnknown(m)
}

var xxx_messageInfo_PublishParameter proto.InternalMessageInfo

func (m *PublishParameter) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *PublishParameter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PublishParameter) GetMandatory() bool {
	if m != nil {
		return m.Mandatory
	}
	return false
}

func (m *PublishParameter) GetImmediate() bool {
	if m != nil {
		return m.Immediate
	}
	return false
}

func (m *PublishParameter) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *PublishParameter) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type GetParameter struct {
	Queue                string   `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParameter) Reset()         { *m = GetParameter{} }
func (m *GetParameter) String() string { return proto.CompactTextString(m) }
func (*GetParameter) ProtoMessage()    {}
func (*GetParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{2}
}

func (m *GetParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParameter.Unmarshal(m, b)
}
func (m *GetParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParameter.Marshal(b, m, deterministic)
}
func (m *GetParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParameter.Merge(m, src)
}
func (m *GetParameter) XXX_Size() int {
	return xxx_messageInfo_GetParameter.Size(m)
}
func (m *GetParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParameter.DiscardUnknown(m)
}

var xxx_messageInfo_GetParameter proto.InternalMessageInfo

func (m *GetParameter) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

type GetResponse struct {
	Error                uint32   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *Data    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *GetResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetResponse) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type Data struct {
	Receipt              string   `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{4}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetReceipt() string {
	if m != nil {
		return m.Receipt
	}
	return ""
}

func (m *Data) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type AckParameter struct {
	Queue                string   `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Receipt              string   `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckParameter) Reset()         { *m = AckParameter{} }
func (m *AckParameter) String() string { return proto.CompactTextString(m) }
func (*AckParameter) ProtoMessage()    {}
func (*AckParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{5}
}

func (m *AckParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckParameter.Unmarshal(m, b)
}
func (m *AckParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckParameter.Marshal(b, m, deterministic)
}
func (m *AckParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckParameter.Merge(m, src)
}
func (m *AckParameter) XXX_Size() int {
	return xxx_messageInfo_AckParameter.Size(m)
}
func (m *AckParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_AckParameter.DiscardUnknown(m)
}

var xxx_messageInfo_AckParameter proto.InternalMessageInfo

func (m *AckParameter) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *AckParameter) GetReceipt() string {
	if m != nil {
		return m.Receipt
	}
	return ""
}

type NackParameter struct {
	Queue                string   `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Receipt              string   `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NackParameter) Reset()         { *m = NackParameter{} }
func (m *NackParameter) String() string { return proto.CompactTextString(m) }
func (*NackParameter) ProtoMessage()    {}
func (*NackParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{6}
}

func (m *NackParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NackParameter.Unmarshal(m, b)
}
func (m *NackParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NackParameter.Marshal(b, m, deterministic)
}
func (m *NackParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NackParameter.Merge(m, src)
}
func (m *NackParameter) XXX_Size() int {
	return xxx_messageInfo_NackParameter.Size(m)
}
func (m *NackParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_NackParameter.DiscardUnknown(m)
}

var xxx_messageInfo_NackParameter proto.InternalMessageInfo

func (m *NackParameter) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *NackParameter) GetReceipt() string {
	if m != nil {
		return m.Receipt
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "amqp.proxy.Response")
	proto.RegisterType((*PublishParameter)(nil), "amqp.proxy.PublishParameter")
	proto.RegisterType((*GetParameter)(nil), "amqp.proxy.GetParameter")
	proto.RegisterType((*GetResponse)(nil), "amqp.proxy.GetResponse")
	proto.RegisterType((*Data)(nil), "amqp.proxy.Data")
	proto.RegisterType((*AckParameter)(nil), "amqp.proxy.AckParameter")
	proto.RegisterType((*NackParameter)(nil), "amqp.proxy.NackParameter")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6b, 0xea, 0x40,
	0x14, 0x7d, 0x31, 0xf1, 0xeb, 0x1a, 0x41, 0x06, 0xe1, 0xcd, 0x13, 0x17, 0x21, 0xb8, 0xc8, 0xca,
	0x85, 0xaf, 0x50, 0xda, 0x45, 0x45, 0x28, 0xb8, 0x2b, 0x32, 0x74, 0xd7, 0xd5, 0x98, 0x5c, 0x34,
	0xa4, 0xc9, 0xc4, 0xc9, 0x04, 0xcc, 0xdf, 0xea, 0xbf, 0xeb, 0xae, 0x24, 0x1a, 0x1d, 0x05, 0x29,
	0xa5, 0xbb, 0x7b, 0xee, 0xe1, 0x9e, 0x39, 0x73, 0xe6, 0x0e, 0xd8, 0x52, 0xe4, 0x0a, 0xe5, 0x34,
	0x95, 0x42, 0x09, 0x02, 0x3c, 0xde, 0xa5, 0x65, 0xbd, 0x2f, 0xdc, 0x19, 0x74, 0x18, 0x66, 0xa9,
	0x48, 0x32, 0x24, 0x43, 0x68, 0xa2, 0x94, 0x42, 0x52, 0xc3, 0x31, 0xbc, 0x3e, 0x3b, 0x00, 0x32,
	0x00, 0x33, 0xce, 0x36, 0xb4, 0xe1, 0x18, 0x5e, 0x97, 0x95, 0xa5, 0xfb, 0x61, 0xc0, 0x60, 0x95,
	0xaf, 0xdf, 0xc3, 0x6c, 0xbb, 0xe2, 0x92, 0xc7, 0xa8, 0x50, 0x92, 0x11, 0x74, 0x70, 0xef, 0x6f,
	0x79, 0xb2, 0xc1, 0x6a, 0xbe, 0xcb, 0x4e, 0xb8, 0x94, 0x88, 0xb0, 0xa8, 0x25, 0x22, 0x2c, 0xc8,
	0x18, 0xba, 0x31, 0x4f, 0x02, 0xae, 0x84, 0x2c, 0xa8, 0xe9, 0x18, 0x5e, 0x87, 0x9d, 0x1b, 0x25,
	0x1b, 0xc6, 0x31, 0x06, 0x21, 0x57, 0x48, 0xad, 0x03, 0x7b, 0x6a, 0x10, 0x07, 0x7a, 0xbe, 0x48,
	0x14, 0x26, 0xea, 0xb5, 0x48, 0x91, 0x36, 0x2b, 0x55, 0xbd, 0x45, 0x08, 0x58, 0x6b, 0x11, 0x14,
	0xb4, 0xe5, 0x18, 0x9e, 0xcd, 0xaa, 0xda, 0x9d, 0x80, 0xbd, 0x44, 0x75, 0xf6, 0x3b, 0x84, 0xe6,
	0x2e, 0xc7, 0xbc, 0x36, 0x7b, 0x00, 0xee, 0x1b, 0xf4, 0x96, 0xa8, 0x7e, 0x9a, 0x08, 0x99, 0x80,
	0x15, 0x70, 0xc5, 0xab, 0x9b, 0xf4, 0x66, 0x83, 0xe9, 0x39, 0xe0, 0xe9, 0x33, 0x57, 0x9c, 0x55,
	0xac, 0x7b, 0x07, 0x56, 0x89, 0x08, 0x85, 0xb6, 0x44, 0x1f, 0xc3, 0x54, 0x1d, 0x0f, 0xaf, 0xe1,
	0xc9, 0x78, 0x43, 0x33, 0xfe, 0x04, 0xf6, 0xc2, 0x8f, 0xbe, 0x31, 0xae, 0x6b, 0x36, 0x2e, 0x34,
	0xdd, 0x39, 0xf4, 0x5f, 0xf8, 0x2f, 0x04, 0x66, 0x9f, 0x06, 0xb4, 0x58, 0xb5, 0x3f, 0x64, 0x0e,
	0xed, 0xe3, 0xc3, 0x93, 0xb1, 0x7e, 0xc9, 0xeb, 0x6d, 0x18, 0x0d, 0x75, 0xb6, 0x8e, 0xd3, 0xfd,
	0x43, 0x1e, 0xc1, 0x5c, 0xa2, 0x22, 0x54, 0xa7, 0xf5, 0x67, 0x19, 0xfd, 0xbd, 0x62, 0xb4, 0xd9,
	0x7b, 0x30, 0x17, 0x7e, 0x74, 0x39, 0xab, 0x27, 0x73, 0xf3, 0xd0, 0x07, 0xb0, 0xca, 0x04, 0xc8,
	0x3f, 0x9d, 0xbf, 0xc8, 0xe4, 0xd6, 0xe8, 0xba, 0x55, 0xfd, 0x98, 0xff, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x44, 0x6d, 0x63, 0x7c, 0x41, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	Publish(ctx context.Context, in *PublishParameter, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *GetParameter, opts ...grpc.CallOption) (*GetResponse, error)
	Ack(ctx context.Context, in *AckParameter, opts ...grpc.CallOption) (*Response, error)
	Nack(ctx context.Context, in *NackParameter, opts ...grpc.CallOption) (*Response, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Publish(ctx context.Context, in *PublishParameter, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/amqp.proxy.Router/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Get(ctx context.Context, in *GetParameter, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/amqp.proxy.Router/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Ack(ctx context.Context, in *AckParameter, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/amqp.proxy.Router/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Nack(ctx context.Context, in *NackParameter, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/amqp.proxy.Router/Nack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	Publish(context.Context, *PublishParameter) (*Response, error)
	Get(context.Context, *GetParameter) (*GetResponse, error)
	Ack(context.Context, *AckParameter) (*Response, error)
	Nack(context.Context, *NackParameter) (*Response, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) Publish(ctx context.Context, req *PublishParameter) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedRouterServer) Get(ctx context.Context, req *GetParameter) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRouterServer) Ack(ctx context.Context, req *AckParameter) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}
func (*UnimplementedRouterServer) Nack(ctx context.Context, req *NackParameter) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nack not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amqp.proxy.Router/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Publish(ctx, req.(*PublishParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amqp.proxy.Router/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Get(ctx, req.(*GetParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amqp.proxy.Router/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Ack(ctx, req.(*AckParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Nack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NackParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Nack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amqp.proxy.Router/Nack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Nack(ctx, req.(*NackParameter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "amqp.proxy.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Router_Publish_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Router_Get_Handler,
		},
		{
			MethodName: "Ack",
			Handler:    _Router_Ack_Handler,
		},
		{
			MethodName: "Nack",
			Handler:    _Router_Nack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router.proto",
}
