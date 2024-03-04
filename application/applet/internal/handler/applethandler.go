package handler

import (
	"net/http"

	"czdemo/application/applet/internal/logic"
	"czdemo/application/applet/internal/svc"
	"czdemo/application/applet/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AppletHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAppletLogic(r.Context(), svcCtx)
		resp, err := l.Applet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
