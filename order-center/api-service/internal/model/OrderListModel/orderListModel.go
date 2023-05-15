package OrderListModel

import (
	"context"
	"coupon-platform/order-center/api-service/internal/types"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderListModel = (*customOrderListModel)(nil)

type (
	// OrderListModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderListModel.
	OrderListModel interface {
		orderListModel
		FindList(ctx context.Context, req *types.OrderListReq) ([]*OrderList, error)
	}

	customOrderListModel struct {
		*defaultOrderListModel
	}
)

// NewOrderListModel returns a model for the database table.
func NewOrderListModel(conn sqlx.SqlConn) OrderListModel {
	return &customOrderListModel{
		defaultOrderListModel: newOrderListModel(conn),
	}
}

//获取列表
func (m *customOrderListModel) FindList(ctx context.Context, req *types.OrderListReq) ([]*OrderList, error) {
	var resp []*OrderList
	query := fmt.Sprintf("select %s from %s where `uid` = ? limit %d, %d order by id desc", orderListRows, m.table, (req.Pg-1)*req.PageSize, req.PageSize)
	
	err := m.conn.QueryRowsCtx(ctx, &resp, query)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
