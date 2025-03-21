package access

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/logic/access"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/svc"
)

func SwaggerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := access.NewSwaggerLogic(r.Context(), svcCtx)
		resp, err := l.Swagger()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
