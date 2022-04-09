package ele

import (
	"net/http"

	"pair/service/pair/api/internal/logic/ele"
	"pair/service/pair/api/internal/svc"
	"pair/service/pair/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ElementHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EleViewReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := ele.NewElementLogic(r.Context(), svcCtx)
		resp, err := l.Element(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
