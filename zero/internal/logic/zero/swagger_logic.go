package zero

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/zero/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwaggerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwaggerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwaggerLogic {
	return &SwaggerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwaggerLogic) Swagger() (resp *types.Response, err error) {
	return
}
