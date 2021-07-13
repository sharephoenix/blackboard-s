package logic

import (
	buglylogicinfo "blackboards/services/buglylog/api/type"
	"fmt"
)

// 需要用到的 rpc client
type CrashLogic struct {}

func (logic CrashLogic)UploadCrashInfo(infos []buglylogicinfo.CrashInfo) error {
	if len(infos) > 0 {
		fmt.Println(infos[0].AppVersion)
	}

	return nil
}

func (logic CrashLogic)UploadCrashDetail(details []buglylogicinfo.CrashDetail) error {
	if len(details) > 0 {
		fmt.Println(details[0].AppBundleId)
	}
	return nil
}