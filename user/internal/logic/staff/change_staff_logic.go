package staff

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeStaffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeStaffLogic {
	return &ChangeStaffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeStaffLogic) ChangeStaff(req *types.ChangeStaffReq) (resp *types.ChangeStaffResp, err error) {
	// todo: add your logic here and delete this line

	return
}
