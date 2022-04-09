// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	home "pair/service/user/api/internal/handler/home"
	profile "pair/service/user/api/internal/handler/profile"
	"pair/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/home",
				Handler: home.HomeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: home.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: home.LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/profile",
				Handler: profile.ProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/profile/edit",
				Handler: profile.EditProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/qiniu/up/token",
				Handler: profile.QiniuUpTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
