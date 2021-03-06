// Code generated by protoc-gen-go.
// source: github.com/vjurenka/hs-proto-go/bnet/attribute/attribute.proto
// DO NOT EDIT!

/*
Package attribute is a generated protocol buffer package.

It is generated from these files:
	github.com/vjurenka/hs-proto-go/bnet/attribute/attribute.proto

It has these top-level messages:
	Variant
	Attribute
	AttributeFilter
*/
package attribute

import proto "github.com/golang/protobuf/proto"
import math "math"
import entity "github.com/vjurenka/hs-proto-go/bnet/entity"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

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

func (m *Variant) Reset()         { *m = Variant{} }
func (m *Variant) String() string { return proto.CompactTextString(m) }
func (*Variant) ProtoMessage()    {}

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

func (m *Attribute) Reset()         { *m = Attribute{} }
func (m *Attribute) String() string { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()    {}

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

func (m *AttributeFilter) Reset()         { *m = AttributeFilter{} }
func (m *AttributeFilter) String() string { return proto.CompactTextString(m) }
func (*AttributeFilter) ProtoMessage()    {}

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
	proto.RegisterEnum("attribute.AttributeFilter_Operation", AttributeFilter_Operation_name, AttributeFilter_Operation_value)
}
