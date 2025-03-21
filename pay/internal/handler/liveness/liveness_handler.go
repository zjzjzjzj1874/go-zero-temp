package liveness

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/logic/liveness"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/svc"
)

func LivenessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := liveness.NewLivenessLogic(r.Context(), svcCtx)
		err := l.Liveness()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
