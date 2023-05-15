package svc

import (
	"coupon-platform/order-center/api-service/internal/config"
	"coupon-platform/order-center/api-service/internal/model/OrderListModel"
	"coupon-platform/user-center/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	OrderListModel OrderListModel.OrderListModel
	UserRpc        usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataBase.MysqlDns)
	return &ServiceContext{
		Config:         c,
		OrderListModel: OrderListModel.NewOrderListModel(conn),
		UserRpc:        usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpc)),
	}
}
