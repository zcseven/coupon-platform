package logic

import (
	"context"

	"coupon-platform/user-center/rpc/internal/svc"
	"coupon-platform/user-center/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *rpc.UserInfoReq) (*rpc.UserInfoRep, error) {
	// todo: add your logic here and delete this line

	return &rpc.UserInfoRep{}, nil
}
