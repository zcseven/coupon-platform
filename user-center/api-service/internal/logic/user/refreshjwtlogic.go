package user

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"coupon-platform/common/bases"
	"coupon-platform/common/util/tool"
	"coupon-platform/user-center/api-service/internal/svc"
	"coupon-platform/user-center/api-service/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshJwtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshJwtLogic {
	return &RefreshJwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshJwtLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *RefreshJwtLogic) RefreshJwt() (resp *types.RefreshJwtResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.RefreshJwtResp)
	uid, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil || uid == 0 {
		return nil, bases.NewCodeError(bases.DefaultCode, err.Error())
	}
	if err != nil {
		return resp, errors.New("参数错误")
	}
	now := time.Now().Unix()
	//token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, uid)

	token, _ := tool.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, uid)

	accessExpire := now + l.svcCtx.Config.Auth.AccessExpire
	refreshAfter := now + l.svcCtx.Config.Auth.AccessExpire/2

	resp.Result = types.RefreshJwtResult{
		Token:        token,
		AccessExpire: accessExpire,
		RefreshAfter: refreshAfter,
	}
	return resp, nil
}
