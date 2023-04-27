package user

import (
	"net/http"

	"coupon-platform/user-center/api-service/internal/logic/user"
	"coupon-platform/user-center/api-service/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshJwtHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewRefreshJwtLogic(r.Context(), svcCtx)
		resp, err := l.RefreshJwt()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
