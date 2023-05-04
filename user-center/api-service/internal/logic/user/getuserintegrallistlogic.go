package user

import (
	"context"
	"encoding/json"

	"coupon-platform/common/bases"
	"coupon-platform/user-center/api-service/internal/svc"
	"coupon-platform/user-center/api-service/internal/types"

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

func (l *GetUserIntegralListLogic) GetUserIntegralList(req *types.UserRightsListReq) (resp *types.UserRightsListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserRightsListResp)
	uid, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil || uid == 0 {
		return nil, bases.NewCodeError(bases.DefaultCode, err.Error())
	}
	list, err := l.svcCtx.RightsHistoryModel.FindList(l.ctx, uid, req.Pg, req.PageSize)
	if err != nil {
		return nil, bases.NewDefaultError(err.Error())
	}

	result := make([]*types.UserRightsList, 0)
	resp.Result.PageList.PageSize = int(req.PageSize)
	resp.Result.PageList.Pg = int(req.Pg)

	for _, v := range list {
		result = append(result, &types.UserRightsList{
			Uid:          v.Uid,
			IntegralType: int(v.IntegralType),
			IngegralNum:  v.IntegralNum,
			CreateAt:     v.CreateAt.Unix(),
		})
	}
	resp.Result.List = result
	return resp, nil
}
