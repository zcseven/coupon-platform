package user

import (
	"context"
	"time"

	"coupon-platform/common/bases"
	"coupon-platform/common/util/tool"
	"coupon-platform/user-center/api-service/internal/model/UserModel"
	"coupon-platform/user-center/api-service/internal/svc"
	"coupon-platform/user-center/api-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	// 判断是否存在
	userinfo, err := l.svcCtx.UserModel.FindOneByTelephone(l.ctx, req.Telephone)
	if err != nil && err != sqlx.ErrNotFound {
		return nil, bases.NewCodeError(1002, err.Error())
	}
	if userinfo.Uid > 0 {
		return nil, bases.NewCodeError(1002, "手机号已存在")
	}

	userinfos, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil && err != UserModel.ErrNotFound {
		return nil, bases.NewCodeError(1002, err.Error())
	}

	if userinfos.Uid > 0 {
		return nil, bases.NewCodeError(1002, "邮箱已存在")
	}

	//入库处理
	data := new(UserModel.UserInfo)
	data.Telephone = req.Telephone
	data.Email = req.Email
	data.Password = tool.MD5(req.Password)
	data.UserName = req.UserName

	status, err := l.svcCtx.UserModel.Insert(l.ctx, data)
	if err != nil {
		return nil, bases.NewCodeError(1002, err.Error())
	}

	uid, _ := status.LastInsertId()
	token, _ := tool.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, uid)
	resp = new(types.RegisterResp)
	resp.Result.Email = req.Email
	resp.Result.Telephone = req.Telephone
	resp.Result.Uid = uid
	resp.Result.Token = token
	return resp, nil
}
