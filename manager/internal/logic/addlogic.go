package logic

import (
	"context"
	"fmt"
	"github.com/1005281342/go-task/manager/internal/metrics"
	"github.com/1005281342/log"
	"github.com/hibiken/asynq"
	"time"

	"github.com/1005281342/go-task/manager/internal/svc"
	"github.com/1005281342/go-task/manager/manager"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		//Logger: logx.WithContext(ctx),
		Logger: log.NewGoZeroELKLoggerWithContext(ctx, log.WithAppName("manager")),
	}
}

func (l *AddLogic) Add(in *manager.AddReq) (*manager.AddRsp, error) {
	metrics.Report(metrics.Add)

	// TODO：时间校验间隔可以配置到远程配置
	if in.Task.At < time.Now().Add(-time.Second).Unix() {
		//if in.Task.At < time.Now().Add(-time.Hour).Unix() {
		metrics.Report(metrics.AddFailed)
		return &manager.AddRsp{}, fmt.Errorf("定时时间不合法")
	}

	var _, err = l.svcCtx.TaskClient.EnqueueContext(l.ctx,
		asynq.NewTask(in.Task.Name, []byte(in.Task.Payload),
			asynq.ProcessAt(time.Unix(in.Task.At, 0)),
			asynq.Queue(in.Task.Queue)))
	if err != nil {
		metrics.Report(metrics.AddFailed)
		return &manager.AddRsp{}, err
	}

	l.Logger.Infof("任务%+v添加成功", in)
	return &manager.AddRsp{}, nil
}
