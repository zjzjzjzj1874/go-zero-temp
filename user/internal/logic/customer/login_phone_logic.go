package customer

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginPhoneLogic {
	return &LoginPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginPhoneLogic) LoginPhone(req *types.LoginPhoneReq) (resp *types.LoginPhoneResp, err error) {
	// todo: add your logic here and delete this line

	return
}
