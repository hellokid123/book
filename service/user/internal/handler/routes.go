// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "book/service/user/internal/handler/user"
	"book/service/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/user"),
	)
}