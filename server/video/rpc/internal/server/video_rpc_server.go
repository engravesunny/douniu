// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"douniu/server/video/rpc/internal/logic"
	"douniu/server/video/rpc/internal/svc"
	"douniu/server/video/rpc/types/pb"
)

type VideoRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedVideoRpcServer
}

func NewVideoRpcServer(svcCtx *svc.ServiceContext) *VideoRpcServer {
	return &VideoRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoRpcServer) GetAuthorId(ctx context.Context, in *pb.GetAuthorIdReq) (*pb.GetAuthorIdResp, error) {
	l := logic.NewGetAuthorIdLogic(ctx, s.svcCtx)
	return l.GetAuthorId(in)
}

func (s *VideoRpcServer) GetVideoListInfo(ctx context.Context, in *pb.GetVideoListInfoReq) (*pb.GetVideoListInfoResp, error) {
	l := logic.NewGetVideoListInfoLogic(ctx, s.svcCtx)
	return l.GetVideoListInfo(in)
}

func (s *VideoRpcServer) PublishVideo(ctx context.Context, in *pb.PublishVideoReq) (*pb.CommonResp, error) {
	l := logic.NewPublishVideoLogic(ctx, s.svcCtx)
	return l.PublishVideo(in)
}
