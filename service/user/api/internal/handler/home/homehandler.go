package home

import (
	"net/http"

	"pair/service/user/api/internal/logic/home"
	"pair/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := home.NewHomeLogic(r.Context(), svcCtx)
		resp, err := l.Home()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
