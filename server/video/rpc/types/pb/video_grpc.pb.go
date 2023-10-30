// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: video.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VideoRpc_GetAuthorId_FullMethodName      = "/video.VideoRpc/GetAuthorId"
	VideoRpc_GetVideoListInfo_FullMethodName = "/video.VideoRpc/GetVideoListInfo"
	VideoRpc_PublishVideo_FullMethodName     = "/video.VideoRpc/PublishVideo"
	VideoRpc_FeedHome_FullMethodName         = "/video.VideoRpc/FeedHome"
	VideoRpc_FeedHot_FullMethodName          = "/video.VideoRpc/FeedHot"
	VideoRpc_FeedUser_FullMethodName         = "/video.VideoRpc/FeedUser"
	VideoRpc_FeedPartition_FullMethodName    = "/video.VideoRpc/FeedPartition"
)

// VideoRpcClient is the client API for VideoRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoRpcClient interface {
	GetAuthorId(ctx context.Context, in *GetAuthorIdReq, opts ...grpc.CallOption) (*GetAuthorIdResp, error)
	GetVideoListInfo(ctx context.Context, in *GetVideoListInfoReq, opts ...grpc.CallOption) (*GetVideoListInfoResp, error)
	PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*CommonResp, error)
	FeedHome(ctx context.Context, in *FeedHomeReq, opts ...grpc.CallOption) (*FeedHomeResp, error)
	FeedHot(ctx context.Context, in *FeedHotReq, opts ...grpc.CallOption) (*FeedHotResp, error)
	FeedUser(ctx context.Context, in *FeedUserReq, opts ...grpc.CallOption) (*FeedResp, error)
	FeedPartition(ctx context.Context, in *FeedPartitionReq, opts ...grpc.CallOption) (*FeedResp, error)
}

type videoRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoRpcClient(cc grpc.ClientConnInterface) VideoRpcClient {
	return &videoRpcClient{cc}
}

func (c *videoRpcClient) GetAuthorId(ctx context.Context, in *GetAuthorIdReq, opts ...grpc.CallOption) (*GetAuthorIdResp, error) {
	out := new(GetAuthorIdResp)
	err := c.cc.Invoke(ctx, VideoRpc_GetAuthorId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) GetVideoListInfo(ctx context.Context, in *GetVideoListInfoReq, opts ...grpc.CallOption) (*GetVideoListInfoResp, error) {
	out := new(GetVideoListInfoResp)
	err := c.cc.Invoke(ctx, VideoRpc_GetVideoListInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, VideoRpc_PublishVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) FeedHome(ctx context.Context, in *FeedHomeReq, opts ...grpc.CallOption) (*FeedHomeResp, error) {
	out := new(FeedHomeResp)
	err := c.cc.Invoke(ctx, VideoRpc_FeedHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) FeedHot(ctx context.Context, in *FeedHotReq, opts ...grpc.CallOption) (*FeedHotResp, error) {
	out := new(FeedHotResp)
	err := c.cc.Invoke(ctx, VideoRpc_FeedHot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) FeedUser(ctx context.Context, in *FeedUserReq, opts ...grpc.CallOption) (*FeedResp, error) {
	out := new(FeedResp)
	err := c.cc.Invoke(ctx, VideoRpc_FeedUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoRpcClient) FeedPartition(ctx context.Context, in *FeedPartitionReq, opts ...grpc.CallOption) (*FeedResp, error) {
	out := new(FeedResp)
	err := c.cc.Invoke(ctx, VideoRpc_FeedPartition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoRpcServer is the server API for VideoRpc service.
// All implementations must embed UnimplementedVideoRpcServer
// for forward compatibility
type VideoRpcServer interface {
	GetAuthorId(context.Context, *GetAuthorIdReq) (*GetAuthorIdResp, error)
	GetVideoListInfo(context.Context, *GetVideoListInfoReq) (*GetVideoListInfoResp, error)
	PublishVideo(context.Context, *PublishVideoReq) (*CommonResp, error)
	FeedHome(context.Context, *FeedHomeReq) (*FeedHomeResp, error)
	FeedHot(context.Context, *FeedHotReq) (*FeedHotResp, error)
	FeedUser(context.Context, *FeedUserReq) (*FeedResp, error)
	FeedPartition(context.Context, *FeedPartitionReq) (*FeedResp, error)
	mustEmbedUnimplementedVideoRpcServer()
}

// UnimplementedVideoRpcServer must be embedded to have forward compatible implementations.
type UnimplementedVideoRpcServer struct {
}

func (UnimplementedVideoRpcServer) GetAuthorId(context.Context, *GetAuthorIdReq) (*GetAuthorIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorId not implemented")
}
func (UnimplementedVideoRpcServer) GetVideoListInfo(context.Context, *GetVideoListInfoReq) (*GetVideoListInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoListInfo not implemented")
}
func (UnimplementedVideoRpcServer) PublishVideo(context.Context, *PublishVideoReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVideo not implemented")
}
func (UnimplementedVideoRpcServer) FeedHome(context.Context, *FeedHomeReq) (*FeedHomeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedHome not implemented")
}
func (UnimplementedVideoRpcServer) FeedHot(context.Context, *FeedHotReq) (*FeedHotResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedHot not implemented")
}
func (UnimplementedVideoRpcServer) FeedUser(context.Context, *FeedUserReq) (*FeedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedUser not implemented")
}
func (UnimplementedVideoRpcServer) FeedPartition(context.Context, *FeedPartitionReq) (*FeedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedPartition not implemented")
}
func (UnimplementedVideoRpcServer) mustEmbedUnimplementedVideoRpcServer() {}

// UnsafeVideoRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoRpcServer will
// result in compilation errors.
type UnsafeVideoRpcServer interface {
	mustEmbedUnimplementedVideoRpcServer()
}

func RegisterVideoRpcServer(s grpc.ServiceRegistrar, srv VideoRpcServer) {
	s.RegisterService(&VideoRpc_ServiceDesc, srv)
}

func _VideoRpc_GetAuthorId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).GetAuthorId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_GetAuthorId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).GetAuthorId(ctx, req.(*GetAuthorIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_GetVideoListInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoListInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).GetVideoListInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_GetVideoListInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).GetVideoListInfo(ctx, req.(*GetVideoListInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_PublishVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishVideoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).PublishVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_PublishVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).PublishVideo(ctx, req.(*PublishVideoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_FeedHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedHomeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).FeedHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_FeedHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).FeedHome(ctx, req.(*FeedHomeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_FeedHot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedHotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).FeedHot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_FeedHot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).FeedHot(ctx, req.(*FeedHotReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_FeedUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).FeedUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_FeedUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).FeedUser(ctx, req.(*FeedUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoRpc_FeedPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedPartitionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRpcServer).FeedPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRpc_FeedPartition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRpcServer).FeedPartition(ctx, req.(*FeedPartitionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoRpc_ServiceDesc is the grpc.ServiceDesc for VideoRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.VideoRpc",
	HandlerType: (*VideoRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthorId",
			Handler:    _VideoRpc_GetAuthorId_Handler,
		},
		{
			MethodName: "GetVideoListInfo",
			Handler:    _VideoRpc_GetVideoListInfo_Handler,
		},
		{
			MethodName: "PublishVideo",
			Handler:    _VideoRpc_PublishVideo_Handler,
		},
		{
			MethodName: "FeedHome",
			Handler:    _VideoRpc_FeedHome_Handler,
		},
		{
			MethodName: "FeedHot",
			Handler:    _VideoRpc_FeedHot_Handler,
		},
		{
			MethodName: "FeedUser",
			Handler:    _VideoRpc_FeedUser_Handler,
		},
		{
			MethodName: "FeedPartition",
			Handler:    _VideoRpc_FeedPartition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
