package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/logic/role"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/svc"
)

func DeleteRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := role.NewDeleteRoleLogic(r.Context(), svcCtx)
		resp, err := l.DeleteRole()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
