package callback

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/logic/callback"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/types"
)

func CallbackWxHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WechatNotifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := callback.NewCallbackWxLogic(r.Context(), svcCtx)
		err := l.CallbackWx(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
