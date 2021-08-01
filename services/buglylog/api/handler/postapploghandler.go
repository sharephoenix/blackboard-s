package handler

import (
	"blackboards/shared/baseresponse"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

type AppLogPostInfoRequest struct {
	Infos []AppLogPostInfo `json:"infos"`
}

type AppLogPostInfo struct {
	Id 				string `json:"id"`	// primarKey
	Mobile          string `json:"mobile"`          // 用户，手机号
	User_id         string `json:"user_id"`         // 用户Id
	Log_url         string `json:"log_url"`         // logUrl
	Message         string `json:"message"`         // 所以信息
	Log_create_time string `json:"log_create_time"` // 创建上传日志的信息
	Update_time     *string `json:"update_time"`
	Create_time     *string `json:"create_time"`
}



func (handler CrashHandler)PostApplogs(w http.ResponseWriter, r *http.Request) {
	var request AppLogPostInfoRequest
	httpx.Parse(r, &request)
	response, err := handler.CrashLogic.PostCrashInfos(request)
	if err != nil {
		baseresponse.FormatResponseWithRequest(nil, err, w, r)
		return
	}
	baseresponse.FormatResponseWithRequest(response, nil, w, r)
}
