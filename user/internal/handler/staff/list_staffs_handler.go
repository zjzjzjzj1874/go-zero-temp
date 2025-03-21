package staff

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/logic/staff"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/types"
)

func ListStaffsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListStaffsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := staff.NewListStaffsLogic(r.Context(), svcCtx)
		resp, err := l.ListStaffs(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
