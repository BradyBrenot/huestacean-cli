// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package libhuestacean

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

type GetDeviceProvidersRequest struct {
	// optional: if set, only return the given providers
	LightProviderId      []uint32 `protobuf:"varint,1,rep,packed,name=light_provider_id,json=lightProviderId,proto3" json:"light_provider_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceProvidersRequest) Reset()         { *m = GetDeviceProvidersRequest{} }
func (m *GetDeviceProvidersRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProvidersRequest) ProtoMessage()    {}
func (*GetDeviceProvidersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{0}
}
func (m *GetDeviceProvidersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProvidersRequest.Unmarshal(m, b)
}
func (m *GetDeviceProvidersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProvidersRequest.Marshal(b, m, deterministic)
}
func (dst *GetDeviceProvidersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProvidersRequest.Merge(dst, src)
}
func (m *GetDeviceProvidersRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProvidersRequest.Size(m)
}
func (m *GetDeviceProvidersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProvidersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProvidersRequest proto.InternalMessageInfo

func (m *GetDeviceProvidersRequest) GetLightProviderId() []uint32 {
	if m != nil {
		return m.LightProviderId
	}
	return nil
}

type GetDeviceProvidersResponse struct {
	Providers            map[uint32]*DeviceProvider `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetDeviceProvidersResponse) Reset()         { *m = GetDeviceProvidersResponse{} }
func (m *GetDeviceProvidersResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProvidersResponse) ProtoMessage()    {}
func (*GetDeviceProvidersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{1}
}
func (m *GetDeviceProvidersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProvidersResponse.Unmarshal(m, b)
}
func (m *GetDeviceProvidersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProvidersResponse.Marshal(b, m, deterministic)
}
func (dst *GetDeviceProvidersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProvidersResponse.Merge(dst, src)
}
func (m *GetDeviceProvidersResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProvidersResponse.Size(m)
}
func (m *GetDeviceProvidersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProvidersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProvidersResponse proto.InternalMessageInfo

func (m *GetDeviceProvidersResponse) GetProviders() map[uint32]*DeviceProvider {
	if m != nil {
		return m.Providers
	}
	return nil
}

type GetDevicesRequest struct {
	// optional: if set, only return *this* device
	DeviceId []uint32 `protobuf:"varint,1,rep,packed,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// optional: if set, only return devices for *this* provider
	ProviderId           []uint32 `protobuf:"varint,2,rep,packed,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDevicesRequest) Reset()         { *m = GetDevicesRequest{} }
func (m *GetDevicesRequest) String() string { return proto.CompactTextString(m) }
func (*GetDevicesRequest) ProtoMessage()    {}
func (*GetDevicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{2}
}
func (m *GetDevicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDevicesRequest.Unmarshal(m, b)
}
func (m *GetDevicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDevicesRequest.Marshal(b, m, deterministic)
}
func (dst *GetDevicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDevicesRequest.Merge(dst, src)
}
func (m *GetDevicesRequest) XXX_Size() int {
	return xxx_messageInfo_GetDevicesRequest.Size(m)
}
func (m *GetDevicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDevicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDevicesRequest proto.InternalMessageInfo

func (m *GetDevicesRequest) GetDeviceId() []uint32 {
	if m != nil {
		return m.DeviceId
	}
	return nil
}

func (m *GetDevicesRequest) GetProviderId() []uint32 {
	if m != nil {
		return m.ProviderId
	}
	return nil
}

type GetDevicesResponse struct {
	Devices              map[uint32]*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetDevicesResponse) Reset()         { *m = GetDevicesResponse{} }
func (m *GetDevicesResponse) String() string { return proto.CompactTextString(m) }
func (*GetDevicesResponse) ProtoMessage()    {}
func (*GetDevicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{3}
}
func (m *GetDevicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDevicesResponse.Unmarshal(m, b)
}
func (m *GetDevicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDevicesResponse.Marshal(b, m, deterministic)
}
func (dst *GetDevicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDevicesResponse.Merge(dst, src)
}
func (m *GetDevicesResponse) XXX_Size() int {
	return xxx_messageInfo_GetDevicesResponse.Size(m)
}
func (m *GetDevicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDevicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDevicesResponse proto.InternalMessageInfo

func (m *GetDevicesResponse) GetDevices() map[uint32]*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type GetRoomsRequest struct {
	// optional: if set, only return *this* room
	RoomId               []uint32 `protobuf:"varint,1,rep,packed,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoomsRequest) Reset()         { *m = GetRoomsRequest{} }
func (m *GetRoomsRequest) String() string { return proto.CompactTextString(m) }
func (*GetRoomsRequest) ProtoMessage()    {}
func (*GetRoomsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{4}
}
func (m *GetRoomsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoomsRequest.Unmarshal(m, b)
}
func (m *GetRoomsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoomsRequest.Marshal(b, m, deterministic)
}
func (dst *GetRoomsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoomsRequest.Merge(dst, src)
}
func (m *GetRoomsRequest) XXX_Size() int {
	return xxx_messageInfo_GetRoomsRequest.Size(m)
}
func (m *GetRoomsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoomsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoomsRequest proto.InternalMessageInfo

func (m *GetRoomsRequest) GetRoomId() []uint32 {
	if m != nil {
		return m.RoomId
	}
	return nil
}

type GetRoomsResponse struct {
	Rooms                map[uint32]*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetRoomsResponse) Reset()         { *m = GetRoomsResponse{} }
func (m *GetRoomsResponse) String() string { return proto.CompactTextString(m) }
func (*GetRoomsResponse) ProtoMessage()    {}
func (*GetRoomsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{5}
}
func (m *GetRoomsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoomsResponse.Unmarshal(m, b)
}
func (m *GetRoomsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoomsResponse.Marshal(b, m, deterministic)
}
func (dst *GetRoomsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoomsResponse.Merge(dst, src)
}
func (m *GetRoomsResponse) XXX_Size() int {
	return xxx_messageInfo_GetRoomsResponse.Size(m)
}
func (m *GetRoomsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoomsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoomsResponse proto.InternalMessageInfo

func (m *GetRoomsResponse) GetRooms() map[uint32]*Room {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type GetDeviceProviderArchetypesRequest struct {
	// optional: if set, only return *this* archetype
	ArchetypeId          []uint32 `protobuf:"varint,1,rep,packed,name=archetype_id,json=archetypeId,proto3" json:"archetype_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceProviderArchetypesRequest) Reset()         { *m = GetDeviceProviderArchetypesRequest{} }
func (m *GetDeviceProviderArchetypesRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProviderArchetypesRequest) ProtoMessage()    {}
func (*GetDeviceProviderArchetypesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{6}
}
func (m *GetDeviceProviderArchetypesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProviderArchetypesRequest.Unmarshal(m, b)
}
func (m *GetDeviceProviderArchetypesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProviderArchetypesRequest.Marshal(b, m, deterministic)
}
func (dst *GetDeviceProviderArchetypesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProviderArchetypesRequest.Merge(dst, src)
}
func (m *GetDeviceProviderArchetypesRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProviderArchetypesRequest.Size(m)
}
func (m *GetDeviceProviderArchetypesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProviderArchetypesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProviderArchetypesRequest proto.InternalMessageInfo

func (m *GetDeviceProviderArchetypesRequest) GetArchetypeId() []uint32 {
	if m != nil {
		return m.ArchetypeId
	}
	return nil
}

type GetDeviceProviderArchetypesResponse struct {
	ProviderArchetypes   map[uint32]*DeviceProviderArchetype `protobuf:"bytes,1,rep,name=provider_archetypes,json=providerArchetypes,proto3" json:"provider_archetypes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *GetDeviceProviderArchetypesResponse) Reset()         { *m = GetDeviceProviderArchetypesResponse{} }
func (m *GetDeviceProviderArchetypesResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProviderArchetypesResponse) ProtoMessage()    {}
func (*GetDeviceProviderArchetypesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{7}
}
func (m *GetDeviceProviderArchetypesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProviderArchetypesResponse.Unmarshal(m, b)
}
func (m *GetDeviceProviderArchetypesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProviderArchetypesResponse.Marshal(b, m, deterministic)
}
func (dst *GetDeviceProviderArchetypesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProviderArchetypesResponse.Merge(dst, src)
}
func (m *GetDeviceProviderArchetypesResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProviderArchetypesResponse.Size(m)
}
func (m *GetDeviceProviderArchetypesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProviderArchetypesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProviderArchetypesResponse proto.InternalMessageInfo

func (m *GetDeviceProviderArchetypesResponse) GetProviderArchetypes() map[uint32]*DeviceProviderArchetype {
	if m != nil {
		return m.ProviderArchetypes
	}
	return nil
}

type GetDeviceArchetypesRequest struct {
	// optional: if set, only return *this* archetype
	ArchetypeId          []uint32 `protobuf:"varint,1,rep,packed,name=archetype_id,json=archetypeId,proto3" json:"archetype_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceArchetypesRequest) Reset()         { *m = GetDeviceArchetypesRequest{} }
func (m *GetDeviceArchetypesRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceArchetypesRequest) ProtoMessage()    {}
func (*GetDeviceArchetypesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{8}
}
func (m *GetDeviceArchetypesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceArchetypesRequest.Unmarshal(m, b)
}
func (m *GetDeviceArchetypesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceArchetypesRequest.Marshal(b, m, deterministic)
}
func (dst *GetDeviceArchetypesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceArchetypesRequest.Merge(dst, src)
}
func (m *GetDeviceArchetypesRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceArchetypesRequest.Size(m)
}
func (m *GetDeviceArchetypesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceArchetypesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceArchetypesRequest proto.InternalMessageInfo

func (m *GetDeviceArchetypesRequest) GetArchetypeId() []uint32 {
	if m != nil {
		return m.ArchetypeId
	}
	return nil
}

type GetDeviceArchetypesResponse struct {
	DeviceArchetypes     map[uint32]*DeviceArchetype `protobuf:"bytes,1,rep,name=device_archetypes,json=deviceArchetypes,proto3" json:"device_archetypes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *GetDeviceArchetypesResponse) Reset()         { *m = GetDeviceArchetypesResponse{} }
func (m *GetDeviceArchetypesResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceArchetypesResponse) ProtoMessage()    {}
func (*GetDeviceArchetypesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{9}
}
func (m *GetDeviceArchetypesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceArchetypesResponse.Unmarshal(m, b)
}
func (m *GetDeviceArchetypesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceArchetypesResponse.Marshal(b, m, deterministic)
}
func (dst *GetDeviceArchetypesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceArchetypesResponse.Merge(dst, src)
}
func (m *GetDeviceArchetypesResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceArchetypesResponse.Size(m)
}
func (m *GetDeviceArchetypesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceArchetypesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceArchetypesResponse proto.InternalMessageInfo

func (m *GetDeviceArchetypesResponse) GetDeviceArchetypes() map[uint32]*DeviceArchetype {
	if m != nil {
		return m.DeviceArchetypes
	}
	return nil
}

type LinkRequest struct {
	DeviceId             uint32   `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkRequest) Reset()         { *m = LinkRequest{} }
func (m *LinkRequest) String() string { return proto.CompactTextString(m) }
func (*LinkRequest) ProtoMessage()    {}
func (*LinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{10}
}
func (m *LinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkRequest.Unmarshal(m, b)
}
func (m *LinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkRequest.Marshal(b, m, deterministic)
}
func (dst *LinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkRequest.Merge(dst, src)
}
func (m *LinkRequest) XXX_Size() int {
	return xxx_messageInfo_LinkRequest.Size(m)
}
func (m *LinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LinkRequest proto.InternalMessageInfo

func (m *LinkRequest) GetDeviceId() uint32 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

type LinkResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkResponse) Reset()         { *m = LinkResponse{} }
func (m *LinkResponse) String() string { return proto.CompactTextString(m) }
func (*LinkResponse) ProtoMessage()    {}
func (*LinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_ef87878282a64c8f, []int{11}
}
func (m *LinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkResponse.Unmarshal(m, b)
}
func (m *LinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkResponse.Marshal(b, m, deterministic)
}
func (dst *LinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkResponse.Merge(dst, src)
}
func (m *LinkResponse) XXX_Size() int {
	return xxx_messageInfo_LinkResponse.Size(m)
}
func (m *LinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LinkResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetDeviceProvidersRequest)(nil), "libhuestacean.GetDeviceProvidersRequest")
	proto.RegisterType((*GetDeviceProvidersResponse)(nil), "libhuestacean.GetDeviceProvidersResponse")
	proto.RegisterMapType((map[uint32]*DeviceProvider)(nil), "libhuestacean.GetDeviceProvidersResponse.ProvidersEntry")
	proto.RegisterType((*GetDevicesRequest)(nil), "libhuestacean.GetDevicesRequest")
	proto.RegisterType((*GetDevicesResponse)(nil), "libhuestacean.GetDevicesResponse")
	proto.RegisterMapType((map[uint32]*Device)(nil), "libhuestacean.GetDevicesResponse.DevicesEntry")
	proto.RegisterType((*GetRoomsRequest)(nil), "libhuestacean.GetRoomsRequest")
	proto.RegisterType((*GetRoomsResponse)(nil), "libhuestacean.GetRoomsResponse")
	proto.RegisterMapType((map[uint32]*Room)(nil), "libhuestacean.GetRoomsResponse.RoomsEntry")
	proto.RegisterType((*GetDeviceProviderArchetypesRequest)(nil), "libhuestacean.GetDeviceProviderArchetypesRequest")
	proto.RegisterType((*GetDeviceProviderArchetypesResponse)(nil), "libhuestacean.GetDeviceProviderArchetypesResponse")
	proto.RegisterMapType((map[uint32]*DeviceProviderArchetype)(nil), "libhuestacean.GetDeviceProviderArchetypesResponse.ProviderArchetypesEntry")
	proto.RegisterType((*GetDeviceArchetypesRequest)(nil), "libhuestacean.GetDeviceArchetypesRequest")
	proto.RegisterType((*GetDeviceArchetypesResponse)(nil), "libhuestacean.GetDeviceArchetypesResponse")
	proto.RegisterMapType((map[uint32]*DeviceArchetype)(nil), "libhuestacean.GetDeviceArchetypesResponse.DeviceArchetypesEntry")
	proto.RegisterType((*LinkRequest)(nil), "libhuestacean.LinkRequest")
	proto.RegisterType((*LinkResponse)(nil), "libhuestacean.LinkResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HuestaceanServerClient is the client API for HuestaceanServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HuestaceanServerClient interface {
	GetDeviceProviders(ctx context.Context, in *GetDeviceProvidersRequest, opts ...grpc.CallOption) (*GetDeviceProvidersResponse, error)
	GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (*GetDevicesResponse, error)
	GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error)
	GetDeviceProviderArchetypes(ctx context.Context, in *GetDeviceProviderArchetypesRequest, opts ...grpc.CallOption) (*GetDeviceProviderArchetypesResponse, error)
	GetDeviceArchetypes(ctx context.Context, in *GetDeviceArchetypesRequest, opts ...grpc.CallOption) (*GetDeviceArchetypesResponse, error)
	Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*LinkResponse, error)
}

type huestaceanServerClient struct {
	cc *grpc.ClientConn
}

func NewHuestaceanServerClient(cc *grpc.ClientConn) HuestaceanServerClient {
	return &huestaceanServerClient{cc}
}

func (c *huestaceanServerClient) GetDeviceProviders(ctx context.Context, in *GetDeviceProvidersRequest, opts ...grpc.CallOption) (*GetDeviceProvidersResponse, error) {
	out := new(GetDeviceProvidersResponse)
	err := c.cc.Invoke(ctx, "/libhuestacean.HuestaceanServer/GetDeviceProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huestaceanServerClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (*GetDevicesResponse, error) {
	out := new(GetDevicesResponse)
	err := c.cc.Invoke(ctx, "/libhuestacean.HuestaceanServer/GetDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huestaceanServerClient) GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsResponse, error) {
	out := new(GetRoomsResponse)
	err := c.cc.Invoke(ctx, "/libhuestacean.HuestaceanServer/GetRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huestaceanServerClient) GetDeviceProviderArchetypes(ctx context.Context, in *GetDeviceProviderArchetypesRequest, opts ...grpc.CallOption) (*GetDeviceProviderArchetypesResponse, error) {
	out := new(GetDeviceProviderArchetypesResponse)
	err := c.cc.Invoke(ctx, "/libhuestacean.HuestaceanServer/GetDeviceProviderArchetypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huestaceanServerClient) GetDeviceArchetypes(ctx context.Context, in *GetDeviceArchetypesRequest, opts ...grpc.CallOption) (*GetDeviceArchetypesResponse, error) {
	out := new(GetDeviceArchetypesResponse)
	err := c.cc.Invoke(ctx, "/libhuestacean.HuestaceanServer/GetDeviceArchetypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huestaceanServerClient) Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*LinkResponse, error) {
	out := new(LinkResponse)
	err := c.cc.Invoke(ctx, "/libhuestacean.HuestaceanServer/Link", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HuestaceanServerServer is the server API for HuestaceanServer service.
type HuestaceanServerServer interface {
	GetDeviceProviders(context.Context, *GetDeviceProvidersRequest) (*GetDeviceProvidersResponse, error)
	GetDevices(context.Context, *GetDevicesRequest) (*GetDevicesResponse, error)
	GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsResponse, error)
	GetDeviceProviderArchetypes(context.Context, *GetDeviceProviderArchetypesRequest) (*GetDeviceProviderArchetypesResponse, error)
	GetDeviceArchetypes(context.Context, *GetDeviceArchetypesRequest) (*GetDeviceArchetypesResponse, error)
	Link(context.Context, *LinkRequest) (*LinkResponse, error)
}

func RegisterHuestaceanServerServer(s *grpc.Server, srv HuestaceanServerServer) {
	s.RegisterService(&_HuestaceanServer_serviceDesc, srv)
}

func _HuestaceanServer_GetDeviceProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuestaceanServerServer).GetDeviceProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libhuestacean.HuestaceanServer/GetDeviceProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuestaceanServerServer).GetDeviceProviders(ctx, req.(*GetDeviceProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuestaceanServer_GetDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuestaceanServerServer).GetDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libhuestacean.HuestaceanServer/GetDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuestaceanServerServer).GetDevices(ctx, req.(*GetDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuestaceanServer_GetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuestaceanServerServer).GetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libhuestacean.HuestaceanServer/GetRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuestaceanServerServer).GetRooms(ctx, req.(*GetRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuestaceanServer_GetDeviceProviderArchetypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceProviderArchetypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuestaceanServerServer).GetDeviceProviderArchetypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libhuestacean.HuestaceanServer/GetDeviceProviderArchetypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuestaceanServerServer).GetDeviceProviderArchetypes(ctx, req.(*GetDeviceProviderArchetypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuestaceanServer_GetDeviceArchetypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceArchetypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuestaceanServerServer).GetDeviceArchetypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libhuestacean.HuestaceanServer/GetDeviceArchetypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuestaceanServerServer).GetDeviceArchetypes(ctx, req.(*GetDeviceArchetypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuestaceanServer_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuestaceanServerServer).Link(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libhuestacean.HuestaceanServer/Link",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuestaceanServerServer).Link(ctx, req.(*LinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HuestaceanServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "libhuestacean.HuestaceanServer",
	HandlerType: (*HuestaceanServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceProviders",
			Handler:    _HuestaceanServer_GetDeviceProviders_Handler,
		},
		{
			MethodName: "GetDevices",
			Handler:    _HuestaceanServer_GetDevices_Handler,
		},
		{
			MethodName: "GetRooms",
			Handler:    _HuestaceanServer_GetRooms_Handler,
		},
		{
			MethodName: "GetDeviceProviderArchetypes",
			Handler:    _HuestaceanServer_GetDeviceProviderArchetypes_Handler,
		},
		{
			MethodName: "GetDeviceArchetypes",
			Handler:    _HuestaceanServer_GetDeviceArchetypes_Handler,
		},
		{
			MethodName: "Link",
			Handler:    _HuestaceanServer_Link_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_rpc_ef87878282a64c8f) }

var fileDescriptor_rpc_ef87878282a64c8f = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xdf, 0x8a, 0xda, 0x40,
	0x14, 0xc6, 0x89, 0xfb, 0xff, 0xa8, 0xbb, 0x7a, 0x64, 0x59, 0x3b, 0xd2, 0x5d, 0x37, 0x85, 0xa2,
	0x16, 0x02, 0x75, 0x7b, 0xb1, 0x94, 0xc2, 0x6e, 0xa1, 0xc5, 0xb5, 0x7f, 0xa0, 0xa6, 0xd0, 0x9b,
	0x5e, 0x2c, 0xae, 0x19, 0x6a, 0x58, 0x35, 0x69, 0xcc, 0x0a, 0xd2, 0xdb, 0xbe, 0x40, 0x9f, 0xa1,
	0x4f, 0xd0, 0x47, 0x29, 0x7d, 0x92, 0xbe, 0x41, 0x49, 0x26, 0x99, 0x99, 0x98, 0x89, 0xba, 0xbd,
	0x4b, 0xe6, 0x9b, 0xf3, 0x4d, 0xce, 0xcf, 0x33, 0x1f, 0xc2, 0x9e, 0xe7, 0x0e, 0x0c, 0xd7, 0x73,
	0x7c, 0x07, 0x8b, 0x23, 0xfb, 0x66, 0x78, 0x47, 0xa7, 0x7e, 0x7f, 0x40, 0xfb, 0x13, 0x52, 0xb0,
	0xe8, 0xcc, 0x1e, 0x50, 0x26, 0x12, 0xf0, 0x1c, 0x67, 0xcc, 0x9e, 0xf5, 0x0e, 0x3c, 0xe8, 0x50,
	0xff, 0x55, 0x28, 0x7f, 0xf0, 0x9c, 0x99, 0x6d, 0x51, 0x6f, 0x6a, 0xd2, 0xaf, 0x41, 0x2d, 0xb6,
	0xa0, 0x3c, 0xb2, 0xbf, 0x0c, 0xfd, 0x6b, 0x37, 0x52, 0xae, 0x6d, 0xab, 0xaa, 0xd5, 0x37, 0x1a,
	0x45, 0xf3, 0x20, 0x14, 0xe2, 0x8a, 0xae, 0xa5, 0xff, 0xd6, 0x80, 0xa8, 0x9c, 0xa6, 0xae, 0x33,
	0x99, 0x52, 0xfc, 0x04, 0x7b, 0xb1, 0xc9, 0x34, 0xb4, 0xc8, 0xb7, 0xcf, 0x8d, 0xc4, 0x47, 0x1a,
	0xd9, 0xd5, 0x06, 0x5f, 0x79, 0x3d, 0xf1, 0xbd, 0xb9, 0x29, 0xac, 0xc8, 0x67, 0xd8, 0x4f, 0x8a,
	0x58, 0x82, 0x8d, 0x5b, 0x3a, 0xaf, 0x6a, 0x75, 0xad, 0x51, 0x34, 0x83, 0x47, 0x3c, 0x83, 0xad,
	0x59, 0x7f, 0x74, 0x47, 0xab, 0xb9, 0xba, 0xd6, 0xc8, 0xb7, 0x1f, 0x2e, 0x9c, 0x9b, 0x3c, 0xd4,
	0x64, 0x7b, 0x9f, 0xe7, 0xce, 0x35, 0xbd, 0x07, 0x65, 0xfe, 0x51, 0x1c, 0x4a, 0x0d, 0xf6, 0x18,
	0x4d, 0x01, 0x63, 0x97, 0x2d, 0x74, 0x2d, 0x3c, 0x81, 0xbc, 0xcc, 0x2a, 0x17, 0xca, 0xe0, 0x0a,
	0x4c, 0xbf, 0x34, 0x40, 0xd9, 0x33, 0xc2, 0x73, 0x05, 0x3b, 0xcc, 0x23, 0x86, 0x63, 0x64, 0xc1,
	0x11, 0x50, 0xa2, 0x77, 0x86, 0x24, 0x2e, 0x27, 0x3d, 0x28, 0xc8, 0x82, 0x02, 0xc7, 0x93, 0x24,
	0x8e, 0x43, 0x25, 0x0e, 0x19, 0x43, 0x0b, 0x0e, 0x3a, 0xd4, 0x37, 0x1d, 0x67, 0xcc, 0x21, 0x1c,
	0xc1, 0x4e, 0x30, 0x44, 0x02, 0xc1, 0x76, 0xf0, 0xda, 0xb5, 0xf4, 0x9f, 0x1a, 0x94, 0xc4, 0xe6,
	0xa8, 0xbb, 0x4b, 0xd8, 0x0a, 0xe4, 0xb8, 0xb7, 0x56, 0xba, 0xb7, 0xc4, 0x7e, 0x23, 0x7c, 0x63,
	0x7d, 0xb1, 0x42, 0xf2, 0x1e, 0x40, 0x2c, 0x2a, 0x7a, 0x6a, 0x26, 0x7b, 0xaa, 0x2c, 0x9c, 0x10,
	0xd4, 0xca, 0x1d, 0x75, 0x40, 0x4f, 0x4d, 0xdb, 0x4b, 0x6f, 0x30, 0xa4, 0xfe, 0xdc, 0x15, 0xbf,
	0xf4, 0x29, 0x14, 0xfa, 0xf1, 0xa2, 0xe8, 0x34, 0xcf, 0xd7, 0xba, 0x96, 0xfe, 0x23, 0x07, 0x8f,
	0x96, 0x3a, 0x45, 0x04, 0xbe, 0x41, 0x85, 0xcf, 0x05, 0xaf, 0x8f, 0x79, 0xbc, 0x59, 0x75, 0x11,
	0xd2, 0x86, 0x46, 0x5a, 0x62, 0xbc, 0xd0, 0x4d, 0x09, 0x64, 0x0c, 0x47, 0x19, 0xdb, 0x15, 0x24,
	0x5f, 0x24, 0x49, 0x3e, 0x5e, 0x7a, 0x59, 0xb8, 0x9d, 0x0c, 0xf7, 0x42, 0x0a, 0x82, 0xff, 0x82,
	0xfa, 0x57, 0x83, 0x9a, 0xd2, 0x21, 0x82, 0x39, 0x86, 0x72, 0x74, 0x03, 0x53, 0x28, 0x2f, 0xb3,
	0x50, 0x2a, 0x10, 0x2e, 0x0a, 0x0c, 0x60, 0xc9, 0x5a, 0x58, 0x26, 0x03, 0x38, 0x54, 0x6e, 0x55,
	0xc0, 0x7b, 0x96, 0x84, 0x77, 0xac, 0x84, 0xa7, 0x84, 0xd6, 0x82, 0xfc, 0x3b, 0x7b, 0x72, 0x9b,
	0x11, 0x32, 0x9a, 0x1c, 0x32, 0xfa, 0x3e, 0x14, 0xd8, 0x5e, 0xd6, 0x48, 0xfb, 0xcf, 0x26, 0x94,
	0xae, 0xf8, 0x29, 0x1f, 0xa9, 0x37, 0xa3, 0x1e, 0xda, 0x52, 0xce, 0xf0, 0x84, 0xc4, 0xc6, 0x1a,
	0x99, 0x1b, 0x7e, 0x01, 0x69, 0xae, 0x9d, 0xce, 0xd8, 0x03, 0x10, 0xf1, 0x84, 0xf5, 0x25, 0xc9,
	0xc5, 0xac, 0x4f, 0x57, 0x66, 0x1b, 0xbe, 0x85, 0xdd, 0x38, 0x15, 0xf0, 0x38, 0x33, 0x2e, 0x98,
	0xdd, 0xc9, 0x8a, 0x38, 0xc1, 0xef, 0xf2, 0x3c, 0xa5, 0x6f, 0x02, 0x3e, 0xbd, 0xcf, 0xfd, 0x63,
	0x67, 0xb6, 0xef, 0x7f, 0x65, 0x71, 0x04, 0x15, 0xc5, 0x38, 0x62, 0x73, 0x9d, 0x91, 0x65, 0xa7,
	0xb6, 0xd6, 0x9f, 0x6e, 0xbc, 0x80, 0xcd, 0x60, 0x48, 0x90, 0x2c, 0xd4, 0x48, 0x53, 0x46, 0x6a,
	0x4a, 0x8d, 0x19, 0xdc, 0x6c, 0x87, 0x7f, 0x10, 0xce, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x72,
	0xbd, 0x2b, 0x8c, 0x56, 0x08, 0x00, 0x00,
}
