package customer

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWechatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWechatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWechatLogic {
	return &LoginWechatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWechatLogic) LoginWechat(req *types.LoginWechatReq) (resp *types.LoginWechatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
