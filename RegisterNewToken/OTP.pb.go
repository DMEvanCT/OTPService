// Code generated by protoc-gen-go. DO NOT EDIT.
// source: OTP.proto

package RegisterNewToken

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

// Get the original token to grab KEY ID
type SetTokenIDReq struct {
	YubikeyOtp           string   `protobuf:"bytes,1,opt,name=yubikey_otp,json=yubikeyOtp,proto3" json:"yubikey_otp,omitempty"`
	Userid               int32    `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTokenIDReq) Reset()         { *m = SetTokenIDReq{} }
func (m *SetTokenIDReq) String() string { return proto.CompactTextString(m) }
func (*SetTokenIDReq) ProtoMessage()    {}
func (*SetTokenIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0d5baf8df589173, []int{0}
}

func (m *SetTokenIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTokenIDReq.Unmarshal(m, b)
}
func (m *SetTokenIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTokenIDReq.Marshal(b, m, deterministic)
}
func (m *SetTokenIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTokenIDReq.Merge(m, src)
}
func (m *SetTokenIDReq) XXX_Size() int {
	return xxx_messageInfo_SetTokenIDReq.Size(m)
}
func (m *SetTokenIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTokenIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetTokenIDReq proto.InternalMessageInfo

func (m *SetTokenIDReq) GetYubikeyOtp() string {
	if m != nil {
		return m.YubikeyOtp
	}
	return ""
}

func (m *SetTokenIDReq) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

// Send a message back to tell the user the token was created
type SetTokenIDResp struct {
	TokenCreated         string   `protobuf:"bytes,1,opt,name=token_created,json=tokenCreated,proto3" json:"token_created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTokenIDResp) Reset()         { *m = SetTokenIDResp{} }
func (m *SetTokenIDResp) String() string { return proto.CompactTextString(m) }
func (*SetTokenIDResp) ProtoMessage()    {}
func (*SetTokenIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0d5baf8df589173, []int{1}
}

func (m *SetTokenIDResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTokenIDResp.Unmarshal(m, b)
}
func (m *SetTokenIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTokenIDResp.Marshal(b, m, deterministic)
}
func (m *SetTokenIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTokenIDResp.Merge(m, src)
}
func (m *SetTokenIDResp) XXX_Size() int {
	return xxx_messageInfo_SetTokenIDResp.Size(m)
}
func (m *SetTokenIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTokenIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_SetTokenIDResp proto.InternalMessageInfo

func (m *SetTokenIDResp) GetTokenCreated() string {
	if m != nil {
		return m.TokenCreated
	}
	return ""
}

func init() {
	proto.RegisterType((*SetTokenIDReq)(nil), "RegisterNewToken.SetTokenIDReq")
	proto.RegisterType((*SetTokenIDResp)(nil), "RegisterNewToken.SetTokenIDResp")
}

func init() {
	proto.RegisterFile("OTP.proto", fileDescriptor_f0d5baf8df589173)
}

var fileDescriptor_f0d5baf8df589173 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf4, 0x0f, 0x09, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x08, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0xf2,
	0x4b, 0x2d, 0x0f, 0xc9, 0xcf, 0x4e, 0xcd, 0x53, 0xf2, 0xe0, 0xe2, 0x0d, 0x4e, 0x2d, 0x01, 0xb3,
	0x3d, 0x5d, 0x82, 0x52, 0x0b, 0x85, 0xe4, 0xb9, 0xb8, 0x2b, 0x4b, 0x93, 0x32, 0xb3, 0x53, 0x2b,
	0xe3, 0xf3, 0x4b, 0x0a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xb8, 0xa0, 0x42, 0xfe, 0x25,
	0x05, 0x42, 0x62, 0x5c, 0x6c, 0xa5, 0xc5, 0xa9, 0x45, 0x99, 0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0xac, 0x41, 0x50, 0x9e, 0x92, 0x29, 0x17, 0x1f, 0xb2, 0x49, 0xc5, 0x05, 0x42, 0xca, 0x5c, 0xbc,
	0x25, 0x20, 0x6e, 0x7c, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x6a, 0x0a, 0xd4, 0x30, 0x1e, 0xb0, 0xa0,
	0x33, 0x44, 0xcc, 0x28, 0x8d, 0x8b, 0x07, 0xac, 0x27, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55,
	0x28, 0x8c, 0x8b, 0x3f, 0x12, 0x62, 0x19, 0xcc, 0xad, 0x42, 0xf2, 0x7a, 0xe8, 0xce, 0xd6, 0x43,
	0x71, 0xb3, 0x94, 0x02, 0x7e, 0x05, 0xc5, 0x05, 0x4a, 0x0c, 0x49, 0x6c, 0xe0, 0x10, 0x30, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xe5, 0xa7, 0xf4, 0x0e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenServiceClient interface {
	YubikeyRegister(ctx context.Context, in *SetTokenIDReq, opts ...grpc.CallOption) (*SetTokenIDResp, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) YubikeyRegister(ctx context.Context, in *SetTokenIDReq, opts ...grpc.CallOption) (*SetTokenIDResp, error) {
	out := new(SetTokenIDResp)
	err := c.cc.Invoke(ctx, "/RegisterNewToken.TokenService/YubikeyRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
type TokenServiceServer interface {
	YubikeyRegister(context.Context, *SetTokenIDReq) (*SetTokenIDResp, error)
}

// UnimplementedTokenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (*UnimplementedTokenServiceServer) YubikeyRegister(ctx context.Context, req *SetTokenIDReq) (*SetTokenIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method YubikeyRegister not implemented")
}

func RegisterTokenServiceServer(s *grpc.Server, srv TokenServiceServer) {
	s.RegisterService(&_TokenService_serviceDesc, srv)
}

func _TokenService_YubikeyRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTokenIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).YubikeyRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegisterNewToken.TokenService/YubikeyRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).YubikeyRegister(ctx, req.(*SetTokenIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RegisterNewToken.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "YubikeyRegister",
			Handler:    _TokenService_YubikeyRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OTP.proto",
}
