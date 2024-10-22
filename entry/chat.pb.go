// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entry/chat.proto

package entry

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

type BroadcastMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Timestamp            int32    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastMessage) Reset()         { *m = BroadcastMessage{} }
func (m *BroadcastMessage) String() string { return proto.CompactTextString(m) }
func (*BroadcastMessage) ProtoMessage()    {}
func (*BroadcastMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_82c3cbe99389f4f3, []int{0}
}

func (m *BroadcastMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastMessage.Unmarshal(m, b)
}
func (m *BroadcastMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastMessage.Marshal(b, m, deterministic)
}
func (m *BroadcastMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastMessage.Merge(m, src)
}
func (m *BroadcastMessage) XXX_Size() int {
	return xxx_messageInfo_BroadcastMessage.Size(m)
}
func (m *BroadcastMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastMessage proto.InternalMessageInfo

func (m *BroadcastMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BroadcastMessage) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *BroadcastMessage) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*BroadcastMessage)(nil), "entry.BroadcastMessage")
}

func init() {
	proto.RegisterFile("entry/chat.proto", fileDescriptor_82c3cbe99389f4f3)
}

var fileDescriptor_82c3cbe99389f4f3 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcd, 0x2b, 0x29,
	0xaa, 0xd4, 0x4f, 0xce, 0x48, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x8b,
	0x28, 0xc5, 0x71, 0x09, 0x38, 0x15, 0xe5, 0x27, 0xa6, 0x24, 0x27, 0x16, 0x97, 0xf8, 0xa6, 0x16,
	0x17, 0x27, 0xa6, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0xe7, 0x42, 0x98, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x30, 0xae, 0x90, 0x10, 0x17, 0x4b, 0x69, 0x71, 0x6a, 0x91, 0x04, 0x13, 0x58, 0x18,
	0xcc, 0x16, 0x92, 0xe1, 0xe2, 0x2c, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60,
	0x56, 0x60, 0xd4, 0x60, 0x0d, 0x42, 0x08, 0x18, 0x85, 0x23, 0x99, 0x1f, 0x9c, 0x5a, 0x54, 0x96,
	0x99, 0x9c, 0x2a, 0xe4, 0xcc, 0xc5, 0x99, 0x04, 0x13, 0x13, 0x12, 0xd7, 0x03, 0x3b, 0x44, 0x0f,
	0xdd, 0x15, 0x52, 0xb8, 0x24, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x9d, 0xe4, 0xa3, 0x64, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xb3, 0xd2, 0xb3, 0xf2, 0xf5, 0xd3,
	0x8b, 0x0a, 0x92, 0x75, 0xd3, 0xf3, 0xf5, 0xc1, 0xfa, 0x92, 0xd8, 0xc0, 0xfe, 0x34, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xf1, 0xb9, 0x8e, 0xfe, 0xfb, 0x00, 0x00, 0x00,
}
