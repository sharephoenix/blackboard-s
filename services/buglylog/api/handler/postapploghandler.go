package handler

import (
	buglylogicinfo "blackboards/services/buglylog/api/type"
	"blackboards/shared/baseresponse"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

func (handler CrashHandler)PostApplogs(w http.ResponseWriter, r *http.Request) {
	var request buglylogicinfo.AppLogPostInfoRequest
	httpx.Parse(r, &request)
	fmt.Println("requestInfo:", request)
	response, err := handler.CrashLogic.PostCrashInfos(request)
	if err != nil {
		baseresponse.FormatResponseWithRequest(nil, err, w, r)
		return
	}
	baseresponse.FormatResponseWithRequest(response, nil, w, r)
}
