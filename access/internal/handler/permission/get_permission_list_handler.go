package permission

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/logic/permission"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/types"
)

func GetPermissionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PermissionListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := permission.NewGetPermissionListLogic(r.Context(), svcCtx)
		resp, err := l.GetPermissionList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
