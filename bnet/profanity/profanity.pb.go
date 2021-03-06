// Code generated by protoc-gen-go.
// source: github.com/vjurenka/hs-proto-go/bnet/profanity/profanity.proto
// DO NOT EDIT!

/*
Package profanity is a generated protocol buffer package.

It is generated from these files:
	github.com/vjurenka/hs-proto-go/bnet/profanity/profanity.proto

It has these top-level messages:
	WordFilter
	WordFilters
*/
package profanity

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type WordFilter struct {
	Type             *string `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Regex            *string `protobuf:"bytes,2,req,name=regex" json:"regex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WordFilter) Reset()         { *m = WordFilter{} }
func (m *WordFilter) String() string { return proto.CompactTextString(m) }
func (*WordFilter) ProtoMessage()    {}

func (m *WordFilter) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *WordFilter) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

type WordFilters struct {
	Filters          []*WordFilter `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *WordFilters) Reset()         { *m = WordFilters{} }
func (m *WordFilters) String() string { return proto.CompactTextString(m) }
func (*WordFilters) ProtoMessage()    {}

func (m *WordFilters) GetFilters() []*WordFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}
