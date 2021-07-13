package config

type BuglyRpcConfig struct {
	MysqlBugly struct {
		DataSource string
		Table struct{
			CrashInfoTable string
			CrashDetailTable string
		}
	}
}
