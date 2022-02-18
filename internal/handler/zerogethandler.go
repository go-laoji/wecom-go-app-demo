package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"wecom-go-app-demo/internal/logic"
	"wecom-go-app-demo/internal/svc"
	"wecom-go-app-demo/internal/types"
)

func zerogetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MsgRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewZerogetLogic(r.Context(), ctx)
		resp, err := l.Zeroget(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			w.Write(resp.Data)
			httpx.Ok(w)
		}
	}
}
