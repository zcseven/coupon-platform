// Code generated by goctl. DO NOT EDIT.

package UserIntegralHistoryModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userIntegralHistoryFieldNames          = builder.RawFieldNames(&UserIntegralHistory{})
	userIntegralHistoryRows                = strings.Join(userIntegralHistoryFieldNames, ",")
	userIntegralHistoryRowsExpectAutoSet   = strings.Join(stringx.Remove(userIntegralHistoryFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userIntegralHistoryRowsWithPlaceHolder = strings.Join(stringx.Remove(userIntegralHistoryFieldNames, "`uid`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userIntegralHistoryModel interface {
		Insert(ctx context.Context, data *UserIntegralHistory) (sql.Result, error)
		FindOne(ctx context.Context, uid int64) (*UserIntegralHistory, error)
		Update(ctx context.Context, data *UserIntegralHistory) error
		Delete(ctx context.Context, uid int64) error
	}

	defaultUserIntegralHistoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserIntegralHistory struct {
		Uid          int64     `db:"uid"`
		IntegralType int64     `db:"integral_type"` // 积分类型:1签到;2查看帮助视频;3完成首单优惠;4召回好友;5邀请新用户
		IntegralNum  int64     `db:"integral_num"`  // 积分量
		CreateAt     time.Time `db:"create_at"`     // 签到时间
		IsActive     int64     `db:"is_active"`     // 有效性
	}
)

func newUserIntegralHistoryModel(conn sqlx.SqlConn) *defaultUserIntegralHistoryModel {
	return &defaultUserIntegralHistoryModel{
		conn:  conn,
		table: "`user_integral_history`",
	}
}

func (m *defaultUserIntegralHistoryModel) Delete(ctx context.Context, uid int64) error {
	query := fmt.Sprintf("delete from %s where `uid` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, uid)
	return err
}

func (m *defaultUserIntegralHistoryModel) FindOne(ctx context.Context, uid int64) (*UserIntegralHistory, error) {
	query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", userIntegralHistoryRows, m.table)
	var resp UserIntegralHistory
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

func (m *defaultUserIntegralHistoryModel) Insert(ctx context.Context, data *UserIntegralHistory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userIntegralHistoryRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.IntegralType, data.IntegralNum, data.IsActive)
	return ret, err
}

func (m *defaultUserIntegralHistoryModel) Update(ctx context.Context, data *UserIntegralHistory) error {
	query := fmt.Sprintf("update %s set %s where `uid` = ?", m.table, userIntegralHistoryRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.IntegralType, data.IntegralNum, data.IsActive, data.Uid)
	return err
}

func (m *defaultUserIntegralHistoryModel) tableName() string {
	return m.table
}
