package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/chat/api/internal/logic"
	"douniu/server/chat/api/internal/svc"
	"douniu/server/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MessageActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMessageActionLogic(r.Context(), svcCtx)
		resp, err := l.MessageAction(&req)
		response.Response(r, w, resp, err)
	}
}
