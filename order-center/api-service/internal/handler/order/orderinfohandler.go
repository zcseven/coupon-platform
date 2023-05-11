package order

import (
	"net/http"

	"coupon-platform/order-center/api-service/internal/logic/order"
	"coupon-platform/order-center/api-service/internal/svc"
	"coupon-platform/order-center/api-service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewOrderInfoLogic(r.Context(), svcCtx)
		resp, err := l.OrderInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
