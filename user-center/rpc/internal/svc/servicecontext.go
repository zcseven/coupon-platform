package svc

import (
	"coupon-platform/user-center/rpc/internal/config"
	"coupon-platform/user-center/rpc/internal/model/UserModel"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel UserModel.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataBase.MysqlDns)
	return &ServiceContext{
		Config:    c,
		UserModel: UserModel.NewUserInfoModel(conn),
	}
}
