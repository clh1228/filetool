package chat

import (
	"net/http"

	"filetool/api/internal/logic/chat"
	"filetool/api/internal/svc"
	"filetool/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatRecordListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRecordListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := chat.NewChatRecordListLogic(r.Context(), svcCtx)
		resp, err := l.ChatRecordList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
