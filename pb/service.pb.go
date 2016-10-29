// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	StatusReq
	StatusResp
	ApplyReq
	ApplyResp
*/
package pb

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

type StatusReq struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
}

func (m *StatusReq) Reset()                    { *m = StatusReq{} }
func (m *StatusReq) String() string            { return proto.CompactTextString(m) }
func (*StatusReq) ProtoMessage()               {}
func (*StatusReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusResp struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *StatusResp) Reset()                    { *m = StatusResp{} }
func (m *StatusResp) String() string            { return proto.CompactTextString(m) }
func (*StatusResp) ProtoMessage()               {}
func (*StatusResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ApplyReq struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Giturl    string `protobuf:"bytes,2,opt,name=giturl" json:"giturl,omitempty"`
	Destroy   bool   `protobuf:"varint,3,opt,name=destroy" json:"destroy,omitempty"`
}

func (m *ApplyReq) Reset()                    { *m = ApplyReq{} }
func (m *ApplyReq) String() string            { return proto.CompactTextString(m) }
func (*ApplyReq) ProtoMessage()               {}
func (*ApplyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ApplyResp struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
}

func (m *ApplyResp) Reset()                    { *m = ApplyResp{} }
func (m *ApplyResp) String() string            { return proto.CompactTextString(m) }
func (*ApplyResp) ProtoMessage()               {}
func (*ApplyResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*StatusReq)(nil), "pb.StatusReq")
	proto.RegisterType((*StatusResp)(nil), "pb.StatusResp")
	proto.RegisterType((*ApplyReq)(nil), "pb.ApplyReq")
	proto.RegisterType((*ApplyResp)(nil), "pb.ApplyResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Provision service

type ProvisionClient interface {
	Status(ctx context.Context, in *StatusReq, opts ...grpc.CallOption) (*StatusResp, error)
	Apply(ctx context.Context, in *ApplyReq, opts ...grpc.CallOption) (*ApplyResp, error)
	Destroy(ctx context.Context, in *ApplyReq, opts ...grpc.CallOption) (*ApplyResp, error)
}

type provisionClient struct {
	cc *grpc.ClientConn
}

func NewProvisionClient(cc *grpc.ClientConn) ProvisionClient {
	return &provisionClient{cc}
}

func (c *provisionClient) Status(ctx context.Context, in *StatusReq, opts ...grpc.CallOption) (*StatusResp, error) {
	out := new(StatusResp)
	err := grpc.Invoke(ctx, "/pb.Provision/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionClient) Apply(ctx context.Context, in *ApplyReq, opts ...grpc.CallOption) (*ApplyResp, error) {
	out := new(ApplyResp)
	err := grpc.Invoke(ctx, "/pb.Provision/Apply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionClient) Destroy(ctx context.Context, in *ApplyReq, opts ...grpc.CallOption) (*ApplyResp, error) {
	out := new(ApplyResp)
	err := grpc.Invoke(ctx, "/pb.Provision/Destroy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Provision service

type ProvisionServer interface {
	Status(context.Context, *StatusReq) (*StatusResp, error)
	Apply(context.Context, *ApplyReq) (*ApplyResp, error)
	Destroy(context.Context, *ApplyReq) (*ApplyResp, error)
}

func RegisterProvisionServer(s *grpc.Server, srv ProvisionServer) {
	s.RegisterService(&_Provision_serviceDesc, srv)
}

func _Provision_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Provision/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionServer).Status(ctx, req.(*StatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provision_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Provision/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionServer).Apply(ctx, req.(*ApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provision_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Provision/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionServer).Destroy(ctx, req.(*ApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Provision_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Provision",
	HandlerType: (*ProvisionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Provision_Status_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Provision_Apply_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _Provision_Destroy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe2,
	0xe2, 0x0c, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x0e, 0x4a, 0x2d, 0x14, 0x92, 0xe5, 0xe2, 0x2a, 0x28,
	0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x89, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2,
	0x84, 0x8a, 0x78, 0xa6, 0x28, 0x39, 0x73, 0x71, 0xc1, 0xd4, 0x16, 0x17, 0x10, 0x50, 0x2c, 0x24,
	0xc6, 0xc5, 0x56, 0x0c, 0x56, 0x2c, 0xc1, 0x04, 0x96, 0x82, 0xf2, 0x94, 0xa2, 0xb9, 0x38, 0x1c,
	0x0b, 0x0a, 0x72, 0x2a, 0x09, 0xdb, 0x07, 0x32, 0x22, 0x3d, 0xb3, 0xa4, 0xb4, 0x28, 0x07, 0x66,
	0x04, 0x84, 0x27, 0x24, 0xc1, 0xc5, 0x9e, 0x92, 0x5a, 0x5c, 0x52, 0x94, 0x5f, 0x29, 0xc1, 0xac,
	0xc0, 0xa8, 0xc1, 0x11, 0x04, 0xe3, 0x82, 0x7c, 0x03, 0x35, 0x9c, 0xa0, 0x03, 0x8d, 0x3a, 0x18,
	0xb9, 0x38, 0x03, 0x8a, 0xf2, 0xcb, 0x32, 0x8b, 0x33, 0xf3, 0xf3, 0x84, 0x34, 0xb9, 0xd8, 0x20,
	0x7e, 0x13, 0xe2, 0xd5, 0x2b, 0x48, 0xd2, 0x83, 0x87, 0x89, 0x14, 0x1f, 0x32, 0xb7, 0xb8, 0x40,
	0x89, 0x41, 0x48, 0x8d, 0x8b, 0x15, 0x6c, 0x89, 0x10, 0x0f, 0x48, 0x0a, 0xe6, 0x19, 0x29, 0x5e,
	0x24, 0x1e, 0x58, 0x9d, 0x06, 0x17, 0xbb, 0x0b, 0xc4, 0x5d, 0x04, 0x54, 0x26, 0xb1, 0x81, 0xe3,
	0xc3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xc7, 0x2a, 0x83, 0xa0, 0x01, 0x00, 0x00,
}
