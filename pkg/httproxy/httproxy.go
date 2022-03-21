package httproxy

import (
	"net"
	"net/http"
	"strings"

	"github.com/1005281342/httproxy/grpchttp"
	"github.com/fullstorydev/grpchan"
	"github.com/fullstorydev/grpchan/httpgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/polaris"
	"google.golang.org/grpc"
)

func Init(c zrpc.RpcServerConf, srv interface{}, gRPCDesc grpc.ServiceDesc) {
	reg := grpchan.HandlerMap{}
	grpchttp.RegisterHandler(reg, srv, &gRPCDesc)

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

	const (
		namespaceZRPC = "default"
		namespaceHTTP = "default"
	)

	// 注册zrpc服务到北极星
	if err = polaris.RegitserService(polaris.NewPolarisConfig(c.ListenOn,
		polaris.WithServiceName(c.Etcd.Key),
		polaris.WithNamespace(namespaceZRPC),
		polaris.WithHeartbeatInervalSec(5))); err != nil {
		panic(err)
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
}
