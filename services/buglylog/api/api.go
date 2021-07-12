package main

import (
	"blackboards/services/buglylog/api/config"
	"blackboards/services/buglylog/api/handler"
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

	engine.Use(handler.First)
	engine.Use(handler.Second)
	engine.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: handler.Handle,
	})
	engine.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/post/crashinfo",
		Handler: handler.PostCrashInfo,
	})
	engine.Start()
}