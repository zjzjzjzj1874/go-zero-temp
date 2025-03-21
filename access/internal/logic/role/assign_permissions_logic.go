package role

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/access/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignPermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignPermissionsLogic {
	return &AssignPermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignPermissionsLogic) AssignPermissions(req *types.RolePermissionReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
