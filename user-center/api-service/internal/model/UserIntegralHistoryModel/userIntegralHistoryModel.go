package UserIntegralHistoryModel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserIntegralHistoryModel = (*customUserIntegralHistoryModel)(nil)

type (
	// UserIntegralHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserIntegralHistoryModel.
	UserIntegralHistoryModel interface {
		userIntegralHistoryModel
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
