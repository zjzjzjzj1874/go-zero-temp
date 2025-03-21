package customer

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginBindPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginBindPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginBindPhoneLogic {
	return &LoginBindPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginBindPhoneLogic) LoginBindPhone(req *types.LoginBindPhoneReq) (resp *types.LoginBindPhoneResp, err error) {
	// todo: add your logic here and delete this line

	return
}
