package handler

import (
	"blackboards/shared/baseresponse"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type CrashInfoRequest struct {
	Versions string `form:"versions"`
}

type CrashInfoResponse struct {
	Data []CrashInfo `json:"data"`
}

type CrashInfo struct {
	IssueId          string `json:"issue_id"`
	ErrorType        string`json:"error_type"`
	AppVersion       string`json:"app_version"`
	AppName          string`json:"app_name"`
	CrashTimes       string`json:"crash_times"`
	CrashDeviceNum   string`json:"crash_device_num"`
	StackFlag        string`json:"stack_flag"`
	CrashDescription string`json:"crash_description"`
	LastCrashTime    string`json:"last_crash_time"`
	Status           string`json:"status"`
	ProcessPerson    string`json:"process_person"`
}

// 获取崩溃信息?viersions=1.1.1,1.2.3
func (handler CrashHandler)GetCrashInfos(w http.ResponseWriter, r *http.Request) {
	var request CrashInfoRequest
	err := httpx.Parse(r, &request)
	if err != nil {
		fmt.Println("request error!!")
		baseresponse.FormatResponseWithRequest(nil, err, w, r)
		return
	}
	fmt.Println("this is request" + request.Versions)
	versions := strings.Split(request.Versions, ",")
	crashInfos, err := handler.CrashLogic.FetchCrashInfos(versions)
	var reskposne []CrashInfo

	for _, item := range crashInfos.Infos {
		context := CrashInfo{
			item.IssueId,
			item.ErrorType,
			item.AppVersion,
			item.AppName,
			item.CrashTimes,
			item.CrashDeviceNum,
			item.StackFlag,
			item.CrashDescription,
			item.LastCrashTime,
			item.Status,
			item.ProcessPerson,
		}
		reskposne = append(reskposne, context)
	}

	baseresponse.FormatResponseWithRequest(CrashInfoResponse{
		reskposne,
	}, nil, w, r)
}
