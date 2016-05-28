// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/attribute/attribute.proto
// DO NOT EDIT!

/*
Package attribute is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/attribute/attribute.proto

It has these top-level messages:
	Variant
	Attribute
	AttributeFilter
*/
package attribute

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import entity "github.com/HearthSim/hs-proto-go/bnet/entity"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type AttributeFilter_Operation int32

const (
	AttributeFilter_MATCH_NONE              AttributeFilter_Operation = 0
	AttributeFilter_MATCH_ANY               AttributeFilter_Operation = 1
	AttributeFilter_MATCH_ALL               AttributeFilter_Operation = 2
	AttributeFilter_MATCH_ALL_MOST_SPECIFIC AttributeFilter_Operation = 3
)

var AttributeFilter_Operation_name = map[int32]string{
	0: "MATCH_NONE",
	1: "MATCH_ANY",
	2: "MATCH_ALL",
	3: "MATCH_ALL_MOST_SPECIFIC",
}
var AttributeFilter_Operation_value = map[string]int32{
	"MATCH_NONE":              0,
	"MATCH_ANY":               1,
	"MATCH_ALL":               2,
	"MATCH_ALL_MOST_SPECIFIC": 3,
}

func (x AttributeFilter_Operation) Enum() *AttributeFilter_Operation {
	p := new(AttributeFilter_Operation)
	*p = x
	return p
}
func (x AttributeFilter_Operation) String() string {
	return proto.EnumName(AttributeFilter_Operation_name, int32(x))
}
func (x *AttributeFilter_Operation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AttributeFilter_Operation_value, data, "AttributeFilter_Operation")
	if err != nil {
		return err
	}
	*x = AttributeFilter_Operation(value)
	return nil
}
func (AttributeFilter_Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Variant struct {
	BoolValue        *bool            `protobuf:"varint,2,opt,name=bool_value" json:"bool_value,omitempty"`
	IntValue         *int64           `protobuf:"varint,3,opt,name=int_value" json:"int_value,omitempty"`
	FloatValue       *float64         `protobuf:"fixed64,4,opt,name=float_value" json:"float_value,omitempty"`
	StringValue      *string          `protobuf:"bytes,5,opt,name=string_value" json:"string_value,omitempty"`
	BlobValue        []byte           `protobuf:"bytes,6,opt,name=blob_value" json:"blob_value,omitempty"`
	MessageValue     []byte           `protobuf:"bytes,7,opt,name=message_value" json:"message_value,omitempty"`
	FourccValue      *string          `protobuf:"bytes,8,opt,name=fourcc_value" json:"fourcc_value,omitempty"`
	UintValue        *uint64          `protobuf:"varint,9,opt,name=uint_value" json:"uint_value,omitempty"`
	EntityidValue    *entity.EntityId `protobuf:"bytes,10,opt,name=entityid_value" json:"entityid_value,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Variant) Reset()                    { *m = Variant{} }
func (m *Variant) String() string            { return proto.CompactTextString(m) }
func (*Variant) ProtoMessage()               {}
func (*Variant) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Variant) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

func (m *Variant) GetIntValue() int64 {
	if m != nil && m.IntValue != nil {
		return *m.IntValue
	}
	return 0
}

func (m *Variant) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Variant) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Variant) GetBlobValue() []byte {
	if m != nil {
		return m.BlobValue
	}
	return nil
}

func (m *Variant) GetMessageValue() []byte {
	if m != nil {
		return m.MessageValue
	}
	return nil
}

func (m *Variant) GetFourccValue() string {
	if m != nil && m.FourccValue != nil {
		return *m.FourccValue
	}
	return ""
}

func (m *Variant) GetUintValue() uint64 {
	if m != nil && m.UintValue != nil {
		return *m.UintValue
	}
	return 0
}

func (m *Variant) GetEntityidValue() *entity.EntityId {
	if m != nil {
		return m.EntityidValue
	}
	return nil
}

type Attribute struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *Variant `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Attribute) Reset()                    { *m = Attribute{} }
func (m *Attribute) String() string            { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()               {}
func (*Attribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Attribute) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Attribute) GetValue() *Variant {
	if m != nil {
		return m.Value
	}
	return nil
}

type AttributeFilter struct {
	Op               *AttributeFilter_Operation `protobuf:"varint,1,req,name=op,enum=attribute.AttributeFilter_Operation" json:"op,omitempty"`
	Attribute        []*Attribute               `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *AttributeFilter) Reset()                    { *m = AttributeFilter{} }
func (m *AttributeFilter) String() string            { return proto.CompactTextString(m) }
func (*AttributeFilter) ProtoMessage()               {}
func (*AttributeFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AttributeFilter) GetOp() AttributeFilter_Operation {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return AttributeFilter_MATCH_NONE
}

func (m *AttributeFilter) GetAttribute() []*Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func init() {
	proto.RegisterType((*Variant)(nil), "attribute.Variant")
	proto.RegisterType((*Attribute)(nil), "attribute.Attribute")
	proto.RegisterType((*AttributeFilter)(nil), "attribute.AttributeFilter")
	proto.RegisterEnum("attribute.AttributeFilter_Operation", AttributeFilter_Operation_name, AttributeFilter_Operation_value)
}

var fileDescriptor0 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xdd, 0x6a, 0xc2, 0x30,
	0x14, 0xc7, 0xd7, 0x56, 0xa7, 0x3d, 0x7e, 0xac, 0xcb, 0x1c, 0x93, 0xed, 0xc6, 0x95, 0xc1, 0xbc,
	0xb1, 0x1d, 0xde, 0x0d, 0x06, 0x43, 0x44, 0x51, 0xf0, 0x63, 0xa0, 0x08, 0xbb, 0x2a, 0xa9, 0xc6,
	0x1a, 0x68, 0x1b, 0x69, 0xd3, 0xc1, 0xde, 0x6f, 0x4f, 0xb1, 0xa7, 0x59, 0xa8, 0xa9, 0x95, 0xb1,
	0x8b, 0x5d, 0xa5, 0xe7, 0x47, 0xfe, 0xbf, 0xd3, 0x73, 0x02, 0xaf, 0x1e, 0xe5, 0xbb, 0xc4, 0xb5,
	0xd6, 0x2c, 0xb0, 0x47, 0x04, 0x47, 0x7c, 0xb7, 0xa0, 0x81, 0xbd, 0x8b, 0x3b, 0xfb, 0x88, 0x71,
	0xd6, 0xf1, 0x98, 0xed, 0x86, 0x84, 0xdb, 0x98, 0xf3, 0x88, 0xba, 0x09, 0x27, 0xf9, 0x97, 0x95,
	0xde, 0x40, 0xfa, 0x11, 0xdc, 0x3e, 0xff, 0xcf, 0x45, 0x42, 0x4e, 0xf9, 0xa7, 0x3c, 0x0e, 0x16,
	0xf3, 0x5b, 0x81, 0xd2, 0x0a, 0x47, 0x14, 0x87, 0x1c, 0x21, 0x00, 0x97, 0x31, 0xdf, 0xf9, 0xc0,
	0x7e, 0x42, 0x9a, 0x6a, 0x4b, 0x69, 0x97, 0xd1, 0x25, 0xe8, 0x34, 0xe4, 0x12, 0x69, 0x02, 0x69,
	0xe8, 0x0a, 0x2a, 0x5b, 0x9f, 0xe1, 0x0c, 0x16, 0x04, 0x54, 0x50, 0x03, 0xaa, 0xb1, 0xf8, 0x9d,
	0xd0, 0x93, 0xb4, 0x28, 0xa8, 0x9e, 0x1a, 0x7d, 0xe6, 0x4a, 0x76, 0x2e, 0x58, 0x15, 0x5d, 0x43,
	0x2d, 0x20, 0x71, 0x8c, 0x3d, 0x22, 0x71, 0x29, 0xc5, 0x42, 0xb0, 0x65, 0x49, 0xb4, 0x5e, 0x4b,
	0x5a, 0xce, 0x04, 0x49, 0xde, 0x5f, 0x17, 0xac, 0x80, 0xda, 0x50, 0x3f, 0x8c, 0x40, 0x37, 0x92,
	0x83, 0xe0, 0x95, 0xae, 0x61, 0xc9, 0xc9, 0x06, 0xe9, 0x31, 0xde, 0x98, 0x2f, 0xa0, 0xf7, 0xb2,
	0x25, 0xa1, 0x2a, 0x14, 0x42, 0x1c, 0x90, 0xa6, 0xd2, 0x52, 0x85, 0xf8, 0x1e, 0x8a, 0xd9, 0x98,
	0xaa, 0xc8, 0x22, 0x2b, 0x5f, 0xaf, 0x5c, 0x87, 0xf9, 0xa5, 0xc0, 0xc5, 0x31, 0x3e, 0xa4, 0x3e,
	0x27, 0x11, 0x7a, 0x02, 0x95, 0xed, 0x53, 0x45, 0xbd, 0xfb, 0x70, 0x92, 0xf9, 0x75, 0xcf, 0x9a,
	0xef, 0x49, 0x84, 0x39, 0x65, 0x21, 0x7a, 0x84, 0xfc, 0xa1, 0x44, 0x33, 0x4d, 0x34, 0x6b, 0xfc,
	0x15, 0x34, 0x57, 0xa0, 0xe7, 0xa9, 0x3a, 0xc0, 0xb4, 0xb7, 0xec, 0x8f, 0x9c, 0xd9, 0x7c, 0x36,
	0x30, 0xce, 0x50, 0x0d, 0xf4, 0x43, 0xdd, 0x9b, 0xbd, 0x1b, 0xca, 0x49, 0x39, 0x99, 0x18, 0x2a,
	0xba, 0x83, 0x9b, 0x63, 0xe9, 0x4c, 0xe7, 0x8b, 0xa5, 0xb3, 0x78, 0x1b, 0xf4, 0xc7, 0xc3, 0x71,
	0xdf, 0xd0, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x69, 0xc7, 0xe0, 0x69, 0x02, 0x00, 0x00,
}
