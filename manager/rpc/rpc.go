// Code generated by goctl. DO NOT EDIT!
// Source: manager.proto

package rpc

import (
	"context"

	"github.com/1005281342/go-task/manager/manager"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddReq   = manager.AddReq
	AddRsp   = manager.AddRsp
	TaskInfo = manager.TaskInfo

	Rpc interface {
		Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRsp, error)
	}

	defaultRpc struct {
		cli zrpc.Client
	}
)

func NewRpc(cli zrpc.Client) Rpc {
	return &defaultRpc{
		cli: cli,
	}
}

func (m *defaultRpc) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRsp, error) {
	client := manager.NewRpcClient(m.cli.Conn())
	return client.Add(ctx, in, opts...)
}
