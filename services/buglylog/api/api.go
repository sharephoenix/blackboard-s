package main

import (
	"blackboards/services/buglylog/api/config"
	"blackboards/services/buglylog/api/handler"
	"blackboards/services/buglylog/api/logic"
	"flag"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
	"net/http"
)

var configFile = flag.String("f", "etc/buglylog.json", "{}")

func main() {
	flag.Parse()
	var c config.BuglyLogConfig
	//conf.MustLoad(*configFile, &c)
	c = config.BuglyLogConfig{
		Port: 9090,
		Timeout: 100,
	}
	engine := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Port:     c.Port,
		Timeout:  c.Timeout,
		MaxConns: 500,
	})
	defer engine.Stop()
	crashLogic := logic.CrashLogic{}
	crashHandler := handler.CrashHandler{crashLogic}

	engine.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/post/crashinfos",
		Handler: crashHandler.PostCrashInfo,
	})
	engine.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/post/crashdetails",
		Handler: crashHandler.PostCrashDetail,
	})
	engine.Start()
}