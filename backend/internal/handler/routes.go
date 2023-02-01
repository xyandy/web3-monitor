// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	task "backend/internal/handler/task"
	user "backend/internal/handler/user"
	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthRequest},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: task.CreateTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: task.ReadTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/",
					Handler: task.DeleteTaskHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/task"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/email",
				Handler: user.ChangeEmailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
	)
}
