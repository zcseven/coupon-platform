package UserModel

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserInfoModel = (*customUserInfoModel)(nil)

type (
	// UserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInfoModel.
	UserInfoModel interface {
		userInfoModel
		FindOneByTelephone(ctx context.Context, telephone string) (*UserInfo, error)
		FindOneByEmail(ctx context.Context, email string) (*UserInfo, error)
	}

	customUserInfoModel struct {
		*defaultUserInfoModel
	}
)

// NewUserInfoModel returns a model for the database table.
func NewUserInfoModel(conn sqlx.SqlConn) UserInfoModel {
	return &customUserInfoModel{
		defaultUserInfoModel: newUserInfoModel(conn),
	}
}

func (m *customUserInfoModel) FindOneByTelephone(ctx context.Context, telephone string) (*UserInfo, error) {
	query := fmt.Sprintf("select %s from %s where is_active and `telephone` = ? limit 1", userInfoRows, m.table)
	var resp UserInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, telephone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserInfoModel) FindOneByEmail(ctx context.Context, email string) (*UserInfo, error) {
	query := fmt.Sprintf("select %s from %s where is_active = 1 and  `email` = ? limit 1", userInfoRows, m.table)
	var resp UserInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
