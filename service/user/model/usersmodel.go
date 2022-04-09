package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`create_time`", "`update_time`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUsersIdPrefix      = "cache:users:id:"
	cacheUsersAccountPrefix = "cache:users:account:"
)

type (
	UsersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindOneByAccount(ctx context.Context, account string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id           int64     `db:"id"`
		Account      string    `db:"account"`
		Name         string    `db:"name"`
		Avatar       string    `db:"avatar"`
		Password     string    `db:"password"`
		Gender       string    `db:"gender"`
		Status       string    `db:"status"`
		Level        string    `db:"level"`
		Birthday     string    `db:"birthday"`
		Contact      string    `db:"contact"`
		Introduction string    `db:"introduction"`
		CreateTime   time.Time `db:"create_time" gorm:"-"`
		UpdateTime   time.Time `db:"update_time" gorm:"-"`
	}
)

func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersAccountKey := fmt.Sprintf("%s%v", cacheUsersAccountPrefix, data.Account)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Account, data.Name, data.Avatar, data.Password, data.Gender, data.Status, data.Level, data.Birthday, data.Contact, data.Introduction)
	}, usersIdKey, usersAccountKey)
	return ret, err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, usersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByAccount(ctx context.Context, account string) (*Users, error) {
	usersAccountKey := fmt.Sprintf("%s%v", cacheUsersAccountPrefix, account)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersAccountKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `account` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, account); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(ctx context.Context, data *Users) error {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersAccountKey := fmt.Sprintf("%s%v", cacheUsersAccountPrefix, data.Account)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Account, data.Name, data.Avatar, data.Password, data.Gender, data.Status, data.Level, data.Birthday, data.Contact, data.Introduction, data.Id)
	}, usersIdKey, usersAccountKey)
	return err
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	usersAccountKey := fmt.Sprintf("%s%v", cacheUsersAccountPrefix, data.Account)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usersIdKey, usersAccountKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
