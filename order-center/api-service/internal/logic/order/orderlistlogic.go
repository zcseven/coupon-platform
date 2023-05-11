package order

import (
	"context"

	"coupon-platform/order-center/api-service/internal/svc"
	"coupon-platform/order-center/api-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.OrderListReq) (resp *types.OrderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
