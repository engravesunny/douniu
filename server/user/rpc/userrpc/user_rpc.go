// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userrpc

import (
	"context"

	"douniu/server/user/rpc/types/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonResp                   = pb.CommonResp
	ModifyUserInfoReq            = pb.ModifyUserInfoReq
	RegisterOrLoginByPasswordReq = pb.RegisterOrLoginByPasswordReq
	RegisterOrLoginByPhoneReq    = pb.RegisterOrLoginByPhoneReq
	RegisterOrLoginResp          = pb.RegisterOrLoginResp
	ResetPassword                = pb.ResetPassword
	SendVerificationCodeReq      = pb.SendVerificationCodeReq
	SendVerificationCodeResp     = pb.SendVerificationCodeResp
	TokenResp                    = pb.TokenResp
	UserInfoItem                 = pb.UserInfoItem
	UserInfoReq                  = pb.UserInfoReq
	UserInfoResp                 = pb.UserInfoResp

	UserRpc interface {
		// 发送验证码
		SendVerificationCode(ctx context.Context, in *SendVerificationCodeReq, opts ...grpc.CallOption) (*SendVerificationCodeResp, error)
		// 使用验证码进行手机号注册或登录
		RegisterOrLoginByPhone(ctx context.Context, in *RegisterOrLoginByPhoneReq, opts ...grpc.CallOption) (*RegisterOrLoginResp, error)
		// 使用密码进行手机号注册或登录
		RegisterOrLoginByPassword(ctx context.Context, in *RegisterOrLoginByPasswordReq, opts ...grpc.CallOption) (*RegisterOrLoginResp, error)
		// 忘记密码并重置密码
		ForgetPassword(ctx context.Context, in *ResetPassword, opts ...grpc.CallOption) (*CommonResp, error)
		// 修改密码
		ChangePassword(ctx context.Context, in *ResetPassword, opts ...grpc.CallOption) (*CommonResp, error)
		// 获取用户信息
		GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		// 修改用户信息
		ModifyUserInfo(ctx context.Context, in *ModifyUserInfoReq, opts ...grpc.CallOption) (*CommonResp, error)
	}

	defaultUserRpc struct {
		cli zrpc.Client
	}
)

func NewUserRpc(cli zrpc.Client) UserRpc {
	return &defaultUserRpc{
		cli: cli,
	}
}

// 发送验证码
func (m *defaultUserRpc) SendVerificationCode(ctx context.Context, in *SendVerificationCodeReq, opts ...grpc.CallOption) (*SendVerificationCodeResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.SendVerificationCode(ctx, in, opts...)
}

// 使用验证码进行手机号注册或登录
func (m *defaultUserRpc) RegisterOrLoginByPhone(ctx context.Context, in *RegisterOrLoginByPhoneReq, opts ...grpc.CallOption) (*RegisterOrLoginResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.RegisterOrLoginByPhone(ctx, in, opts...)
}

// 使用密码进行手机号注册或登录
func (m *defaultUserRpc) RegisterOrLoginByPassword(ctx context.Context, in *RegisterOrLoginByPasswordReq, opts ...grpc.CallOption) (*RegisterOrLoginResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.RegisterOrLoginByPassword(ctx, in, opts...)
}

// 忘记密码并重置密码
func (m *defaultUserRpc) ForgetPassword(ctx context.Context, in *ResetPassword, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.ForgetPassword(ctx, in, opts...)
}

// 修改密码
func (m *defaultUserRpc) ChangePassword(ctx context.Context, in *ResetPassword, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.ChangePassword(ctx, in, opts...)
}

// 获取用户信息
func (m *defaultUserRpc) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

// 修改用户信息
func (m *defaultUserRpc) ModifyUserInfo(ctx context.Context, in *ModifyUserInfoReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.ModifyUserInfo(ctx, in, opts...)
}
