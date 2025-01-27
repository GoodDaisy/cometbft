// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/services/block/v1/block_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("tendermint/services/block/v1/block_service.proto", fileDescriptor_1488dadaa3ae41e3)
}

var fileDescriptor_1488dadaa3ae41e3 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x28, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6,
	0x4f, 0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x2f, 0x33, 0x84, 0x30, 0xe2, 0xa1, 0xe2, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x32, 0x08, 0x1d, 0x7a, 0x30, 0x1d, 0x7a, 0x60, 0x85, 0x7a, 0x65, 0x86,
	0x52, 0x1a, 0x84, 0xcd, 0x83, 0x98, 0x63, 0xf4, 0x99, 0x89, 0x8b, 0xc7, 0x09, 0xc4, 0x0f, 0x86,
	0x28, 0x13, 0x2a, 0xe2, 0xe2, 0x76, 0x4f, 0x2d, 0x71, 0xaa, 0xf4, 0x48, 0xcd, 0x4c, 0xcf, 0x28,
	0x11, 0x32, 0xd0, 0xc3, 0x67, 0x91, 0x1e, 0x92, 0xd2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x29, 0x43, 0x12, 0x74, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xe5, 0x70, 0x71, 0xba, 0xa7,
	0x96, 0xf8, 0x24, 0x96, 0xa4, 0x16, 0x97, 0x08, 0xe9, 0x11, 0xd4, 0x0f, 0x51, 0x08, 0xb3, 0x4f,
	0x9f, 0x68, 0xf5, 0x50, 0xdb, 0x1a, 0x18, 0xb9, 0xf8, 0xe1, 0xa2, 0x50, 0x6f, 0x9a, 0x10, 0x69,
	0x08, 0xaa, 0x57, 0x4d, 0x49, 0xd4, 0x05, 0x71, 0x80, 0x01, 0xa3, 0x53, 0xe4, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xd9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0x27, 0xe7, 0xe7, 0xa6, 0x96, 0x24, 0xa5, 0x95, 0x20, 0x18, 0xe0, 0x38, 0xd3,
	0xc7, 0x17, 0xb9, 0x49, 0x6c, 0x60, 0x35, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xf6,
	0xc7, 0x53, 0x53, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockServiceClient is the client API for BlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockServiceClient interface {
	// GetBlock retrieves the block information at a particular height.
	GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error)
	// GetLatest retrieves the latest block.
	GetLatest(ctx context.Context, in *GetLatestRequest, opts ...grpc.CallOption) (*GetLatestResponse, error)
	// GetLatestHeight returns a stream of the latest block heights committed by
	// the network. This is a long-lived stream that is only terminated by the
	// server if an error occurs. The caller is expected to handle such
	// disconnections and automatically reconnect.
	GetLatestHeight(ctx context.Context, in *GetLatestHeightRequest, opts ...grpc.CallOption) (BlockService_GetLatestHeightClient, error)
}

type blockServiceClient struct {
	cc grpc1.ClientConn
}

func NewBlockServiceClient(cc grpc1.ClientConn) BlockServiceClient {
	return &blockServiceClient{cc}
}

func (c *blockServiceClient) GetByHeight(ctx context.Context, in *GetByHeightRequest, opts ...grpc.CallOption) (*GetByHeightResponse, error) {
	out := new(GetByHeightResponse)
	err := c.cc.Invoke(ctx, "/tendermint.services.block.v1.BlockService/GetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockServiceClient) GetLatest(ctx context.Context, in *GetLatestRequest, opts ...grpc.CallOption) (*GetLatestResponse, error) {
	out := new(GetLatestResponse)
	err := c.cc.Invoke(ctx, "/tendermint.services.block.v1.BlockService/GetLatest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockServiceClient) GetLatestHeight(ctx context.Context, in *GetLatestHeightRequest, opts ...grpc.CallOption) (BlockService_GetLatestHeightClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlockService_serviceDesc.Streams[0], "/tendermint.services.block.v1.BlockService/GetLatestHeight", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockServiceGetLatestHeightClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlockService_GetLatestHeightClient interface {
	Recv() (*GetLatestHeightResponse, error)
	grpc.ClientStream
}

type blockServiceGetLatestHeightClient struct {
	grpc.ClientStream
}

func (x *blockServiceGetLatestHeightClient) Recv() (*GetLatestHeightResponse, error) {
	m := new(GetLatestHeightResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlockServiceServer is the server API for BlockService service.
type BlockServiceServer interface {
	// GetBlock retrieves the block information at a particular height.
	GetByHeight(context.Context, *GetByHeightRequest) (*GetByHeightResponse, error)
	// GetLatest retrieves the latest block.
	GetLatest(context.Context, *GetLatestRequest) (*GetLatestResponse, error)
	// GetLatestHeight returns a stream of the latest block heights committed by
	// the network. This is a long-lived stream that is only terminated by the
	// server if an error occurs. The caller is expected to handle such
	// disconnections and automatically reconnect.
	GetLatestHeight(*GetLatestHeightRequest, BlockService_GetLatestHeightServer) error
}

// UnimplementedBlockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlockServiceServer struct {
}

func (*UnimplementedBlockServiceServer) GetByHeight(ctx context.Context, req *GetByHeightRequest) (*GetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHeight not implemented")
}
func (*UnimplementedBlockServiceServer) GetLatest(ctx context.Context, req *GetLatestRequest) (*GetLatestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatest not implemented")
}
func (*UnimplementedBlockServiceServer) GetLatestHeight(req *GetLatestHeightRequest, srv BlockService_GetLatestHeightServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLatestHeight not implemented")
}

func RegisterBlockServiceServer(s grpc1.Server, srv BlockServiceServer) {
	s.RegisterService(&_BlockService_serviceDesc, srv)
}

func _BlockService_GetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).GetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.services.block.v1.BlockService/GetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).GetByHeight(ctx, req.(*GetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockService_GetLatest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockServiceServer).GetLatest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.services.block.v1.BlockService/GetLatest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockServiceServer).GetLatest(ctx, req.(*GetLatestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockService_GetLatestHeight_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetLatestHeightRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlockServiceServer).GetLatestHeight(m, &blockServiceGetLatestHeightServer{stream})
}

type BlockService_GetLatestHeightServer interface {
	Send(*GetLatestHeightResponse) error
	grpc.ServerStream
}

type blockServiceGetLatestHeightServer struct {
	grpc.ServerStream
}

func (x *blockServiceGetLatestHeightServer) Send(m *GetLatestHeightResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BlockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.services.block.v1.BlockService",
	HandlerType: (*BlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByHeight",
			Handler:    _BlockService_GetByHeight_Handler,
		},
		{
			MethodName: "GetLatest",
			Handler:    _BlockService_GetLatest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLatestHeight",
			Handler:       _BlockService_GetLatestHeight_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tendermint/services/block/v1/block_service.proto",
}
