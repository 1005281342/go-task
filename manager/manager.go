package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
	"net/http"
	"strings"

	"github.com/1005281342/go-task/manager/internal/config"
	"github.com/1005281342/go-task/manager/internal/server"
	"github.com/1005281342/go-task/manager/internal/svc"
	"github.com/1005281342/go-task/manager/manager"

	"github.com/1005281342/httproxy/grpchttp"
	"github.com/fullstorydev/grpchan"
	"github.com/fullstorydev/grpchan/httpgrpc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/polaris"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/manager.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewRpcServer(ctx)

	reg := grpchan.HandlerMap{}
	grpchttp.RegisterHandler(reg, srv, &manager.Rpc_ServiceDesc)

	var mux http.ServeMux
	httpgrpc.HandleServices(mux.HandleFunc, "/", reg, nil, nil)
	lis, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		panic(err)
	}
	var a = strings.Split(lis.Addr().String(), ":")
	var sPort = a[len(a)-1]
	logx.Infof("http port: %s", sPort)

	httpServer := http.Server{Handler: &mux}
	go httpServer.Serve(lis)
	defer httpServer.Close()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		manager.RegisterRpcServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	const (
		namespaceZRPC = "default"
		namespaceHTTP = "default"
	)

	// 注册zrpc服务到北极星
	if err = polaris.RegitserService(polaris.NewPolarisConfig(c.ListenOn,
		polaris.WithServiceName(c.Etcd.Key),
		polaris.WithNamespace(namespaceZRPC),
		polaris.WithHeartbeatInervalSec(5))); err != nil {
		logx.Errorf("注册zrpc到Polaris失败")
	}

	// 注册http服务到北极星
	var lo = "0.0.0.0:" + sPort
	if err = polaris.RegitserService(polaris.NewPolarisConfig(lo,
		polaris.WithServiceName(c.Etcd.Key+"-http"),
		polaris.WithNamespace(namespaceHTTP),
		polaris.WithHeartbeatInervalSec(5),
		polaris.WithProtocol("http"))); err != nil {
		panic(err)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
