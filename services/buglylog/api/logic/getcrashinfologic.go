package logic

import (
	"blackboards/services/buglylog/rpc/pb"
	"context"
)

// get
func (logic CrashLogic)FetchCrashInfos(versions []string) (*pb.BuglyInfoResponse, error) {
	response, err := logic.BuglyRpcClient.GetBuglyInfo(context.TODO(), &pb.BuglyInfoIssueRequest{
		Versions: versions,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

