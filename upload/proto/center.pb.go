// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/center.proto

package upload

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Packet struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4acfdbf4d8b001b, []int{0}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadResult struct {
	Checksum             string   `protobuf:"bytes,1,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResult) Reset()         { *m = UploadResult{} }
func (m *UploadResult) String() string { return proto.CompactTextString(m) }
func (*UploadResult) ProtoMessage()    {}
func (*UploadResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4acfdbf4d8b001b, []int{1}
}

func (m *UploadResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResult.Unmarshal(m, b)
}
func (m *UploadResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResult.Marshal(b, m, deterministic)
}
func (m *UploadResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResult.Merge(m, src)
}
func (m *UploadResult) XXX_Size() int {
	return xxx_messageInfo_UploadResult.Size(m)
}
func (m *UploadResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResult.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResult proto.InternalMessageInfo

func (m *UploadResult) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *UploadResult) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*Packet)(nil), "upload.Packet")
	proto.RegisterType((*UploadResult)(nil), "upload.UploadResult")
}

func init() {
	proto.RegisterFile("proto/center.proto", fileDescriptor_b4acfdbf4d8b001b)
}

var fileDescriptor_b4acfdbf4d8b001b = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4e, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x03, 0x73, 0x84, 0xd8, 0x4a, 0x0b, 0x72,
	0xf2, 0x13, 0x53, 0x94, 0x64, 0xb8, 0xd8, 0x02, 0x12, 0x93, 0xb3, 0x53, 0x4b, 0x84, 0x84, 0xb8,
	0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c, 0x25, 0x3b,
	0x2e, 0x9e, 0x50, 0xb0, 0xba, 0xa0, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0x21, 0x29, 0x2e, 0x8e, 0xe4,
	0x8c, 0xd4, 0xe4, 0xec, 0xe2, 0xd2, 0x5c, 0xb0, 0x3a, 0xce, 0x20, 0x38, 0x1f, 0xa4, 0xbf, 0x38,
	0xb3, 0x2a, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0xcc, 0x36, 0x72, 0xe0, 0xe2, 0x72,
	0x49, 0x2c, 0x49, 0x74, 0x06, 0xdb, 0x2c, 0x64, 0xc4, 0xc5, 0x06, 0x31, 0x4d, 0x88, 0x4f, 0x0f,
	0x62, 0xbd, 0x1e, 0xc4, 0x6e, 0x29, 0x11, 0x18, 0x1f, 0xd9, 0x36, 0x25, 0x06, 0x0d, 0xc6, 0x24,
	0x36, 0xb0, 0x73, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xbf, 0x27, 0xe3, 0xc4, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataCenterClient is the client API for DataCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataCenterClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (DataCenter_UploadClient, error)
}

type dataCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewDataCenterClient(cc grpc.ClientConnInterface) DataCenterClient {
	return &dataCenterClient{cc}
}

func (c *dataCenterClient) Upload(ctx context.Context, opts ...grpc.CallOption) (DataCenter_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataCenter_serviceDesc.Streams[0], "/upload.DataCenter/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataCenterUploadClient{stream}
	return x, nil
}

type DataCenter_UploadClient interface {
	Send(*Packet) error
	CloseAndRecv() (*UploadResult, error)
	grpc.ClientStream
}

type dataCenterUploadClient struct {
	grpc.ClientStream
}

func (x *dataCenterUploadClient) Send(m *Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataCenterUploadClient) CloseAndRecv() (*UploadResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataCenterServer is the server API for DataCenter service.
type DataCenterServer interface {
	Upload(DataCenter_UploadServer) error
}

// UnimplementedDataCenterServer can be embedded to have forward compatible implementations.
type UnimplementedDataCenterServer struct {
}

func (*UnimplementedDataCenterServer) Upload(srv DataCenter_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterDataCenterServer(s *grpc.Server, srv DataCenterServer) {
	s.RegisterService(&_DataCenter_serviceDesc, srv)
}

func _DataCenter_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataCenterServer).Upload(&dataCenterUploadServer{stream})
}

type DataCenter_UploadServer interface {
	SendAndClose(*UploadResult) error
	Recv() (*Packet, error)
	grpc.ServerStream
}

type dataCenterUploadServer struct {
	grpc.ServerStream
}

func (x *dataCenterUploadServer) SendAndClose(m *UploadResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataCenterUploadServer) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataCenter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "upload.DataCenter",
	HandlerType: (*DataCenterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _DataCenter_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/center.proto",
}
