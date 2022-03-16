package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/zrpc/registry/polaris"

	"github.com/1005281342/go-task/manager/internal/config"
	"github.com/1005281342/go-task/manager/internal/server"
	"github.com/1005281342/go-task/manager/internal/svc"
	"github.com/1005281342/go-task/manager/manager"
	"github.com/1005281342/httproxy/grpchttp"

	"github.com/rookie-ninja/rk-entry/v2/entry"
	"github.com/rookie-ninja/rk-zero/boot"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/manager.yaml", "the config file")

//go:embed boot.yaml
var boot []byte

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewRpcServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		manager.RegisterRpcServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//httproxy.Init(c.RpcServerConf, srv, manager.Rpc_ServiceDesc)
	var httpPort uint64
	go func() {
		// Bootstrap preload entries
		rkentry.BootstrapPreloadEntryYAML(boot)

		// Bootstrap zero entry from boot config
		res := rkzero.RegisterZeroEntryYAML(boot)

		// Get ZeroEntry
		zeroEntry := res["go-zero"].(*rkzero.ZeroEntry)
		httpPort = zeroEntry.Port

		// 注册路由
		grpchttp.RegisterWithGoZero(srv, &manager.Rpc_ServiceDesc, zeroEntry.Server)

		// Bootstrap zero entry
		zeroEntry.Bootstrap(context.Background())

		// Wait for shutdown signal
		rkentry.GlobalAppCtx.WaitForShutdownSig()

		// Interrupt zero entry
		zeroEntry.Interrupt(context.Background())
	}()

	const (
		namespaceZRPC = "default"
		namespaceHTTP = "default"
	)

	var err error
	// 注册zrpc服务
	if err = polaris.RegitserService(polaris.NewPolarisConfig(c.ListenOn,
		polaris.WithServiceName(c.Etcd.Key),
		polaris.WithNamespace(namespaceZRPC),
		polaris.WithHeartbeatInervalSec(5))); err != nil {
		logx.Errorf("注册zrpc到Polaris失败")
	}

	// 注册http服务
	var lo = fmt.Sprintf("0.0.0.0:%d", httpPort)
	fmt.Printf("http: %s \n", lo)
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
