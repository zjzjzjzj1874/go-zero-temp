package staff

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStaffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStaffLogic {
	return &AddStaffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStaffLogic) AddStaff(req *types.AddStaffReq) (resp *types.AddStaffResp, err error) {
	// todo: add your logic here and delete this line

	return
}
