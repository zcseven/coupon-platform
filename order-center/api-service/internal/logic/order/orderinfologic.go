package order

import (
	"context"

	"coupon-platform/order-center/api-service/internal/svc"
	"coupon-platform/order-center/api-service/internal/types"
	"coupon-platform/user-center/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderInfoLogic {
	return &OrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderInfoLogic) OrderInfo(req *types.OrderInfoReq) (resp *types.OrderInfoResp, err error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &usercenter.UserInfoReq{
		Uid: req.Uid,
	})
	if err != nil {
		return nil, err
	}

	return &types.OrderInfoResp{
		Uid:        userInfo.Uid,
		UserName:   userInfo.Telephone,
		OrderSn:    "11121212",
		TradeNo:    "11121212",
		IsActive:   1,
		CreateTime: 11121212,
	}, nil
}
