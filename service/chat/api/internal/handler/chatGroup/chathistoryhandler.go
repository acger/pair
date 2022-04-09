package chatGroup

import (
	"net/http"

	"pair/service/chat/api/internal/logic/chatGroup"
	"pair/service/chat/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chatGroup.NewChatHistoryLogic(r.Context(), svcCtx)
		resp, err := l.ChatHistory()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
