// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	"douniu/server/video/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/publish",
				Handler: PublishVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/feed/home",
				Handler: FeedHomeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/feed/hot",
				Handler: FeedHotHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/feed/user",
				Handler: FeedUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/feed/follow",
				Handler: FeedFollowHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/feed/partition",
				Handler: FeedPartitionHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: DeleteVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search",
				Handler: SearchVideoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JWTAuth.AccessSecret),
		rest.WithPrefix("/douniu/video"),
		rest.WithTimeout(9999999000*time.Millisecond),
	)
}
