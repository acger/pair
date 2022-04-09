package chatGroup

import (
	"net/http"

	"pair/service/chat/api/internal/logic/chatGroup"
	"pair/service/chat/api/internal/svc"
	"pair/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatMessageSaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatMessageSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := chatGroup.NewChatMessageSaveLogic(r.Context(), svcCtx)
		resp, err := l.ChatMessageSave(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
