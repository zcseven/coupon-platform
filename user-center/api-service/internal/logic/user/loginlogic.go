package user

import (
	"context"
	"strconv"
	"time"

	"coupon-platform/user-center/api-service/internal/svc"
	"coupon-platform/user-center/api-service/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	telephone, _ := strconv.Atoi(req.Telephone)
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(telephone))
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	token, _ := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	//accessExpire = now + l.svcCtx.Config.Auth.AccessExpire   // 过期时间
	//refreshAfter = now + l.svcCtx.Config.Auth.AccessExpire/2  // 告知前端什么时候刷新 token
	return &types.LoginResp{
		Uid:      userInfo.Id,
		Username: userInfo.UserName,
		Headpic:  userInfo.HeadPic,
		Email:    userInfo.Email,
		Token:    token,
	}, nil
}
