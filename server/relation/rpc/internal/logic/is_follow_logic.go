package logic

import (
	"context"
	"douniu/server/common/consts"
	"strconv"

	"douniu/server/relation/rpc/internal/svc"
	"douniu/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowLogic {
	return &IsFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFollowLogic) IsFollow(in *pb.IsFollowRequest) (resp *pb.IsFollowResponse, err error) {
	userIdStr := strconv.FormatInt(in.UserId, 10)
	ToUserIdStr := strconv.FormatInt(in.TargetUserId, 10)
	isFollow, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts.UserFollowPrefix+userIdStr, ToUserIdStr)
	if err != nil {
		l.Errorf("redis sismember err: %v", err)
		return nil, err
	}

	resp = new(pb.IsFollowResponse)
	resp.IsFollow = isFollow != 0

	return
}
