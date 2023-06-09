// Code generated by goctl. DO NOT EDIT.

package UserRightsModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userRightsFieldNames          = builder.RawFieldNames(&UserRights{})
	userRightsRows                = strings.Join(userRightsFieldNames, ",")
	userRightsRowsExpectAutoSet   = strings.Join(stringx.Remove(userRightsFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRightsRowsWithPlaceHolder = strings.Join(stringx.Remove(userRightsFieldNames, "`uid`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userRightsModel interface {
		Insert(ctx context.Context, data *UserRights) (sql.Result, error)
		FindOne(ctx context.Context, uid int64) (*UserRights, error)
		Update(ctx context.Context, data *UserRights) error
		Delete(ctx context.Context, uid int64) error
	}

	defaultUserRightsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserRights struct {
		Uid                 int64   `db:"uid"`
		IntegralNum         int64   `db:"integral_num"`         // 签到积分量
		MembershipLevel     int64   `db:"membership_level"`     // 会员等级
		FansNum             int64   `db:"fans_num"`             // 粉丝数量
		WithdrawalAmount    float64 `db:"withdrawal_amount"`    // 提现金额
		RedEnvelope         float64 `db:"red_envelope"`         // 未到位的红包金额
		PromotionCommission float64 `db:"promotion_commission"` // 推广提成金额
		HistoryWithdrawal   float64 `db:"history_withdrawal"`   // 历史提现金额
		IsActive            int64   `db:"is_active"`            // 有效性
	}
)

func newUserRightsModel(conn sqlx.SqlConn) *defaultUserRightsModel {
	return &defaultUserRightsModel{
		conn:  conn,
		table: "`user_rights`",
	}
}

func (m *defaultUserRightsModel) Delete(ctx context.Context, uid int64) error {
	query := fmt.Sprintf("delete from %s where `uid` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, uid)
	return err
}

func (m *defaultUserRightsModel) FindOne(ctx context.Context, uid int64) (*UserRights, error) {
	query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", userRightsRows, m.table)
	var resp UserRights
	err := m.conn.QueryRowCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserRightsModel) Insert(ctx context.Context, data *UserRights) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRightsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.IntegralNum, data.MembershipLevel, data.FansNum, data.WithdrawalAmount, data.RedEnvelope, data.PromotionCommission, data.HistoryWithdrawal, data.IsActive)
	return ret, err
}

func (m *defaultUserRightsModel) Update(ctx context.Context, data *UserRights) error {
	query := fmt.Sprintf("update %s set %s where `uid` = ?", m.table, userRightsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.IntegralNum, data.MembershipLevel, data.FansNum, data.WithdrawalAmount, data.RedEnvelope, data.PromotionCommission, data.HistoryWithdrawal, data.IsActive, data.Uid)
	return err
}

func (m *defaultUserRightsModel) tableName() string {
	return m.table
}
