package user

import (
	"context"
	"errors"

	"coupon-platform/user-center/api-service/internal/svc"
	"coupon-platform/user-center/api-service/internal/types"

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

	if req.Uid == 0 {
		return nil, errors.New("参数错误")
	}

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Uid)
	if err != nil {
		return nil, err
	}

	userRights, err := l.svcCtx.UserRightsModel.FindOne(l.ctx, req.Uid)

	return &types.UserResp{
		Uid:                 req.Uid,
		Username:            userInfo.UserName,
		Headpic:             userInfo.HeadPic,
		WithdrawalAmount:    userRights.WithdrawalAmount,
		FansNum:             userRights.FansNum,
		IntegralNum:         userRights.IntegralNum,
		RedEnvelope:         userRights.RedEnvelope,
		PromotionCommission: userRights.PromotionCommission,
		HistoryWithdrawal:   userRights.HistoryWithdrawal,
		MembershipLevel:     userRights.MembershipLevel,
	}, nil
}
