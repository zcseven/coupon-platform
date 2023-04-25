package UserRightsModel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserRightsModel = (*customUserRightsModel)(nil)

type (
	// UserRightsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRightsModel.
	UserRightsModel interface {
		userRightsModel
	}

	customUserRightsModel struct {
		*defaultUserRightsModel
	}
)

// NewUserRightsModel returns a model for the database table.
func NewUserRightsModel(conn sqlx.SqlConn) UserRightsModel {
	return &customUserRightsModel{
		defaultUserRightsModel: newUserRightsModel(conn),
	}
}
