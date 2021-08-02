package logic

import (
	"blackboards/services/buglylog/rpc/pb"
	"context"
	"fmt"
)

// get
func (logic CrashLogic)FetchCrashInfos(versions []string) (*pb.BuglyInfoResponse, error) {
	response, err := logic.BuglyRpcClient.GetBuglyInfo(context.TODO(), &pb.BuglyInfoIssueRequest{
		Versions: versions,
	})

	if err != nil {
		fmt.Println("FetchCrashInfos-error:", err.Error())
		return nil, err
	}
	return response, nil
}

