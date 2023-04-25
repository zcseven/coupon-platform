package user

import (
	"net/http"

	"coupon-platform/user-center/api-service/internal/logic/user"
	"coupon-platform/user-center/api-service/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserIntegralListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserIntegralListLogic(r.Context(), svcCtx)
		err := l.GetUserIntegralList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
