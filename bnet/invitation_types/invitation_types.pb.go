// Code generated by protoc-gen-go.
// source: github.com/vjurenka/hs-proto-go/bnet/invitation_types/invitation_types.proto
// DO NOT EDIT!

/*
Package invitation_types is a generated protocol buffer package.

It is generated from these files:
	github.com/vjurenka/hs-proto-go/bnet/invitation_types/invitation_types.proto

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
import math "math"
import entity "github.com/vjurenka/hs-proto-go/bnet/entity"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

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

func (m *Invitation) Reset()         { *m = Invitation{} }
func (m *Invitation) String() string { return proto.CompactTextString(m) }
func (*Invitation) ProtoMessage()    {}

var extRange_Invitation = []proto.ExtensionRange{
	{100, 10000},
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

func (m *Suggestion) Reset()         { *m = Suggestion{} }
func (m *Suggestion) String() string { return proto.CompactTextString(m) }
func (*Suggestion) ProtoMessage()    {}

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

func (m *InvitationTarget) Reset()         { *m = InvitationTarget{} }
func (m *InvitationTarget) String() string { return proto.CompactTextString(m) }
func (*InvitationTarget) ProtoMessage()    {}

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

func (m *InvitationParams) Reset()         { *m = InvitationParams{} }
func (m *InvitationParams) String() string { return proto.CompactTextString(m) }
func (*InvitationParams) ProtoMessage()    {}

var extRange_InvitationParams = []proto.ExtensionRange{
	{100, 10000},
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

func (m *SendInvitationRequest) Reset()         { *m = SendInvitationRequest{} }
func (m *SendInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*SendInvitationRequest) ProtoMessage()    {}

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

func (m *SendInvitationResponse) Reset()         { *m = SendInvitationResponse{} }
func (m *SendInvitationResponse) String() string { return proto.CompactTextString(m) }
func (*SendInvitationResponse) ProtoMessage()    {}

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

func (m *UpdateInvitationRequest) Reset()         { *m = UpdateInvitationRequest{} }
func (m *UpdateInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInvitationRequest) ProtoMessage()    {}

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

func (m *GenericRequest) Reset()         { *m = GenericRequest{} }
func (m *GenericRequest) String() string { return proto.CompactTextString(m) }
func (*GenericRequest) ProtoMessage()    {}

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
