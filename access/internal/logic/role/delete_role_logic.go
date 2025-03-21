package role

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/access/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoleLogic) DeleteRole() (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
