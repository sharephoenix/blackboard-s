package logic

import (
	"blackboards/services/buglylog/rpc/pb"
	"context"
)

// get
func (logic CrashLogic)FetchAppLogs(mobiles []string) (*pb.AppLogsInfo, error) {
	response, err := logic.BuglyRpcClient.GetAppLogInfos(context.TODO(), &pb.GetAppLogsRequest{
		Mobles: mobiles,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}