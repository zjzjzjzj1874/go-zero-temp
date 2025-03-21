package pay

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/pay/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayResultLogic {
	return &PayResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayResultLogic) PayResult(req *types.PayQueryReq) (resp *types.PayResultResp, err error) {
	// todo: add your logic here and delete this line

	return
}
