package handler

import (
	"blackboards/demo/apizero/gozeroapi/api/svc"
	"github.com/tal-tech/go-zero/rest"
	"net/http"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.HelloWorldContext) {
	engine.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/hello/world",
			Handler: addApplyHandler(serverCtx),
		},
		{
			Method:  http.MethodPost,
			Path:    "/hello/worldp",
			Handler: addApplyHandlerP(serverCtx),
		},
	})
}

