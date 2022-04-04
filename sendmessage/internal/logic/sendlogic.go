package logic

import (
	"context"
	"github.com/1005281342/go-task/sendmessage/internal/service/ability"

	"github.com/1005281342/go-task/sendmessage/internal/svc"
	"github.com/1005281342/go-task/sendmessage/sendmessage"

	"github.com/1005281342/log"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		//Logger: logx.WithContext(ctx),
		Logger: log.NewGoZeroELKLoggerWithContext(ctx, log.WithAppName("sendmessage"), log.WithFuncName("Send")),
	}
}

func (l *SendLogic) Send(in *sendmessage.SendReq) (*sendmessage.SendRsp, error) {

	var (
		out = &sendmessage.SendRsp{}
		a   = ability.NewASend(in, out, l.Logger)
		err error
	)

	defer func() {
		if err != nil {
			l.Logger.Errorf("Send failed: %+v", err)
		}
	}()

	if err = a.Check(l.ctx); err != nil {
		return nil, err
	}
	if err = a.Process(l.ctx); err != nil {
		return nil, err
	}
	return out, nil
}
