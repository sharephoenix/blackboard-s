package config

type Table struct {
	Myjob string
}

type DemoMysql struct {
	DataSource string
	TableName Table
}
