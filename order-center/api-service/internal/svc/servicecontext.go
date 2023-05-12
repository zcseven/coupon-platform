package svc

import (
	"coupon-platform/order-center/api-service/internal/config"
	"coupon-platform/order-center/api-service/internal/model/OrderListModel"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	OrderListModel OrderListModel.OrderListModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataBase.MysqlDns)
	return &ServiceContext{
		Config:         c,
		OrderListModel: OrderListModel.NewOrderListModel(conn),
	}
}
