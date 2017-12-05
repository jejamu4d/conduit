// Code generated by protoc-gen-go. DO NOT EDIT.
// source: controller/tap/tap.proto

/*
Package conduit_controller_tap is a generated protocol buffer package.

It is generated from these files:
	controller/tap/tap.proto

It has these top-level messages:
*/
package conduit_controller_tap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import conduit_common "github.com/runconduit/conduit/controller/gen/common"
import conduit_public "github.com/runconduit/conduit/controller/gen/public"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tap service

type TapClient interface {
	Tap(ctx context.Context, in *conduit_public.TapRequest, opts ...grpc.CallOption) (Tap_TapClient, error)
}

type tapClient struct {
	cc *grpc.ClientConn
}

func NewTapClient(cc *grpc.ClientConn) TapClient {
	return &tapClient{cc}
}

func (c *tapClient) Tap(ctx context.Context, in *conduit_public.TapRequest, opts ...grpc.CallOption) (Tap_TapClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Tap_serviceDesc.Streams[0], c.cc, "/conduit.controller.tap.Tap/Tap", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapTapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tap_TapClient interface {
	Recv() (*conduit_common.TapEvent, error)
	grpc.ClientStream
}

type tapTapClient struct {
	grpc.ClientStream
}

func (x *tapTapClient) Recv() (*conduit_common.TapEvent, error) {
	m := new(conduit_common.TapEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Tap service

type TapServer interface {
	Tap(*conduit_public.TapRequest, Tap_TapServer) error
}

func RegisterTapServer(s *grpc.Server, srv TapServer) {
	s.RegisterService(&_Tap_serviceDesc, srv)
}

func _Tap_Tap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(conduit_public.TapRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TapServer).Tap(m, &tapTapServer{stream})
}

type Tap_TapServer interface {
	Send(*conduit_common.TapEvent) error
	grpc.ServerStream
}

type tapTapServer struct {
	grpc.ServerStream
}

func (x *tapTapServer) Send(m *conduit_common.TapEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _Tap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "conduit.controller.tap.Tap",
	HandlerType: (*TapServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tap",
			Handler:       _Tap_Tap_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "controller/tap/tap.proto",
}

func init() { proto.RegisterFile("controller/tap/tap.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0xcf, 0x2b,
	0x29, 0xca, 0xcf, 0xc9, 0x49, 0x2d, 0xd2, 0x2f, 0x49, 0x2c, 0x00, 0x61, 0xbd, 0x82, 0xa2, 0xfc,
	0x92, 0x7c, 0x21, 0xb1, 0xe4, 0xfc, 0xbc, 0x94, 0xd2, 0xcc, 0x12, 0x3d, 0x84, 0x0a, 0xbd, 0x92,
	0xc4, 0x02, 0x29, 0xe1, 0xe4, 0xfc, 0xdc, 0xdc, 0xfc, 0x3c, 0x7d, 0x08, 0x05, 0x51, 0x2c, 0x25,
	0x50, 0x50, 0x9a, 0x94, 0x93, 0x99, 0xac, 0x9f, 0x58, 0x90, 0x09, 0x11, 0x31, 0x72, 0xe3, 0x62,
	0x0e, 0x49, 0x2c, 0x10, 0xb2, 0x87, 0x50, 0x52, 0x7a, 0x30, 0xd3, 0x20, 0x0a, 0xf5, 0x42, 0x12,
	0x0b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0x24, 0xf4, 0x10, 0x36, 0x81, 0x8d, 0x0c,
	0x49, 0x2c, 0x70, 0x2d, 0x4b, 0xcd, 0x2b, 0x51, 0x62, 0x30, 0x60, 0x4c, 0x62, 0x03, 0x1b, 0x67,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x17, 0x0c, 0x64, 0xf7, 0xa9, 0x00, 0x00, 0x00,
}
