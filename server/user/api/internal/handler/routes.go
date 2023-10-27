// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"douniu/server/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/sendcode",
				Handler: SendcodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login/phone",
				Handler: LoginByPhoneHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login/password",
				Handler: LoginByPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/password/forget",
				Handler: ForgetPasswordHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douniu/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/password/change",
				Handler: ChangePasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/userinfo",
				Handler: userInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/refreshToken",
				Handler: RefreshTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/douniu/user"),
	)
}
