// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Category int32

const (
	Category_CMD    Category = 0
	Category_SENSOR Category = 1
)

var Category_name = map[int32]string{
	0: "CMD",
	1: "SENSOR",
}

var Category_value = map[string]int32{
	"CMD":    0,
	"SENSOR": 1,
}

func (x Category) String() string {
	return proto.EnumName(Category_name, int32(x))
}

func (Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

type Device int32

const (
	Device_LAMP   Device = 0
	Device_HEATER Device = 1
)

var Device_name = map[int32]string{
	0: "LAMP",
	1: "HEATER",
}

var Device_value = map[string]int32{
	"LAMP":   0,
	"HEATER": 1,
}

func (x Device) String() string {
	return proto.EnumName(Device_name, int32(x))
}

func (Device) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

type Operation int32

const (
	Operation_OFF Operation = 0
	Operation_ON  Operation = 1
)

var Operation_name = map[int32]string{
	0: "OFF",
	1: "ON",
}

var Operation_value = map[string]int32{
	"OFF": 0,
	"ON":  1,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}

func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}

//message是固定的。UserInfo是类名，可以随意指定，符合规范即可
type UplinkFrame struct {
	FrameType            int32    `protobuf:"varint,1,opt,name=FrameType,proto3" json:"FrameType,omitempty"`
	DevAddr              []byte   `protobuf:"bytes,2,opt,name=DevAddr,proto3" json:"DevAddr,omitempty"`
	FrameNum             uint32   `protobuf:"varint,3,opt,name=FrameNum,proto3" json:"FrameNum,omitempty"`
	Port                 uint32   `protobuf:"varint,4,opt,name=Port,proto3" json:"Port,omitempty"`
	PhyPayload           []byte   `protobuf:"bytes,5,opt,name=PhyPayload,proto3" json:"PhyPayload,omitempty"`
	UplinkId             uint32   `protobuf:"varint,6,opt,name=UplinkId,proto3" json:"UplinkId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UplinkFrame) Reset()         { *m = UplinkFrame{} }
func (m *UplinkFrame) String() string { return proto.CompactTextString(m) }
func (*UplinkFrame) ProtoMessage()    {}
func (*UplinkFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

func (m *UplinkFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkFrame.Unmarshal(m, b)
}
func (m *UplinkFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkFrame.Marshal(b, m, deterministic)
}
func (m *UplinkFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkFrame.Merge(m, src)
}
func (m *UplinkFrame) XXX_Size() int {
	return xxx_messageInfo_UplinkFrame.Size(m)
}
func (m *UplinkFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkFrame.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkFrame proto.InternalMessageInfo

func (m *UplinkFrame) GetFrameType() int32 {
	if m != nil {
		return m.FrameType
	}
	return 0
}

func (m *UplinkFrame) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *UplinkFrame) GetFrameNum() uint32 {
	if m != nil {
		return m.FrameNum
	}
	return 0
}

func (m *UplinkFrame) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *UplinkFrame) GetPhyPayload() []byte {
	if m != nil {
		return m.PhyPayload
	}
	return nil
}

func (m *UplinkFrame) GetUplinkId() uint32 {
	if m != nil {
		return m.UplinkId
	}
	return 0
}

type DownlinkFrame struct {
	FrameType            int32    `protobuf:"varint,1,opt,name=FrameType,proto3" json:"FrameType,omitempty"`
	DevAddr              []byte   `protobuf:"bytes,2,opt,name=DevAddr,proto3" json:"DevAddr,omitempty"`
	FrameNum             uint32   `protobuf:"varint,3,opt,name=FrameNum,proto3" json:"FrameNum,omitempty"`
	Port                 uint32   `protobuf:"varint,4,opt,name=Port,proto3" json:"Port,omitempty"`
	PhyPayload           []byte   `protobuf:"bytes,5,opt,name=PhyPayload,proto3" json:"PhyPayload,omitempty"`
	DownlinkId           uint32   `protobuf:"varint,6,opt,name=DownlinkId,proto3" json:"DownlinkId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkFrame) Reset()         { *m = DownlinkFrame{} }
func (m *DownlinkFrame) String() string { return proto.CompactTextString(m) }
func (*DownlinkFrame) ProtoMessage()    {}
func (*DownlinkFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

func (m *DownlinkFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkFrame.Unmarshal(m, b)
}
func (m *DownlinkFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkFrame.Marshal(b, m, deterministic)
}
func (m *DownlinkFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkFrame.Merge(m, src)
}
func (m *DownlinkFrame) XXX_Size() int {
	return xxx_messageInfo_DownlinkFrame.Size(m)
}
func (m *DownlinkFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkFrame.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkFrame proto.InternalMessageInfo

func (m *DownlinkFrame) GetFrameType() int32 {
	if m != nil {
		return m.FrameType
	}
	return 0
}

func (m *DownlinkFrame) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *DownlinkFrame) GetFrameNum() uint32 {
	if m != nil {
		return m.FrameNum
	}
	return 0
}

func (m *DownlinkFrame) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *DownlinkFrame) GetPhyPayload() []byte {
	if m != nil {
		return m.PhyPayload
	}
	return nil
}

func (m *DownlinkFrame) GetDownlinkId() uint32 {
	if m != nil {
		return m.DownlinkId
	}
	return 0
}

type Payload struct {
	Kind                 uint32   `protobuf:"varint,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Key                  uint32   `protobuf:"varint,2,opt,name=key,proto3" json:"key,omitempty"`
	Val                  []byte   `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetKind() uint32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *Payload) GetKey() uint32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Payload) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.Category", Category_name, Category_value)
	proto.RegisterEnum("proto.Device", Device_name, Device_value)
	proto.RegisterEnum("proto.Operation", Operation_name, Operation_value)
	proto.RegisterType((*UplinkFrame)(nil), "proto.UplinkFrame")
	proto.RegisterType((*DownlinkFrame)(nil), "proto.DownlinkFrame")
	proto.RegisterType((*Payload)(nil), "proto.Payload")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xd1, 0x6a, 0xc2, 0x30,
	0x14, 0x86, 0x8d, 0xd6, 0xaa, 0x67, 0x16, 0xca, 0xb9, 0x0a, 0x43, 0x9c, 0x78, 0x25, 0x5e, 0xec,
	0x66, 0x4f, 0x50, 0xac, 0xb2, 0xc1, 0x6c, 0x4b, 0x74, 0x0f, 0x90, 0xad, 0xc1, 0x95, 0xd6, 0xa6,
	0x84, 0xae, 0x92, 0xc7, 0xda, 0xd5, 0x5e, 0x6f, 0x24, 0xa5, 0xce, 0x47, 0xd8, 0x55, 0xfe, 0xf3,
	0x9d, 0xc3, 0xcf, 0x77, 0x11, 0xf0, 0x4e, 0xbc, 0x16, 0x17, 0xae, 0x1f, 0x2b, 0x25, 0x6b, 0x89,
	0x43, 0xfb, 0x2c, 0xbf, 0x09, 0xdc, 0xbd, 0x55, 0x45, 0x56, 0xe6, 0x3b, 0xc5, 0xcf, 0x02, 0x67,
	0x30, 0xb1, 0xe1, 0xa8, 0x2b, 0x41, 0xc9, 0x82, 0xac, 0x86, 0xec, 0x0f, 0x20, 0x85, 0x51, 0x28,
	0x9a, 0x20, 0x4d, 0x15, 0xed, 0x2f, 0xc8, 0x6a, 0xca, 0xba, 0x11, 0xef, 0x61, 0x6c, 0xcf, 0xa2,
	0xaf, 0x33, 0x1d, 0x2c, 0xc8, 0xca, 0x63, 0xd7, 0x19, 0x11, 0x9c, 0x44, 0xaa, 0x9a, 0x3a, 0x96,
	0xdb, 0x8c, 0x73, 0x80, 0xe4, 0x53, 0x27, 0x5c, 0x17, 0x92, 0xa7, 0x74, 0x68, 0xcb, 0x6e, 0x88,
	0xe9, 0x6b, 0xb5, 0x5e, 0x52, 0xea, 0xb6, 0x7d, 0xdd, 0xbc, 0xfc, 0x21, 0xe0, 0x85, 0xf2, 0x52,
	0xfe, 0x3f, 0xeb, 0x39, 0x40, 0x27, 0x76, 0xf5, 0xbe, 0x21, 0xcb, 0x00, 0x46, 0xdd, 0x29, 0x82,
	0x93, 0x67, 0x65, 0x6a, 0x6d, 0x3d, 0x66, 0x33, 0xfa, 0x30, 0xc8, 0x85, 0xb6, 0x92, 0x1e, 0x33,
	0xd1, 0x90, 0x86, 0x17, 0xd6, 0x6d, 0xca, 0x4c, 0x5c, 0x3f, 0xc0, 0x78, 0xc3, 0x6b, 0x71, 0x92,
	0x4a, 0xe3, 0x08, 0x06, 0x9b, 0x7d, 0xe8, 0xf7, 0x10, 0xc0, 0x3d, 0x6c, 0xa3, 0x43, 0xcc, 0x7c,
	0xb2, 0x9e, 0x83, 0x1b, 0x8a, 0x26, 0xfb, 0x10, 0x38, 0x06, 0xe7, 0x35, 0xd8, 0x27, 0xed, 0xfe,
	0x79, 0x1b, 0x1c, 0xb7, 0x66, 0x3f, 0x83, 0x49, 0x5c, 0x09, 0xc5, 0xeb, 0x4c, 0x96, 0xa6, 0x21,
	0xde, 0xed, 0xfc, 0x1e, 0xba, 0xd0, 0x8f, 0x23, 0x9f, 0xbc, 0xbb, 0xf6, 0x5b, 0x3c, 0xfd, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xec, 0x22, 0x6e, 0x33, 0x2e, 0x02, 0x00, 0x00,
}