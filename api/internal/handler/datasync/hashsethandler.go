package handler

import (
	"net/http"

	logic "filetool/api/internal/logic/datasync"
	"filetool/api/internal/svc"
	"filetool/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HashSetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetHashReq
		// log.Infof("request:>>>", r)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHashSetLogic(r.Context(), ctx)
		resp, err := l.HashSet(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
