package permission

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/access/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionListLogic {
	return &GetPermissionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermissionListLogic) GetPermissionList(req *types.PermissionListReq) (resp *types.PermissionListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
