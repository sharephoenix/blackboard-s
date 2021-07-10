package handler

import (
	"blackboards/demo/apizero/gozeroapi/api/logic"
	"blackboards/demo/apizero/gozeroapi/api/svc"
	"blackboards/demo/apizero/gozeroapi/rpc/pb"
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

func addApplyHandler(ctx *svc.HelloWorldContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "first")
		// 调用 各种 logic 实现数据拼装，最后返回
		var req logic.HellWorldRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			fmt.Println("requestName-error:", err.Error())
			HttpError(w, 200, -2, "参数错误",err.Error())
			return
		}
		fmt.Println("requestName:", req.Name)
		data, err := ctx.HelloWorldClient.SayHello(context.TODO(), &pb.HelloRequest{
			Name: req.Name,
		})
		FormatResponseWithRequest(data, err, w, r)
	}
}

func addApplyHandlerP(ctx *svc.HelloWorldContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "first")
		// 调用 各种 logic 实现数据拼装，最后返回
		var req logic.HellWorldRequestP
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fmt.Println("requestName-error:", err.Error())
			HttpError(w, 200, -2, "参数错误",err.Error())
			return
		}
		fmt.Println("requestName:", req.Name)
		data, err := ctx.HelloWorldClient.SayHello(context.TODO(), &pb.HelloRequest{
			Name: req.Name,
		})
		FormatResponseWithRequest(data, err, w, r)
	}
}

func FormatResponseWithRequest(data interface{}, err error, w http.ResponseWriter, r *http.Request) {
	httpx.WriteJson(w, http.StatusOK, response{
		Data: data,
	})
}

func HttpError(w http.ResponseWriter, httpCode int, appCode int64, desc string, message interface{}) {
	httpx.WriteJson(w, httpCode, response{
		Code:    appCode,
		Desc:    desc,
		Message: message,
	})
}


type response struct {
	Code    int64       `json:"code"`
	Desc    string      `json:"desc,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}