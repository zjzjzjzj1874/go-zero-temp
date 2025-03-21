package callback

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/svc"
)

type CallbackAliLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallbackAliLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackAliLogic {
	return &CallbackAliLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CallbackAliLogic) CallbackAli() error {
	// todo: add your logic here and delete this line

	return nil
}
