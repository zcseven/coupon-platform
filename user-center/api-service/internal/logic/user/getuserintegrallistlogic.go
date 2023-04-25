package user

import (
	"context"

	"coupon-platform/user-center/api-service/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserIntegralListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserIntegralListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserIntegralListLogic {
	return &GetUserIntegralListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserIntegralListLogic) GetUserIntegralList() error {
	// todo: add your logic here and delete this line

	return nil
}
