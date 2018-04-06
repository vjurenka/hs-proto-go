// Code generated by protoc-gen-go.
// source: github.com/vjurenka/hs-proto-go/bnet/chat_types/chat_types.proto
// DO NOT EDIT!

/*
Package chat_types is a generated protocol buffer package.

It is generated from these files:
	github.com/vjurenka/hs-proto-go/bnet/chat_types/chat_types.proto

It has these top-level messages:
	ChannelState
*/
package chat_types

import proto "github.com/golang/protobuf/proto"
import math "math"
import channel_types "github.com/vjurenka/hs-proto-go/bnet/channel_types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ChannelState struct {
	Identity         *string `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	Program          *uint32 `protobuf:"fixed32,2,opt,name=program" json:"program,omitempty"`
	Locale           *uint32 `protobuf:"fixed32,3,opt,name=locale" json:"locale,omitempty"`
	Public           *bool   `protobuf:"varint,4,opt,name=public,def=0" json:"public,omitempty"`
	BucketIndex      *uint32 `protobuf:"varint,5,opt,name=bucket_index" json:"bucket_index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChannelState) Reset()         { *m = ChannelState{} }
func (m *ChannelState) String() string { return proto.CompactTextString(m) }
func (*ChannelState) ProtoMessage()    {}

const Default_ChannelState_Public bool = false

func (m *ChannelState) GetIdentity() string {
	if m != nil && m.Identity != nil {
		return *m.Identity
	}
	return ""
}

func (m *ChannelState) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *ChannelState) GetLocale() uint32 {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return 0
}

func (m *ChannelState) GetPublic() bool {
	if m != nil && m.Public != nil {
		return *m.Public
	}
	return Default_ChannelState_Public
}

func (m *ChannelState) GetBucketIndex() uint32 {
	if m != nil && m.BucketIndex != nil {
		return *m.BucketIndex
	}
	return 0
}

var E_ChannelState_Chat = &proto.ExtensionDesc{
	ExtendedType:  (*channel_types.ChannelState)(nil),
	ExtensionType: (*ChannelState)(nil),
	Field:         100,
	Name:          "chat_types.ChannelState.chat",
	Tag:           "bytes,100,opt,name=chat",
}

func init() {
	proto.RegisterExtension(E_ChannelState_Chat)
}