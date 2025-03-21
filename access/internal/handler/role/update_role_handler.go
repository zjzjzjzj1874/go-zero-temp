package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/logic/role"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/types"
)

func UpdateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewUpdateRoleLogic(r.Context(), svcCtx)
		resp, err := l.UpdateRole(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
