package main

import (
	pb "blackboards/demo/apizero/gozeroapi/rpc/pb"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}
const (
	port = ":50051"
)

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("######### get client request name :"+in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

var configFile = flag.String("f", "etc/config.json", "{}")

func main() {
	flag.Parse()

	var c zrpc.RpcServerConf
	conf.MustLoad(*configFile, &c)

	server := zrpc.MustNewServer(c, func(grpcServer *grpc.Server) {
		pb.RegisterHelloWorldServer(grpcServer, &server{})//RegisterGraceServiceServer(grpcServer, NewGracefulServer())
	})
	defer server.Stop()
	server.Start()
}
