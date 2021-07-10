package main

import (
	"blackboards/demo/apizero/gozeroapi/api/config"
	"blackboards/demo/apizero/gozeroapi/api/handler"
	"blackboards/demo/apizero/gozeroapi/api/svc"
	"blackboards/demo/apizero/gozeroapi/rpc/clientrpc"
	"flag"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
	"net/http"
)

var configFile = flag.String("f", "etc/buglylog.json", "{}")

func main() {
	flag.Parse()
	//var c config.BuglyLogConfig
	c := config.BuglyLogConfig{
		Port: 9090,
		Timeout: 100,
	}
	//conf.MustLoad(*configFile, &c)
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
	// go-zero 代码
	helloClient := clientrpc.HelloWorldInternalClient{}
	clientnew := helloClient.NewClient()

	svc := svc.HelloWorldContext{
		*clientnew,
	}
	handler.RegisterHandlers(engine, &svc)
	engine.Start()
}