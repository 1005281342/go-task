package logic

import (
	"context"
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
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *manager.AddReq) (*manager.AddRsp, error) {

	var _, err = l.svcCtx.TaskClient.EnqueueContext(l.ctx, asynq.NewTask(in.Task.Name, []byte(in.Task.Payload),
		asynq.ProcessAt(time.Unix(in.Task.At, 0)), asynq.Queue(in.Task.Queue)))

	return &manager.AddRsp{}, err
}
