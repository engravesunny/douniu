// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.10
// source: user.proto

package pb

import (
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

// 请求发送验证码
type SendVerificationCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *SendVerificationCodeReq) Reset() {
	*x = SendVerificationCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerificationCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationCodeReq) ProtoMessage() {}

func (x *SendVerificationCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerificationCodeReq.ProtoReflect.Descriptor instead.
func (*SendVerificationCodeReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *SendVerificationCodeReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

// 包含验证码的响应
type SendVerificationCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerificationCode string `protobuf:"bytes,1,opt,name=verification_code,json=verificationCode,proto3" json:"verification_code,omitempty"`
}

func (x *SendVerificationCodeResp) Reset() {
	*x = SendVerificationCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerificationCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationCodeResp) ProtoMessage() {}

func (x *SendVerificationCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerificationCodeResp.ProtoReflect.Descriptor instead.
func (*SendVerificationCodeResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *SendVerificationCodeResp) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

// 用户注册或登录后的响应
type RegisterOrLoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RegisterOrLoginResp) Reset() {
	*x = RegisterOrLoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterOrLoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOrLoginResp) ProtoMessage() {}

func (x *RegisterOrLoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOrLoginResp.ProtoReflect.Descriptor instead.
func (*RegisterOrLoginResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterOrLoginResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RegisterOrLoginResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RegisterOrLoginResp) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RegisterOrLoginByPhoneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone            string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	VerificationCode string `protobuf:"bytes,2,opt,name=verification_code,json=verificationCode,proto3" json:"verification_code,omitempty"`
}

func (x *RegisterOrLoginByPhoneReq) Reset() {
	*x = RegisterOrLoginByPhoneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterOrLoginByPhoneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOrLoginByPhoneReq) ProtoMessage() {}

func (x *RegisterOrLoginByPhoneReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOrLoginByPhoneReq.ProtoReflect.Descriptor instead.
func (*RegisterOrLoginByPhoneReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterOrLoginByPhoneReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegisterOrLoginByPhoneReq) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

// 使用密码进行登录的请求
type RegisterOrLoginByPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone    string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterOrLoginByPasswordReq) Reset() {
	*x = RegisterOrLoginByPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterOrLoginByPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOrLoginByPasswordReq) ProtoMessage() {}

func (x *RegisterOrLoginByPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOrLoginByPasswordReq.ProtoReflect.Descriptor instead.
func (*RegisterOrLoginByPasswordReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterOrLoginByPasswordReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegisterOrLoginByPasswordReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 重置密码的请求
type ResetPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPassword string `protobuf:"bytes,1,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	UserId      int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ResetPassword) Reset() {
	*x = ResetPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPassword) ProtoMessage() {}

func (x *ResetPassword) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPassword.ProtoReflect.Descriptor instead.
func (*ResetPassword) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *ResetPassword) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *ResetPassword) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 获取用户信息的请求
type UserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FromUserId int64 `protobuf:"varint,2,opt,name=from_user_id,json=fromUserId,proto3" json:"from_user_id,omitempty"`
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserInfoReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfoReq) GetFromUserId() int64 {
	if x != nil {
		return x.FromUserId
	}
	return 0
}

// 包含用户信息的响应
type UserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userinfo *UserInfoItem `protobuf:"bytes,1,opt,name=userinfo,proto3" json:"userinfo,omitempty"`
}

func (x *UserInfoResp) Reset() {
	*x = UserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResp) ProtoMessage() {}

func (x *UserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResp.ProtoReflect.Descriptor instead.
func (*UserInfoResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserInfoResp) GetUserinfo() *UserInfoItem {
	if x != nil {
		return x.Userinfo
	}
	return nil
}

// 用户信息项
type UserInfoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone           string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Nickname        string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	IsFollow        bool   `protobuf:"varint,4,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
	Avatar          string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	BackgroundImage string `protobuf:"bytes,6,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"`
	Signature       string `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
	TotalFavorited  int64  `protobuf:"varint,8,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"`
	WorkCount       int64  `protobuf:"varint,9,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`
	FavoriteCount   int64  `protobuf:"varint,10,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`
	FollowCount     int64  `protobuf:"varint,11,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`
	FollowerCount   int64  `protobuf:"varint,12,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
}

func (x *UserInfoItem) Reset() {
	*x = UserInfoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoItem) ProtoMessage() {}

func (x *UserInfoItem) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoItem.ProtoReflect.Descriptor instead.
func (*UserInfoItem) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserInfoItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoItem) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserInfoItem) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserInfoItem) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *UserInfoItem) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfoItem) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *UserInfoItem) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *UserInfoItem) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *UserInfoItem) GetWorkCount() int64 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

func (x *UserInfoItem) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *UserInfoItem) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *UserInfoItem) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

// 包含访问令牌和刷新令牌的响应
type TokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *TokenResp) Reset() {
	*x = TokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResp) ProtoMessage() {}

func (x *TokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResp.ProtoReflect.Descriptor instead.
func (*TokenResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *TokenResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenResp) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// 通用响应
type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *CommonResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x2f, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x22, 0x47, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2b, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x76, 0x0a, 0x13,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5e, 0x0a, 0x19, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x50, 0x0a, 0x1c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4b, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x87, 0x03,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3a, 0x0a, 0x0a,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xba, 0x03, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x70, 0x63, 0x12, 0x55, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x16, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x5a, 0x0a, 0x19, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x22,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x4f, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a,
	0x0e, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x10, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_user_proto_goTypes = []interface{}{
	(*SendVerificationCodeReq)(nil),      // 0: user.SendVerificationCodeReq
	(*SendVerificationCodeResp)(nil),     // 1: user.SendVerificationCodeResp
	(*RegisterOrLoginResp)(nil),          // 2: user.RegisterOrLoginResp
	(*RegisterOrLoginByPhoneReq)(nil),    // 3: user.RegisterOrLoginByPhoneReq
	(*RegisterOrLoginByPasswordReq)(nil), // 4: user.RegisterOrLoginByPasswordReq
	(*ResetPassword)(nil),                // 5: user.ResetPassword
	(*UserInfoReq)(nil),                  // 6: user.UserInfoReq
	(*UserInfoResp)(nil),                 // 7: user.UserInfoResp
	(*UserInfoItem)(nil),                 // 8: user.UserInfoItem
	(*TokenResp)(nil),                    // 9: user.TokenResp
	(*CommonResp)(nil),                   // 10: user.CommonResp
}
var file_user_proto_depIdxs = []int32{
	8,  // 0: user.UserInfoResp.userinfo:type_name -> user.UserInfoItem
	0,  // 1: user.UserRpc.SendVerificationCode:input_type -> user.SendVerificationCodeReq
	3,  // 2: user.UserRpc.RegisterOrLoginByPhone:input_type -> user.RegisterOrLoginByPhoneReq
	4,  // 3: user.UserRpc.RegisterOrLoginByPassword:input_type -> user.RegisterOrLoginByPasswordReq
	5,  // 4: user.UserRpc.ForgetPassword:input_type -> user.ResetPassword
	5,  // 5: user.UserRpc.ChangePassword:input_type -> user.ResetPassword
	6,  // 6: user.UserRpc.GetUserInfo:input_type -> user.UserInfoReq
	1,  // 7: user.UserRpc.SendVerificationCode:output_type -> user.SendVerificationCodeResp
	2,  // 8: user.UserRpc.RegisterOrLoginByPhone:output_type -> user.RegisterOrLoginResp
	2,  // 9: user.UserRpc.RegisterOrLoginByPassword:output_type -> user.RegisterOrLoginResp
	10, // 10: user.UserRpc.ForgetPassword:output_type -> user.CommonResp
	10, // 11: user.UserRpc.ChangePassword:output_type -> user.CommonResp
	7,  // 12: user.UserRpc.GetUserInfo:output_type -> user.UserInfoResp
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerificationCodeReq); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerificationCodeResp); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterOrLoginResp); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterOrLoginByPhoneReq); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterOrLoginByPasswordReq); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPassword); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReq); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResp); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoItem); i {
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
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResp); i {
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
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}