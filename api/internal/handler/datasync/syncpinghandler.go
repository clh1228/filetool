package handler

import (
	"net/http"

	logic "filetool/api/internal/logic/datasync"
	"filetool/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SyncPingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSyncPingLogic(r.Context(), ctx)
		err := l.SyncPing()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
