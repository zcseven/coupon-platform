package user

import (
	"context"
	"coupon-platform/common/bases"
	"coupon-platform/user-center/api-service/internal/svc"
	"coupon-platform/user-center/api-service/internal/types"
	"encoding/json"

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
	resp = new(types.UserResp)
	uid, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil || uid == 0 {
		return nil, bases.NewCodeError(1002, err.Error())
	}
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, err
	}

	userRights, _ := l.svcCtx.UserRightsModel.FindOne(l.ctx, uid)

	resp.Result = types.UserRespResult{
		Uid:                 userInfo.Uid,
		Username:            userInfo.UserName,
		Headpic:             userInfo.HeadPic,
		WithdrawalAmount:    userRights.WithdrawalAmount,
		FansNum:             userRights.FansNum,
		IntegralNum:         userRights.IntegralNum,
		RedEnvelope:         userRights.RedEnvelope,
		PromotionCommission: userRights.PromotionCommission,
		HistoryWithdrawal:   userRights.HistoryWithdrawal,
		MembershipLevel:     userRights.MembershipLevel,
	}
	return resp, nil
}
