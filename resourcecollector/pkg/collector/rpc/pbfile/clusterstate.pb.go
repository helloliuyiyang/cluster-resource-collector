// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clusterstate.proto

package pbfile

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

type ClusterState struct {
	NameSpace            string   `protobuf:"bytes,1,opt,name=NameSpace,proto3" json:"NameSpace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	State                int64    `protobuf:"varint,3,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterState) Reset()         { *m = ClusterState{} }
func (m *ClusterState) String() string { return proto.CompactTextString(m) }
func (*ClusterState) ProtoMessage()    {}
func (*ClusterState) Descriptor() ([]byte, []int) {
	return fileDescriptor_clusterstate_fb3f89b746513e0c, []int{0}
}
func (m *ClusterState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterState.Unmarshal(m, b)
}
func (m *ClusterState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterState.Marshal(b, m, deterministic)
}
func (dst *ClusterState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterState.Merge(dst, src)
}
func (m *ClusterState) XXX_Size() int {
	return xxx_messageInfo_ClusterState.Size(m)
}
func (m *ClusterState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterState.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterState proto.InternalMessageInfo

func (m *ClusterState) GetNameSpace() string {
	if m != nil {
		return m.NameSpace
	}
	return ""
}

func (m *ClusterState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterState) GetState() int64 {
	if m != nil {
		return m.State
	}
	return 0
}

// message from ClusterController
type ReturnMessageClusterState struct {
	NameSpace            string   `protobuf:"bytes,1,opt,name=NameSpace,proto3" json:"NameSpace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ReturnCode           int64    `protobuf:"varint,3,opt,name=ReturnCode,proto3" json:"ReturnCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnMessageClusterState) Reset()         { *m = ReturnMessageClusterState{} }
func (m *ReturnMessageClusterState) String() string { return proto.CompactTextString(m) }
func (*ReturnMessageClusterState) ProtoMessage()    {}
func (*ReturnMessageClusterState) Descriptor() ([]byte, []int) {
	return fileDescriptor_clusterstate_fb3f89b746513e0c, []int{1}
}
func (m *ReturnMessageClusterState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnMessageClusterState.Unmarshal(m, b)
}
func (m *ReturnMessageClusterState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnMessageClusterState.Marshal(b, m, deterministic)
}
func (dst *ReturnMessageClusterState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnMessageClusterState.Merge(dst, src)
}
func (m *ReturnMessageClusterState) XXX_Size() int {
	return xxx_messageInfo_ReturnMessageClusterState.Size(m)
}
func (m *ReturnMessageClusterState) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnMessageClusterState.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnMessageClusterState proto.InternalMessageInfo

func (m *ReturnMessageClusterState) GetNameSpace() string {
	if m != nil {
		return m.NameSpace
	}
	return ""
}

func (m *ReturnMessageClusterState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReturnMessageClusterState) GetReturnCode() int64 {
	if m != nil {
		return m.ReturnCode
	}
	return 0
}

func init() {
	proto.RegisterType((*ClusterState)(nil), "pbfile.ClusterState")
	proto.RegisterType((*ReturnMessageClusterState)(nil), "pbfile.ReturnMessageClusterState")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceCollectorProtocolClient is the client API for ResourceCollectorProtocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceCollectorProtocolClient interface {
	UpdateClusterStatus(ctx context.Context, in *ClusterState, opts ...grpc.CallOption) (*ReturnMessageClusterState, error)
}

type resourceCollectorProtocolClient struct {
	cc *grpc.ClientConn
}

func NewResourceCollectorProtocolClient(cc *grpc.ClientConn) ResourceCollectorProtocolClient {
	return &resourceCollectorProtocolClient{cc}
}

func (c *resourceCollectorProtocolClient) UpdateClusterStatus(ctx context.Context, in *ClusterState, opts ...grpc.CallOption) (*ReturnMessageClusterState, error) {
	out := new(ReturnMessageClusterState)
	err := c.cc.Invoke(ctx, "/pbfile.ResourceCollectorProtocol/UpdateClusterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceCollectorProtocolServer is the server API for ResourceCollectorProtocol service.
type ResourceCollectorProtocolServer interface {
	UpdateClusterStatus(context.Context, *ClusterState) (*ReturnMessageClusterState, error)
}

func RegisterResourceCollectorProtocolServer(s *grpc.Server, srv ResourceCollectorProtocolServer) {
	s.RegisterService(&_ResourceCollectorProtocol_serviceDesc, srv)
}

func _ResourceCollectorProtocol_UpdateClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceCollectorProtocolServer).UpdateClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfile.ResourceCollectorProtocol/UpdateClusterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceCollectorProtocolServer).UpdateClusterStatus(ctx, req.(*ClusterState))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceCollectorProtocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbfile.ResourceCollectorProtocol",
	HandlerType: (*ResourceCollectorProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateClusterStatus",
			Handler:    _ResourceCollectorProtocol_UpdateClusterStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clusterstate.proto",
}

func init() { proto.RegisterFile("clusterstate.proto", fileDescriptor_clusterstate_fb3f89b746513e0c) }

var fileDescriptor_clusterstate_fb3f89b746513e0c = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x50, 0xbb, 0x6e, 0x83, 0x40,
	0x10, 0x0c, 0x21, 0x41, 0x62, 0x95, 0x6a, 0x43, 0x41, 0xa2, 0x28, 0xc2, 0x54, 0x54, 0x14, 0xf6,
	0x27, 0x50, 0xdb, 0x42, 0x20, 0xbb, 0x3f, 0x8e, 0xb5, 0x65, 0xe9, 0xf0, 0x9d, 0xee, 0xf1, 0xff,
	0x16, 0x77, 0x42, 0xa6, 0x71, 0xe5, 0x6e, 0x67, 0x46, 0xbb, 0x33, 0x3b, 0x80, 0x5c, 0x38, 0x63,
	0x49, 0x1b, 0xcb, 0x2c, 0xd5, 0x4a, 0x4b, 0x2b, 0x31, 0x51, 0xc3, 0xf9, 0x2a, 0xa8, 0x3c, 0xc1,
	0x57, 0x13, 0xd4, 0x7e, 0x56, 0xf1, 0x0f, 0xd2, 0x03, 0x9b, 0xa8, 0x57, 0x8c, 0x53, 0x1e, 0x15,
	0x51, 0x95, 0x76, 0x0f, 0x02, 0x11, 0x3e, 0x66, 0x90, 0xbf, 0x7b, 0xc1, 0xcf, 0x98, 0xc1, 0xa7,
	0x5f, 0xcd, 0xe3, 0x22, 0xaa, 0xe2, 0x2e, 0x80, 0x72, 0x82, 0x9f, 0x8e, 0xac, 0xd3, 0xb7, 0x3d,
	0x19, 0xc3, 0x2e, 0xf4, 0xa2, 0xc9, 0x3f, 0x40, 0x38, 0xd7, 0xc8, 0x71, 0x71, 0x5a, 0x31, 0x5b,
	0x6f, 0x67, 0xa4, 0xd3, 0x9c, 0x1a, 0x29, 0x04, 0x71, 0x2b, 0x75, 0x3b, 0x3f, 0xca, 0xa5, 0xc0,
	0x16, 0xbe, 0x8f, 0x6a, 0x64, 0x76, 0x1d, 0xc2, 0x19, 0xcc, 0xea, 0xd0, 0x41, 0xbd, 0xce, 0xf6,
	0xbb, 0x59, 0xd8, 0xa7, 0xf1, 0xcb, 0xb7, 0x21, 0xf1, 0x25, 0xee, 0xee, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x7c, 0x57, 0xa4, 0x7f, 0x5a, 0x01, 0x00, 0x00,
}
