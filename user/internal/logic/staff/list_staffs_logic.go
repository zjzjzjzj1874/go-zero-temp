package staff

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListStaffsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListStaffsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStaffsLogic {
	return &ListStaffsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListStaffsLogic) ListStaffs(req *types.ListStaffsReq) (resp *types.ListStaffsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
