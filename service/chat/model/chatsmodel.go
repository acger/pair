package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
	"time"
)

var (
	chatsFieldNames          = builder.RawFieldNames(&Chats{})
	chatsRows                = strings.Join(chatsFieldNames, ",")
	chatsRowsExpectAutoSet   = strings.Join(stringx.Remove(chatsFieldNames, "`create_time`", "`update_time`"), ",")
	chatsRowsWithPlaceHolder = strings.Join(stringx.Remove(chatsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheChatsIdPrefix = "cache:chats:id:"
)

type (
	ChatsModel interface {
		Insert(ctx context.Context, data *Chats) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Chats, error)
		Update(ctx context.Context, data *Chats) error
		Delete(ctx context.Context, id int64) error
	}

	defaultChatsModel struct {
		sqlc.CachedConn
		table string
	}

	ClientChatMsg struct {
		Uid       string `json:"uid"`
		ToUid     string `json:"to_uid"`
		Message   string `json:"message"`
		ClientNum int    `json:"client_num"`
	}

	Chats struct {
		Id         int64     `db:"id" json:"id"`
		CreateTime time.Time `db:"create_time" json:"create_time" gorm:"-"`
		UpdateTime time.Time `db:"update_time" json:"update_time" gorm:"-"`
		Uid        int64     `db:"uid" json:"uid"`
		ToUid      int64     `db:"to_uid" json:"to_uid"`
		Message    string    `db:"message" json:"message"`
		Status     int64     `db:"status" json:"status"`
	}
)

func NewChatsModel(conn sqlx.SqlConn, c cache.CacheConf) ChatsModel {
	return &defaultChatsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`chats`",
	}
}

func (m *defaultChatsModel) Insert(ctx context.Context, data *Chats) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, chatsRowsExpectAutoSet)
	ret, err := m.ExecNoCacheCtx(ctx, query, data.Id, data.Uid, data.ToUid, data.Message, data.Status)

	return ret, err
}

func (m *defaultChatsModel) FindOne(ctx context.Context, id int64) (*Chats, error) {
	chatsIdKey := fmt.Sprintf("%s%v", cacheChatsIdPrefix, id)
	var resp Chats
	err := m.QueryRowCtx(ctx, &resp, chatsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatsRows, m.table)
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

func (m *defaultChatsModel) Update(ctx context.Context, data *Chats) error {
	chatsIdKey := fmt.Sprintf("%s%v", cacheChatsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, chatsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.ToUid, data.Message, data.Status, data.Id)
	}, chatsIdKey)
	return err
}

func (m *defaultChatsModel) Delete(ctx context.Context, id int64) error {
	chatsIdKey := fmt.Sprintf("%s%v", cacheChatsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, chatsIdKey)
	return err
}

func (m *defaultChatsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheChatsIdPrefix, primary)
}

func (m *defaultChatsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
