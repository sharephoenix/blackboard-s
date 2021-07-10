package main

import (
	"encoding/json"
	"fmt"
	pb "blackboards/demo/apizero/gozeroapi/rpc/pb"
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
	"golang.org/x/net/context"
	"testing"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func TestRpc(t *testing.T) {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			// etcd 服务器
			Hosts: []string{"0.0.0.0:2379"},
			Key: "/testdir/testkey",
		},
		App: "iOS",
	})

	clientnew := pb.NewHelloWorldClient(conn.Conn())

	request := pb.HelloRequest{
		Name: "alex luan",
	}
	response, err := clientnew.SayHello(context.Background(),&request)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		btys, err := json.Marshal(response)
		if err != nil {
			fmt.Println("err")
		} else {
			fmt.Println(string(btys))
		}
	}
}
