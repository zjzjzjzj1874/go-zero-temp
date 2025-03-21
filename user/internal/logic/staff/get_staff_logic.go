package staff

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStaffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStaffLogic {
	return &GetStaffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStaffLogic) GetStaff(req *types.GetStaffReq) (resp *types.GetStaffResp, err error) {
	// todo: add your logic here and delete this line

	return
}
