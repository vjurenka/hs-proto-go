// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/invitation_types/invitation_types.proto
// DO NOT EDIT!

/*
Package invitation_types is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/invitation_types/invitation_types.proto

It has these top-level messages:
	Invitation
	Suggestion
	InvitationTarget
	InvitationParams
	SendInvitationRequest
	SendInvitationResponse
	UpdateInvitationRequest
	GenericRequest
*/
package invitation_types

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

type Invitation struct {
	Id                *uint64                   `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	InviterIdentity   *entity.Identity          `protobuf:"bytes,2,req,name=inviter_identity" json:"inviter_identity,omitempty"`
	InviteeIdentity   *entity.Identity          `protobuf:"bytes,3,req,name=invitee_identity" json:"invitee_identity,omitempty"`
	InviterName       *string                   `protobuf:"bytes,4,opt,name=inviter_name" json:"inviter_name,omitempty"`
	InviteeName       *string                   `protobuf:"bytes,5,opt,name=invitee_name" json:"invitee_name,omitempty"`
	InvitationMessage *string                   `protobuf:"bytes,6,opt,name=invitation_message" json:"invitation_message,omitempty"`
	CreationTime      *uint64                   `protobuf:"varint,7,opt,name=creation_time" json:"creation_time,omitempty"`
	ExpirationTime    *uint64                   `protobuf:"varint,8,opt,name=expiration_time" json:"expiration_time,omitempty"`
	XXX_extensions    map[int32]proto.Extension `json:"-"`
	XXX_unrecognized  []byte                    `json:"-"`
}

func (m *Invitation) Reset()                    { *m = Invitation{} }
func (m *Invitation) String() string            { return proto.CompactTextString(m) }
func (*Invitation) ProtoMessage()               {}
func (*Invitation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

var extRange_Invitation = []proto.ExtensionRange{
	{103, 105},
}

func (*Invitation) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Invitation
}
func (m *Invitation) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Invitation) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Invitation) GetInviterIdentity() *entity.Identity {
	if m != nil {
		return m.InviterIdentity
	}
	return nil
}

func (m *Invitation) GetInviteeIdentity() *entity.Identity {
	if m != nil {
		return m.InviteeIdentity
	}
	return nil
}

func (m *Invitation) GetInviterName() string {
	if m != nil && m.InviterName != nil {
		return *m.InviterName
	}
	return ""
}

func (m *Invitation) GetInviteeName() string {
	if m != nil && m.InviteeName != nil {
		return *m.InviteeName
	}
	return ""
}

func (m *Invitation) GetInvitationMessage() string {
	if m != nil && m.InvitationMessage != nil {
		return *m.InvitationMessage
	}
	return ""
}

func (m *Invitation) GetCreationTime() uint64 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

func (m *Invitation) GetExpirationTime() uint64 {
	if m != nil && m.ExpirationTime != nil {
		return *m.ExpirationTime
	}
	return 0
}

type Suggestion struct {
	ChannelId          *entity.EntityId `protobuf:"bytes,1,opt,name=channel_id" json:"channel_id,omitempty"`
	SuggesterId        *entity.EntityId `protobuf:"bytes,2,req,name=suggester_id" json:"suggester_id,omitempty"`
	SuggesteeId        *entity.EntityId `protobuf:"bytes,3,req,name=suggestee_id" json:"suggestee_id,omitempty"`
	SuggesterName      *string          `protobuf:"bytes,4,opt,name=suggester_name" json:"suggester_name,omitempty"`
	SuggesteeName      *string          `protobuf:"bytes,5,opt,name=suggestee_name" json:"suggestee_name,omitempty"`
	SuggesterAccountId *entity.EntityId `protobuf:"bytes,6,opt,name=suggester_account_id" json:"suggester_account_id,omitempty"`
	XXX_unrecognized   []byte           `json:"-"`
}

func (m *Suggestion) Reset()                    { *m = Suggestion{} }
func (m *Suggestion) String() string            { return proto.CompactTextString(m) }
func (*Suggestion) ProtoMessage()               {}
func (*Suggestion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Suggestion) GetChannelId() *entity.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *Suggestion) GetSuggesterId() *entity.EntityId {
	if m != nil {
		return m.SuggesterId
	}
	return nil
}

func (m *Suggestion) GetSuggesteeId() *entity.EntityId {
	if m != nil {
		return m.SuggesteeId
	}
	return nil
}

func (m *Suggestion) GetSuggesterName() string {
	if m != nil && m.SuggesterName != nil {
		return *m.SuggesterName
	}
	return ""
}

func (m *Suggestion) GetSuggesteeName() string {
	if m != nil && m.SuggesteeName != nil {
		return *m.SuggesteeName
	}
	return ""
}

func (m *Suggestion) GetSuggesterAccountId() *entity.EntityId {
	if m != nil {
		return m.SuggesterAccountId
	}
	return nil
}

type InvitationTarget struct {
	Identity         *entity.Identity `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	Email            *string          `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	BattleTag        *string          `protobuf:"bytes,3,opt,name=battle_tag" json:"battle_tag,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *InvitationTarget) Reset()                    { *m = InvitationTarget{} }
func (m *InvitationTarget) String() string            { return proto.CompactTextString(m) }
func (*InvitationTarget) ProtoMessage()               {}
func (*InvitationTarget) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InvitationTarget) GetIdentity() *entity.Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *InvitationTarget) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *InvitationTarget) GetBattleTag() string {
	if m != nil && m.BattleTag != nil {
		return *m.BattleTag
	}
	return ""
}

type InvitationParams struct {
	InvitationMessage *string                   `protobuf:"bytes,1,opt,name=invitation_message" json:"invitation_message,omitempty"`
	ExpirationTime    *uint64                   `protobuf:"varint,2,opt,name=expiration_time,def=0" json:"expiration_time,omitempty"`
	XXX_extensions    map[int32]proto.Extension `json:"-"`
	XXX_unrecognized  []byte                    `json:"-"`
}

func (m *InvitationParams) Reset()                    { *m = InvitationParams{} }
func (m *InvitationParams) String() string            { return proto.CompactTextString(m) }
func (*InvitationParams) ProtoMessage()               {}
func (*InvitationParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

var extRange_InvitationParams = []proto.ExtensionRange{
	{103, 105},
}

func (*InvitationParams) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_InvitationParams
}
func (m *InvitationParams) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_InvitationParams_ExpirationTime uint64 = 0

func (m *InvitationParams) GetInvitationMessage() string {
	if m != nil && m.InvitationMessage != nil {
		return *m.InvitationMessage
	}
	return ""
}

func (m *InvitationParams) GetExpirationTime() uint64 {
	if m != nil && m.ExpirationTime != nil {
		return *m.ExpirationTime
	}
	return Default_InvitationParams_ExpirationTime
}

type SendInvitationRequest struct {
	AgentIdentity    *entity.Identity    `protobuf:"bytes,1,opt,name=agent_identity" json:"agent_identity,omitempty"`
	TargetId         *entity.EntityId    `protobuf:"bytes,2,req,name=target_id" json:"target_id,omitempty"`
	Params           *InvitationParams   `protobuf:"bytes,3,req,name=params" json:"params,omitempty"`
	AgentInfo        *entity.AccountInfo `protobuf:"bytes,4,opt,name=agent_info" json:"agent_info,omitempty"`
	Target           *InvitationTarget   `protobuf:"bytes,5,opt,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *SendInvitationRequest) Reset()                    { *m = SendInvitationRequest{} }
func (m *SendInvitationRequest) String() string            { return proto.CompactTextString(m) }
func (*SendInvitationRequest) ProtoMessage()               {}
func (*SendInvitationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SendInvitationRequest) GetAgentIdentity() *entity.Identity {
	if m != nil {
		return m.AgentIdentity
	}
	return nil
}

func (m *SendInvitationRequest) GetTargetId() *entity.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *SendInvitationRequest) GetParams() *InvitationParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *SendInvitationRequest) GetAgentInfo() *entity.AccountInfo {
	if m != nil {
		return m.AgentInfo
	}
	return nil
}

func (m *SendInvitationRequest) GetTarget() *InvitationTarget {
	if m != nil {
		return m.Target
	}
	return nil
}

type SendInvitationResponse struct {
	Invitation       *Invitation `protobuf:"bytes,2,opt,name=invitation" json:"invitation,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SendInvitationResponse) Reset()                    { *m = SendInvitationResponse{} }
func (m *SendInvitationResponse) String() string            { return proto.CompactTextString(m) }
func (*SendInvitationResponse) ProtoMessage()               {}
func (*SendInvitationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SendInvitationResponse) GetInvitation() *Invitation {
	if m != nil {
		return m.Invitation
	}
	return nil
}

type UpdateInvitationRequest struct {
	AgentIdentity    *entity.Identity  `protobuf:"bytes,1,opt,name=agent_identity" json:"agent_identity,omitempty"`
	InvitationId     *uint64           `protobuf:"fixed64,2,req,name=invitation_id" json:"invitation_id,omitempty"`
	Params           *InvitationParams `protobuf:"bytes,3,req,name=params" json:"params,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *UpdateInvitationRequest) Reset()                    { *m = UpdateInvitationRequest{} }
func (m *UpdateInvitationRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateInvitationRequest) ProtoMessage()               {}
func (*UpdateInvitationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateInvitationRequest) GetAgentIdentity() *entity.Identity {
	if m != nil {
		return m.AgentIdentity
	}
	return nil
}

func (m *UpdateInvitationRequest) GetInvitationId() uint64 {
	if m != nil && m.InvitationId != nil {
		return *m.InvitationId
	}
	return 0
}

func (m *UpdateInvitationRequest) GetParams() *InvitationParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type GenericRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	TargetId         *entity.EntityId `protobuf:"bytes,2,opt,name=target_id" json:"target_id,omitempty"`
	InvitationId     *uint64          `protobuf:"fixed64,3,req,name=invitation_id" json:"invitation_id,omitempty"`
	InviteeName      *string          `protobuf:"bytes,4,opt,name=invitee_name" json:"invitee_name,omitempty"`
	InviterName      *string          `protobuf:"bytes,5,opt,name=inviter_name" json:"inviter_name,omitempty"`
	PreviousRole     []uint32         `protobuf:"varint,6,rep,packed,name=previous_role" json:"previous_role,omitempty"`
	DesiredRole      []uint32         `protobuf:"varint,7,rep,packed,name=desired_role" json:"desired_role,omitempty"`
	Reason           *uint32          `protobuf:"varint,8,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GenericRequest) Reset()                    { *m = GenericRequest{} }
func (m *GenericRequest) String() string            { return proto.CompactTextString(m) }
func (*GenericRequest) ProtoMessage()               {}
func (*GenericRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GenericRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *GenericRequest) GetTargetId() *entity.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *GenericRequest) GetInvitationId() uint64 {
	if m != nil && m.InvitationId != nil {
		return *m.InvitationId
	}
	return 0
}

func (m *GenericRequest) GetInviteeName() string {
	if m != nil && m.InviteeName != nil {
		return *m.InviteeName
	}
	return ""
}

func (m *GenericRequest) GetInviterName() string {
	if m != nil && m.InviterName != nil {
		return *m.InviterName
	}
	return ""
}

func (m *GenericRequest) GetPreviousRole() []uint32 {
	if m != nil {
		return m.PreviousRole
	}
	return nil
}

func (m *GenericRequest) GetDesiredRole() []uint32 {
	if m != nil {
		return m.DesiredRole
	}
	return nil
}

func (m *GenericRequest) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

func init() {
	proto.RegisterType((*Invitation)(nil), "invitation_types.Invitation")
	proto.RegisterType((*Suggestion)(nil), "invitation_types.Suggestion")
	proto.RegisterType((*InvitationTarget)(nil), "invitation_types.InvitationTarget")
	proto.RegisterType((*InvitationParams)(nil), "invitation_types.InvitationParams")
	proto.RegisterType((*SendInvitationRequest)(nil), "invitation_types.SendInvitationRequest")
	proto.RegisterType((*SendInvitationResponse)(nil), "invitation_types.SendInvitationResponse")
	proto.RegisterType((*UpdateInvitationRequest)(nil), "invitation_types.UpdateInvitationRequest")
	proto.RegisterType((*GenericRequest)(nil), "invitation_types.GenericRequest")
}

var fileDescriptor0 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x96, 0x93, 0xd4, 0x4d, 0x27, 0x1f, 0x6f, 0xb4, 0x6f, 0x3f, 0x42, 0xc5, 0xa1, 0x32, 0x08,
	0xa2, 0x4a, 0x4d, 0xaa, 0xdc, 0xe0, 0x06, 0x12, 0x82, 0x20, 0x21, 0x55, 0x04, 0x0e, 0x9c, 0xa2,
	0x8d, 0x3d, 0x38, 0x8b, 0xe2, 0x5d, 0xb3, 0xbb, 0xa9, 0xe8, 0x9f, 0xe0, 0x3f, 0xc2, 0x0f, 0xe0,
	0xca, 0x95, 0xf5, 0xda, 0x89, 0x8d, 0x9d, 0xb4, 0x08, 0x4e, 0x96, 0x66, 0x66, 0x9f, 0x7d, 0x3e,
	0xbc, 0x03, 0x6f, 0x42, 0xa6, 0x17, 0xab, 0xf9, 0xd0, 0x17, 0xd1, 0xe8, 0x15, 0x52, 0xa9, 0x17,
	0x53, 0x16, 0x8d, 0x16, 0xea, 0x22, 0x96, 0x42, 0x8b, 0x8b, 0x50, 0x8c, 0xe6, 0x1c, 0xf5, 0x88,
	0xf1, 0x6b, 0xa6, 0xa9, 0x66, 0x82, 0xcf, 0xf4, 0x4d, 0x8c, 0xaa, 0x52, 0x18, 0xda, 0x79, 0xd2,
	0x2b, 0xd7, 0x4f, 0x9f, 0xfc, 0xd9, 0x05, 0xc8, 0x35, 0xd3, 0x37, 0xd9, 0x27, 0x05, 0xf3, 0x7e,
	0x38, 0x00, 0x93, 0x0d, 0x1e, 0x01, 0xa8, 0xb1, 0xa0, 0xef, 0x9c, 0xd5, 0x06, 0x2e, 0x39, 0x87,
	0xf4, 0x26, 0x94, 0x33, 0x16, 0xa4, 0x87, 0xfa, 0x35, 0xd3, 0x69, 0x8d, 0x7b, 0xc3, 0x0c, 0x63,
	0x92, 0xd5, 0xf3, 0x59, 0xcc, 0x67, 0xeb, 0x3b, 0x66, 0x0f, 0xa1, 0xbd, 0xc6, 0xe5, 0x34, 0xc2,
	0x7e, 0xe3, 0xcc, 0x19, 0x1c, 0xe4, 0x55, 0x4c, 0xab, 0x7b, 0xb6, 0x7a, 0x0a, 0xa4, 0xa0, 0x36,
	0x42, 0xa5, 0x68, 0x88, 0x7d, 0xd7, 0xf6, 0x8e, 0xa0, 0xe3, 0x4b, 0xcc, 0x7c, 0x60, 0xe6, 0xc8,
	0xbe, 0x29, 0x37, 0xc8, 0x09, 0xfc, 0x87, 0x5f, 0x62, 0x26, 0x0b, 0x8d, 0x66, 0xd2, 0x38, 0x6f,
	0x34, 0xc3, 0xde, 0x27, 0xef, 0xbb, 0x11, 0x3c, 0x5d, 0x85, 0x21, 0x2a, 0x2b, 0xf8, 0x21, 0x80,
	0xbf, 0xa0, 0x9c, 0xe3, 0x72, 0x66, 0x85, 0x3b, 0x45, 0xca, 0x2f, 0xec, 0x67, 0x12, 0x90, 0x47,
	0xd0, 0x56, 0xe9, 0x19, 0x6b, 0x46, 0xd9, 0x86, 0x6d, 0x73, 0x89, 0x11, 0x65, 0x0b, 0x36, 0x73,
	0xc7, 0xd0, 0xcd, 0xf1, 0x0a, 0x26, 0x14, 0xea, 0xbf, 0xd9, 0x30, 0x84, 0xc3, 0x7c, 0x9e, 0xfa,
	0xbe, 0x58, 0x71, 0x9d, 0xe0, 0xbb, 0xdb, 0xf9, 0x7a, 0x1f, 0xa0, 0x97, 0x87, 0xfa, 0x8e, 0xca,
	0x10, 0x35, 0xf1, 0xa0, 0xb9, 0x89, 0xa6, 0xa4, 0x73, 0x13, 0x4d, 0x07, 0xf6, 0x30, 0xa2, 0x6c,
	0x69, 0x04, 0x26, 0xd7, 0x9a, 0xdf, 0x61, 0x4e, 0xb5, 0x5e, 0xe2, 0x4c, 0xd3, 0xd0, 0x88, 0x31,
	0x35, 0xef, 0xaa, 0x08, 0x7d, 0x45, 0x25, 0x8d, 0xd4, 0x8e, 0x94, 0x9c, 0x2c, 0xc1, 0x4a, 0x1c,
	0x09, 0x78, 0xe3, 0xa9, 0x73, 0x99, 0x25, 0xf2, 0xd3, 0x81, 0xa3, 0x29, 0xf2, 0x20, 0x87, 0x7d,
	0x8b, 0x9f, 0x57, 0x46, 0x2b, 0x19, 0x40, 0xd7, 0x00, 0x59, 0xa9, 0x77, 0x10, 0x7f, 0x00, 0x07,
	0xda, 0xca, 0xbc, 0x2d, 0x9d, 0x31, 0xb8, 0xb1, 0x25, 0x9c, 0xe5, 0xe2, 0x0d, 0x2b, 0x2f, 0xac,
	0x22, 0xed, 0x31, 0x40, 0x46, 0x81, 0x7f, 0x14, 0x36, 0xa5, 0xd6, 0xf8, 0xff, 0x35, 0xf2, 0xb3,
	0x34, 0x89, 0x89, 0x69, 0x25, 0xe0, 0x29, 0x03, 0x1b, 0xd9, 0x1d, 0xe0, 0x69, 0x24, 0xde, 0x6b,
	0x38, 0x2e, 0x0b, 0x57, 0xb1, 0xe0, 0x0a, 0xc9, 0x25, 0x40, 0x7e, 0xdc, 0x1a, 0xd6, 0x1a, 0xdf,
	0xbf, 0x0d, 0xd1, 0xfb, 0xea, 0xc0, 0xc9, 0xfb, 0x38, 0xa0, 0x1a, 0xff, 0xc5, 0x47, 0xf3, 0xa6,
	0x0a, 0x97, 0x64, 0x5e, 0xba, 0x7f, 0xe3, 0x9c, 0xf7, 0xcd, 0x81, 0xee, 0x4b, 0xe4, 0x28, 0x99,
	0xbf, 0xe6, 0x61, 0x7e, 0xc1, 0x35, 0x8f, 0x9d, 0x4f, 0xad, 0x94, 0xe4, 0xf6, 0xa1, 0x0a, 0xcd,
	0xba, 0xa5, 0x59, 0xde, 0x21, 0xa5, 0xcd, 0x22, 0x8b, 0x4f, 0xea, 0x1e, 0x74, 0x62, 0x89, 0xd7,
	0x4c, 0xac, 0xd4, 0x4c, 0x8a, 0x65, 0xb2, 0x54, 0xea, 0x83, 0xce, 0xf3, 0x5a, 0xcf, 0x21, 0x7d,
	0x68, 0x07, 0xa8, 0x98, 0xc4, 0x20, 0xed, 0xec, 0x6f, 0x3a, 0x5d, 0x70, 0xcd, 0xc6, 0x51, 0x26,
	0x92, 0x64, 0xa5, 0x74, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x10, 0xc1, 0x0c, 0x97, 0xda, 0x05,
	0x00, 0x00,
}
