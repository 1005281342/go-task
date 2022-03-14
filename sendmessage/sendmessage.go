package main

import (
	"flag"
	"fmt"
	"github.com/1005281342/go-task/pkg/httproxy"

	"github.com/1005281342/go-task/sendmessage/internal/config"
	"github.com/1005281342/go-task/sendmessage/internal/server"
	"github.com/1005281342/go-task/sendmessage/internal/svc"
	"github.com/1005281342/go-task/sendmessage/sendmessage"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/sendmessage.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewRpcServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sendmessage.RegisterRpcServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	httproxy.Init(c.RpcServerConf, srv, sendmessage.Rpc_ServiceDesc)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
