// Code generated by protoc-gen-go. DO NOT EDIT.
// source: say.proto

package say

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

type Text struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dad255d0361606e, []int{0}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Speech struct {
	Audio                []byte   `protobuf:"bytes,1,opt,name=Audio,proto3" json:"Audio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Speech) Reset()         { *m = Speech{} }
func (m *Speech) String() string { return proto.CompactTextString(m) }
func (*Speech) ProtoMessage()    {}
func (*Speech) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dad255d0361606e, []int{1}
}

func (m *Speech) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Speech.Unmarshal(m, b)
}
func (m *Speech) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Speech.Marshal(b, m, deterministic)
}
func (m *Speech) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Speech.Merge(m, src)
}
func (m *Speech) XXX_Size() int {
	return xxx_messageInfo_Speech.Size(m)
}
func (m *Speech) XXX_DiscardUnknown() {
	xxx_messageInfo_Speech.DiscardUnknown(m)
}

var xxx_messageInfo_Speech proto.InternalMessageInfo

func (m *Speech) GetAudio() []byte {
	if m != nil {
		return m.Audio
	}
	return nil
}

func init() {
	proto.RegisterType((*Text)(nil), "say.Text")
	proto.RegisterType((*Speech)(nil), "say.Speech")
}

func init() { proto.RegisterFile("say.proto", fileDescriptor_6dad255d0361606e) }

var fileDescriptor_6dad255d0361606e = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x4e, 0xac, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x4e, 0xac, 0x54, 0x92, 0xe2, 0x62, 0x09, 0x49,
	0xad, 0x28, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad, 0x28, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0xe4, 0xb8, 0xd8, 0x82, 0x0b, 0x52, 0x53, 0x93, 0x33, 0x84, 0x44, 0xb8,
	0x58, 0x1d, 0x4b, 0x53, 0x32, 0xf3, 0xc1, 0xd2, 0x3c, 0x41, 0x10, 0x8e, 0x91, 0x3e, 0x17, 0x0f,
	0x48, 0x6f, 0x48, 0x3e, 0x54, 0x95, 0x3c, 0x17, 0x73, 0x70, 0x62, 0xa5, 0x10, 0xa7, 0x1e, 0xc8,
	0x0e, 0x90, 0x8c, 0x14, 0x37, 0x98, 0x09, 0x91, 0x56, 0x62, 0x48, 0x62, 0x03, 0x5b, 0x6c, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xab, 0xb3, 0x00, 0x19, 0x85, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TextToSpeechClient is the client API for TextToSpeech service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TextToSpeechClient interface {
	Say(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Speech, error)
}

type textToSpeechClient struct {
	cc *grpc.ClientConn
}

func NewTextToSpeechClient(cc *grpc.ClientConn) TextToSpeechClient {
	return &textToSpeechClient{cc}
}
    
func (c *textToSpeechClient) Say(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Speech, error) {
	out := new(Speech)
	err := c.cc.Invoke(ctx, "/say.TextToSpeech/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextToSpeechServer is the server API for TextToSpeech service.
type TextToSpeechServer interface {
	Say(context.Context, *Text) (*Speech, error)
}

// UnimplementedTextToSpeechServer can be embedded to have forward compatible implementations.
type UnimplementedTextToSpeechServer struct {
}

func (*UnimplementedTextToSpeechServer) Say(ctx context.Context, req *Text) (*Speech, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}

func RegisterTextToSpeechServer(s *grpc.Server, srv TextToSpeechServer) {
	s.RegisterService(&_TextToSpeech_serviceDesc, srv)
}

func _TextToSpeech_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextToSpeechServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/say.TextToSpeech/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextToSpeechServer).Say(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

var _TextToSpeech_serviceDesc = grpc.ServiceDesc{
	ServiceName: "say.TextToSpeech",
	HandlerType: (*TextToSpeechServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _TextToSpeech_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "say.proto",
}
