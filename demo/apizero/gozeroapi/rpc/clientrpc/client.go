package clientrpc

import (
	"blackboards/demo/apizero/gozeroapi/rpc/pb"
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
)

type HelloWorldInternalClient struct {}

func (c HelloWorldInternalClient)NewClient() *pb.HelloWorldClient {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			// etcd 服务器
			Hosts: []string{"0.0.0.0:2379"},
			Key: "/testdir/testkey",
		},
		App: "iOS",
	})

	clientnew := pb.NewHelloWorldClient(conn.Conn())
	return &clientnew
}

