// Code generated by protoc-gen-go.
// source: github.com/vjurenka/hs-proto-go/bnet/game_master_types/game_master_types.proto
// DO NOT EDIT!

/*
Package game_master_types is a generated protocol buffer package.

It is generated from these files:
	github.com/vjurenka/hs-proto-go/bnet/game_master_types/game_master_types.proto

It has these top-level messages:
	Player
	ConnectInfo
	GameStatsBucket
	GameFactoryDescription
	GameHandle
	CancelGameEntryRequest
*/
package game_master_types

import proto "github.com/golang/protobuf/proto"
import math "math"
import attribute "github.com/vjurenka/hs-proto-go/bnet/attribute"
import entity "github.com/vjurenka/hs-proto-go/bnet/entity"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Player struct {
	Identity         *entity.Identity       `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}

func (m *Player) GetIdentity() *entity.Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Player) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type ConnectInfo struct {
	MemberId         *entity.EntityId       `protobuf:"bytes,1,req,name=member_id" json:"member_id,omitempty"`
	Host             *string                `protobuf:"bytes,2,req,name=host" json:"host,omitempty"`
	Port             *int32                 `protobuf:"varint,3,req,name=port" json:"port,omitempty"`
	Token            []byte                 `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,5,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ConnectInfo) Reset()         { *m = ConnectInfo{} }
func (m *ConnectInfo) String() string { return proto.CompactTextString(m) }
func (*ConnectInfo) ProtoMessage()    {}

func (m *ConnectInfo) GetMemberId() *entity.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *ConnectInfo) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *ConnectInfo) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ConnectInfo) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *ConnectInfo) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type GameStatsBucket struct {
	BucketMin                  *float32 `protobuf:"fixed32,1,opt,name=bucket_min,def=0" json:"bucket_min,omitempty"`
	BucketMax                  *float32 `protobuf:"fixed32,2,opt,name=bucket_max,def=4.294967e+009" json:"bucket_max,omitempty"`
	WaitMilliseconds           *uint32  `protobuf:"varint,3,opt,name=wait_milliseconds,def=0" json:"wait_milliseconds,omitempty"`
	GamesPerHour               *uint32  `protobuf:"varint,4,opt,name=games_per_hour,def=0" json:"games_per_hour,omitempty"`
	ActiveGames                *uint32  `protobuf:"varint,5,opt,name=active_games,def=0" json:"active_games,omitempty"`
	ActivePlayers              *uint32  `protobuf:"varint,6,opt,name=active_players,def=0" json:"active_players,omitempty"`
	FormingGames               *uint32  `protobuf:"varint,7,opt,name=forming_games,def=0" json:"forming_games,omitempty"`
	WaitingPlayers             *uint32  `protobuf:"varint,8,opt,name=waiting_players,def=0" json:"waiting_players,omitempty"`
	OpenJoinableGames          *uint32  `protobuf:"varint,9,opt,name=open_joinable_games,def=0" json:"open_joinable_games,omitempty"`
	PlayersInOpenJoinableGames *uint32  `protobuf:"varint,10,opt,name=players_in_open_joinable_games,def=0" json:"players_in_open_joinable_games,omitempty"`
	OpenGamesTotal             *uint32  `protobuf:"varint,11,opt,name=open_games_total,def=0" json:"open_games_total,omitempty"`
	PlayersInOpenGamesTotal    *uint32  `protobuf:"varint,12,opt,name=players_in_open_games_total,def=0" json:"players_in_open_games_total,omitempty"`
	XXX_unrecognized           []byte   `json:"-"`
}

func (m *GameStatsBucket) Reset()         { *m = GameStatsBucket{} }
func (m *GameStatsBucket) String() string { return proto.CompactTextString(m) }
func (*GameStatsBucket) ProtoMessage()    {}

const Default_GameStatsBucket_BucketMin float32 = 0
const Default_GameStatsBucket_BucketMax float32 = 4.294967e+009
const Default_GameStatsBucket_WaitMilliseconds uint32 = 0
const Default_GameStatsBucket_GamesPerHour uint32 = 0
const Default_GameStatsBucket_ActiveGames uint32 = 0
const Default_GameStatsBucket_ActivePlayers uint32 = 0
const Default_GameStatsBucket_FormingGames uint32 = 0
const Default_GameStatsBucket_WaitingPlayers uint32 = 0
const Default_GameStatsBucket_OpenJoinableGames uint32 = 0
const Default_GameStatsBucket_PlayersInOpenJoinableGames uint32 = 0
const Default_GameStatsBucket_OpenGamesTotal uint32 = 0
const Default_GameStatsBucket_PlayersInOpenGamesTotal uint32 = 0

func (m *GameStatsBucket) GetBucketMin() float32 {
	if m != nil && m.BucketMin != nil {
		return *m.BucketMin
	}
	return Default_GameStatsBucket_BucketMin
}

func (m *GameStatsBucket) GetBucketMax() float32 {
	if m != nil && m.BucketMax != nil {
		return *m.BucketMax
	}
	return Default_GameStatsBucket_BucketMax
}

func (m *GameStatsBucket) GetWaitMilliseconds() uint32 {
	if m != nil && m.WaitMilliseconds != nil {
		return *m.WaitMilliseconds
	}
	return Default_GameStatsBucket_WaitMilliseconds
}

func (m *GameStatsBucket) GetGamesPerHour() uint32 {
	if m != nil && m.GamesPerHour != nil {
		return *m.GamesPerHour
	}
	return Default_GameStatsBucket_GamesPerHour
}

func (m *GameStatsBucket) GetActiveGames() uint32 {
	if m != nil && m.ActiveGames != nil {
		return *m.ActiveGames
	}
	return Default_GameStatsBucket_ActiveGames
}

func (m *GameStatsBucket) GetActivePlayers() uint32 {
	if m != nil && m.ActivePlayers != nil {
		return *m.ActivePlayers
	}
	return Default_GameStatsBucket_ActivePlayers
}

func (m *GameStatsBucket) GetFormingGames() uint32 {
	if m != nil && m.FormingGames != nil {
		return *m.FormingGames
	}
	return Default_GameStatsBucket_FormingGames
}

func (m *GameStatsBucket) GetWaitingPlayers() uint32 {
	if m != nil && m.WaitingPlayers != nil {
		return *m.WaitingPlayers
	}
	return Default_GameStatsBucket_WaitingPlayers
}

func (m *GameStatsBucket) GetOpenJoinableGames() uint32 {
	if m != nil && m.OpenJoinableGames != nil {
		return *m.OpenJoinableGames
	}
	return Default_GameStatsBucket_OpenJoinableGames
}

func (m *GameStatsBucket) GetPlayersInOpenJoinableGames() uint32 {
	if m != nil && m.PlayersInOpenJoinableGames != nil {
		return *m.PlayersInOpenJoinableGames
	}
	return Default_GameStatsBucket_PlayersInOpenJoinableGames
}

func (m *GameStatsBucket) GetOpenGamesTotal() uint32 {
	if m != nil && m.OpenGamesTotal != nil {
		return *m.OpenGamesTotal
	}
	return Default_GameStatsBucket_OpenGamesTotal
}

func (m *GameStatsBucket) GetPlayersInOpenGamesTotal() uint32 {
	if m != nil && m.PlayersInOpenGamesTotal != nil {
		return *m.PlayersInOpenGamesTotal
	}
	return Default_GameStatsBucket_PlayersInOpenGamesTotal
}

type GameFactoryDescription struct {
	Id               *uint64                `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	Name             *string                `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,3,rep,name=attribute" json:"attribute,omitempty"`
	StatsBucket      []*GameStatsBucket     `protobuf:"bytes,4,rep,name=stats_bucket" json:"stats_bucket,omitempty"`
	UnseededId       *uint64                `protobuf:"fixed64,5,opt,name=unseeded_id,def=0" json:"unseeded_id,omitempty"`
	AllowQueueing    *bool                  `protobuf:"varint,6,opt,name=allow_queueing,def=1" json:"allow_queueing,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *GameFactoryDescription) Reset()         { *m = GameFactoryDescription{} }
func (m *GameFactoryDescription) String() string { return proto.CompactTextString(m) }
func (*GameFactoryDescription) ProtoMessage()    {}

const Default_GameFactoryDescription_UnseededId uint64 = 0
const Default_GameFactoryDescription_AllowQueueing bool = true

func (m *GameFactoryDescription) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GameFactoryDescription) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GameFactoryDescription) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *GameFactoryDescription) GetStatsBucket() []*GameStatsBucket {
	if m != nil {
		return m.StatsBucket
	}
	return nil
}

func (m *GameFactoryDescription) GetUnseededId() uint64 {
	if m != nil && m.UnseededId != nil {
		return *m.UnseededId
	}
	return Default_GameFactoryDescription_UnseededId
}

func (m *GameFactoryDescription) GetAllowQueueing() bool {
	if m != nil && m.AllowQueueing != nil {
		return *m.AllowQueueing
	}
	return Default_GameFactoryDescription_AllowQueueing
}

type GameHandle struct {
	FactoryId        *uint64          `protobuf:"fixed64,1,req,name=factory_id" json:"factory_id,omitempty"`
	GameId           *entity.EntityId `protobuf:"bytes,2,req,name=game_id" json:"game_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GameHandle) Reset()         { *m = GameHandle{} }
func (m *GameHandle) String() string { return proto.CompactTextString(m) }
func (*GameHandle) ProtoMessage()    {}

func (m *GameHandle) GetFactoryId() uint64 {
	if m != nil && m.FactoryId != nil {
		return *m.FactoryId
	}
	return 0
}

func (m *GameHandle) GetGameId() *entity.EntityId {
	if m != nil {
		return m.GameId
	}
	return nil
}

type CancelGameEntryRequest struct {
	RequestId        *uint64   `protobuf:"fixed64,1,req,name=request_id" json:"request_id,omitempty"`
	FactoryId        *uint64   `protobuf:"fixed64,2,opt,name=factory_id" json:"factory_id,omitempty"`
	Player           []*Player `protobuf:"bytes,3,rep,name=player" json:"player,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *CancelGameEntryRequest) Reset()         { *m = CancelGameEntryRequest{} }
func (m *CancelGameEntryRequest) String() string { return proto.CompactTextString(m) }
func (*CancelGameEntryRequest) ProtoMessage()    {}

func (m *CancelGameEntryRequest) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

func (m *CancelGameEntryRequest) GetFactoryId() uint64 {
	if m != nil && m.FactoryId != nil {
		return *m.FactoryId
	}
	return 0
}

func (m *CancelGameEntryRequest) GetPlayer() []*Player {
	if m != nil {
		return m.Player
	}
	return nil
}
