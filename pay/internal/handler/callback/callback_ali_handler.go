package callback

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/logic/callback"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/svc"
)

func CallbackAliHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := callback.NewCallbackAliLogic(r.Context(), svcCtx)
		err := l.CallbackAli()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
