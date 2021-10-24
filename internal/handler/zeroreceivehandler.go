package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"zero-svr/internal/logic"
	"zero-svr/internal/svc"
	"zero-svr/internal/types"
)

func zeroreceiveHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MsgRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewZeroreceiveLogic(r.Context(), ctx)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			httpx.Error(w, err)
		}
		resp, err := l.Zeroreceive(req, body)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
