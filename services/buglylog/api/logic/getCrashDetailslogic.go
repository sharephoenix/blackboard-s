package logic

import (
	"blackboards/services/buglylog/rpc/pb"
	"context"
)

// get
func (logic CrashLogic)FetchCrashDetails(mobiles []string) (*pb.BuglyDetailResponse, error) {
	response, err := logic.BuglyRpcClient.GetBuglyDetails(context.TODO(), &pb.BuglyDetailIssueRequest{
		Mobiles: mobiles,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}