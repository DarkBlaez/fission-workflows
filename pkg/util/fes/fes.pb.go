// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/util/fes/fes.proto

/*
Package fes is a generated protocol buffer package.

It is generated from these files:
	pkg/util/fes/fes.proto

It has these top-level messages:
	Aggregate
	Event
*/
package fes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Aggregate struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *Aggregate) Reset()                    { *m = Aggregate{} }
func (m *Aggregate) String() string            { return proto.CompactTextString(m) }
func (*Aggregate) ProtoMessage()               {}
func (*Aggregate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Aggregate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Aggregate) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Event struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type      string                     `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Aggregate *Aggregate                 `protobuf:"bytes,3,opt,name=aggregate" json:"aggregate,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Data      []byte                     `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetAggregate() *Aggregate {
	if m != nil {
		return m.Aggregate
	}
	return nil
}

func (m *Event) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Aggregate)(nil), "Aggregate")
	proto.RegisterType((*Event)(nil), "Event")
}

func init() { proto.RegisterFile("pkg/util/fes/fes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xc8, 0x4e, 0xd7,
	0x2f, 0x2d, 0xc9, 0xcc, 0xd1, 0x4f, 0x4b, 0x2d, 0x06, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c,
	0x29, 0xf9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x30, 0x2f, 0xa9, 0x34, 0x4d, 0xbf, 0x24,
	0x33, 0x37, 0xb5, 0xb8, 0x24, 0x31, 0xb7, 0x00, 0xa2, 0x40, 0x49, 0x9f, 0x8b, 0xd3, 0x31, 0x3d,
	0xbd, 0x28, 0x35, 0x3d, 0xb1, 0x24, 0x55, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b, 0xa5, 0xa4, 0xb2, 0x20, 0x55, 0x82,
	0x09, 0x2c, 0x02, 0x66, 0x2b, 0x2d, 0x66, 0xe4, 0x62, 0x75, 0x2d, 0x4b, 0xcd, 0x2b, 0x21, 0x46,
	0xb5, 0x90, 0x06, 0x17, 0x67, 0x22, 0xcc, 0x78, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x2e,
	0x3d, 0xb8, 0x85, 0x41, 0x08, 0x49, 0x21, 0x0b, 0x2e, 0x4e, 0xb8, 0xdb, 0x24, 0x58, 0xc0, 0x2a,
	0xa5, 0xf4, 0x20, 0xae, 0xd7, 0x83, 0xb9, 0x5e, 0x2f, 0x04, 0xa6, 0x22, 0x08, 0xa1, 0x18, 0x64,
	0x6f, 0x4a, 0x62, 0x49, 0xa2, 0x04, 0xab, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0x9d, 0xc4, 0x06,
	0xd6, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x63, 0x7b, 0x1a, 0xd6, 0x18, 0x01, 0x00, 0x00,
}
