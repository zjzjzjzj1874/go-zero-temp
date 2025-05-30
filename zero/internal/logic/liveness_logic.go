package logic

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/zero/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LivenessLogic) Liveness(req *types.LivenessRequest) (resp *types.LivenessResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
