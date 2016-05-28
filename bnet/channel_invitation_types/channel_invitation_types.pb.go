// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/channel_invitation_types/channel_invitation_types.proto
// DO NOT EDIT!

/*
Package channel_invitation_types is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/channel_invitation_types/channel_invitation_types.proto

It has these top-level messages:
	ChannelInvitation
	ChannelInvitationParams
	InvitationCollection
*/
package channel_invitation_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import channel_types "github.com/HearthSim/hs-proto-go/bnet/channel_types"
import entity "github.com/HearthSim/hs-proto-go/bnet/entity"
import invitation_types "github.com/HearthSim/hs-proto-go/bnet/invitation_types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ChannelInvitation struct {
	ChannelDescription *channel_types.ChannelDescription `protobuf:"bytes,1,req,name=channel_description" json:"channel_description,omitempty"`
	Reserved           *bool                             `protobuf:"varint,2,opt,name=reserved,def=0" json:"reserved,omitempty"`
	Rejoin             *bool                             `protobuf:"varint,3,opt,name=rejoin,def=0" json:"rejoin,omitempty"`
	ServiceType        *uint32                           `protobuf:"varint,4,req,name=service_type" json:"service_type,omitempty"`
	XXX_unrecognized   []byte                            `json:"-"`
}

func (m *ChannelInvitation) Reset()                    { *m = ChannelInvitation{} }
func (m *ChannelInvitation) String() string            { return proto.CompactTextString(m) }
func (*ChannelInvitation) ProtoMessage()               {}
func (*ChannelInvitation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ChannelInvitation_Reserved bool = false
const Default_ChannelInvitation_Rejoin bool = false

func (m *ChannelInvitation) GetChannelDescription() *channel_types.ChannelDescription {
	if m != nil {
		return m.ChannelDescription
	}
	return nil
}

func (m *ChannelInvitation) GetReserved() bool {
	if m != nil && m.Reserved != nil {
		return *m.Reserved
	}
	return Default_ChannelInvitation_Reserved
}

func (m *ChannelInvitation) GetRejoin() bool {
	if m != nil && m.Rejoin != nil {
		return *m.Rejoin
	}
	return Default_ChannelInvitation_Rejoin
}

func (m *ChannelInvitation) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

var E_ChannelInvitation_ChannelInvitation = &proto.ExtensionDesc{
	ExtendedType:  (*invitation_types.Invitation)(nil),
	ExtensionType: (*ChannelInvitation)(nil),
	Field:         105,
	Name:          "channel_invitation_types.ChannelInvitation.channel_invitation",
	Tag:           "bytes,105,opt,name=channel_invitation",
}

type ChannelInvitationParams struct {
	ChannelId        *entity.EntityId `protobuf:"bytes,1,req,name=channel_id" json:"channel_id,omitempty"`
	Reserved         *bool            `protobuf:"varint,2,opt,name=reserved" json:"reserved,omitempty"`
	Rejoin           *bool            `protobuf:"varint,3,opt,name=rejoin" json:"rejoin,omitempty"`
	ServiceType      *uint32          `protobuf:"varint,4,req,name=service_type" json:"service_type,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ChannelInvitationParams) Reset()                    { *m = ChannelInvitationParams{} }
func (m *ChannelInvitationParams) String() string            { return proto.CompactTextString(m) }
func (*ChannelInvitationParams) ProtoMessage()               {}
func (*ChannelInvitationParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChannelInvitationParams) GetChannelId() *entity.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *ChannelInvitationParams) GetReserved() bool {
	if m != nil && m.Reserved != nil {
		return *m.Reserved
	}
	return false
}

func (m *ChannelInvitationParams) GetRejoin() bool {
	if m != nil && m.Rejoin != nil {
		return *m.Rejoin
	}
	return false
}

func (m *ChannelInvitationParams) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

var E_ChannelInvitationParams_ChannelParams = &proto.ExtensionDesc{
	ExtendedType:  (*invitation_types.InvitationParams)(nil),
	ExtensionType: (*ChannelInvitationParams)(nil),
	Field:         105,
	Name:          "channel_invitation_types.ChannelInvitationParams.channel_params",
	Tag:           "bytes,105,opt,name=channel_params",
}

type InvitationCollection struct {
	ServiceType            *uint32                        `protobuf:"varint,1,opt,name=service_type" json:"service_type,omitempty"`
	MaxReceivedInvitations *uint32                        `protobuf:"varint,2,opt,name=max_received_invitations" json:"max_received_invitations,omitempty"`
	ObjectId               *uint64                        `protobuf:"varint,3,opt,name=object_id" json:"object_id,omitempty"`
	ReceivedInvitation     []*invitation_types.Invitation `protobuf:"bytes,4,rep,name=received_invitation" json:"received_invitation,omitempty"`
	XXX_unrecognized       []byte                         `json:"-"`
}

func (m *InvitationCollection) Reset()                    { *m = InvitationCollection{} }
func (m *InvitationCollection) String() string            { return proto.CompactTextString(m) }
func (*InvitationCollection) ProtoMessage()               {}
func (*InvitationCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InvitationCollection) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

func (m *InvitationCollection) GetMaxReceivedInvitations() uint32 {
	if m != nil && m.MaxReceivedInvitations != nil {
		return *m.MaxReceivedInvitations
	}
	return 0
}

func (m *InvitationCollection) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *InvitationCollection) GetReceivedInvitation() []*invitation_types.Invitation {
	if m != nil {
		return m.ReceivedInvitation
	}
	return nil
}

func init() {
	proto.RegisterType((*ChannelInvitation)(nil), "channel_invitation_types.ChannelInvitation")
	proto.RegisterType((*ChannelInvitationParams)(nil), "channel_invitation_types.ChannelInvitationParams")
	proto.RegisterType((*InvitationCollection)(nil), "channel_invitation_types.InvitationCollection")
	proto.RegisterExtension(E_ChannelInvitation_ChannelInvitation)
	proto.RegisterExtension(E_ChannelInvitationParams_ChannelParams)
}

var fileDescriptor0 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x6a, 0xea, 0x40,
	0x14, 0xc7, 0x89, 0x7a, 0x2f, 0xde, 0xe3, 0x55, 0x74, 0xf4, 0x62, 0x90, 0xbb, 0xb0, 0xa1, 0x0b,
	0xa1, 0x98, 0x50, 0x77, 0x76, 0xd1, 0x8d, 0x2d, 0xad, 0x8b, 0x42, 0xa1, 0xeb, 0x22, 0x31, 0x39,
	0x35, 0x23, 0x49, 0x26, 0x4c, 0xa6, 0x52, 0x77, 0x7d, 0x90, 0x3e, 0x5f, 0x9f, 0xa3, 0x93, 0x0f,
	0x3f, 0x63, 0x44, 0x57, 0x03, 0x67, 0xe6, 0xff, 0xf1, 0x9b, 0x03, 0xaf, 0x33, 0x2a, 0x9c, 0xf7,
	0xa9, 0x6e, 0x31, 0xcf, 0x78, 0x44, 0x93, 0x0b, 0xe7, 0x85, 0x7a, 0x86, 0x13, 0xf6, 0x03, 0xce,
	0x04, 0xeb, 0xcf, 0x98, 0x31, 0xf5, 0x51, 0x18, 0x96, 0x63, 0xfa, 0x3e, 0xba, 0x13, 0xea, 0x2f,
	0xa8, 0x30, 0x05, 0x65, 0xfe, 0x44, 0x2c, 0x03, 0x0c, 0x73, 0x2f, 0xf4, 0x58, 0x4f, 0xd4, 0xbc,
	0xfb, 0xce, 0xc3, 0x79, 0xc1, 0xbb, 0x69, 0x5b, 0x11, 0x9d, 0xe1, 0x69, 0x46, 0xe8, 0x0b, 0x2a,
	0x96, 0xe9, 0x91, 0x4a, 0x9f, 0x4e, 0x93, 0x66, 0xa0, 0x0f, 0xc3, 0x6a, 0x9f, 0x05, 0x68, 0x8c,
	0x92, 0x86, 0xe3, 0xf5, 0x0b, 0x72, 0x0b, 0xcd, 0x55, 0x6d, 0x1b, 0x43, 0x8b, 0xd3, 0x20, 0x1a,
	0xab, 0x4a, 0xb7, 0xd0, 0xab, 0x0c, 0x2e, 0xf4, 0x5d, 0xa4, 0x54, 0x7e, 0xb7, 0x79, 0x48, 0xda,
	0x50, 0xe6, 0x18, 0x22, 0x5f, 0xa0, 0xad, 0x16, 0xba, 0x4a, 0xaf, 0x7c, 0xf3, 0xeb, 0xcd, 0x74,
	0x43, 0x24, 0xff, 0xe0, 0x37, 0xc7, 0x39, 0xa3, 0xbe, 0x5a, 0xdc, 0x1e, 0xb7, 0xe0, 0x6f, 0xf4,
	0x9a, 0x5a, 0x18, 0x7b, 0xaa, 0x25, 0x19, 0x54, 0x1d, 0x20, 0x90, 0xec, 0x2a, 0xc8, 0x7f, 0x3d,
	0x83, 0xb2, 0x69, 0xae, 0x52, 0x69, 0x5b, 0x19, 0x5c, 0xe9, 0xb9, 0x3b, 0xce, 0xc0, 0x6a, 0xdf,
	0x0a, 0xb4, 0x33, 0xd3, 0x67, 0x93, 0x9b, 0x5e, 0x48, 0x2e, 0x01, 0xd6, 0x4e, 0x76, 0xca, 0x5f,
	0xd7, 0xd3, 0x85, 0xdc, 0xc7, 0xc7, 0xd8, 0x26, 0xf5, 0x7d, 0x5c, 0x52, 0xdb, 0xe5, 0xcc, 0x01,
	0xf4, 0xa0, 0xb6, 0x72, 0x0f, 0x92, 0x3c, 0xed, 0x18, 0x5c, 0xd2, 0x29, 0x45, 0xbc, 0x3e, 0x03,
	0x31, 0x11, 0x6a, 0x5f, 0x0a, 0xb4, 0x36, 0xc3, 0x11, 0x73, 0x5d, 0xb4, 0xe2, 0x2f, 0xdd, 0x6f,
	0xa7, 0xc8, 0x84, 0x2a, 0xe9, 0x82, 0xea, 0x99, 0x1f, 0x13, 0x8e, 0x16, 0x52, 0x49, 0xb6, 0x95,
	0x13, 0xc6, 0x94, 0x55, 0xd2, 0x80, 0x3f, 0x6c, 0x3a, 0x97, 0x26, 0xd1, 0xe7, 0x44, 0xa0, 0x25,
	0x32, 0x84, 0xe6, 0x01, 0x81, 0xe4, 0x2d, 0xca, 0xce, 0x47, 0x57, 0xf7, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xb8, 0xc6, 0x47, 0x66, 0xd7, 0x03, 0x00, 0x00,
}
