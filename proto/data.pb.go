// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

/*
Package data is a generated protocol buffer package.

It is generated from these files:
	data.proto

It has these top-level messages:
	ClientInfo
*/
package data

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

type ClientInfo struct {
	Hostname    string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	VendorClass string `protobuf:"bytes,2,opt,name=vendor_class,json=vendorClass" json:"vendor_class,omitempty"`
}

func (m *ClientInfo) Reset()                    { *m = ClientInfo{} }
func (m *ClientInfo) String() string            { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()               {}
func (*ClientInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ClientInfo) GetVendorClass() string {
	if m != nil {
		return m.VendorClass
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientInfo)(nil), "ClientInfo")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xf2, 0xe6, 0xe2, 0x72, 0xce, 0xc9, 0x4c, 0xcd, 0x2b,
	0xf1, 0xcc, 0x4b, 0xcb, 0x17, 0x92, 0xe2, 0xe2, 0xc8, 0xc8, 0x2f, 0x2e, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x14, 0xb9, 0x78, 0xca, 0x52, 0xf3,
	0x52, 0xf2, 0x8b, 0xe2, 0x93, 0x73, 0x12, 0x8b, 0x8b, 0x25, 0x98, 0xc0, 0xf2, 0xdc, 0x10, 0x31,
	0x67, 0x90, 0x50, 0x12, 0x1b, 0xd8, 0x4c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x97,
	0x5b, 0x51, 0x61, 0x00, 0x00, 0x00,
}