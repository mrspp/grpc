// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: proto/rank.proto

package rank

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index       string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ProductName string `protobuf:"bytes,2,opt,name=productName,proto3" json:"productName,omitempty"`
	Rating      string `protobuf:"bytes,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *RankRequest) Reset() {
	*x = RankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankRequest) ProtoMessage() {}

func (x *RankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankRequest.ProtoReflect.Descriptor instead.
func (*RankRequest) Descriptor() ([]byte, []int) {
	return file_proto_rank_proto_rawDescGZIP(), []int{0}
}

func (x *RankRequest) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *RankRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *RankRequest) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

type RankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Popularity string `protobuf:"bytes,1,opt,name=popularity,proto3" json:"popularity,omitempty"`
}

func (x *RankResponse) Reset() {
	*x = RankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankResponse) ProtoMessage() {}

func (x *RankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankResponse.ProtoReflect.Descriptor instead.
func (*RankResponse) Descriptor() ([]byte, []int) {
	return file_proto_rank_proto_rawDescGZIP(), []int{1}
}

func (x *RankResponse) GetPopularity() string {
	if x != nil {
		return x.Popularity
	}
	return ""
}

var File_proto_rank_proto protoreflect.FileDescriptor

var file_proto_rank_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x5d, 0x0a, 0x0b, 0x72, 0x61, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x2e, 0x0a, 0x0c, 0x72, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c,
	0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x70,
	0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x32, 0x45, 0x0a, 0x0b, 0x72, 0x61, 0x6e, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x11, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e,
	0x72, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x72, 0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rank_proto_rawDescOnce sync.Once
	file_proto_rank_proto_rawDescData = file_proto_rank_proto_rawDesc
)

func file_proto_rank_proto_rawDescGZIP() []byte {
	file_proto_rank_proto_rawDescOnce.Do(func() {
		file_proto_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rank_proto_rawDescData)
	})
	return file_proto_rank_proto_rawDescData
}

var file_proto_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_rank_proto_goTypes = []interface{}{
	(*RankRequest)(nil),  // 0: rank.rankRequest
	(*RankResponse)(nil), // 1: rank.rankResponse
}
var file_proto_rank_proto_depIdxs = []int32{
	0, // 0: rank.rankService.RankService:input_type -> rank.rankRequest
	1, // 1: rank.rankService.RankService:output_type -> rank.rankResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_rank_proto_init() }
func file_proto_rank_proto_init() {
	if File_proto_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rank_proto_goTypes,
		DependencyIndexes: file_proto_rank_proto_depIdxs,
		MessageInfos:      file_proto_rank_proto_msgTypes,
	}.Build()
	File_proto_rank_proto = out.File
	file_proto_rank_proto_rawDesc = nil
	file_proto_rank_proto_goTypes = nil
	file_proto_rank_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RankServiceClient is the client API for RankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RankServiceClient interface {
	RankService(ctx context.Context, in *RankRequest, opts ...grpc.CallOption) (*RankResponse, error)
}

type rankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRankServiceClient(cc grpc.ClientConnInterface) RankServiceClient {
	return &rankServiceClient{cc}
}

func (c *rankServiceClient) RankService(ctx context.Context, in *RankRequest, opts ...grpc.CallOption) (*RankResponse, error) {
	out := new(RankResponse)
	err := c.cc.Invoke(ctx, "/rank.rankService/RankService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankServiceServer is the server API for RankService service.
type RankServiceServer interface {
	RankService(context.Context, *RankRequest) (*RankResponse, error)
}

// UnimplementedRankServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRankServiceServer struct {
}

func (*UnimplementedRankServiceServer) RankService(context.Context, *RankRequest) (*RankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RankService not implemented")
}

func RegisterRankServiceServer(s *grpc.Server, srv RankServiceServer) {
	s.RegisterService(&_RankService_serviceDesc, srv)
}

func _RankService_RankService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServiceServer).RankService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rank.rankService/RankService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServiceServer).RankService(ctx, req.(*RankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RankService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rank.rankService",
	HandlerType: (*RankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RankService",
			Handler:    _RankService_RankService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rank.proto",
}
