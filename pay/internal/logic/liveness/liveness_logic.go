package liveness

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/svc"
)

type LivenessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLivenessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LivenessLogic {
	return &LivenessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LivenessLogic) Liveness() error {
	// todo: add your logic here and delete this line

	return nil
}
