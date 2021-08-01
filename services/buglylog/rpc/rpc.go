package main

import (
	config2 "blackboards/services/buglylog/rpc/config"
	logic2 "blackboards/services/buglylog/rpc/logic"
	model2 "blackboards/services/buglylog/rpc/model"
	"blackboards/services/buglylog/rpc/pb"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var config = flag.String("f", "etc/rpc.json", "{}")

func main() {
	flag.Parse()
	var c config2.BuglyRpcConfig
	conf.MustLoad(*config, &c)


	var csv zrpc.RpcServerConf
	conf.MustLoad(*config, &csv)

	bytes, _ := json.Marshal(c)
	fmt.Println(string(bytes))

	bytes, _ = json.Marshal(csv)
	fmt.Println(string(bytes))

	dytMysql := sqlx.NewMysql(c.MysqlBugly.DataSource)

	model := model2.CrashModel{
		dytMysql,
		c.MysqlBugly.Table.CrashInfoTable,
		c.MysqlBugly.Table.CrashDetailTable,
		c.MysqlBugly.Table.CrashLogTable,
	}

	logicServer := logic2.BuglyRpcLogic{
		model,
	}

	server := zrpc.MustNewServer(csv, func(grpcServer *grpc.Server) {
		pb.RegisterBuglyRpcServiceServer(grpcServer, &logicServer)
	})
	defer server.Stop()
	server.Start()
}
