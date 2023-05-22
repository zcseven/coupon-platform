package goods

import (
	"net/http"

	"coupon-platform/taobao-api/internal/logic/goods"
	"coupon-platform/taobao-api/internal/svc"
	"coupon-platform/taobao-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goods.NewHomeListLogic(r.Context(), svcCtx)
		resp, err := l.HomeList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
