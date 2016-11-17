// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/pegasus/client/client.proto
// DO NOT EDIT!

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/pegasus/client/client.proto

It has these top-level messages:
	SessionRecord
*/
package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type SessionRecordType int32

const (
	SessionRecordType_ARENA        SessionRecordType = 0
	SessionRecordType_TAVERN_BRAWL SessionRecordType = 1
)

var SessionRecordType_name = map[int32]string{
	0: "ARENA",
	1: "TAVERN_BRAWL",
}
var SessionRecordType_value = map[string]int32{
	"ARENA":        0,
	"TAVERN_BRAWL": 1,
}

func (x SessionRecordType) Enum() *SessionRecordType {
	p := new(SessionRecordType)
	*p = x
	return p
}
func (x SessionRecordType) String() string {
	return proto.EnumName(SessionRecordType_name, int32(x))
}
func (x *SessionRecordType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SessionRecordType_value, data, "SessionRecordType")
	if err != nil {
		return err
	}
	*x = SessionRecordType(value)
	return nil
}
func (SessionRecordType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SessionRecord struct {
	Wins              *uint32            `protobuf:"varint,1,req,name=wins" json:"wins,omitempty"`
	Losses            *uint32            `protobuf:"varint,2,req,name=losses" json:"losses,omitempty"`
	RunFinished       *bool              `protobuf:"varint,3,req,name=run_finished" json:"run_finished,omitempty"`
	SessionRecordType *SessionRecordType `protobuf:"varint,4,req,name=session_record_type,enum=client.SessionRecordType" json:"session_record_type,omitempty"`
	XXX_unrecognized  []byte             `json:"-"`
}

func (m *SessionRecord) Reset()                    { *m = SessionRecord{} }
func (m *SessionRecord) String() string            { return proto.CompactTextString(m) }
func (*SessionRecord) ProtoMessage()               {}
func (*SessionRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SessionRecord) GetWins() uint32 {
	if m != nil && m.Wins != nil {
		return *m.Wins
	}
	return 0
}

func (m *SessionRecord) GetLosses() uint32 {
	if m != nil && m.Losses != nil {
		return *m.Losses
	}
	return 0
}

func (m *SessionRecord) GetRunFinished() bool {
	if m != nil && m.RunFinished != nil {
		return *m.RunFinished
	}
	return false
}

func (m *SessionRecord) GetSessionRecordType() SessionRecordType {
	if m != nil && m.SessionRecordType != nil {
		return *m.SessionRecordType
	}
	return SessionRecordType_ARENA
}

func init() {
	proto.RegisterType((*SessionRecord)(nil), "client.SessionRecord")
	proto.RegisterEnum("client.SessionRecordType", SessionRecordType_name, SessionRecordType_value)
}

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xf7, 0x48, 0x4d, 0x2c, 0x2a, 0xc9, 0x08, 0xce, 0xcc,
	0xd5, 0xcf, 0x28, 0xd6, 0x2d, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4d, 0xcf, 0xd7, 0x2f, 0x48, 0x4d,
	0x4f, 0x2c, 0x2e, 0x2d, 0xd6, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x81, 0x52, 0x7a, 0x60, 0x69,
	0x21, 0x36, 0x08, 0x4f, 0xa9, 0x9a, 0x8b, 0x37, 0x38, 0xb5, 0xb8, 0x38, 0x33, 0x3f, 0x2f, 0x28,
	0x35, 0x39, 0xbf, 0x28, 0x45, 0x88, 0x87, 0x8b, 0xa5, 0x3c, 0x33, 0xaf, 0x58, 0x82, 0x51, 0x81,
	0x49, 0x83, 0x57, 0x88, 0x8f, 0x8b, 0x2d, 0x27, 0xbf, 0xb8, 0x38, 0xb5, 0x58, 0x82, 0x09, 0xcc,
	0x17, 0xe1, 0xe2, 0x29, 0x2a, 0xcd, 0x8b, 0x4f, 0xcb, 0xcc, 0xcb, 0x2c, 0xce, 0x48, 0x4d, 0x91,
	0x60, 0x06, 0x8a, 0x72, 0x08, 0x99, 0x71, 0x09, 0x17, 0x43, 0x0c, 0x89, 0x2f, 0x02, 0x9b, 0x12,
	0x5f, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0x02, 0x94, 0xe4, 0x33, 0x92, 0xd4, 0x83, 0x5a, 0x8c, 0x62,
	0x4f, 0x08, 0x50, 0x81, 0x96, 0x01, 0x97, 0x20, 0x86, 0xa0, 0x10, 0x27, 0x17, 0xab, 0x63, 0x90,
	0xab, 0x9f, 0xa3, 0x00, 0x83, 0x90, 0x00, 0x17, 0x4f, 0x88, 0x63, 0x98, 0x6b, 0x90, 0x5f, 0xbc,
	0x53, 0x90, 0x63, 0xb8, 0x8f, 0x00, 0x23, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x49, 0x8f, 0x3d,
	0xf5, 0x00, 0x00, 0x00,
}
