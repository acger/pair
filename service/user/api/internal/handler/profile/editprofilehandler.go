package profile

import (
	"net/http"

	"pair/service/user/api/internal/logic/profile"
	"pair/service/user/api/internal/svc"
	"pair/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := profile.NewEditProfileLogic(r.Context(), svcCtx)
		resp, err := l.EditProfile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
