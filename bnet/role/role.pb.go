// Code generated by protoc-gen-go.
// source: github.com/vjurenka/hs-proto-go/bnet/role/role.proto
// DO NOT EDIT!

/*
Package role is a generated protocol buffer package.

It is generated from these files:
	github.com/vjurenka/hs-proto-go/bnet/role/role.proto

It has these top-level messages:
	Privilege
	RoleSetConfig
	Role
	RoleSet
*/
package role

import proto "github.com/golang/protobuf/proto"
import math "math"
import attribute "github.com/vjurenka/hs-proto-go/bnet/attribute"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Privilege struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *uint32 `protobuf:"varint,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Privilege) Reset()         { *m = Privilege{} }
func (m *Privilege) String() string { return proto.CompactTextString(m) }
func (*Privilege) ProtoMessage()    {}

func (m *Privilege) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Privilege) GetValue() uint32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type RoleSetConfig struct {
	Privilege        []*Privilege `protobuf:"bytes,1,rep,name=privilege" json:"privilege,omitempty"`
	RoleSet          *RoleSet     `protobuf:"bytes,2,req,name=role_set" json:"role_set,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *RoleSetConfig) Reset()         { *m = RoleSetConfig{} }
func (m *RoleSetConfig) String() string { return proto.CompactTextString(m) }
func (*RoleSetConfig) ProtoMessage()    {}

func (m *RoleSetConfig) GetPrivilege() []*Privilege {
	if m != nil {
		return m.Privilege
	}
	return nil
}

func (m *RoleSetConfig) GetRoleSet() *RoleSet {
	if m != nil {
		return m.RoleSet
	}
	return nil
}

type Role struct {
	Id               *uint32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name             *string                `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Privilege        []string               `protobuf:"bytes,3,rep,name=privilege" json:"privilege,omitempty"`
	AssignableRole   []uint32               `protobuf:"varint,4,rep,packed,name=assignable_role" json:"assignable_role,omitempty"`
	Required         *bool                  `protobuf:"varint,5,opt,name=required,def=0" json:"required,omitempty"`
	Unique           *bool                  `protobuf:"varint,6,opt,name=unique,def=0" json:"unique,omitempty"`
	RelegationRole   *uint32                `protobuf:"varint,7,opt,name=relegation_role" json:"relegation_role,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,8,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}

const Default_Role_Required bool = false
const Default_Role_Unique bool = false

func (m *Role) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Role) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Role) GetPrivilege() []string {
	if m != nil {
		return m.Privilege
	}
	return nil
}

func (m *Role) GetAssignableRole() []uint32 {
	if m != nil {
		return m.AssignableRole
	}
	return nil
}

func (m *Role) GetRequired() bool {
	if m != nil && m.Required != nil {
		return *m.Required
	}
	return Default_Role_Required
}

func (m *Role) GetUnique() bool {
	if m != nil && m.Unique != nil {
		return *m.Unique
	}
	return Default_Role_Unique
}

func (m *Role) GetRelegationRole() uint32 {
	if m != nil && m.RelegationRole != nil {
		return *m.RelegationRole
	}
	return 0
}

func (m *Role) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type RoleSet struct {
	Program          *string                `protobuf:"bytes,1,req,name=program" json:"program,omitempty"`
	Service          *string                `protobuf:"bytes,2,req,name=service" json:"service,omitempty"`
	Subtype          *string                `protobuf:"bytes,3,opt,name=subtype,def=default" json:"subtype,omitempty"`
	Role             []*Role                `protobuf:"bytes,4,rep,name=role" json:"role,omitempty"`
	DefaultRole      []uint32               `protobuf:"varint,5,rep,packed,name=default_role" json:"default_role,omitempty"`
	MaxMembers       *int32                 `protobuf:"varint,6,opt,name=max_members" json:"max_members,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,7,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *RoleSet) Reset()         { *m = RoleSet{} }
func (m *RoleSet) String() string { return proto.CompactTextString(m) }
func (*RoleSet) ProtoMessage()    {}

const Default_RoleSet_Subtype string = "default"

func (m *RoleSet) GetProgram() string {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return ""
}

func (m *RoleSet) GetService() string {
	if m != nil && m.Service != nil {
		return *m.Service
	}
	return ""
}

func (m *RoleSet) GetSubtype() string {
	if m != nil && m.Subtype != nil {
		return *m.Subtype
	}
	return Default_RoleSet_Subtype
}

func (m *RoleSet) GetRole() []*Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *RoleSet) GetDefaultRole() []uint32 {
	if m != nil {
		return m.DefaultRole
	}
	return nil
}

func (m *RoleSet) GetMaxMembers() int32 {
	if m != nil && m.MaxMembers != nil {
		return *m.MaxMembers
	}
	return 0
}

func (m *RoleSet) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}
