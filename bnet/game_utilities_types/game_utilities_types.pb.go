// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/game_utilities_types/game_utilities_types.proto
// DO NOT EDIT!

/*
Package game_utilities_types is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/game_utilities_types/game_utilities_types.proto

It has these top-level messages:
	PlayerVariables
*/
package game_utilities_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import attribute "github.com/HearthSim/hs-proto-go/bnet/attribute"
import entity "github.com/HearthSim/hs-proto-go/bnet/entity"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type PlayerVariables struct {
	Identity         *entity.Identity       `protobuf:"bytes,1,req,name=identity" json:"identity,omitempty"`
	Rating           *float64               `protobuf:"fixed64,2,opt,name=rating" json:"rating,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,3,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *PlayerVariables) Reset()                    { *m = PlayerVariables{} }
func (m *PlayerVariables) String() string            { return proto.CompactTextString(m) }
func (*PlayerVariables) ProtoMessage()               {}
func (*PlayerVariables) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PlayerVariables) GetIdentity() *entity.Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *PlayerVariables) GetRating() float64 {
	if m != nil && m.Rating != nil {
		return *m.Rating
	}
	return 0
}

func (m *PlayerVariables) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func init() {
	proto.RegisterType((*PlayerVariables)(nil), "game_utilities_types.PlayerVariables")
}

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x0a, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xf7, 0x48, 0x4d, 0x2c, 0x2a, 0xc9, 0x08, 0xce, 0xcc,
	0xd5, 0xcf, 0x28, 0xd6, 0x2d, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4d, 0xcf, 0xd7, 0x4f, 0xca, 0x4b,
	0x2d, 0xd1, 0x4f, 0x4f, 0xcc, 0x4d, 0x8d, 0x2f, 0x2d, 0xc9, 0xcc, 0xc9, 0x2c, 0xc9, 0x4c, 0x2d,
	0x8e, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xc6, 0x2a, 0xa8, 0x07, 0xd6, 0x27, 0x24, 0x82, 0x4d, 0x4e,
	0xca, 0x9e, 0x38, 0xcb, 0x12, 0x4b, 0x4a, 0x8a, 0x32, 0x93, 0x4a, 0x4b, 0x52, 0x11, 0x2c, 0x88,
	0xb1, 0x52, 0x96, 0xc4, 0x19, 0x90, 0x9a, 0x57, 0x92, 0x59, 0x52, 0x09, 0xa5, 0x20, 0x5a, 0x95,
	0xf2, 0xb8, 0xf8, 0x03, 0x72, 0x12, 0x2b, 0x53, 0x8b, 0xc2, 0x12, 0x8b, 0x32, 0x13, 0x93, 0x72,
	0x52, 0x8b, 0x85, 0x94, 0xb8, 0x38, 0x32, 0x53, 0x20, 0x8a, 0x24, 0x18, 0x15, 0x98, 0x34, 0xb8,
	0x8d, 0x04, 0xf4, 0xa0, 0x7a, 0x3c, 0xa1, 0xe2, 0x42, 0x7c, 0x5c, 0x6c, 0x45, 0x89, 0x25, 0x99,
	0x79, 0xe9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x8c, 0x42, 0xea, 0x5c, 0x9c, 0x70, 0x47, 0x49, 0x30,
	0x2b, 0x30, 0x03, 0x35, 0x89, 0xe8, 0x21, 0x9c, 0xe9, 0x08, 0x63, 0x01, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0x3c, 0x86, 0x37, 0x59, 0x01, 0x00, 0x00,
}
