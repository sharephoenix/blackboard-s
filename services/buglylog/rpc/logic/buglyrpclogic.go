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