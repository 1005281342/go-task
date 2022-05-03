package ability

import (
	"context"
	"encoding/json"
	"github.com/1005281342/go-task/pkg/apollo"
	"github.com/1005281342/go-task/pkg/errs"
	"github.com/1005281342/go-task/sendmessage/sendmessage"
	"github.com/zeromicro/go-zero/core/logx"
)

type ASend struct {
	req    *sendmessage.SendReq
	rsp    *sendmessage.SendRsp
	logger logx.Logger
}

type Users struct {
	Black bool     `json:"black"`
	Users []string `json:"users"`
}

func NewASend(req *sendmessage.SendReq,
	rsp *sendmessage.SendRsp,
	logger logx.Logger) *ASend {

	return &ASend{req: req, rsp: rsp, logger: logger}
}

func (a *ASend) Check(ctx context.Context) error {
	var err error
	if err = a.checkSender(ctx); err != nil {
		return err
	}
	if err = a.checkReceiver(ctx); err != nil {
		return err
	}

	return nil
}

func (a *ASend) checkSender(ctx context.Context) error {
	var err = a.checkAccount(ctx, a.req.GetSender().GetId())
	if err != nil {
		return err
	}

	a.logger.Infof("account: %s, 发送者账号校验通过", a.req.GetSender().GetId())
	return nil
}

func (a *ASend) checkReceiver(ctx context.Context) error {
	var err = a.checkAccount(ctx, a.req.GetReceiver().GetId())
	if err != nil {
		return err
	}

	a.logger.Infof("account: %s, 接收者账号校验通过", a.req.GetReceiver().GetId())
	return nil
}

func (a *ASend) checkAccount(ctx context.Context, id string) error {
	if id == "" {
		a.logger.Errorf("req: %+v err: %+v", a.req, errs.ErrAccountIllegal)
		return errs.ErrAccountIllegal
	}

	var s = apollo.Config().GetStringValue("users", "")
	if s == "" {
		return nil
	}
	var us Users
	var err = json.Unmarshal([]byte(s), &us)
	if err != nil {
		return err
	}

	a.logger.Infof("us: %+v", us)

	for _, u := range us.Users {
		if u != id {
			continue
		}
		if us.Black {
			return errs.ErrAccountBlack
		}
		return nil
	}

	if !us.Black {
		return errs.ErrAccountNonWhitelist
	}
	return nil
}

func (a *ASend) Process(ctx context.Context) error {
	defer a.render(ctx)

	return a.handler(ctx)
}

func (a *ASend) handler(ctx context.Context) error {
	if a.req.GetMsg().GetType() == sendmessage.MessageType_markdown {
		a.req.Msg.Message = "格式是Markdown，内容为：" + a.req.Msg.Message
	}
	a.logger.Infof("msg: %s", a.req.GetMsg().GetMessage())
	return nil
}

func (a *ASend) render(ctx context.Context) {

}
