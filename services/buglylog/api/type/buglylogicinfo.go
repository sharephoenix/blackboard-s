package buglylogicinfo

type CrashInfo struct {
	IssueId          string `json:"issue_id"`          // 问题 Id
	ErrorType        string `json:"error_type"`        // 错误类型
	AppVersion       string `json:"app_version"`       // 软件版本
	AppName          string `json:"app_name"`          // 产品名称
	CrashTimes       string `json:"crash_times"`       // 异常次数
	CrashDeviceNum   string `json:"crash_device_num"`  // 异常设备数量
	StackFlag        string `json:"stack_flag"`        // 堆栈特征
	CrashDescription string `json:"crash_description"` // crash 描述
	LastCrashTime    string `json:"last_crash_time"`   // 最后一次 crash 时间
	Status           string `json:"status"`            // 处理状态
	ProcessPerson    string `json:"process_person"`    // 处理人
}

type CrashDetail struct {
	CrashHash       string `json:"crash_hash"`       // crashHash
	IssueId         string `json:"issue_id"`         // issueId
	CrashId         string `json:"crash_id"`         // crashId
	UserId          string `json:"user_id"`          // 用户ID
	DeviceId        string `json:"device_id"`        // 设备ID
	UploadTime      string `json:"upload_time"`      // 上报时间
	CrashTime       string `json:"crash_time"`       // 发生时间
	AppBundleId     string `json:"app_bundle_id"`    // 应用包名
	AppVersion      string `json:"app_version"`      // 应用版本
	DeviceModel     string `json:"device_model"`     // 设备机型
	SystemVersion   string `json:"system_version"`   // 系统版本
	RomDetail       string `json:"rom_detail"`       // ROM详情
	CpuArchitecture string `json:"cpu_architecture"` // CPU架构
	IsJump          string `json:"is_jump"`          // 是否越狱
	MemorySize      string `json:"memory_size"`      // 可用内存大小
	StoreSizse      string `json:"store_sizse"`      // 可用存储空间
	SdSizse         string `json:"sd_sizse"`         // 可用SD卡大小
}

type AppLogPostInfoRequest struct {
	Infos []AppLogPostInfo `json:"infos"`
}

type AppLogPostInfo struct {
	Id 				string `json:"id"`	// primarKey
	Mobile          string `json:"mobile"`          // 用户，手机号
	User_id         string `json:"user_id"`         // 用户Id
	Log_url         string `json:"log_url"`         // logUrl
	Message         string `json:"message"`         // 所以信息
	Log_create_time string `json:"log_create_time"` // 创建上传日志的信息
	//Update_time     *string `json:"update_time"`
	//Create_time     *string `json:"create_time"`
}