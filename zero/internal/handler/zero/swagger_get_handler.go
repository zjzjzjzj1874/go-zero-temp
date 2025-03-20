package zero

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/logic/zero"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/svc"
)

func SwaggerGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := zero.NewSwaggerGetLogic(r.Context(), svcCtx)
		resp, err := l.SwaggerGet()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			_, _ = w.Write(resp)
			httpx.Ok(w)
		}
	}
}
