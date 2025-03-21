package swagger

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/logic/swagger"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/svc"
)

func SwaggerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := swagger.NewSwaggerLogic(r.Context(), svcCtx)
		err := l.Swagger()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
