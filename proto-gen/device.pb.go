// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

package libhuestacean

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeviceProvider_State int32

const (
	DeviceProvider_Disconnected DeviceProvider_State = 0
	DeviceProvider_Connecting   DeviceProvider_State = 1
	DeviceProvider_PendingLink  DeviceProvider_State = 2
	DeviceProvider_Connected    DeviceProvider_State = 3
)

var DeviceProvider_State_name = map[int32]string{
	0: "Disconnected",
	1: "Connecting",
	2: "PendingLink",
	3: "Connected",
}
var DeviceProvider_State_value = map[string]int32{
	"Disconnected": 0,
	"Connecting":   1,
	"PendingLink":  2,
	"Connected":    3,
}

func (x DeviceProvider_State) String() string {
	return proto.EnumName(DeviceProvider_State_name, int32(x))
}
func (DeviceProvider_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_device_b2ce4af9d5a0629f, []int{3, 0}
}

// Describes a type of light provider known by the daemon
type DeviceProviderArchetype struct {
	// Archetype's name, e.g. "Philips Hue bridge"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Whether the provider supports low latency light updates
	LowLatencySupported bool `protobuf:"varint,2,opt,name=low_latency_supported,json=lowLatencySupported,proto3" json:"low_latency_supported,omitempty"`
	// if > 0, how many low-latency lights can be present at once; if <= 0, no limit
	MaxLowLatency        uint32   `protobuf:"varint,3,opt,name=max_low_latency,json=maxLowLatency,proto3" json:"max_low_latency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceProviderArchetype) Reset()         { *m = DeviceProviderArchetype{} }
func (m *DeviceProviderArchetype) String() string { return proto.CompactTextString(m) }
func (*DeviceProviderArchetype) ProtoMessage()    {}
func (*DeviceProviderArchetype) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_b2ce4af9d5a0629f, []int{0}
}
func (m *DeviceProviderArchetype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProviderArchetype.Unmarshal(m, b)
}
func (m *DeviceProviderArchetype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProviderArchetype.Marshal(b, m, deterministic)
}
func (dst *DeviceProviderArchetype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProviderArchetype.Merge(dst, src)
}
func (m *DeviceProviderArchetype) XXX_Size() int {
	return xxx_messageInfo_DeviceProviderArchetype.Size(m)
}
func (m *DeviceProviderArchetype) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProviderArchetype.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProviderArchetype proto.InternalMessageInfo

func (m *DeviceProviderArchetype) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceProviderArchetype) GetLowLatencySupported() bool {
	if m != nil {
		return m.LowLatencySupported
	}
	return false
}

func (m *DeviceProviderArchetype) GetMaxLowLatency() uint32 {
	if m != nil {
		return m.MaxLowLatency
	}
	return 0
}

// A light contained within a device archetype, with its default location
type DeviceArchetypeLight struct {
	DefaultLocation      *LightLocation `protobuf:"bytes,1,opt,name=default_location,json=defaultLocation,proto3" json:"default_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeviceArchetypeLight) Reset()         { *m = DeviceArchetypeLight{} }
func (m *DeviceArchetypeLight) String() string { return proto.CompactTextString(m) }
func (*DeviceArchetypeLight) ProtoMessage()    {}
func (*DeviceArchetypeLight) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_b2ce4af9d5a0629f, []int{1}
}
func (m *DeviceArchetypeLight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceArchetypeLight.Unmarshal(m, b)
}
func (m *DeviceArchetypeLight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceArchetypeLight.Marshal(b, m, deterministic)
}
func (dst *DeviceArchetypeLight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceArchetypeLight.Merge(dst, src)
}
func (m *DeviceArchetypeLight) XXX_Size() int {
	return xxx_messageInfo_DeviceArchetypeLight.Size(m)
}
func (m *DeviceArchetypeLight) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceArchetypeLight.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceArchetypeLight proto.InternalMessageInfo

func (m *DeviceArchetypeLight) GetDefaultLocation() *LightLocation {
	if m != nil {
		return m.DefaultLocation
	}
	return nil
}

// Describes a type of device known by the daemon
type DeviceArchetype struct {
	// This device's lights, with their default location.
	// the key, an int, is unique only _within_ this device; e.g. two Devices may have the same Light0
	Lights               map[uint32]*DeviceArchetypeLight `protobuf:"bytes,2,rep,name=lights,proto3" json:"lights,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *DeviceArchetype) Reset()         { *m = DeviceArchetype{} }
func (m *DeviceArchetype) String() string { return proto.CompactTextString(m) }
func (*DeviceArchetype) ProtoMessage()    {}
func (*DeviceArchetype) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_b2ce4af9d5a0629f, []int{2}
}
func (m *DeviceArchetype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceArchetype.Unmarshal(m, b)
}
func (m *DeviceArchetype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceArchetype.Marshal(b, m, deterministic)
}
func (dst *DeviceArchetype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceArchetype.Merge(dst, src)
}
func (m *DeviceArchetype) XXX_Size() int {
	return xxx_messageInfo_DeviceArchetype.Size(m)
}
func (m *DeviceArchetype) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceArchetype.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceArchetype proto.InternalMessageInfo

func (m *DeviceArchetype) GetLights() map[uint32]*DeviceArchetypeLight {
	if m != nil {
		return m.Lights
	}
	return nil
}

// DeviceProvider is a service the daemon is using to connect to lights,
// e.g. it may be a single Hue bridge, or it may be the Razer Chroma SDK
type DeviceProvider struct {
	// What sort of light provider is this? This is the underlying tech / API being used.
	ArchetypeId uint32 `protobuf:"varint,1,opt,name=archetype_id,json=archetypeId,proto3" json:"archetype_id,omitempty"`
	// Name of the light provider. Disambiguates between providers of the same archetype.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Current state this DeviceProvider is in
	State                DeviceProvider_State `protobuf:"varint,3,opt,name=state,proto3,enum=libhuestacean.DeviceProvider_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeviceProvider) Reset()         { *m = DeviceProvider{} }
func (m *DeviceProvider) String() string { return proto.CompactTextString(m) }
func (*DeviceProvider) ProtoMessage()    {}
func (*DeviceProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_b2ce4af9d5a0629f, []int{3}
}
func (m *DeviceProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProvider.Unmarshal(m, b)
}
func (m *DeviceProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProvider.Marshal(b, m, deterministic)
}
func (dst *DeviceProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProvider.Merge(dst, src)
}
func (m *DeviceProvider) XXX_Size() int {
	return xxx_messageInfo_DeviceProvider.Size(m)
}
func (m *DeviceProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProvider.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProvider proto.InternalMessageInfo

func (m *DeviceProvider) GetArchetypeId() uint32 {
	if m != nil {
		return m.ArchetypeId
	}
	return 0
}

func (m *DeviceProvider) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceProvider) GetState() DeviceProvider_State {
	if m != nil {
		return m.State
	}
	return DeviceProvider_Disconnected
}

// Color for one light. Color is in HSLuv, *not* the common RGB HSL
type LightColor struct {
	H                    float64  `protobuf:"fixed64,1,opt,name=h,proto3" json:"h,omitempty"`
	S                    float64  `protobuf:"fixed64,2,opt,name=s,proto3" json:"s,omitempty"`
	L                    float64  `protobuf:"fixed64,3,opt,name=l,proto3" json:"l,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LightColor) Reset()         { *m = LightColor{} }
func (m *LightColor) String() string { return proto.CompactTextString(m) }
func (*LightColor) ProtoMessage()    {}
func (*LightColor) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_b2ce4af9d5a0629f, []int{4}
}
func (m *LightColor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightColor.Unmarshal(m, b)
}
func (m *LightColor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightColor.Marshal(b, m, deterministic)
}
func (dst *LightColor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightColor.Merge(dst, src)
}
func (m *LightColor) XXX_Size() int {
	return xxx_messageInfo_LightColor.Size(m)
}
func (m *LightColor) XXX_DiscardUnknown() {
	xxx_messageInfo_LightColor.DiscardUnknown(m)
}

var xxx_messageInfo_LightColor proto.InternalMessageInfo

func (m *LightColor) GetH() float64 {
	if m != nil {
		return m.H
	}
	return 0
}

func (m *LightColor) GetS() float64 {
	if m != nil {
		return m.S
	}
	return 0
}

func (m *LightColor) GetL() float64 {
	if m != nil {
		return m.L
	}
	return 0
}

// Devices are a thing in the real-world that emit light, and that Huestacean can control.
type Device struct {
	// What sort of light is this? This is the underlying tech / API being used.
	ArchetypeId uint32 `protobuf:"varint,1,opt,name=archetype_id,json=archetypeId,proto3" json:"archetype_id,omitempty"`
	// The device's friendly name, if it has one
	Name uint32 `protobuf:"varint,2,opt,name=name,proto3" json:"name,omitempty"`
	// All of this device's lights' colors
	Lights               map[uint32]*LightColor `protobuf:"bytes,3,rep,name=lights,proto3" json:"lights,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_b2ce4af9d5a0629f, []int{5}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (dst *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(dst, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetArchetypeId() uint32 {
	if m != nil {
		return m.ArchetypeId
	}
	return 0
}

func (m *Device) GetName() uint32 {
	if m != nil {
		return m.Name
	}
	return 0
}

func (m *Device) GetLights() map[uint32]*LightColor {
	if m != nil {
		return m.Lights
	}
	return nil
}

// The position of a light in the room is given as an axis-aligned box.
// This position is given as relative to its device.
type LightLocation struct {
	MinX                 float32  `protobuf:"fixed32,1,opt,name=min_x,json=minX,proto3" json:"min_x,omitempty"`
	MinY                 float32  `protobuf:"fixed32,2,opt,name=min_y,json=minY,proto3" json:"min_y,omitempty"`
	MinZ                 float32  `protobuf:"fixed32,3,opt,name=min_z,json=minZ,proto3" json:"min_z,omitempty"`
	MaxX                 float32  `protobuf:"fixed32,4,opt,name=max_x,json=maxX,proto3" json:"max_x,omitempty"`
	MaxY                 float32  `protobuf:"fixed32,5,opt,name=max_y,json=maxY,proto3" json:"max_y,omitempty"`
	MaxZ                 float32  `protobuf:"fixed32,6,opt,name=max_z,json=maxZ,proto3" json:"max_z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LightLocation) Reset()         { *m = LightLocation{} }
func (m *LightLocation) String() string { return proto.CompactTextString(m) }
func (*LightLocation) ProtoMessage()    {}
func (*LightLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_b2ce4af9d5a0629f, []int{6}
}
func (m *LightLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightLocation.Unmarshal(m, b)
}
func (m *LightLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightLocation.Marshal(b, m, deterministic)
}
func (dst *LightLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightLocation.Merge(dst, src)
}
func (m *LightLocation) XXX_Size() int {
	return xxx_messageInfo_LightLocation.Size(m)
}
func (m *LightLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_LightLocation.DiscardUnknown(m)
}

var xxx_messageInfo_LightLocation proto.InternalMessageInfo

func (m *LightLocation) GetMinX() float32 {
	if m != nil {
		return m.MinX
	}
	return 0
}

func (m *LightLocation) GetMinY() float32 {
	if m != nil {
		return m.MinY
	}
	return 0
}

func (m *LightLocation) GetMinZ() float32 {
	if m != nil {
		return m.MinZ
	}
	return 0
}

func (m *LightLocation) GetMaxX() float32 {
	if m != nil {
		return m.MaxX
	}
	return 0
}

func (m *LightLocation) GetMaxY() float32 {
	if m != nil {
		return m.MaxY
	}
	return 0
}

func (m *LightLocation) GetMaxZ() float32 {
	if m != nil {
		return m.MaxZ
	}
	return 0
}

// Devices do not emit light, their coordinates are given as a (3D) point
// and scaling factor that's applied uniformly to the contained lights
type DeviceLocation struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	ScaleX               float32  `protobuf:"fixed32,4,opt,name=scale_x,json=scaleX,proto3" json:"scale_x,omitempty"`
	ScaleY               float32  `protobuf:"fixed32,5,opt,name=scale_y,json=scaleY,proto3" json:"scale_y,omitempty"`
	ScaleZ               float32  `protobuf:"fixed32,6,opt,name=scale_z,json=scaleZ,proto3" json:"scale_z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceLocation) Reset()         { *m = DeviceLocation{} }
func (m *DeviceLocation) String() string { return proto.CompactTextString(m) }
func (*DeviceLocation) ProtoMessage()    {}
func (*DeviceLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_b2ce4af9d5a0629f, []int{7}
}
func (m *DeviceLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceLocation.Unmarshal(m, b)
}
func (m *DeviceLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceLocation.Marshal(b, m, deterministic)
}
func (dst *DeviceLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceLocation.Merge(dst, src)
}
func (m *DeviceLocation) XXX_Size() int {
	return xxx_messageInfo_DeviceLocation.Size(m)
}
func (m *DeviceLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceLocation.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceLocation proto.InternalMessageInfo

func (m *DeviceLocation) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *DeviceLocation) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *DeviceLocation) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *DeviceLocation) GetScaleX() float32 {
	if m != nil {
		return m.ScaleX
	}
	return 0
}

func (m *DeviceLocation) GetScaleY() float32 {
	if m != nil {
		return m.ScaleY
	}
	return 0
}

func (m *DeviceLocation) GetScaleZ() float32 {
	if m != nil {
		return m.ScaleZ
	}
	return 0
}

func init() {
	proto.RegisterType((*DeviceProviderArchetype)(nil), "libhuestacean.DeviceProviderArchetype")
	proto.RegisterType((*DeviceArchetypeLight)(nil), "libhuestacean.DeviceArchetypeLight")
	proto.RegisterType((*DeviceArchetype)(nil), "libhuestacean.DeviceArchetype")
	proto.RegisterMapType((map[uint32]*DeviceArchetypeLight)(nil), "libhuestacean.DeviceArchetype.LightsEntry")
	proto.RegisterType((*DeviceProvider)(nil), "libhuestacean.DeviceProvider")
	proto.RegisterType((*LightColor)(nil), "libhuestacean.LightColor")
	proto.RegisterType((*Device)(nil), "libhuestacean.Device")
	proto.RegisterMapType((map[uint32]*LightColor)(nil), "libhuestacean.Device.LightsEntry")
	proto.RegisterType((*LightLocation)(nil), "libhuestacean.LightLocation")
	proto.RegisterType((*DeviceLocation)(nil), "libhuestacean.DeviceLocation")
	proto.RegisterEnum("libhuestacean.DeviceProvider_State", DeviceProvider_State_name, DeviceProvider_State_value)
}

func init() { proto.RegisterFile("device.proto", fileDescriptor_device_b2ce4af9d5a0629f) }

var fileDescriptor_device_b2ce4af9d5a0629f = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x71, 0xba, 0x16, 0xf6, 0xd2, 0xac, 0x91, 0x07, 0x5a, 0x40, 0x1c, 0xba, 0x20, 0xa1,
	0x8a, 0x43, 0x91, 0x8a, 0x84, 0x18, 0x37, 0xe8, 0x10, 0x9a, 0xd4, 0xc3, 0xe4, 0x71, 0x68, 0x7b,
	0x20, 0xf2, 0x12, 0xd3, 0x5a, 0x73, 0x9d, 0x2a, 0x71, 0xbb, 0xa4, 0x77, 0x0e, 0x1c, 0xf8, 0x38,
	0x7c, 0x0c, 0x24, 0x3e, 0x12, 0x8a, 0x53, 0x67, 0x6d, 0x55, 0x40, 0xdc, 0xf2, 0xfe, 0xef, 0x3d,
	0xbf, 0xdf, 0xf3, 0xdf, 0x0a, 0x34, 0x23, 0xb6, 0xe4, 0x21, 0xeb, 0xce, 0x93, 0x58, 0xc5, 0xd8,
	0x11, 0xfc, 0x7a, 0xba, 0x60, 0xa9, 0xa2, 0x21, 0xa3, 0xd2, 0xff, 0x86, 0xe0, 0xe4, 0x5c, 0xe7,
	0x2f, 0x93, 0x78, 0xc9, 0x23, 0x96, 0xbc, 0x4b, 0xc2, 0x29, 0x53, 0xf9, 0x9c, 0x61, 0x0c, 0x07,
	0x92, 0xce, 0x98, 0x87, 0xda, 0xa8, 0x73, 0x48, 0xf4, 0x37, 0xee, 0xc1, 0x23, 0x11, 0xdf, 0x06,
	0x82, 0x2a, 0x26, 0xc3, 0x3c, 0x48, 0x17, 0xf3, 0x79, 0x9c, 0x28, 0x16, 0x79, 0x56, 0x1b, 0x75,
	0x1e, 0x90, 0x63, 0x11, 0xdf, 0x0e, 0xca, 0xdc, 0x95, 0x49, 0xe1, 0xe7, 0xd0, 0x9a, 0xd1, 0x2c,
	0xd8, 0xe8, 0xf3, 0x6a, 0x6d, 0xd4, 0x71, 0x88, 0x33, 0xa3, 0xd9, 0xa0, 0x6a, 0xf0, 0x03, 0x78,
	0x58, 0xa2, 0x54, 0x08, 0x03, 0x3e, 0x99, 0x2a, 0xfc, 0x11, 0xdc, 0x88, 0x7d, 0xa1, 0x0b, 0xa1,
	0x02, 0x11, 0x87, 0x54, 0xf1, 0x58, 0x6a, 0x26, 0xbb, 0xf7, 0xb4, 0xbb, 0xb5, 0x4d, 0x57, 0xd7,
	0x0f, 0xd6, 0x35, 0xa4, 0xb5, 0xee, 0x32, 0x82, 0xff, 0x03, 0x41, 0x6b, 0x67, 0x02, 0x7e, 0x0f,
	0x0d, 0x51, 0x74, 0xa5, 0x9e, 0xd5, 0xae, 0x75, 0xec, 0xde, 0x8b, 0x9d, 0x23, 0x77, 0xea, 0xcb,
	0x11, 0xe9, 0x07, 0xa9, 0x92, 0x9c, 0xac, 0x3b, 0x9f, 0x7c, 0x06, 0x7b, 0x43, 0xc6, 0x2e, 0xd4,
	0x6e, 0x58, 0xae, 0x11, 0x1d, 0x52, 0x7c, 0xe2, 0x33, 0xa8, 0x2f, 0xa9, 0x58, 0x30, 0x7d, 0x4b,
	0x76, 0xef, 0xd9, 0xdf, 0x67, 0xe8, 0xb3, 0x48, 0xd9, 0xf1, 0xd6, 0x7a, 0x83, 0xfc, 0x9f, 0x08,
	0x8e, 0xb6, 0x4d, 0xc2, 0xa7, 0xd0, 0xa4, 0xa6, 0x3e, 0xe0, 0xd1, 0x7a, 0x98, 0x5d, 0x69, 0x17,
	0x51, 0x65, 0x9f, 0xb5, 0x61, 0xdf, 0x19, 0xd4, 0x53, 0x45, 0x15, 0xd3, 0x06, 0x1c, 0xfd, 0x01,
	0xc4, 0x0c, 0xe9, 0x5e, 0x15, 0xa5, 0xa4, 0xec, 0xf0, 0x2f, 0xa0, 0xae, 0x63, 0xec, 0x42, 0xf3,
	0x9c, 0xa7, 0x61, 0x2c, 0x25, 0x0b, 0x15, 0x8b, 0xdc, 0x7b, 0xf8, 0x08, 0xa0, 0x5f, 0x86, 0x5c,
	0x4e, 0x5c, 0x84, 0x5b, 0x60, 0x5f, 0x32, 0x19, 0x71, 0x39, 0x19, 0x70, 0x79, 0xe3, 0x5a, 0xd8,
	0x81, 0xc3, 0x7e, 0x55, 0x5f, 0xf3, 0x5f, 0x03, 0xe8, 0x1d, 0xfb, 0xb1, 0x88, 0x13, 0xdc, 0x04,
	0x34, 0xd5, 0xfc, 0x88, 0xa0, 0x69, 0x11, 0xa5, 0x1a, 0x19, 0x11, 0x94, 0x16, 0x91, 0xd0, 0xac,
	0x88, 0x20, 0xe1, 0xff, 0x42, 0xd0, 0x28, 0x11, 0xff, 0x77, 0x7f, 0xa7, 0xda, 0xdf, 0xb8, 0x5d,
	0xd3, 0x6e, 0x9f, 0xee, 0xbd, 0x80, 0xbd, 0x26, 0x7f, 0xfa, 0x97, 0xc9, 0x2f, 0xb7, 0x4d, 0x7e,
	0xbc, 0xef, 0x6d, 0xea, 0x8d, 0x37, 0xad, 0xfd, 0x8e, 0xc0, 0xd9, 0x7a, 0xb5, 0xf8, 0x18, 0xea,
	0x33, 0x2e, 0x83, 0x4c, 0x1f, 0x6d, 0x91, 0x83, 0x19, 0x97, 0x43, 0x23, 0xe6, 0xfa, 0xec, 0x52,
	0x1c, 0x19, 0x71, 0xa5, 0x2f, 0xa8, 0x14, 0xc7, 0x5a, 0xa4, 0x59, 0x90, 0x79, 0x07, 0x6b, 0x91,
	0x66, 0x43, 0x23, 0xe6, 0x5e, 0xbd, 0x12, 0x47, 0x46, 0x5c, 0x79, 0x8d, 0x4a, 0x1c, 0xfb, 0x5f,
	0xab, 0xa7, 0x56, 0x01, 0x35, 0x01, 0x19, 0x18, 0x94, 0x15, 0x91, 0xa1, 0x40, 0x79, 0x11, 0x99,
	0xf1, 0x68, 0x85, 0x4f, 0xe0, 0x7e, 0x1a, 0x52, 0xc1, 0xaa, 0xe9, 0x0d, 0x1d, 0x0e, 0xef, 0x12,
	0x86, 0xa0, 0x4c, 0x8c, 0xee, 0x12, 0x86, 0xa2, 0x4c, 0x8c, 0xaf, 0x1b, 0xfa, 0x6f, 0xf5, 0xea,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x17, 0xce, 0x70, 0xbd, 0x04, 0x00, 0x00,
}
