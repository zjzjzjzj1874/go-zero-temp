package callback

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/pay/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackWxLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallbackWxLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackWxLogic {
	return &CallbackWxLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CallbackWxLogic) CallbackWx(req *types.WechatNotifyReq) error {
	// todo: add your logic here and delete this line

	return nil
}
