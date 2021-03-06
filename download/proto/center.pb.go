// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/center.proto

package download

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

func init() {
	proto.RegisterType((*Packet)(nil), "download.Packet")
}

func init() {
	proto.RegisterFile("proto/center.proto", fileDescriptor_b4acfdbf4d8b001b)
}

var fileDescriptor_b4acfdbf4d8b001b = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4e, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x03, 0x73, 0x84, 0x38, 0x52, 0xf2, 0xcb,
	0xf3, 0x72, 0xf2, 0x13, 0x53, 0xa4, 0xa4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xe2,
	0x49, 0xa5, 0x69, 0xfa, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x10, 0x65, 0x4a, 0x32, 0x5c, 0x6c, 0x01,
	0x89, 0xc9, 0xd9, 0xa9, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x91, 0x1b, 0x17, 0x97, 0x4b, 0x62, 0x49, 0xa2, 0x33, 0xd8,
	0x60, 0x21, 0x0b, 0x2e, 0x0e, 0x17, 0xa8, 0xa1, 0x42, 0x62, 0x7a, 0x10, 0x53, 0xf5, 0x60, 0xa6,
	0xea, 0xb9, 0x82, 0x4c, 0x95, 0x12, 0xd0, 0x83, 0xd9, 0xab, 0x07, 0x31, 0x57, 0x89, 0xc1, 0x80,
	0x31, 0x89, 0x0d, 0xac, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x5a, 0xfd, 0x37, 0xa9,
	0x00, 0x00, 0x00,
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
	Download(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (DataCenter_DownloadClient, error)
}

type dataCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewDataCenterClient(cc grpc.ClientConnInterface) DataCenterClient {
	return &dataCenterClient{cc}
}

func (c *dataCenterClient) Download(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (DataCenter_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataCenter_serviceDesc.Streams[0], "/download.DataCenter/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataCenterDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataCenter_DownloadClient interface {
	Recv() (*Packet, error)
	grpc.ClientStream
}

type dataCenterDownloadClient struct {
	grpc.ClientStream
}

func (x *dataCenterDownloadClient) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataCenterServer is the server API for DataCenter service.
type DataCenterServer interface {
	Download(*empty.Empty, DataCenter_DownloadServer) error
}

// UnimplementedDataCenterServer can be embedded to have forward compatible implementations.
type UnimplementedDataCenterServer struct {
}

func (*UnimplementedDataCenterServer) Download(req *empty.Empty, srv DataCenter_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}

func RegisterDataCenterServer(s *grpc.Server, srv DataCenterServer) {
	s.RegisterService(&_DataCenter_serviceDesc, srv)
}

func _DataCenter_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataCenterServer).Download(m, &dataCenterDownloadServer{stream})
}

type DataCenter_DownloadServer interface {
	Send(*Packet) error
	grpc.ServerStream
}

type dataCenterDownloadServer struct {
	grpc.ServerStream
}

func (x *dataCenterDownloadServer) Send(m *Packet) error {
	return x.ServerStream.SendMsg(m)
}

var _DataCenter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "download.DataCenter",
	HandlerType: (*DataCenterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _DataCenter_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/center.proto",
}
