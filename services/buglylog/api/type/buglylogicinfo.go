package buglylogicinfo

type CrashInfo struct {
	IssueId string `json:"issue_id"`	// 问题 Id
	ErrorType string `json:"error_type"` // 错误类型
	AppVersion string `json:"app_version"`	// 软件版本
	AppName string `json:"app_name"`	// 产品名称
	CrashTimes string `json:"crash_times"`	// 异常次数
	CrashDeviceNum string `json:"crash_device_num"`	// 异常设备数量
	StackFlag string `json:"stack_flag"`	// 堆栈特征
	CrashDescription string	`json:"crash_description"`	// crash 描述
	LastCrashTime string `json:"last_crash_time"` // 最后一次 crash 时间
	Status string `json:"status"`	// 处理状态
	ProcessPerson string `json:"process_person"` // 处理人
}

