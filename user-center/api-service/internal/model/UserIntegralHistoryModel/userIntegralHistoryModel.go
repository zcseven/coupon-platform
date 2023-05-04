package UserIntegralHistoryModel

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserIntegralHistoryModel = (*customUserIntegralHistoryModel)(nil)

type (
	// UserIntegralHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserIntegralHistoryModel.
	UserIntegralHistoryModel interface {
		userIntegralHistoryModel
		FindList(ctx context.Context, uid int64, pg int64, pagesize int64) ([]*UserIntegralHistory, error)
	}

	customUserIntegralHistoryModel struct {
		*defaultUserIntegralHistoryModel
	}
)

// NewUserIntegralHistoryModel returns a model for the database table.
func NewUserIntegralHistoryModel(conn sqlx.SqlConn) UserIntegralHistoryModel {
	return &customUserIntegralHistoryModel{
		defaultUserIntegralHistoryModel: newUserIntegralHistoryModel(conn),
	}
}

func (m *customUserIntegralHistoryModel) FindList(ctx context.Context, uid, pg, pagesize int64) ([]*UserIntegralHistory, error) {
	query := fmt.Sprintf("select %s from %s where is_active = 1 and `uid` = ? limit %d, %d ", userIntegralHistoryRows, m.table, (pg-1)*pagesize, pagesize)
	var resp []*UserIntegralHistory
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
