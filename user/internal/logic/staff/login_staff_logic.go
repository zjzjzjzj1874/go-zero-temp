package staff

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginStaffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginStaffLogic {
	return &LoginStaffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginStaffLogic) LoginStaff(req *types.LoginStaffReq) (resp *types.LoginStaffResp, err error) {
	// todo: add your logic here and delete this line

	return
}
