// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package usercenter

import (
	"context"

	"coupon-platform/user-center/rpc/rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserInfoRep = rpc.UserInfoRep
	UserInfoReq = rpc.UserInfoReq

	Usercenter interface {
		GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRep, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRep, error) {
	client := rpc.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}
