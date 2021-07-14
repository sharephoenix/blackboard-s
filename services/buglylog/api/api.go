package main

import (
	"blackboards/services/buglylog/api/config"
	"blackboards/services/buglylog/api/handler"
	"blackboards/services/buglylog/api/logic"
	pb2 "blackboards/services/buglylog/rpc/pb"
	"flag"
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
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

	rpcClient := newBuglyClient()
	rpcBuglyClien := newBuglyRpcClient(rpcClient)
	crashLogic := logic.CrashLogic{rpcBuglyClien,}

	crashHandler := handler.CrashHandler{
		crashLogic,
	}

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

func newBuglyClient() zrpc.Client {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			// etcd 服务器
			Hosts: []string{"0.0.0.0:2379"},
			Key: "/bugly/rpc",
		},
		App: "iOS",
	})
	return conn
}

func newBuglyRpcClient(conn zrpc.Client) pb2.BuglyRpcServiceClient {
	clientnew := pb2.NewBuglyRpcServiceClient(conn.Conn())
	return clientnew
}