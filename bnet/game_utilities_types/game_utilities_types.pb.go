// Code generated by protoc-gen-go.
// source: github.com/vjurenka/hs-proto-go/bnet/game_utilities_types/game_utilities_types.proto
// DO NOT EDIT!

/*
Package game_utilities_types is a generated protocol buffer package.

It is generated from these files:
	github.com/vjurenka/hs-proto-go/bnet/game_utilities_types/game_utilities_types.proto

It has these top-level messages:
	PlayerVariables
*/
package game_utilities_types

import proto "github.com/golang/protobuf/proto"
import math "math"
import attribute "github.com/vjurenka/hs-proto-go/bnet/attribute"
import entity "github.com/vjurenka/hs-proto-go/bnet/entity"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type PlayerVariables struct {
	Identity         *entity.Identity       `protobuf:"bytes,1,req,name=identity" json:"identity,omitempty"`
	Rating           *float64               `protobuf:"fixed64,2,opt,name=rating" json:"rating,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,3,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *PlayerVariables) Reset()         { *m = PlayerVariables{} }
func (m *PlayerVariables) String() string { return proto.CompactTextString(m) }
func (*PlayerVariables) ProtoMessage()    {}

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
