// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: relation.proto

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
	RelationRpc_FollowAction_FullMethodName         = "/relation.RelationRpc/FollowAction"
	RelationRpc_GetFollowList_FullMethodName        = "/relation.RelationRpc/GetFollowList"
	RelationRpc_GetFollowerList_FullMethodName      = "/relation.RelationRpc/GetFollowerList"
	RelationRpc_GetUserFollowCount_FullMethodName   = "/relation.RelationRpc/GetUserFollowCount"
	RelationRpc_GetUserFollowerCount_FullMethodName = "/relation.RelationRpc/GetUserFollowerCount"
	RelationRpc_GetFriendList_FullMethodName        = "/relation.RelationRpc/GetFriendList"
	RelationRpc_IsFollow_FullMethodName             = "/relation.RelationRpc/IsFollow"
)

// RelationRpcClient is the client API for RelationRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationRpcClient interface {
	FollowAction(ctx context.Context, in *FollowActionRequest, opts ...grpc.CallOption) (*FollowActionResponse, error)
	GetFollowList(ctx context.Context, in *GetFollowListRequest, opts ...grpc.CallOption) (*GetFollowListResponse, error)
	GetFollowerList(ctx context.Context, in *GetFollowerListRequest, opts ...grpc.CallOption) (*GetFollowerListResponse, error)
	GetUserFollowCount(ctx context.Context, in *GetUserFollowCountRequest, opts ...grpc.CallOption) (*GetUserFollowCountResponse, error)
	GetUserFollowerCount(ctx context.Context, in *GetUserFollowerCountRequest, opts ...grpc.CallOption) (*GetUserFollowerCountResponse, error)
	GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error)
	IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowResponse, error)
}

type relationRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationRpcClient(cc grpc.ClientConnInterface) RelationRpcClient {
	return &relationRpcClient{cc}
}

func (c *relationRpcClient) FollowAction(ctx context.Context, in *FollowActionRequest, opts ...grpc.CallOption) (*FollowActionResponse, error) {
	out := new(FollowActionResponse)
	err := c.cc.Invoke(ctx, RelationRpc_FollowAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationRpcClient) GetFollowList(ctx context.Context, in *GetFollowListRequest, opts ...grpc.CallOption) (*GetFollowListResponse, error) {
	out := new(GetFollowListResponse)
	err := c.cc.Invoke(ctx, RelationRpc_GetFollowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationRpcClient) GetFollowerList(ctx context.Context, in *GetFollowerListRequest, opts ...grpc.CallOption) (*GetFollowerListResponse, error) {
	out := new(GetFollowerListResponse)
	err := c.cc.Invoke(ctx, RelationRpc_GetFollowerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationRpcClient) GetUserFollowCount(ctx context.Context, in *GetUserFollowCountRequest, opts ...grpc.CallOption) (*GetUserFollowCountResponse, error) {
	out := new(GetUserFollowCountResponse)
	err := c.cc.Invoke(ctx, RelationRpc_GetUserFollowCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationRpcClient) GetUserFollowerCount(ctx context.Context, in *GetUserFollowerCountRequest, opts ...grpc.CallOption) (*GetUserFollowerCountResponse, error) {
	out := new(GetUserFollowerCountResponse)
	err := c.cc.Invoke(ctx, RelationRpc_GetUserFollowerCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationRpcClient) GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error) {
	out := new(GetFriendListResponse)
	err := c.cc.Invoke(ctx, RelationRpc_GetFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationRpcClient) IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowResponse, error) {
	out := new(IsFollowResponse)
	err := c.cc.Invoke(ctx, RelationRpc_IsFollow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationRpcServer is the server API for RelationRpc service.
// All implementations must embed UnimplementedRelationRpcServer
// for forward compatibility
type RelationRpcServer interface {
	FollowAction(context.Context, *FollowActionRequest) (*FollowActionResponse, error)
	GetFollowList(context.Context, *GetFollowListRequest) (*GetFollowListResponse, error)
	GetFollowerList(context.Context, *GetFollowerListRequest) (*GetFollowerListResponse, error)
	GetUserFollowCount(context.Context, *GetUserFollowCountRequest) (*GetUserFollowCountResponse, error)
	GetUserFollowerCount(context.Context, *GetUserFollowerCountRequest) (*GetUserFollowerCountResponse, error)
	GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListResponse, error)
	IsFollow(context.Context, *IsFollowRequest) (*IsFollowResponse, error)
	mustEmbedUnimplementedRelationRpcServer()
}

// UnimplementedRelationRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRelationRpcServer struct {
}

func (UnimplementedRelationRpcServer) FollowAction(context.Context, *FollowActionRequest) (*FollowActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowAction not implemented")
}
func (UnimplementedRelationRpcServer) GetFollowList(context.Context, *GetFollowListRequest) (*GetFollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowList not implemented")
}
func (UnimplementedRelationRpcServer) GetFollowerList(context.Context, *GetFollowerListRequest) (*GetFollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerList not implemented")
}
func (UnimplementedRelationRpcServer) GetUserFollowCount(context.Context, *GetUserFollowCountRequest) (*GetUserFollowCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollowCount not implemented")
}
func (UnimplementedRelationRpcServer) GetUserFollowerCount(context.Context, *GetUserFollowerCountRequest) (*GetUserFollowerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollowerCount not implemented")
}
func (UnimplementedRelationRpcServer) GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}
func (UnimplementedRelationRpcServer) IsFollow(context.Context, *IsFollowRequest) (*IsFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFollow not implemented")
}
func (UnimplementedRelationRpcServer) mustEmbedUnimplementedRelationRpcServer() {}

// UnsafeRelationRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationRpcServer will
// result in compilation errors.
type UnsafeRelationRpcServer interface {
	mustEmbedUnimplementedRelationRpcServer()
}

func RegisterRelationRpcServer(s grpc.ServiceRegistrar, srv RelationRpcServer) {
	s.RegisterService(&RelationRpc_ServiceDesc, srv)
}

func _RelationRpc_FollowAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationRpcServer).FollowAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationRpc_FollowAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationRpcServer).FollowAction(ctx, req.(*FollowActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationRpc_GetFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationRpcServer).GetFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationRpc_GetFollowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationRpcServer).GetFollowList(ctx, req.(*GetFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationRpc_GetFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationRpcServer).GetFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationRpc_GetFollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationRpcServer).GetFollowerList(ctx, req.(*GetFollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationRpc_GetUserFollowCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFollowCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationRpcServer).GetUserFollowCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationRpc_GetUserFollowCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationRpcServer).GetUserFollowCount(ctx, req.(*GetUserFollowCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationRpc_GetUserFollowerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFollowerCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationRpcServer).GetUserFollowerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationRpc_GetUserFollowerCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationRpcServer).GetUserFollowerCount(ctx, req.(*GetUserFollowerCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationRpc_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationRpcServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationRpc_GetFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationRpcServer).GetFriendList(ctx, req.(*GetFriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationRpc_IsFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationRpcServer).IsFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationRpc_IsFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationRpcServer).IsFollow(ctx, req.(*IsFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationRpc_ServiceDesc is the grpc.ServiceDesc for RelationRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relation.RelationRpc",
	HandlerType: (*RelationRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FollowAction",
			Handler:    _RelationRpc_FollowAction_Handler,
		},
		{
			MethodName: "GetFollowList",
			Handler:    _RelationRpc_GetFollowList_Handler,
		},
		{
			MethodName: "GetFollowerList",
			Handler:    _RelationRpc_GetFollowerList_Handler,
		},
		{
			MethodName: "GetUserFollowCount",
			Handler:    _RelationRpc_GetUserFollowCount_Handler,
		},
		{
			MethodName: "GetUserFollowerCount",
			Handler:    _RelationRpc_GetUserFollowerCount_Handler,
		},
		{
			MethodName: "GetFriendList",
			Handler:    _RelationRpc_GetFriendList_Handler,
		},
		{
			MethodName: "IsFollow",
			Handler:    _RelationRpc_IsFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation.proto",
}