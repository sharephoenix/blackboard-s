package config

type Table struct {
	Myjob string
	Pingyin string
}

type DemoMysql struct {
	DataSource string
	TableName Table
}
