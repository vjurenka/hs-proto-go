// Code generated by protoc-gen-go.
// source: github.com/vjurenka/hs-proto-go/bnet/channel_invitation_service/channel_invitation_service.proto
// DO NOT EDIT!

/*
Package channel_invitation_service is a generated protocol buffer package.

It is generated from these files:
	github.com/vjurenka/hs-proto-go/bnet/channel_invitation_service/channel_invitation_service.proto

It has these top-level messages:
	AcceptInvitationRequest
	AcceptInvitationResponse
	SubscribeRequest
	SubscribeResponse
	UnsubscribeRequest
	SuggestInvitationRequest
	RevokeInvitationRequest
	HasRoomForInvitationRequest
	ChannelCountDescription
	IncrementChannelCountRequest
	IncrementChannelCountResponse
	DecrementChannelCountRequest
	UpdateChannelCountRequest
	ListChannelCountRequest
	ChannelCount
	ListChannelCountResponse
	InvitationAddedNotification
	InvitationRemovedNotification
	SuggestionAddedNotification
*/
package channel_invitation_service

import proto "github.com/golang/protobuf/proto"
import math "math"
import channel_invitation_types "github.com/vjurenka/hs-proto-go/bnet/channel_invitation_types"
import channel_types "github.com/vjurenka/hs-proto-go/bnet/channel_types"
import entity "github.com/vjurenka/hs-proto-go/bnet/entity"
import invitation_types "github.com/vjurenka/hs-proto-go/bnet/invitation_types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type AcceptInvitationRequest struct {
	AgentId          *entity.EntityId           `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	MemberState      *channel_types.MemberState `protobuf:"bytes,2,opt,name=member_state" json:"member_state,omitempty"`
	InvitationId     *uint64                    `protobuf:"fixed64,3,req,name=invitation_id" json:"invitation_id,omitempty"`
	ObjectId         *uint64                    `protobuf:"varint,4,req,name=object_id" json:"object_id,omitempty"`
	ChannelId        *entity.EntityId           `protobuf:"bytes,5,opt,name=channel_id" json:"channel_id,omitempty"`
	ServiceType      *uint32                    `protobuf:"varint,6,opt,name=service_type" json:"service_type,omitempty"`
	LocalSubscriber  *bool                      `protobuf:"varint,7,opt,name=local_subscriber,def=1" json:"local_subscriber,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *AcceptInvitationRequest) Reset()         { *m = AcceptInvitationRequest{} }
func (m *AcceptInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*AcceptInvitationRequest) ProtoMessage()    {}

const Default_AcceptInvitationRequest_LocalSubscriber bool = true

func (m *AcceptInvitationRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *AcceptInvitationRequest) GetMemberState() *channel_types.MemberState {
	if m != nil {
		return m.MemberState
	}
	return nil
}

func (m *AcceptInvitationRequest) GetInvitationId() uint64 {
	if m != nil && m.InvitationId != nil {
		return *m.InvitationId
	}
	return 0
}

func (m *AcceptInvitationRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *AcceptInvitationRequest) GetChannelId() *entity.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *AcceptInvitationRequest) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

func (m *AcceptInvitationRequest) GetLocalSubscriber() bool {
	if m != nil && m.LocalSubscriber != nil {
		return *m.LocalSubscriber
	}
	return Default_AcceptInvitationRequest_LocalSubscriber
}

type AcceptInvitationResponse struct {
	ObjectId         *uint64 `protobuf:"varint,1,req,name=object_id" json:"object_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AcceptInvitationResponse) Reset()         { *m = AcceptInvitationResponse{} }
func (m *AcceptInvitationResponse) String() string { return proto.CompactTextString(m) }
func (*AcceptInvitationResponse) ProtoMessage()    {}

func (m *AcceptInvitationResponse) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

type SubscribeRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	ObjectId         *uint64          `protobuf:"varint,2,req,name=object_id" json:"object_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}

func (m *SubscribeRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SubscribeRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

type SubscribeResponse struct {
	Collection         []*channel_invitation_types.InvitationCollection `protobuf:"bytes,1,rep,name=collection" json:"collection,omitempty"`
	ReceivedInvitation []*invitation_types.Invitation                   `protobuf:"bytes,2,rep,name=received_invitation" json:"received_invitation,omitempty"`
	XXX_unrecognized   []byte                                           `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}

func (m *SubscribeResponse) GetCollection() []*channel_invitation_types.InvitationCollection {
	if m != nil {
		return m.Collection
	}
	return nil
}

func (m *SubscribeResponse) GetReceivedInvitation() []*invitation_types.Invitation {
	if m != nil {
		return m.ReceivedInvitation
	}
	return nil
}

type UnsubscribeRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *UnsubscribeRequest) Reset()         { *m = UnsubscribeRequest{} }
func (m *UnsubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeRequest) ProtoMessage()    {}

func (m *UnsubscribeRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

type SuggestInvitationRequest struct {
	AgentId          *entity.EntityId    `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	ChannelId        *entity.EntityId    `protobuf:"bytes,2,req,name=channel_id" json:"channel_id,omitempty"`
	TargetId         *entity.EntityId    `protobuf:"bytes,3,req,name=target_id" json:"target_id,omitempty"`
	ApprovalId       *entity.EntityId    `protobuf:"bytes,4,opt,name=approval_id" json:"approval_id,omitempty"`
	AgentIdentity    *entity.Identity    `protobuf:"bytes,5,opt,name=agent_identity" json:"agent_identity,omitempty"`
	AgentInfo        *entity.AccountInfo `protobuf:"bytes,6,opt,name=agent_info" json:"agent_info,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *SuggestInvitationRequest) Reset()         { *m = SuggestInvitationRequest{} }
func (m *SuggestInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*SuggestInvitationRequest) ProtoMessage()    {}

func (m *SuggestInvitationRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SuggestInvitationRequest) GetChannelId() *entity.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *SuggestInvitationRequest) GetTargetId() *entity.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *SuggestInvitationRequest) GetApprovalId() *entity.EntityId {
	if m != nil {
		return m.ApprovalId
	}
	return nil
}

func (m *SuggestInvitationRequest) GetAgentIdentity() *entity.Identity {
	if m != nil {
		return m.AgentIdentity
	}
	return nil
}

func (m *SuggestInvitationRequest) GetAgentInfo() *entity.AccountInfo {
	if m != nil {
		return m.AgentInfo
	}
	return nil
}

type RevokeInvitationRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	TargetId         *entity.EntityId `protobuf:"bytes,2,opt,name=target_id" json:"target_id,omitempty"`
	InvitationId     *uint64          `protobuf:"fixed64,3,req,name=invitation_id" json:"invitation_id,omitempty"`
	ChannelId        *entity.EntityId `protobuf:"bytes,4,req,name=channel_id" json:"channel_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *RevokeInvitationRequest) Reset()         { *m = RevokeInvitationRequest{} }
func (m *RevokeInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeInvitationRequest) ProtoMessage()    {}

func (m *RevokeInvitationRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *RevokeInvitationRequest) GetTargetId() *entity.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *RevokeInvitationRequest) GetInvitationId() uint64 {
	if m != nil && m.InvitationId != nil {
		return *m.InvitationId
	}
	return 0
}

func (m *RevokeInvitationRequest) GetChannelId() *entity.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

type HasRoomForInvitationRequest struct {
	ServiceType      *uint32 `protobuf:"varint,1,req,name=service_type" json:"service_type,omitempty"`
	Program          *uint32 `protobuf:"fixed32,2,opt,name=program" json:"program,omitempty"`
	ChannelType      *string `protobuf:"bytes,3,opt,name=channel_type,def=default" json:"channel_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HasRoomForInvitationRequest) Reset()         { *m = HasRoomForInvitationRequest{} }
func (m *HasRoomForInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*HasRoomForInvitationRequest) ProtoMessage()    {}

const Default_HasRoomForInvitationRequest_ChannelType string = "default"

func (m *HasRoomForInvitationRequest) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

func (m *HasRoomForInvitationRequest) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *HasRoomForInvitationRequest) GetChannelType() string {
	if m != nil && m.ChannelType != nil {
		return *m.ChannelType
	}
	return Default_HasRoomForInvitationRequest_ChannelType
}

type ChannelCountDescription struct {
	ServiceType      *uint32          `protobuf:"varint,1,req,name=service_type" json:"service_type,omitempty"`
	Program          *uint32          `protobuf:"fixed32,2,req,name=program" json:"program,omitempty"`
	ChannelType      *string          `protobuf:"bytes,3,opt,name=channel_type,def=default" json:"channel_type,omitempty"`
	ChannelId        *entity.EntityId `protobuf:"bytes,4,opt,name=channel_id" json:"channel_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ChannelCountDescription) Reset()         { *m = ChannelCountDescription{} }
func (m *ChannelCountDescription) String() string { return proto.CompactTextString(m) }
func (*ChannelCountDescription) ProtoMessage()    {}

const Default_ChannelCountDescription_ChannelType string = "default"

func (m *ChannelCountDescription) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

func (m *ChannelCountDescription) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *ChannelCountDescription) GetChannelType() string {
	if m != nil && m.ChannelType != nil {
		return *m.ChannelType
	}
	return Default_ChannelCountDescription_ChannelType
}

func (m *ChannelCountDescription) GetChannelId() *entity.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

type IncrementChannelCountRequest struct {
	AgentId          *entity.EntityId           `protobuf:"bytes,1,req,name=agent_id" json:"agent_id,omitempty"`
	Descriptions     []*ChannelCountDescription `protobuf:"bytes,2,rep,name=descriptions" json:"descriptions,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IncrementChannelCountRequest) Reset()         { *m = IncrementChannelCountRequest{} }
func (m *IncrementChannelCountRequest) String() string { return proto.CompactTextString(m) }
func (*IncrementChannelCountRequest) ProtoMessage()    {}

func (m *IncrementChannelCountRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *IncrementChannelCountRequest) GetDescriptions() []*ChannelCountDescription {
	if m != nil {
		return m.Descriptions
	}
	return nil
}

type IncrementChannelCountResponse struct {
	ReservationTokens []uint64 `protobuf:"varint,1,rep,name=reservation_tokens" json:"reservation_tokens,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *IncrementChannelCountResponse) Reset()         { *m = IncrementChannelCountResponse{} }
func (m *IncrementChannelCountResponse) String() string { return proto.CompactTextString(m) }
func (*IncrementChannelCountResponse) ProtoMessage()    {}

func (m *IncrementChannelCountResponse) GetReservationTokens() []uint64 {
	if m != nil {
		return m.ReservationTokens
	}
	return nil
}

type DecrementChannelCountRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,req,name=agent_id" json:"agent_id,omitempty"`
	ChannelId        *entity.EntityId `protobuf:"bytes,2,opt,name=channel_id" json:"channel_id,omitempty"`
	ReservationToken *uint64          `protobuf:"varint,3,opt,name=reservation_token" json:"reservation_token,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *DecrementChannelCountRequest) Reset()         { *m = DecrementChannelCountRequest{} }
func (m *DecrementChannelCountRequest) String() string { return proto.CompactTextString(m) }
func (*DecrementChannelCountRequest) ProtoMessage()    {}

func (m *DecrementChannelCountRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *DecrementChannelCountRequest) GetChannelId() *entity.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *DecrementChannelCountRequest) GetReservationToken() uint64 {
	if m != nil && m.ReservationToken != nil {
		return *m.ReservationToken
	}
	return 0
}

type UpdateChannelCountRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,req,name=agent_id" json:"agent_id,omitempty"`
	ReservationToken *uint64          `protobuf:"varint,2,opt,name=reservation_token" json:"reservation_token,omitempty"`
	ChannelId        *entity.EntityId `protobuf:"bytes,3,req,name=channel_id" json:"channel_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *UpdateChannelCountRequest) Reset()         { *m = UpdateChannelCountRequest{} }
func (m *UpdateChannelCountRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateChannelCountRequest) ProtoMessage()    {}

func (m *UpdateChannelCountRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateChannelCountRequest) GetReservationToken() uint64 {
	if m != nil && m.ReservationToken != nil {
		return *m.ReservationToken
	}
	return 0
}

func (m *UpdateChannelCountRequest) GetChannelId() *entity.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

type ListChannelCountRequest struct {
	MemberId         *entity.EntityId `protobuf:"bytes,1,req,name=member_id" json:"member_id,omitempty"`
	ServiceType      *uint32          `protobuf:"varint,2,req,name=service_type" json:"service_type,omitempty"`
	Program          *uint32          `protobuf:"fixed32,3,opt,name=program" json:"program,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ListChannelCountRequest) Reset()         { *m = ListChannelCountRequest{} }
func (m *ListChannelCountRequest) String() string { return proto.CompactTextString(m) }
func (*ListChannelCountRequest) ProtoMessage()    {}

func (m *ListChannelCountRequest) GetMemberId() *entity.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *ListChannelCountRequest) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

func (m *ListChannelCountRequest) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

type ChannelCount struct {
	ChannelId        *entity.EntityId `protobuf:"bytes,1,opt,name=channel_id" json:"channel_id,omitempty"`
	ChannelType      *string          `protobuf:"bytes,2,opt,name=channel_type,def=default" json:"channel_type,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ChannelCount) Reset()         { *m = ChannelCount{} }
func (m *ChannelCount) String() string { return proto.CompactTextString(m) }
func (*ChannelCount) ProtoMessage()    {}

const Default_ChannelCount_ChannelType string = "default"

func (m *ChannelCount) GetChannelId() *entity.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *ChannelCount) GetChannelType() string {
	if m != nil && m.ChannelType != nil {
		return *m.ChannelType
	}
	return Default_ChannelCount_ChannelType
}

type ListChannelCountResponse struct {
	Channel          []*ChannelCount `protobuf:"bytes,1,rep,name=channel" json:"channel,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ListChannelCountResponse) Reset()         { *m = ListChannelCountResponse{} }
func (m *ListChannelCountResponse) String() string { return proto.CompactTextString(m) }
func (*ListChannelCountResponse) ProtoMessage()    {}

func (m *ListChannelCountResponse) GetChannel() []*ChannelCount {
	if m != nil {
		return m.Channel
	}
	return nil
}

type InvitationAddedNotification struct {
	Invitation       *invitation_types.Invitation `protobuf:"bytes,1,req,name=invitation" json:"invitation,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *InvitationAddedNotification) Reset()         { *m = InvitationAddedNotification{} }
func (m *InvitationAddedNotification) String() string { return proto.CompactTextString(m) }
func (*InvitationAddedNotification) ProtoMessage()    {}

func (m *InvitationAddedNotification) GetInvitation() *invitation_types.Invitation {
	if m != nil {
		return m.Invitation
	}
	return nil
}

type InvitationRemovedNotification struct {
	Invitation       *invitation_types.Invitation `protobuf:"bytes,1,req,name=invitation" json:"invitation,omitempty"`
	Reason           *uint32                      `protobuf:"varint,2,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *InvitationRemovedNotification) Reset()         { *m = InvitationRemovedNotification{} }
func (m *InvitationRemovedNotification) String() string { return proto.CompactTextString(m) }
func (*InvitationRemovedNotification) ProtoMessage()    {}

func (m *InvitationRemovedNotification) GetInvitation() *invitation_types.Invitation {
	if m != nil {
		return m.Invitation
	}
	return nil
}

func (m *InvitationRemovedNotification) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

type SuggestionAddedNotification struct {
	Suggestion       *invitation_types.Suggestion `protobuf:"bytes,1,req,name=suggestion" json:"suggestion,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *SuggestionAddedNotification) Reset()         { *m = SuggestionAddedNotification{} }
func (m *SuggestionAddedNotification) String() string { return proto.CompactTextString(m) }
func (*SuggestionAddedNotification) ProtoMessage()    {}

func (m *SuggestionAddedNotification) GetSuggestion() *invitation_types.Suggestion {
	if m != nil {
		return m.Suggestion
	}
	return nil
}
