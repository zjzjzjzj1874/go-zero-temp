package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/logic"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/types"
)

func livenessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LivenessRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLivenessLogic(r.Context(), svcCtx)
		resp, err := l.Liveness(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
