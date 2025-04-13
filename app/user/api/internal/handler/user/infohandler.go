package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"kkim/app/user/api/internal/logic/user"
	"kkim/app/user/api/internal/svc"
	"kkim/app/user/api/internal/types"
)

// 获取用户信息
func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
