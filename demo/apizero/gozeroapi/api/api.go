package main

import (
	"blackboards/demo/apizero/gozeroapi/api/config"
	"blackboards/demo/apizero/gozeroapi/api/handler"
	"flag"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
	"net/http"
)

var configFile = flag.String("f", "etc/buglylog.json", "{}")

func main() {
	flag.Parse()
	var c config.BuglyLogConfig
	conf.MustLoad(*configFile, &c)
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
	engine.Start()
}