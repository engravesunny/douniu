package handler

import (
	"douniu/server/common/response"
	"net/http"

	"douniu/server/favorite/api/internal/logic"
	"douniu/server/favorite/api/internal/svc"
	"douniu/server/favorite/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteLikeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFavoriteActionLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteAction(&req)
		response.Response(r, w, resp, err)
	}
}
