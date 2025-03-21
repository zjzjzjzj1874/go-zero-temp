package customer

import (
	"context"

	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCellPhoneAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCellPhoneAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCellPhoneAuthLogic {
	return &GetCellPhoneAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCellPhoneAuthLogic) GetCellPhoneAuth(req *types.GetCellPhoneAuthReq) (resp *types.GetCellPhoneAuthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
