// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/connection_service/connection_service.proto
// DO NOT EDIT!

/*
Package connection_service is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/connection_service/connection_service.proto

It has these top-level messages:
	ConnectRequest
	ConnectionMeteringContentHandles
	ConnectResponse
	BoundService
	BindRequest
	BindResponse
	EchoRequest
	EchoResponse
	DisconnectRequest
	DisconnectNotification
	EncryptRequest
*/
package connection_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import content_handle "github.com/HearthSim/hs-proto-go/bnet/content_handle"
import rpc "github.com/HearthSim/hs-proto-go/bnet/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ConnectRequest struct {
	ClientId         *rpc.ProcessId `protobuf:"bytes,1,opt,name=client_id" json:"client_id,omitempty"`
	BindRequest      *BindRequest   `protobuf:"bytes,2,opt,name=bind_request" json:"bind_request,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ConnectRequest) Reset()                    { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string            { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()               {}
func (*ConnectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConnectRequest) GetClientId() *rpc.ProcessId {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *ConnectRequest) GetBindRequest() *BindRequest {
	if m != nil {
		return m.BindRequest
	}
	return nil
}

type ConnectionMeteringContentHandles struct {
	ContentHandle    []*content_handle.ContentHandle `protobuf:"bytes,1,rep,name=content_handle" json:"content_handle,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *ConnectionMeteringContentHandles) Reset()         { *m = ConnectionMeteringContentHandles{} }
func (m *ConnectionMeteringContentHandles) String() string { return proto.CompactTextString(m) }
func (*ConnectionMeteringContentHandles) ProtoMessage()    {}
func (*ConnectionMeteringContentHandles) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1}
}

func (m *ConnectionMeteringContentHandles) GetContentHandle() []*content_handle.ContentHandle {
	if m != nil {
		return m.ContentHandle
	}
	return nil
}

type ConnectResponse struct {
	ServerId           *rpc.ProcessId                    `protobuf:"bytes,1,req,name=server_id" json:"server_id,omitempty"`
	ClientId           *rpc.ProcessId                    `protobuf:"bytes,2,opt,name=client_id" json:"client_id,omitempty"`
	BindResult         *uint32                           `protobuf:"varint,3,opt,name=bind_result" json:"bind_result,omitempty"`
	BindResponse       *BindResponse                     `protobuf:"bytes,4,opt,name=bind_response" json:"bind_response,omitempty"`
	ContentHandleArray *ConnectionMeteringContentHandles `protobuf:"bytes,5,opt,name=content_handle_array" json:"content_handle_array,omitempty"`
	ServerTime         *uint64                           `protobuf:"varint,6,opt,name=server_time" json:"server_time,omitempty"`
	XXX_unrecognized   []byte                            `json:"-"`
}

func (m *ConnectResponse) Reset()                    { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string            { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()               {}
func (*ConnectResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ConnectResponse) GetServerId() *rpc.ProcessId {
	if m != nil {
		return m.ServerId
	}
	return nil
}

func (m *ConnectResponse) GetClientId() *rpc.ProcessId {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *ConnectResponse) GetBindResult() uint32 {
	if m != nil && m.BindResult != nil {
		return *m.BindResult
	}
	return 0
}

func (m *ConnectResponse) GetBindResponse() *BindResponse {
	if m != nil {
		return m.BindResponse
	}
	return nil
}

func (m *ConnectResponse) GetContentHandleArray() *ConnectionMeteringContentHandles {
	if m != nil {
		return m.ContentHandleArray
	}
	return nil
}

func (m *ConnectResponse) GetServerTime() uint64 {
	if m != nil && m.ServerTime != nil {
		return *m.ServerTime
	}
	return 0
}

type BoundService struct {
	Hash             *uint32 `protobuf:"fixed32,1,req,name=hash" json:"hash,omitempty"`
	Id               *uint32 `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BoundService) Reset()                    { *m = BoundService{} }
func (m *BoundService) String() string            { return proto.CompactTextString(m) }
func (*BoundService) ProtoMessage()               {}
func (*BoundService) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BoundService) GetHash() uint32 {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return 0
}

func (m *BoundService) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type BindRequest struct {
	ImportedServiceHash []uint32        `protobuf:"fixed32,1,rep,name=imported_service_hash" json:"imported_service_hash,omitempty"`
	ExportedService     []*BoundService `protobuf:"bytes,2,rep,name=exported_service" json:"exported_service,omitempty"`
	XXX_unrecognized    []byte          `json:"-"`
}

func (m *BindRequest) Reset()                    { *m = BindRequest{} }
func (m *BindRequest) String() string            { return proto.CompactTextString(m) }
func (*BindRequest) ProtoMessage()               {}
func (*BindRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BindRequest) GetImportedServiceHash() []uint32 {
	if m != nil {
		return m.ImportedServiceHash
	}
	return nil
}

func (m *BindRequest) GetExportedService() []*BoundService {
	if m != nil {
		return m.ExportedService
	}
	return nil
}

type BindResponse struct {
	ImportedServiceId []uint32 `protobuf:"varint,1,rep,packed,name=imported_service_id" json:"imported_service_id,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *BindResponse) Reset()                    { *m = BindResponse{} }
func (m *BindResponse) String() string            { return proto.CompactTextString(m) }
func (*BindResponse) ProtoMessage()               {}
func (*BindResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BindResponse) GetImportedServiceId() []uint32 {
	if m != nil {
		return m.ImportedServiceId
	}
	return nil
}

type EchoRequest struct {
	Time             *uint64 `protobuf:"fixed64,1,opt,name=time" json:"time,omitempty"`
	NetworkOnly      *bool   `protobuf:"varint,2,opt,name=network_only,def=0" json:"network_only,omitempty"`
	Payload          []byte  `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

const Default_EchoRequest_NetworkOnly bool = false

func (m *EchoRequest) GetTime() uint64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *EchoRequest) GetNetworkOnly() bool {
	if m != nil && m.NetworkOnly != nil {
		return *m.NetworkOnly
	}
	return Default_EchoRequest_NetworkOnly
}

func (m *EchoRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type EchoResponse struct {
	Time             *uint64 `protobuf:"fixed64,1,opt,name=time" json:"time,omitempty"`
	Payload          []byte  `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *EchoResponse) GetTime() uint64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *EchoResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type DisconnectRequest struct {
	ErrorCode        *uint32 `protobuf:"varint,1,req,name=error_code" json:"error_code,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DisconnectRequest) Reset()                    { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string            { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()               {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DisconnectRequest) GetErrorCode() uint32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

type DisconnectNotification struct {
	ErrorCode        *uint32 `protobuf:"varint,1,req,name=error_code" json:"error_code,omitempty"`
	Reason           *string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DisconnectNotification) Reset()                    { *m = DisconnectNotification{} }
func (m *DisconnectNotification) String() string            { return proto.CompactTextString(m) }
func (*DisconnectNotification) ProtoMessage()               {}
func (*DisconnectNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DisconnectNotification) GetErrorCode() uint32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *DisconnectNotification) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

type EncryptRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *EncryptRequest) Reset()                    { *m = EncryptRequest{} }
func (m *EncryptRequest) String() string            { return proto.CompactTextString(m) }
func (*EncryptRequest) ProtoMessage()               {}
func (*EncryptRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*ConnectRequest)(nil), "connection_service.ConnectRequest")
	proto.RegisterType((*ConnectionMeteringContentHandles)(nil), "connection_service.ConnectionMeteringContentHandles")
	proto.RegisterType((*ConnectResponse)(nil), "connection_service.ConnectResponse")
	proto.RegisterType((*BoundService)(nil), "connection_service.BoundService")
	proto.RegisterType((*BindRequest)(nil), "connection_service.BindRequest")
	proto.RegisterType((*BindResponse)(nil), "connection_service.BindResponse")
	proto.RegisterType((*EchoRequest)(nil), "connection_service.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "connection_service.EchoResponse")
	proto.RegisterType((*DisconnectRequest)(nil), "connection_service.DisconnectRequest")
	proto.RegisterType((*DisconnectNotification)(nil), "connection_service.DisconnectNotification")
	proto.RegisterType((*EncryptRequest)(nil), "connection_service.EncryptRequest")
}

var fileDescriptor0 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x4e,
	0x10, 0xc5, 0x15, 0x27, 0x4d, 0xff, 0x9d, 0x38, 0x69, 0xea, 0xfc, 0x41, 0x56, 0x51, 0xd5, 0xe0,
	0x0b, 0xb9, 0x24, 0x96, 0x0a, 0x08, 0xa9, 0xe2, 0x94, 0x50, 0xa9, 0x39, 0x80, 0xa0, 0x3d, 0x71,
	0xb2, 0x36, 0xeb, 0x69, 0xbc, 0xe0, 0xec, 0x9a, 0xdd, 0x0d, 0x90, 0x2f, 0xc1, 0x67, 0x66, 0xbd,
	0x76, 0xda, 0xb8, 0x75, 0x45, 0x0f, 0x39, 0x64, 0x77, 0xde, 0x9b, 0xf7, 0x9b, 0x59, 0xc3, 0x97,
	0x25, 0xd3, 0xc9, 0x7a, 0x31, 0xa1, 0x62, 0x15, 0x5e, 0x22, 0x91, 0x3a, 0xb9, 0x66, 0xab, 0x30,
	0x51, 0xe3, 0x4c, 0x0a, 0x2d, 0xc6, 0x4b, 0x11, 0x2e, 0x38, 0xea, 0x90, 0x0a, 0xce, 0x91, 0x6a,
	0x26, 0x78, 0xa4, 0x50, 0xfe, 0x64, 0x14, 0x6b, 0x8e, 0x26, 0x56, 0xe3, 0x79, 0x0f, 0x6f, 0x8e,
	0xe7, 0x4f, 0x6e, 0xa3, 0x91, 0xeb, 0x28, 0x21, 0x3c, 0x4e, 0xf1, 0xde, 0xdf, 0xc2, 0xfe, 0xf8,
	0xf5, 0xd3, 0xac, 0x64, 0x46, 0xf3, 0x5f, 0x21, 0x0a, 0xbe, 0x41, 0x6f, 0x56, 0xa4, 0xba, 0xc2,
	0x1f, 0x6b, 0x54, 0xda, 0x7b, 0x09, 0x07, 0x34, 0x65, 0xb9, 0x3b, 0x8b, 0xfd, 0xc6, 0xb0, 0x31,
	0xea, 0x9c, 0xf5, 0x26, 0xb9, 0xe0, 0xb3, 0x14, 0x14, 0x95, 0x9a, 0xc7, 0xde, 0x5b, 0x70, 0x17,
	0x8c, 0xc7, 0x91, 0x2c, 0x24, 0xbe, 0x63, 0xab, 0x4e, 0x27, 0x35, 0xe4, 0x53, 0x53, 0x57, 0x3a,
	0x07, 0x5f, 0x61, 0x38, 0xbb, 0xad, 0xf8, 0x88, 0x1a, 0x25, 0xe3, 0xcb, 0x59, 0x81, 0x72, 0x69,
	0x49, 0x94, 0xb1, 0xee, 0x55, 0xe1, 0x4c, 0x84, 0xa6, 0x31, 0x3f, 0x99, 0xdc, 0x63, 0xae, 0xe8,
	0x82, 0x3f, 0x0e, 0x1c, 0xde, 0x72, 0xa8, 0x4c, 0x70, 0x85, 0x39, 0x48, 0x9e, 0x02, 0x65, 0x01,
	0xe2, 0xd4, 0x80, 0x54, 0x58, 0x9d, 0x5a, 0xd6, 0x01, 0x74, 0x4a, 0x56, 0xb5, 0x4e, 0xb5, 0xdf,
	0x34, 0x45, 0x5d, 0xef, 0x1d, 0x74, 0xb7, 0x87, 0xb6, 0x97, 0xdf, 0xb2, 0xda, 0xe1, 0xe3, 0x13,
	0x28, 0x33, 0x5d, 0xc1, 0xff, 0x55, 0x8e, 0x88, 0x48, 0x49, 0x36, 0xfe, 0x9e, 0xd5, 0xbf, 0xa9,
	0xd3, 0xff, 0x73, 0x64, 0x26, 0x61, 0xc9, 0xa9, 0xd9, 0x0a, 0xfd, 0xb6, 0xb1, 0x6a, 0x05, 0x23,
	0x70, 0xa7, 0x62, 0xcd, 0xe3, 0xeb, 0xc2, 0xc5, 0x73, 0xa1, 0x95, 0x10, 0x95, 0xd8, 0x39, 0xec,
	0x7b, 0x00, 0x8e, 0x05, 0x76, 0x46, 0xdd, 0x20, 0x81, 0xce, 0xce, 0x92, 0xbc, 0x13, 0x78, 0xc6,
	0x56, 0x99, 0x90, 0x1a, 0xe3, 0x6d, 0x84, 0xa8, 0x54, 0x36, 0x8d, 0xf2, 0x1c, 0xfa, 0xf8, 0xbb,
	0x7a, 0x6d, 0x7c, 0x9a, 0x8f, 0xc2, 0xef, 0x64, 0x08, 0x42, 0x93, 0x69, 0x77, 0x18, 0xa7, 0x30,
	0x78, 0xd0, 0xca, 0xae, 0xaa, 0x39, 0xea, 0x4e, 0x9d, 0x7e, 0x23, 0x98, 0x43, 0xe7, 0x82, 0x26,
	0x62, 0x1b, 0xcd, 0x30, 0x58, 0xc2, 0xfc, 0x51, 0xb6, 0xbd, 0x17, 0xe0, 0x9a, 0xe7, 0xfc, 0x4b,
	0xc8, 0xef, 0x91, 0xe0, 0xe9, 0xc6, 0xae, 0xef, 0xbf, 0xf3, 0xbd, 0x1b, 0x92, 0x1a, 0xeb, 0x43,
	0xd8, 0xcf, 0xc8, 0x26, 0x15, 0x24, 0xb6, 0x1b, 0x73, 0x83, 0x31, 0xb8, 0x85, 0x55, 0xd9, 0xbb,
	0xea, 0xb5, 0x53, 0xee, 0xd8, 0xf2, 0x57, 0x70, 0xf4, 0x81, 0x29, 0x5a, 0xfd, 0x32, 0xcc, 0xd8,
	0x50, 0x4a, 0x21, 0x23, 0x2a, 0x62, 0xb4, 0x93, 0xec, 0x06, 0xef, 0xe1, 0xf9, 0x5d, 0xe1, 0x27,
	0xa1, 0xd9, 0x0d, 0xa3, 0x24, 0x1f, 0x41, 0x5d, 0xb5, 0xd7, 0x83, 0xb6, 0x44, 0xa2, 0x04, 0xb7,
	0x6d, 0x0e, 0x82, 0x3e, 0xf4, 0x2e, 0x38, 0x95, 0x9b, 0x6c, 0xdb, 0xe3, 0x6c, 0x00, 0x47, 0x77,
	0x0b, 0x2f, 0x07, 0xf7, 0x37, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xa0, 0x9a, 0x20, 0x8c, 0x04, 0x00,
	0x00,
}
