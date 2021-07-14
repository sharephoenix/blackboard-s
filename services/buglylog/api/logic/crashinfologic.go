package logic

import (
	buglylogicinfo "blackboards/services/buglylog/api/type"
	"blackboards/services/buglylog/rpc/pb"
	"context"
)

// 需要用到的 rpc client
type CrashLogic struct {
	BuglyRpcClient pb.BuglyRpcServiceClient
}

func (logic CrashLogic)UploadCrashInfo(infos []buglylogicinfo.CrashInfo) error {
	buglyInfos := []*pb.BuglyInfo{}

	for _, info := range infos {
		buglyInfo := pb.BuglyInfo{
			IssueId: info.IssueId,
			ErrorType: info.ErrorType,
			AppVersion: info.AppVersion,
			AppName: info.AppName,
			CrashTimes: info.CrashTimes,
			CrashDeviceNum: info.CrashDeviceNum,
			StackFlag: info.StackFlag,
			CrashDescription: info.CrashDescription,
			LastCrashTime: info.LastCrashTime,
			Status: info.Status,
			ProcessPerson: info.ProcessPerson,
		}
		buglyInfos = append(buglyInfos, &buglyInfo)
	}
	request := pb.BuglyInfoRequest{
		Infos: buglyInfos,
	}
	logic.BuglyRpcClient.UploadBuglyInfo(context.TODO(), &request)
	return nil
}

func (logic CrashLogic)UploadCrashDetail(details []buglylogicinfo.CrashDetail) error {
	buglyDetales := []*pb.BugylDetail{}
	for _, info := range details {
		buglyDetale := pb.BugylDetail{
			CrashHash: info.CrashHash,
			IssueId: info.IssueId,
			CrashId: info.CrashId,
			UserId: info.UserId,
			DeviceId: info.DeviceId,
			UploadTime: info.UploadTime,
			CrashTime: info.CrashTime,
			AppBundleId: info.AppBundleId,
			AppVersion: info.AppVersion,
			DeviceModel: info.DeviceModel,
			SystemVersion: info.SystemVersion,
			RomDetail: info.RomDetail,
			CpuArchitecture: info.CpuArchitecture,
			IsJump: info.IsJump,
			MemorySize: info.MemorySize,
			StoreSizse: info.StoreSizse,
			SdSizse: info.SdSizse,
		}
		buglyDetales = append(buglyDetales, &buglyDetale)
	}
	request := pb.BuglyDetailRequest{
		Infos: buglyDetales,
	}
	_, err := logic.BuglyRpcClient.UploadBuglyDetails(context.TODO(), &request)
	return err
}