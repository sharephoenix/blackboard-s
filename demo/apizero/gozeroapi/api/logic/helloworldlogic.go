package logic

import (
	"blackboards/demo/apizero/gozeroapi/api/svc"
	"blackboards/demo/apizero/gozeroapi/rpc/pb"
	"golang.org/x/net/context"
)

type HellWorldRequest struct {
	Name string `form:"name,optional"`
}

type HellWorldRequestP struct {
	Name string `json:"name,optional"`
}

type HelloWorldLogic struct {
	ctx context.Context
	HelloWorldClient pb.HelloWorldClient
}

func NewHelloWorldLogic(ctx context.Context, svcCtx *svc.HelloWorldContext) HelloWorldLogic {
	return HelloWorldLogic{
		ctx: ctx,
		HelloWorldClient: svcCtx.HelloWorldClient,
	}
}