package chatGroup

import (
	"net/http"

	"pair/service/chat/api/internal/logic/chatGroup"
	"pair/service/chat/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHistoryNumberHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chatGroup.NewChatHistoryNumberLogic(r.Context(), svcCtx)
		resp, err := l.ChatHistoryNumber()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
