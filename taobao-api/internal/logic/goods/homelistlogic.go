package goods

import (
	"context"

	"coupon-platform/taobao-api/internal/svc"
	"coupon-platform/taobao-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeListLogic {
	return &HomeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeListLogic) HomeList(req *types.GoodsReq) (resp *types.GoodsResp, err error) {
	// todo: add your logic here and delete this line
	//taobaoke.TopClient。

	return
}
