package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/video/api/internal/logic"
	"douniu/server/video/api/internal/svc"
	"douniu/server/video/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserVideoListLogic(r.Context(), svcCtx)
		resp, err := l.UserVideoList(&req)
		response.Response(r, w, resp, err) //②
	}
}
