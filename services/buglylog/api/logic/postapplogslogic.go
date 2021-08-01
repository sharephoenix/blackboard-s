package logic

import (
	"blackboards/services/buglylog/api/handler"
	"blackboards/services/buglylog/rpc/pb"
	"context"
)

type PostCrashInfoLogicResponse struct {
	Success []string `json:"success"`
	Fails []string `json:"fails"`
}

// post
func (logic CrashLogic)PostCrashInfos(request handler.AppLogPostInfoRequest) (PostCrashInfoLogicResponse, error) {
	var success []string
	var fails []string
	for _, item := range request.Infos {
		_, err := logic.BuglyRpcClient.InsertAppLogInfos(context.TODO(), &pb.AppLogsRequest{
			Id: item.Id,
			Mobile: item.Mobile,
			UserId: item.User_id,
			LogUrl: item.Log_url,
			Message: item.Message,
			CreateTime: item.Log_create_time,
		})
		if err != nil {
			fails = append(fails, item.Mobile)
		} else {
			success = append(success, item.Mobile)
		}
	}
	return PostCrashInfoLogicResponse{
		Success: success,
		Fails: fails,
	}, nil
}

