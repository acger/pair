package handler

import (
	"net/http"

	"pair/service/pair/api/internal/logic"
	"pair/service/pair/api/internal/svc"
	"pair/service/pair/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func pairHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EleListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPairLogic(r.Context(), svcCtx)
		resp, err := l.Pair(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
