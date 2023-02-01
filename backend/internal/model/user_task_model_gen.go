// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userTaskFieldNames          = builder.RawFieldNames(&UserTask{})
	userTaskRows                = strings.Join(userTaskFieldNames, ",")
	userTaskRowsExpectAutoSet   = strings.Join(stringx.Remove(userTaskFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), ",")
	userTaskRowsWithPlaceHolder = strings.Join(stringx.Remove(userTaskFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), "=?,") + "=?"
)

type (
	userTaskModel interface {
		Insert(ctx context.Context, data *UserTask) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserTask, error)
		FindByAccount(ctx context.Context, account string) ([]UserTask, error)
		FindByAccountAndChainAndAddressAndType(ctx context.Context,
			account string, chain int32, address string, addressType string) ([]UserTask, error)
		Update(ctx context.Context, data *UserTask) error
		BatchUpdateStatusById(ctx context.Context, status string, id []int64) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserTaskModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserTask struct {
		Id         int64     `db:"id"`
		Account    string    `db:"account"`
		Chain      int64     `db:"chain"`
		Address    string    `db:"address"`
		Type       string    `db:"type"`
		Status     string    `db:"status"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUserTaskModel(conn sqlx.SqlConn) *defaultUserTaskModel {
	return &defaultUserTaskModel{
		conn:  conn,
		table: "`user_task`",
	}
}

func (m *defaultUserTaskModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserTaskModel) FindOne(ctx context.Context, id int64) (*UserTask, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userTaskRows, m.table)
	var resp UserTask
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultUserTaskModel) FindByAccount(ctx context.Context, account string) ([]UserTask, error) {
	query := fmt.Sprintf("select %s from %s where `account` = ?", userTaskRows, m.table)
	var resp []UserTask
	err := m.conn.QueryRowsCtx(ctx, &resp, query, account)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultUserTaskModel) FindByAccountAndChainAndAddressAndType(ctx context.Context,
	account string, chain int32, address string, addressType string) ([]UserTask, error) {
	query := fmt.Sprintf("select %s from %s where `account` = ? and `chain` = ? and `address` = ? and `type` = ?", userTaskRows, m.table)
	var resp []UserTask
	err := m.conn.QueryRowsCtx(ctx, &resp, query, account, chain, address, addressType)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultUserTaskModel) Insert(ctx context.Context, data *UserTask) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userTaskRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Account, data.Chain, data.Address, data.Type, data.Status)
	return ret, err
}

func (m *defaultUserTaskModel) Update(ctx context.Context, data *UserTask) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userTaskRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Account, data.Chain, data.Address, data.Type, data.Status, data.Id)
	return err
}

func (m *defaultUserTaskModel) BatchUpdateStatusById(ctx context.Context, status string, idArray []int64) error {
	if len(idArray) == 0 {
		return nil
	}
	b, _ := json.Marshal(idArray)
	str := string(b)
	str = str[1:len(str)-1]
	query := fmt.Sprintf("update %s set `status` = ? where `id` in (%s)", m.table, str)
	_, err := m.conn.ExecCtx(ctx, query, status)
	return err
}

func (m *defaultUserTaskModel) tableName() string {
	return m.table
}
