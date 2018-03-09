// Code generated by protoc-gen-go. DO NOT EDIT.
// source: invite.proto

/*
Package webpb is a generated protocol buffer package.

It is generated from these files:
	invite.proto

It has these top-level messages:
	InviteRequest
	InviteResponse
	User
*/
package webpb

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

type InviteRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Page   int32  `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	Size   int32  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (m *InviteRequest) Reset()                    { *m = InviteRequest{} }
func (m *InviteRequest) String() string            { return proto.CompactTextString(m) }
func (*InviteRequest) ProtoMessage()               {}
func (*InviteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InviteRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *InviteRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *InviteRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type InviteResponse struct {
	Page int32   `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	Size int32   `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Data []*User `protobuf:"bytes,3,rep,name=data" json:"data,omitempty"`
}

func (m *InviteResponse) Reset()                    { *m = InviteResponse{} }
func (m *InviteResponse) String() string            { return proto.CompactTextString(m) }
func (*InviteResponse) ProtoMessage()               {}
func (*InviteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InviteResponse) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *InviteResponse) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *InviteResponse) GetData() []*User {
	if m != nil {
		return m.Data
	}
	return nil
}

type User struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*InviteRequest)(nil), "InviteRequest")
	proto.RegisterType((*InviteResponse)(nil), "InviteResponse")
	proto.RegisterType((*User)(nil), "User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for InviteService service

type InviteServiceClient interface {
	// Sends a greeting
	InviteList(ctx context.Context, in *InviteRequest, opts ...grpc.CallOption) (*InviteResponse, error)
}

type inviteServiceClient struct {
	cc *grpc.ClientConn
}

func NewInviteServiceClient(cc *grpc.ClientConn) InviteServiceClient {
	return &inviteServiceClient{cc}
}

func (c *inviteServiceClient) InviteList(ctx context.Context, in *InviteRequest, opts ...grpc.CallOption) (*InviteResponse, error) {
	out := new(InviteResponse)
	err := grpc.Invoke(ctx, "/InviteService/InviteList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InviteService service

type InviteServiceServer interface {
	// Sends a greeting
	InviteList(context.Context, *InviteRequest) (*InviteResponse, error)
}

func RegisterInviteServiceServer(s *grpc.Server, srv InviteServiceServer) {
	s.RegisterService(&_InviteService_serviceDesc, srv)
}

func _InviteService_InviteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InviteServiceServer).InviteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InviteService/InviteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InviteServiceServer).InviteList(ctx, req.(*InviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InviteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "InviteService",
	HandlerType: (*InviteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InviteList",
			Handler:    _InviteService_InviteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invite.proto",
}

func init() { proto.RegisterFile("invite.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x14, 0xc4, 0xcd, 0x9f, 0x56, 0xfa, 0xd4, 0x08, 0x7b, 0x31, 0x7a, 0x0a, 0x39, 0x85, 0x22, 0xbb,
	0x58, 0x8f, 0x5e, 0xc4, 0x5b, 0xc1, 0x83, 0x6c, 0xf1, 0xe2, 0x45, 0x76, 0x9b, 0x47, 0xbb, 0x87,
	0xec, 0xc6, 0xec, 0x26, 0x42, 0x3f, 0xbd, 0xec, 0x86, 0x56, 0x02, 0xbd, 0xcd, 0xfc, 0x60, 0x1e,
	0x33, 0x0f, 0xae, 0x95, 0x1e, 0x94, 0x43, 0xda, 0x76, 0xc6, 0x99, 0xf2, 0x03, 0x6e, 0xd6, 0xc1,
	0x73, 0xfc, 0xe9, 0xd1, 0x3a, 0x72, 0x07, 0x97, 0xbd, 0xc5, 0xee, 0x5b, 0xd5, 0x79, 0x54, 0x44,
	0xd5, 0x82, 0xcf, 0xbd, 0x5d, 0xd7, 0x84, 0x40, 0xda, 0x8a, 0x1d, 0xe6, 0x71, 0x11, 0x55, 0x33,
	0x1e, 0xb4, 0x67, 0x56, 0x1d, 0x30, 0x4f, 0x46, 0xe6, 0x75, 0xb9, 0x81, 0xec, 0x78, 0xd1, 0xb6,
	0x46, 0x5b, 0x3c, 0x25, 0xa3, 0x33, 0xc9, 0xf8, 0x3f, 0x49, 0xee, 0x21, 0xad, 0x85, 0x13, 0x79,
	0x52, 0x24, 0xd5, 0xd5, 0x6a, 0x46, 0x3f, 0x2d, 0x76, 0x3c, 0xa0, 0x72, 0x09, 0xa9, 0x77, 0x24,
	0x83, 0xf8, 0x54, 0x2c, 0x56, 0xa1, 0x94, 0x16, 0xcd, 0x78, 0x66, 0xc1, 0x83, 0x5e, 0xbd, 0x1e,
	0x27, 0x6d, 0xb0, 0x1b, 0xd4, 0x16, 0x09, 0x03, 0x18, 0xc1, 0xbb, 0xb2, 0x8e, 0x64, 0x74, 0x32,
	0xf8, 0xe1, 0x96, 0x4e, 0xeb, 0x96, 0x17, 0x6f, 0x8f, 0x5f, 0xcb, 0x9d, 0x72, 0xfb, 0x5e, 0xd2,
	0xad, 0x69, 0x98, 0x44, 0xa1, 0x0f, 0x7b, 0xd3, 0x33, 0x27, 0x74, 0x8d, 0x8d, 0x61, 0xe1, 0x73,
	0x6c, 0x78, 0x7a, 0xf9, 0x45, 0xd9, 0x4a, 0x39, 0x0f, 0xfe, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0xfd, 0x9d, 0xa4, 0xe6, 0x59, 0x01, 0x00, 0x00,
}
