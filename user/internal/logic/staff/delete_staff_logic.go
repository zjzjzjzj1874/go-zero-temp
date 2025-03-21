package staff

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStaffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStaffLogic {
	return &DeleteStaffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStaffLogic) DeleteStaff(req *types.DeleteStaffReq) (resp *types.DeleteStaffResp, err error) {
	// todo: add your logic here and delete this line

	return
}
