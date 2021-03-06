// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cluster.proto

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

type ReturnMessageClusterProfile_CodeType int32

const (
	ReturnMessageClusterProfile_Error ReturnMessageClusterProfile_CodeType = 0
	ReturnMessageClusterProfile_OK    ReturnMessageClusterProfile_CodeType = 1
)

var ReturnMessageClusterProfile_CodeType_name = map[int32]string{
	0: "Error",
	1: "OK",
}
var ReturnMessageClusterProfile_CodeType_value = map[string]int32{
	"Error": 0,
	"OK":    1,
}

func (x ReturnMessageClusterProfile_CodeType) String() string {
	return proto.EnumName(ReturnMessageClusterProfile_CodeType_name, int32(x))
}
func (ReturnMessageClusterProfile_CodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cluster_66603368d6b0a134, []int{1, 0}
}

// message from ClusterController to ResourceCollector
type ClusterProfile struct {
	ClusterNameSpace     string                          `protobuf:"bytes,1,opt,name=ClusterNameSpace,proto3" json:"ClusterNameSpace,omitempty"`
	ClusterName          string                          `protobuf:"bytes,2,opt,name=ClusterName,proto3" json:"ClusterName,omitempty"`
	ClusterSpec          *ClusterProfile_ClusterSpecInfo `protobuf:"bytes,3,opt,name=ClusterSpec,proto3" json:"ClusterSpec,omitempty"`
	Status               string                          `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ClusterProfile) Reset()         { *m = ClusterProfile{} }
func (m *ClusterProfile) String() string { return proto.CompactTextString(m) }
func (*ClusterProfile) ProtoMessage()    {}
func (*ClusterProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_66603368d6b0a134, []int{0}
}
func (m *ClusterProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterProfile.Unmarshal(m, b)
}
func (m *ClusterProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterProfile.Marshal(b, m, deterministic)
}
func (dst *ClusterProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterProfile.Merge(dst, src)
}
func (m *ClusterProfile) XXX_Size() int {
	return xxx_messageInfo_ClusterProfile.Size(m)
}
func (m *ClusterProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterProfile proto.InternalMessageInfo

func (m *ClusterProfile) GetClusterNameSpace() string {
	if m != nil {
		return m.ClusterNameSpace
	}
	return ""
}

func (m *ClusterProfile) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterProfile) GetClusterSpec() *ClusterProfile_ClusterSpecInfo {
	if m != nil {
		return m.ClusterSpec
	}
	return nil
}

func (m *ClusterProfile) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ClusterProfile_ClusterSpecInfo struct {
	ClusterIpAddress     string                                          `protobuf:"bytes,1,opt,name=ClusterIpAddress,proto3" json:"ClusterIpAddress,omitempty"`
	GeoLocation          *ClusterProfile_ClusterSpecInfo_GeoLocationInfo `protobuf:"bytes,2,opt,name=GeoLocation,proto3" json:"GeoLocation,omitempty"`
	Region               *ClusterProfile_ClusterSpecInfo_RegionInfo      `protobuf:"bytes,3,opt,name=Region,proto3" json:"Region,omitempty"`
	Operator             *ClusterProfile_ClusterSpecInfo_OperatorInfo    `protobuf:"bytes,4,opt,name=Operator,proto3" json:"Operator,omitempty"`
	Flavor               []*ClusterProfile_ClusterSpecInfo_FlavorInfo    `protobuf:"bytes,5,rep,name=Flavor,proto3" json:"Flavor,omitempty"`
	Storage              []*ClusterProfile_ClusterSpecInfo_StorageInfo   `protobuf:"bytes,6,rep,name=Storage,proto3" json:"Storage,omitempty"`
	EipCapacity          int64                                           `protobuf:"varint,7,opt,name=EipCapacity,proto3" json:"EipCapacity,omitempty"`
	CPUCapacity          int64                                           `protobuf:"varint,8,opt,name=CPUCapacity,proto3" json:"CPUCapacity,omitempty"`
	MemCapacity          int64                                           `protobuf:"varint,9,opt,name=MemCapacity,proto3" json:"MemCapacity,omitempty"`
	ServerPrice          int64                                           `protobuf:"varint,10,opt,name=ServerPrice,proto3" json:"ServerPrice,omitempty"`
	HomeScheduler        string                                          `protobuf:"bytes,11,opt,name=HomeScheduler,proto3" json:"HomeScheduler,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *ClusterProfile_ClusterSpecInfo) Reset()         { *m = ClusterProfile_ClusterSpecInfo{} }
func (m *ClusterProfile_ClusterSpecInfo) String() string { return proto.CompactTextString(m) }
func (*ClusterProfile_ClusterSpecInfo) ProtoMessage()    {}
func (*ClusterProfile_ClusterSpecInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_66603368d6b0a134, []int{0, 0}
}
func (m *ClusterProfile_ClusterSpecInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo.Unmarshal(m, b)
}
func (m *ClusterProfile_ClusterSpecInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo.Marshal(b, m, deterministic)
}
func (dst *ClusterProfile_ClusterSpecInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo.Merge(dst, src)
}
func (m *ClusterProfile_ClusterSpecInfo) XXX_Size() int {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo.Size(m)
}
func (m *ClusterProfile_ClusterSpecInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterProfile_ClusterSpecInfo proto.InternalMessageInfo

func (m *ClusterProfile_ClusterSpecInfo) GetClusterIpAddress() string {
	if m != nil {
		return m.ClusterIpAddress
	}
	return ""
}

func (m *ClusterProfile_ClusterSpecInfo) GetGeoLocation() *ClusterProfile_ClusterSpecInfo_GeoLocationInfo {
	if m != nil {
		return m.GeoLocation
	}
	return nil
}

func (m *ClusterProfile_ClusterSpecInfo) GetRegion() *ClusterProfile_ClusterSpecInfo_RegionInfo {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *ClusterProfile_ClusterSpecInfo) GetOperator() *ClusterProfile_ClusterSpecInfo_OperatorInfo {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *ClusterProfile_ClusterSpecInfo) GetFlavor() []*ClusterProfile_ClusterSpecInfo_FlavorInfo {
	if m != nil {
		return m.Flavor
	}
	return nil
}

func (m *ClusterProfile_ClusterSpecInfo) GetStorage() []*ClusterProfile_ClusterSpecInfo_StorageInfo {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *ClusterProfile_ClusterSpecInfo) GetEipCapacity() int64 {
	if m != nil {
		return m.EipCapacity
	}
	return 0
}

func (m *ClusterProfile_ClusterSpecInfo) GetCPUCapacity() int64 {
	if m != nil {
		return m.CPUCapacity
	}
	return 0
}

func (m *ClusterProfile_ClusterSpecInfo) GetMemCapacity() int64 {
	if m != nil {
		return m.MemCapacity
	}
	return 0
}

func (m *ClusterProfile_ClusterSpecInfo) GetServerPrice() int64 {
	if m != nil {
		return m.ServerPrice
	}
	return 0
}

func (m *ClusterProfile_ClusterSpecInfo) GetHomeScheduler() string {
	if m != nil {
		return m.HomeScheduler
	}
	return ""
}

type ClusterProfile_ClusterSpecInfo_GeoLocationInfo struct {
	City                 string   `protobuf:"bytes,1,opt,name=City,proto3" json:"City,omitempty"`
	Province             string   `protobuf:"bytes,2,opt,name=Province,proto3" json:"Province,omitempty"`
	Area                 string   `protobuf:"bytes,3,opt,name=Area,proto3" json:"Area,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=Country,proto3" json:"Country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) Reset() {
	*m = ClusterProfile_ClusterSpecInfo_GeoLocationInfo{}
}
func (m *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) String() string {
	return proto.CompactTextString(m)
}
func (*ClusterProfile_ClusterSpecInfo_GeoLocationInfo) ProtoMessage() {}
func (*ClusterProfile_ClusterSpecInfo_GeoLocationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_66603368d6b0a134, []int{0, 0, 0}
}
func (m *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_GeoLocationInfo.Unmarshal(m, b)
}
func (m *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_GeoLocationInfo.Marshal(b, m, deterministic)
}
func (dst *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo_GeoLocationInfo.Merge(dst, src)
}
func (m *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) XXX_Size() int {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_GeoLocationInfo.Size(m)
}
func (m *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo_GeoLocationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterProfile_ClusterSpecInfo_GeoLocationInfo proto.InternalMessageInfo

func (m *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *ClusterProfile_ClusterSpecInfo_GeoLocationInfo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type ClusterProfile_ClusterSpecInfo_RegionInfo struct {
	Region               string   `protobuf:"bytes,1,opt,name=Region,proto3" json:"Region,omitempty"`
	AvailabilityZone     string   `protobuf:"bytes,2,opt,name=AvailabilityZone,proto3" json:"AvailabilityZone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterProfile_ClusterSpecInfo_RegionInfo) Reset() {
	*m = ClusterProfile_ClusterSpecInfo_RegionInfo{}
}
func (m *ClusterProfile_ClusterSpecInfo_RegionInfo) String() string {
	return proto.CompactTextString(m)
}
func (*ClusterProfile_ClusterSpecInfo_RegionInfo) ProtoMessage() {}
func (*ClusterProfile_ClusterSpecInfo_RegionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_66603368d6b0a134, []int{0, 0, 1}
}
func (m *ClusterProfile_ClusterSpecInfo_RegionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_RegionInfo.Unmarshal(m, b)
}
func (m *ClusterProfile_ClusterSpecInfo_RegionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_RegionInfo.Marshal(b, m, deterministic)
}
func (dst *ClusterProfile_ClusterSpecInfo_RegionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo_RegionInfo.Merge(dst, src)
}
func (m *ClusterProfile_ClusterSpecInfo_RegionInfo) XXX_Size() int {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_RegionInfo.Size(m)
}
func (m *ClusterProfile_ClusterSpecInfo_RegionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo_RegionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterProfile_ClusterSpecInfo_RegionInfo proto.InternalMessageInfo

func (m *ClusterProfile_ClusterSpecInfo_RegionInfo) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *ClusterProfile_ClusterSpecInfo_RegionInfo) GetAvailabilityZone() string {
	if m != nil {
		return m.AvailabilityZone
	}
	return ""
}

type ClusterProfile_ClusterSpecInfo_OperatorInfo struct {
	Operator             string   `protobuf:"bytes,1,opt,name=Operator,proto3" json:"Operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterProfile_ClusterSpecInfo_OperatorInfo) Reset() {
	*m = ClusterProfile_ClusterSpecInfo_OperatorInfo{}
}
func (m *ClusterProfile_ClusterSpecInfo_OperatorInfo) String() string {
	return proto.CompactTextString(m)
}
func (*ClusterProfile_ClusterSpecInfo_OperatorInfo) ProtoMessage() {}
func (*ClusterProfile_ClusterSpecInfo_OperatorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_66603368d6b0a134, []int{0, 0, 2}
}
func (m *ClusterProfile_ClusterSpecInfo_OperatorInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_OperatorInfo.Unmarshal(m, b)
}
func (m *ClusterProfile_ClusterSpecInfo_OperatorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_OperatorInfo.Marshal(b, m, deterministic)
}
func (dst *ClusterProfile_ClusterSpecInfo_OperatorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo_OperatorInfo.Merge(dst, src)
}
func (m *ClusterProfile_ClusterSpecInfo_OperatorInfo) XXX_Size() int {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_OperatorInfo.Size(m)
}
func (m *ClusterProfile_ClusterSpecInfo_OperatorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo_OperatorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterProfile_ClusterSpecInfo_OperatorInfo proto.InternalMessageInfo

func (m *ClusterProfile_ClusterSpecInfo_OperatorInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ClusterProfile_ClusterSpecInfo_FlavorInfo struct {
	FlavorID             string   `protobuf:"bytes,1,opt,name=FlavorID,proto3" json:"FlavorID,omitempty"`
	TotalCapacity        int64    `protobuf:"varint,2,opt,name=TotalCapacity,proto3" json:"TotalCapacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterProfile_ClusterSpecInfo_FlavorInfo) Reset() {
	*m = ClusterProfile_ClusterSpecInfo_FlavorInfo{}
}
func (m *ClusterProfile_ClusterSpecInfo_FlavorInfo) String() string {
	return proto.CompactTextString(m)
}
func (*ClusterProfile_ClusterSpecInfo_FlavorInfo) ProtoMessage() {}
func (*ClusterProfile_ClusterSpecInfo_FlavorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_66603368d6b0a134, []int{0, 0, 3}
}
func (m *ClusterProfile_ClusterSpecInfo_FlavorInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_FlavorInfo.Unmarshal(m, b)
}
func (m *ClusterProfile_ClusterSpecInfo_FlavorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_FlavorInfo.Marshal(b, m, deterministic)
}
func (dst *ClusterProfile_ClusterSpecInfo_FlavorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo_FlavorInfo.Merge(dst, src)
}
func (m *ClusterProfile_ClusterSpecInfo_FlavorInfo) XXX_Size() int {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_FlavorInfo.Size(m)
}
func (m *ClusterProfile_ClusterSpecInfo_FlavorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo_FlavorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterProfile_ClusterSpecInfo_FlavorInfo proto.InternalMessageInfo

func (m *ClusterProfile_ClusterSpecInfo_FlavorInfo) GetFlavorID() string {
	if m != nil {
		return m.FlavorID
	}
	return ""
}

func (m *ClusterProfile_ClusterSpecInfo_FlavorInfo) GetTotalCapacity() int64 {
	if m != nil {
		return m.TotalCapacity
	}
	return 0
}

type ClusterProfile_ClusterSpecInfo_StorageInfo struct {
	TypeID               string   `protobuf:"bytes,1,opt,name=TypeID,proto3" json:"TypeID,omitempty"`
	StorageCapacity      int64    `protobuf:"varint,2,opt,name=StorageCapacity,proto3" json:"StorageCapacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterProfile_ClusterSpecInfo_StorageInfo) Reset() {
	*m = ClusterProfile_ClusterSpecInfo_StorageInfo{}
}
func (m *ClusterProfile_ClusterSpecInfo_StorageInfo) String() string {
	return proto.CompactTextString(m)
}
func (*ClusterProfile_ClusterSpecInfo_StorageInfo) ProtoMessage() {}
func (*ClusterProfile_ClusterSpecInfo_StorageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_66603368d6b0a134, []int{0, 0, 4}
}
func (m *ClusterProfile_ClusterSpecInfo_StorageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_StorageInfo.Unmarshal(m, b)
}
func (m *ClusterProfile_ClusterSpecInfo_StorageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_StorageInfo.Marshal(b, m, deterministic)
}
func (dst *ClusterProfile_ClusterSpecInfo_StorageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo_StorageInfo.Merge(dst, src)
}
func (m *ClusterProfile_ClusterSpecInfo_StorageInfo) XXX_Size() int {
	return xxx_messageInfo_ClusterProfile_ClusterSpecInfo_StorageInfo.Size(m)
}
func (m *ClusterProfile_ClusterSpecInfo_StorageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterProfile_ClusterSpecInfo_StorageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterProfile_ClusterSpecInfo_StorageInfo proto.InternalMessageInfo

func (m *ClusterProfile_ClusterSpecInfo_StorageInfo) GetTypeID() string {
	if m != nil {
		return m.TypeID
	}
	return ""
}

func (m *ClusterProfile_ClusterSpecInfo_StorageInfo) GetStorageCapacity() int64 {
	if m != nil {
		return m.StorageCapacity
	}
	return 0
}

// Message from ResourceCollector, Cluster Controller should get response from ResourceCollector.
type ReturnMessageClusterProfile struct {
	ClusterNameSpace     string                               `protobuf:"bytes,1,opt,name=ClusterNameSpace,proto3" json:"ClusterNameSpace,omitempty"`
	ClusterName          string                               `protobuf:"bytes,2,opt,name=ClusterName,proto3" json:"ClusterName,omitempty"`
	ReturnCode           ReturnMessageClusterProfile_CodeType `protobuf:"varint,3,opt,name=ReturnCode,proto3,enum=pbfile.ReturnMessageClusterProfile_CodeType" json:"ReturnCode,omitempty"`
	Message              string                               `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ReturnMessageClusterProfile) Reset()         { *m = ReturnMessageClusterProfile{} }
func (m *ReturnMessageClusterProfile) String() string { return proto.CompactTextString(m) }
func (*ReturnMessageClusterProfile) ProtoMessage()    {}
func (*ReturnMessageClusterProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_66603368d6b0a134, []int{1}
}
func (m *ReturnMessageClusterProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnMessageClusterProfile.Unmarshal(m, b)
}
func (m *ReturnMessageClusterProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnMessageClusterProfile.Marshal(b, m, deterministic)
}
func (dst *ReturnMessageClusterProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnMessageClusterProfile.Merge(dst, src)
}
func (m *ReturnMessageClusterProfile) XXX_Size() int {
	return xxx_messageInfo_ReturnMessageClusterProfile.Size(m)
}
func (m *ReturnMessageClusterProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnMessageClusterProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnMessageClusterProfile proto.InternalMessageInfo

func (m *ReturnMessageClusterProfile) GetClusterNameSpace() string {
	if m != nil {
		return m.ClusterNameSpace
	}
	return ""
}

func (m *ReturnMessageClusterProfile) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ReturnMessageClusterProfile) GetReturnCode() ReturnMessageClusterProfile_CodeType {
	if m != nil {
		return m.ReturnCode
	}
	return ReturnMessageClusterProfile_Error
}

func (m *ReturnMessageClusterProfile) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ClusterProfile)(nil), "pbfile.ClusterProfile")
	proto.RegisterType((*ClusterProfile_ClusterSpecInfo)(nil), "pbfile.ClusterProfile.ClusterSpecInfo")
	proto.RegisterType((*ClusterProfile_ClusterSpecInfo_GeoLocationInfo)(nil), "pbfile.ClusterProfile.ClusterSpecInfo.GeoLocationInfo")
	proto.RegisterType((*ClusterProfile_ClusterSpecInfo_RegionInfo)(nil), "pbfile.ClusterProfile.ClusterSpecInfo.RegionInfo")
	proto.RegisterType((*ClusterProfile_ClusterSpecInfo_OperatorInfo)(nil), "pbfile.ClusterProfile.ClusterSpecInfo.OperatorInfo")
	proto.RegisterType((*ClusterProfile_ClusterSpecInfo_FlavorInfo)(nil), "pbfile.ClusterProfile.ClusterSpecInfo.FlavorInfo")
	proto.RegisterType((*ClusterProfile_ClusterSpecInfo_StorageInfo)(nil), "pbfile.ClusterProfile.ClusterSpecInfo.StorageInfo")
	proto.RegisterType((*ReturnMessageClusterProfile)(nil), "pbfile.ReturnMessageClusterProfile")
	proto.RegisterEnum("pbfile.ReturnMessageClusterProfile_CodeType", ReturnMessageClusterProfile_CodeType_name, ReturnMessageClusterProfile_CodeType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterProtocolClient is the client API for ClusterProtocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterProtocolClient interface {
	SendClusterProfile(ctx context.Context, in *ClusterProfile, opts ...grpc.CallOption) (*ReturnMessageClusterProfile, error)
}

type clusterProtocolClient struct {
	cc *grpc.ClientConn
}

func NewClusterProtocolClient(cc *grpc.ClientConn) ClusterProtocolClient {
	return &clusterProtocolClient{cc}
}

func (c *clusterProtocolClient) SendClusterProfile(ctx context.Context, in *ClusterProfile, opts ...grpc.CallOption) (*ReturnMessageClusterProfile, error) {
	out := new(ReturnMessageClusterProfile)
	err := c.cc.Invoke(ctx, "/pbfile.ClusterProtocol/SendClusterProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterProtocolServer is the server API for ClusterProtocol service.
type ClusterProtocolServer interface {
	SendClusterProfile(context.Context, *ClusterProfile) (*ReturnMessageClusterProfile, error)
}

func RegisterClusterProtocolServer(s *grpc.Server, srv ClusterProtocolServer) {
	s.RegisterService(&_ClusterProtocol_serviceDesc, srv)
}

func _ClusterProtocol_SendClusterProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterProtocolServer).SendClusterProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfile.ClusterProtocol/SendClusterProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterProtocolServer).SendClusterProfile(ctx, req.(*ClusterProfile))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterProtocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbfile.ClusterProtocol",
	HandlerType: (*ClusterProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendClusterProfile",
			Handler:    _ClusterProtocol_SendClusterProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor_cluster_66603368d6b0a134) }

var fileDescriptor_cluster_66603368d6b0a134 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0xad, 0x93, 0xd6, 0x49, 0xae, 0xd7, 0x26, 0xe8, 0xa1, 0x18, 0x8f, 0x41, 0xc8, 0xc6, 0x08,
	0x61, 0x18, 0xe6, 0xc2, 0xde, 0x43, 0xd6, 0xad, 0x61, 0x69, 0x13, 0xec, 0x0e, 0xc6, 0xde, 0x14,
	0x47, 0xe9, 0x0c, 0x8e, 0x65, 0x64, 0x25, 0x90, 0x5f, 0xd8, 0xbf, 0xed, 0x67, 0xf6, 0x05, 0x43,
	0x92, 0xad, 0x2a, 0xee, 0x18, 0xd9, 0x43, 0xdf, 0x74, 0x4f, 0xce, 0x3d, 0xf7, 0x5e, 0xdf, 0x23,
	0x05, 0xce, 0xe3, 0x74, 0x5b, 0x70, 0xc2, 0xfc, 0x9c, 0x51, 0x4e, 0x91, 0x9d, 0x2f, 0xd7, 0x49,
	0x4a, 0x06, 0xbf, 0x3a, 0x70, 0x31, 0x51, 0xbf, 0x2c, 0x18, 0x15, 0x10, 0x1a, 0x41, 0xaf, 0x44,
	0xee, 0xf0, 0x86, 0x44, 0x39, 0x8e, 0x89, 0x6b, 0xf5, 0xad, 0x61, 0x27, 0x7c, 0x82, 0xa3, 0x3e,
	0x38, 0x06, 0xe6, 0x36, 0x24, 0xcd, 0x84, 0xd0, 0x8d, 0x66, 0x44, 0x39, 0x89, 0xdd, 0x66, 0xdf,
	0x1a, 0x3a, 0xc1, 0x5b, 0x5f, 0x95, 0xf7, 0x0f, 0x4b, 0xfb, 0x06, 0x73, 0x9a, 0xad, 0x69, 0x68,
	0xa6, 0xa2, 0x4b, 0xb0, 0x23, 0x8e, 0xf9, 0xb6, 0x70, 0x4f, 0x65, 0x99, 0x32, 0xf2, 0x7e, 0xb6,
	0xa1, 0x5b, 0x4b, 0x34, 0x66, 0x98, 0xe6, 0xe3, 0xd5, 0x8a, 0x91, 0xa2, 0xa8, 0xcd, 0xa0, 0x71,
	0xf4, 0x0d, 0x9c, 0xcf, 0x84, 0xce, 0x68, 0x8c, 0x79, 0x42, 0x33, 0x39, 0x83, 0x13, 0x7c, 0x38,
	0xae, 0x43, 0xdf, 0xc8, 0x54, 0x1d, 0x1b, 0x00, 0x9a, 0x82, 0x1d, 0x92, 0x07, 0x21, 0xaa, 0xc6,
	0x7e, 0x7f, 0xa4, 0xa8, 0x4a, 0x92, 0x7a, 0xa5, 0x00, 0x9a, 0x43, 0x7b, 0x9e, 0x13, 0x86, 0x39,
	0x65, 0x72, 0x7c, 0x27, 0xb8, 0x3a, 0x52, 0xac, 0x4a, 0x93, 0x72, 0x5a, 0x44, 0xf4, 0xf6, 0x29,
	0xc5, 0x3b, 0xca, 0xdc, 0xb3, 0x7e, 0xf3, 0x3f, 0x7a, 0x53, 0x49, 0xaa, 0x37, 0x75, 0x46, 0x33,
	0x68, 0x45, 0x9c, 0x32, 0xfc, 0x40, 0x5c, 0x5b, 0x6a, 0x05, 0x47, 0x6a, 0x95, 0x59, 0x52, 0xac,
	0x92, 0x10, 0x96, 0xba, 0x4e, 0xf2, 0x09, 0xce, 0x71, 0x9c, 0xf0, 0xbd, 0xdb, 0xea, 0x5b, 0xc3,
	0x66, 0x68, 0x42, 0xd2, 0x74, 0x8b, 0xaf, 0x9a, 0xd1, 0x56, 0x0c, 0x03, 0x12, 0x8c, 0x5b, 0xb2,
	0xd1, 0x8c, 0x8e, 0x62, 0x18, 0x90, 0x60, 0x44, 0x84, 0xed, 0x44, 0x6f, 0x49, 0x4c, 0x5c, 0x50,
	0x0c, 0x03, 0x42, 0x6f, 0xe0, 0xfc, 0x86, 0x6e, 0x48, 0x14, 0xff, 0x20, 0xab, 0x6d, 0x4a, 0x98,
	0xeb, 0x48, 0xff, 0x1c, 0x82, 0x1e, 0x85, 0x6e, 0xcd, 0x02, 0x08, 0xc1, 0xe9, 0x44, 0x54, 0x55,
	0x7e, 0x93, 0x67, 0xe4, 0x41, 0x7b, 0xc1, 0xe8, 0x2e, 0xc9, 0xe2, 0xea, 0x92, 0xe8, 0x58, 0xf0,
	0xc7, 0x8c, 0x60, 0xe9, 0x91, 0x4e, 0x28, 0xcf, 0xc8, 0x85, 0xd6, 0x84, 0x6e, 0x33, 0xce, 0xf6,
	0xa5, 0xd9, 0xab, 0xd0, 0x5b, 0x00, 0x3c, 0xda, 0x43, 0xdc, 0x89, 0xd2, 0x61, 0xaa, 0x5a, 0x65,
	0x97, 0x11, 0xf4, 0xc6, 0x3b, 0x9c, 0xa4, 0x78, 0x99, 0xa4, 0x09, 0xdf, 0x7f, 0xa7, 0x59, 0x55,
	0xf7, 0x09, 0xee, 0x8d, 0xe0, 0x85, 0xe9, 0x11, 0xd1, 0xab, 0xb6, 0x9a, 0x52, 0xd5, 0xb1, 0x77,
	0x07, 0xf0, 0x68, 0x00, 0xc1, 0x2c, 0xa3, 0x8f, 0x15, 0xb3, 0x8a, 0xc5, 0xe7, 0xbb, 0xa7, 0x1c,
	0xa7, 0x7a, 0x09, 0x0d, 0xf9, 0x89, 0x0f, 0x41, 0x6f, 0x0e, 0x8e, 0x61, 0x02, 0x31, 0xce, 0xfd,
	0x3e, 0x27, 0x5a, 0xae, 0x8c, 0xd0, 0x10, 0xba, 0x25, 0xad, 0x26, 0x57, 0x87, 0x07, 0xbf, 0x2d,
	0x78, 0x19, 0x12, 0xbe, 0x65, 0xd9, 0x2d, 0x29, 0x0a, 0xf1, 0xcb, 0x73, 0x3e, 0x6e, 0x33, 0xb1,
	0x0c, 0x51, 0x6c, 0x42, 0x57, 0x44, 0x2e, 0xf0, 0x22, 0x78, 0x57, 0x99, 0xff, 0x1f, 0x6d, 0xf8,
	0x82, 0x2f, 0x66, 0x0b, 0x8d, 0x7c, 0xb1, 0xf4, 0x92, 0x5d, 0x2d, 0xbd, 0x0c, 0x07, 0xaf, 0xa0,
	0x5d, 0x65, 0xa0, 0x0e, 0x9c, 0x5d, 0x33, 0x46, 0x59, 0xef, 0x04, 0xd9, 0xd0, 0x98, 0x7f, 0xe9,
	0x59, 0xc1, 0x5a, 0x3f, 0x80, 0x0b, 0xf1, 0xb8, 0xc7, 0x34, 0x45, 0x11, 0xa0, 0x88, 0x64, 0xab,
	0xda, 0xf4, 0x97, 0x7f, 0xbf, 0x98, 0xde, 0xeb, 0x23, 0x7a, 0x1e, 0x9c, 0x2c, 0x6d, 0xf9, 0xdf,
	0x71, 0xf5, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x28, 0xe6, 0x59, 0x4c, 0x06, 0x00, 0x00,
}
