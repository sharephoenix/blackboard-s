package model

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
	"log"
	"strings"
)

type CrashModel struct {
	sqlx.SqlConn
	CrashInfoTable string
	CrashDetailTable string
}

func (model CrashModel)InsertCrashInfos(infos []CrashInfoTable) error {
	var err error
	for _, info := range infos {
		fieldNames          := builderx.FieldNames(&CrashInfoTable{})
		rows                := strings.Join(fieldNames, ",")
		values := []string{
			info.IssueId, info.ErrorType, info.AppVersion, info.AppName,
			info.CrashTimes, info.CrashDeviceNum, info.StackFlag, info.CrashDescription,
			info.LastCrashTime, info.Status, info.ProcessPerson,
		}
		valueString := strings.Join(values, "\",\"")

		sql := `insert into ` + model.CrashInfoTable + ` (` + rows + `) values ` + `("` + valueString +`") `
		sql += `ON DUPLICATE KEY UPDATE crash_times="` + info.CrashTimes + `"`
		sql += `, last_crash_time= "` + info.LastCrashTime + `"`
		_, err = model.Exec(sql)
		if err != nil {
			log.Fatalln(sql)
			return err
		}
	}
	return err
}

func (model CrashModel)InsertCrashDetails(infos []CrashDetailTable) error {
	var err error
	for _, info := range infos {
		fieldNames          := builderx.FieldNames(&CrashDetailTable{})
		rows                := strings.Join(fieldNames, ",")

		values := []string{
			info.CrashHash,info.IssueId,info.CrashId,
			info.UserId,info.DeviceId,info.UploadTime,
			info.CrashTime,info.AppBundleId,info.AppVersion,
			info.DeviceModel,info.SystemVersion,info.RomDetail,
			info.CpuArchitecture,info.IsJump,info.MemorySize,
			info.StoreSizse,info.SdSize,
		}
		valueString := strings.Join(values, "\",\"")

		sql := `insert into ` + model.CrashDetailTable + ` (` + rows + `) values ` + `("` + valueString +`") `
		sql += `ON DUPLICATE KEY UPDATE crash_time="` + info.CrashTime + `"`
		_, err = model.Exec(sql)

		if err != nil {
			return err
		}
	}
	return err
}

type CrashInfoTable struct {
	IssueId          string `db:"issue_id"`          // 问题 Id
	ErrorType        string `db:"error_type"`        // 错误类型
	AppVersion       string `db:"app_version"`       // 软件版本
	AppName          string `db:"app_name"`          // 产品名称
	CrashTimes       string `db:"crash_times"`       // 异常次数
	CrashDeviceNum   string `db:"crash_device_num"`  // 异常设备数量
	StackFlag        string `db:"stack_flag"`        // 堆栈特征
	CrashDescription string `db:"crash_description"` // crash 描述
	LastCrashTime    string `db:"last_crash_time"`   // 最后一次 crash 时间
	Status           string `db:"status"`            // 处理状态
	ProcessPerson    string `db:"process_person"`    // 处理人
}

type CrashDetailTable struct {
	CrashHash       string `db:"crash_hash"`       // crashHash
	IssueId         string `db:"issue_id"`         // issueId
	CrashId         string `db:"crash_id"`         // crashId
	UserId          string `db:"user_id"`          // 用户ID
	DeviceId        string `db:"device_id"`        // 设备ID
	UploadTime      string `db:"upload_time"`      // 上报时间
	CrashTime       string `db:"crash_time"`       // 发生时间
	AppBundleId     string `db:"app_bundle_id"`    // 应用包名
	AppVersion      string `db:"app_version"`      // 应用版本
	DeviceModel     string `db:"device_model"`     // 设备机型
	SystemVersion   string `db:"system_version"`   // 系统版本
	RomDetail       string `db:"rom_detail"`       // ROM详情
	CpuArchitecture string `db:"cpu_architecture"` // CPU架构
	IsJump          string `db:"is_jump"`          // 是否越狱
	MemorySize      string `db:"memory_size"`      // 可用内存大小
	StoreSizse      string `db:"store_sizse"`      // 可用存储空间
	SdSize         string `db:"sd_size"`         // 可用SD卡大小
}