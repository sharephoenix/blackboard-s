package logic

import (
	"blackboards/services/buglylog/rpc/model"
	"blackboards/services/buglylog/rpc/pb"
	"context"
)

type BuglyRpcLogic struct {
	Model model.CrashModel
}

/// 将数据插入数据库
func (logic BuglyRpcLogic)UploadBuglyInfo(ctx context.Context, infos *pb.BuglyInfoRequest) (*pb.Response, error) {
	tables := []model.CrashInfoTable{}

	for _, value := range infos.Infos {
		table := model.CrashInfoTable{
			value.IssueId,
			value.ErrorType,
			value.AppVersion,
			value.AppName,
			value.CrashTimes,
			value.CrashDeviceNum,
			value.StackFlag,
			value.CrashDescription,
			value.LastCrashTime,
			value.Status,
			value.ProcessPerson,
		}
		tables = append(tables, table)
	}

	err := logic.Model.InsertCrashInfos(tables)
	return &pb.Response{
		Code: 0,
	}, err
}

/// 将数据插入数据库
func (logic BuglyRpcLogic)UploadBuglyDetails(ctx context.Context, infos *pb.BuglyDetailRequest) (*pb.Response, error) {
	tables := []model.CrashDetailTable{}

	for _, value := range infos.Infos {
		table := model.CrashDetailTable{
			value.CrashHash,
			value.IssueId,
			value.CrashId,
			value.UserId,
			value.DeviceId,
			value.UploadTime,
			value.CrashTime,
			value.AppBundleId,
			value.AppVersion,
			value.DeviceModel,
			value.SystemVersion,
			value.RomDetail,
			value.CpuArchitecture,
			value.IsJump,
			value.MemorySize,
			value.StoreSizse,
			value.SdSizse,
		}
		tables = append(tables, table)
	}
	err := logic.Model.InsertCrashDetails(tables)
	return &pb.Response{
		Code: 0,
	}, err
}

func (logic BuglyRpcLogic)InsertAppLogInfos(ctx context.Context, request *pb.AppLogsRequest) (*pb.Response, error) {
	table := model.CrashLogModelTable{
		request.Id,
		request.Mobile,
		request.UserId,
		request.UserId,
		request.LogUrl,
		request.CreateTime,
		nil,
		nil,
	}
	err := logic.Model.InsertCrashLog(table)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Code: 1}, nil
}

func (logic BuglyRpcLogic)GetAppLogInfos(ctx context.Context, request *pb.GetAppLogsRequest) (*pb.AppLogsInfo, error) {
	infos, err := logic.Model.GetCrashLogs(request.Mobles)
	if err != nil {
		return nil, err
	}
	responseInfos := []*pb.AppLogsRequest{}
	for _, log := range infos {
		info := pb.AppLogsRequest{
			Mobile: log.Mobile,
			UserId: log.User_id,
			LogUrl: log.Log_url,
			Message: log.Message,
		}
		responseInfos = append(responseInfos, &info)
	}

	return &pb.AppLogsInfo {
		Infos: responseInfos,
	}, nil
}
func (logic BuglyRpcLogic)GetBuglyInfo(ctx context.Context, request *pb.BuglyInfoIssueRequest) (*pb.BuglyInfoResponse, error) {
	crashInfos, err := logic.Model.GetCrashInfos(request.Versions)
	if err != nil {
		return nil, err
	}
	var buglyInfos []*pb.BuglyInfo
	for _, info := range crashInfos {
		infoCurrent := pb.BuglyInfo{
			IssueId: info.IssueId,
			ErrorType: info.ErrorType,
			AppVersion: info.AppVersion,
			AppName: info.AppName,
			CrashDescription: info.CrashDescription,
			Status: info.Status,
			CrashDeviceNum: info.CrashTimes,
			StackFlag: info.StackFlag,
			LastCrashTime: info.LastCrashTime,
			ProcessPerson: info.ProcessPerson,
		}
		buglyInfos = append(buglyInfos, &infoCurrent)
	}
	return &pb.BuglyInfoResponse{
		Infos: buglyInfos,
	}, nil
}

func (logic BuglyRpcLogic)GetBuglyDetails(ctx context.Context, request *pb.BuglyDetailIssueRequest) (*pb.BuglyDetailResponse, error) {
	crashInfos, err := logic.Model.GetCrashDetailInfos(request.Mobiles)
	if err != nil {
		return nil, err
	}
	var buglyInfos []*pb.BugylDetail
	for _, info := range crashInfos {
		infoCurrent := pb.BugylDetail{
			IssueId: info.IssueId,
			UserId: info.UserId,
			CrashId: info.CrashId,
			DeviceModel: info.DeviceModel,
			DeviceId: info.DeviceId,
			RomDetail: info.RomDetail,
			UploadTime: info.UploadTime,
			CrashHash: info.CrashHash,
			CrashTime: info.CrashTime,
			AppVersion: info.AppVersion,
			AppBundleId: info.AppBundleId,
			SystemVersion: info.SystemVersion,
		}
		buglyInfos = append(buglyInfos, &infoCurrent)
	}
	return &pb.BuglyDetailResponse{
		Detals: buglyInfos,
	}, nil
}



