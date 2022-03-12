package svc

import (
	"github.com/1005281342/go-task/dispatcher/sdk"
	"github.com/1005281342/go-task/manager/internal/config"
	"github.com/hibiken/asynq"
)

type ServiceContext struct {
	Config     config.Config
	TaskClient *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:     c,
		TaskClient: sdk.NewClient(sdk.DefaultRedisClientOpt),
	}
}
