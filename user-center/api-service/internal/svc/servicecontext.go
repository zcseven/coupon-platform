package svc

import (
	"coupon-platform/user-center/api-service/internal/config"
	"coupon-platform/user-center/api-service/internal/middleware"
	"coupon-platform/user-center/api-service/internal/model/UserIntegralHistoryModel"
	"coupon-platform/user-center/api-service/internal/model/UserModel"
	"coupon-platform/user-center/api-service/internal/model/UserRightsModel"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	//Request rest.Middleware
	//Cors            rest.Middleware
	AuthToken          rest.Middleware
	UserModel          UserModel.UserInfoModel
	UserRightsModel    UserRightsModel.UserRightsModel
	RightsHistoryModel UserIntegralHistoryModel.UserIntegralHistoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataBase.MysqlDns)
	return &ServiceContext{
		Config: c,
		//middleware
		//Request: middleware.NewRequestMiddleware().Handle,
		//Cors:    middleware.NewCorsMiddleware().Handle,
		AuthToken: middleware.NewAuthTokenMiddleware().Handle,

		//model
		UserModel:          UserModel.NewUserInfoModel(conn),
		UserRightsModel:    UserRightsModel.NewUserRightsModel(conn),
		RightsHistoryModel: UserIntegralHistoryModel.NewUserIntegralHistoryModel(conn),
	}
}
