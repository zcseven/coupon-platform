// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"coupon-platform/user-center/rpc/internal/logic"
	"coupon-platform/user-center/rpc/internal/svc"
	"coupon-platform/user-center/rpc/rpc"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) GetUserInfo(ctx context.Context, in *rpc.UserInfoReq) (*rpc.UserInfoRep, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}
