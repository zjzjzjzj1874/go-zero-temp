package permission

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/logic/permission"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/svc"
)

func DeletePermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := permission.NewDeletePermissionLogic(r.Context(), svcCtx)
		resp, err := l.DeletePermission()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
