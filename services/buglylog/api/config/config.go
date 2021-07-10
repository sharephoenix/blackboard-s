package config

type BuglyLogConfig struct {
	Port int `json:"port"`
	Timeout int64 `json:"timeout"`
}