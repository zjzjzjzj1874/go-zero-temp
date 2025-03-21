package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/logic"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
)

func LivenessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLivenessLogic(r.Context(), svcCtx)
		err := l.Liveness()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
