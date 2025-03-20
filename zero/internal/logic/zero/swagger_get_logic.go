package zero

import (
	"context"
	"io/ioutil"

	"github/zjzjzjzj1874/go-zero-temp/zero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwaggerGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwaggerGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwaggerGetLogic {
	return &SwaggerGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwaggerGetLogic) SwaggerGet() ([]byte, error) {
	logx.Infof("读取swagger文件")
	return ioutil.ReadFile(l.svcCtx.Config.SwaggerPath)
}
