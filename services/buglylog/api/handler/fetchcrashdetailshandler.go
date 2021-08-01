package handler

import (
	"blackboards/shared/baseresponse"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type CrashDetalRequest struct {
	Mobiles string `form:"mobiles"`
}

type CrashDetailResponse struct {
	Data []CrashDetail `json:"data"`
}

type CrashDetail struct {
	CrashHash       string `json:"crash_hash"`
	IssueId         string `json:"issue_id"`
	CrashId         string `json:"crash_id"`
	UserId          string `json:"user_id"`
	DeviceId        string `json:"device_id"`
	UploadTime      string `json:"upload_time"`
	CrashTime       string `json:"crash_time"`
	AppBundleId     string `json:"app_bundle_id"`
	AppVersion      string `json:"app_version"`
	DeviceModel     string `json:"device_model"`
	SystemVersion   string `json:"system_version"`
	RomDetail       string `json:"rom_detail"`
	CpuArchitecture string `json:"cpu_architecture"`
	IsJump          string `json:"is_jump"`
	MemorySize      string `json:"memory_size"`
	StoreSizse      string `json:"store_sizse"`
	SdSizse         string `json:"sd_sizse"`
}

// 获取崩溃详情?mobiles=110,112,113
func (handler CrashHandler) GetCrashDetail(w http.ResponseWriter, r *http.Request) {
	var request CrashDetalRequest
	err := httpx.Parse(r, request)
	if err != nil {
		baseresponse.FormatResponseWithRequest(nil, err, w, r)
		return
	}
	mobiles := strings.Split(request.Mobiles, ",")
	logsInfos, err := handler.CrashLogic.FetchCrashDetails(mobiles)
	var reskposne []CrashDetail

	for _, item := range logsInfos.Detals {
		context := CrashDetail{
			item.CrashHash,
			item.IssueId,
			item.CrashId,
			item.UserId,
			item.DeviceId,
			item.UploadTime,
			item.CrashTime,
			item.AppBundleId,
			item.AppVersion,
			item.DeviceModel,
			item.SystemVersion,
			item.RomDetail,
			item.CpuArchitecture,
			item.IsJump,
			item.MemorySize,
			item.StoreSizse,
			item.SdSizse,
		}
		reskposne = append(reskposne, context)
	}

	baseresponse.FormatResponseWithRequest(CrashDetailResponse{
		reskposne,
	}, nil, w, r)
}
