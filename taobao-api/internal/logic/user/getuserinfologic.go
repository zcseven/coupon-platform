package user

import (
	"context"

	"coupon-platform/taobao-api/internal/svc"
	"coupon-platform/taobao-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	list := map[int64]string{
		1: "zc",
		2: "zby",
		3: "zk",
	}

	name := "unknow"
	if na, ok := list[int64(req.Uid)]; ok {
		name = na
	}

	return &types.UserResp{
		Uid:      req.Uid,
		Username: name,
	}, nil
}
