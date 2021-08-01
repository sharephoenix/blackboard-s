package model

func (model CrashModel) InsertCrashLog(log CrashLogModelTable) error {
	sqlString := `insert into ` + model.CrashLogTable + ` (id, mobile, user_id, log_url, messge, log_create_time) values (`
	sqlString += `"` + string(log.Id) + `"` + `,`
	sqlString += `"` + log.Mobile + `"` + `,`
	sqlString += `"` + log.User_id + `"` + `,`
	sqlString += `"` + log.Log_url + `"` + `,`
	sqlString += `"` + log.Message + `"` + `,`
	sqlString += `"` + log.Log_create_time + `"`
	sqlString += `)`
	sqlString += ` ON DUPLICATE KEY UPDATE log_url="` + log.Log_url + `"`
	sqlString += `, mobile="` + log.Mobile + `"`
	sqlString += `, user_id="` + log.User_id + `"`
	sqlString += `, messge="` + log.Message + `"`
	sqlString += `, log_create_time="` + log.Log_create_time + `"`
	_, err := model.Exec(sqlString)
	if err != nil {
		return err
	}
	return nil
}

func (model CrashModel) GetCrashLogs(mobiles []string) ([]CrashLogModelTable, error) {
	var logs []CrashLogModelTable
	if len(mobiles) < 1 {
		return nil, CustomError{"没有输入手机号码"}
	}
	sqlString := `select * from ` + model.CrashLogTable+ ` where mobile=` + `"` + mobiles[0] + `"`
	for _, mobile := range mobiles[1:] {
		sqlString += ` or mobile=` + `"` + mobile + `"`
	}
	err := model.QueryRows(&logs, sqlString)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

type CrashLogModelTable struct {
	Id 		string `db:"id"`	// primarKey
	Mobile          string `db:"mobile"`          // 用户，手机号
	User_id         string `db:"user_id"`         // 用户Id
	Log_url         string `db:"log_url"`         // logUrl
	Message         string `db:"message"`         // 所以信息
	Log_create_time string `db:"log_create_time"` // 创建上传日志的信息
	Update_time     *string `db:"update_time"`
	Create_time     *string `db:"create_time"`
}

type CustomError struct {
	ErrString string
}
func (e CustomError)Error() string {
	return e.ErrString
}

