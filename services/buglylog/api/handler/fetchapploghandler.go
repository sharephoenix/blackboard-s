package handler

import (
	"blackboards/shared/baseresponse"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type ApplogsRequest struct {
	Mobiles string `form:"mobiles"`
}

type ApplogsResposne struct {
	Data []ApplogsContext `json:"data"`
}

type ApplogsContext struct {
	Id         string `json:"id"`
	Mobile     string	`json:"mobile"`
	UserId     string	`json:"user_id"`
	LogUrl     string	`json:"log_url"`
	Message    string	`json:"message"`
	CreateTime string	`json:"create_time"`
}

// 获取崩溃详情?mobiles=110,112,113
func (handler CrashHandler)GetCrashApplogs(w http.ResponseWriter, r *http.Request) {
	var request ApplogsRequest
	err := httpx.Parse(r, &request)
	if err != nil {
		baseresponse.FormatResponseWithRequest(nil, err, w, r)
		return
	}
	mobiles := strings.Split(request.Mobiles, ",")
	logsInfos, err := handler.CrashLogic.FetchAppLogs(mobiles)
	var reskposne []ApplogsContext

	for _, item := range logsInfos.Infos {
		context := ApplogsContext{
			item.Id,
			item.Mobile,
			item.UserId,
			item.LogUrl,
			item.Message,
			item.CreateTime,
		}
		reskposne = append(reskposne, context)
	}

	baseresponse.FormatResponseWithRequest(ApplogsResposne{
		reskposne,
	}, nil, w, r)
}