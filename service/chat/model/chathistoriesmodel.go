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
	chatHistoriesFieldNames          = builder.RawFieldNames(&ChatHistories{})
	chatHistoriesRows                = strings.Join(chatHistoriesFieldNames, ",")
	chatHistoriesRowsExpectAutoSet   = strings.Join(stringx.Remove(chatHistoriesFieldNames, "`create_time`", "`update_time`"), ",")
	chatHistoriesRowsWithPlaceHolder = strings.Join(stringx.Remove(chatHistoriesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheChatHistoriesIdPrefix = "cache:chatHistories:id:"
)

type (
	ChatHistoriesModel interface {
		Insert(ctx context.Context, data *ChatHistories) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ChatHistories, error)
		Update(ctx context.Context, data *ChatHistories) error
		Delete(ctx context.Context, id int64) error
	}

	defaultChatHistoriesModel struct {
		sqlc.CachedConn
		table string
	}

	ChatHistories struct {
		Id         int64     `db:"id" json:"id"`
		CreateTime time.Time `db:"create_time" json:"create_time" gorm:"-"`
		UpdateTime time.Time `db:"update_time" json:"update_time" gorm:"-"`
		Uid        int64     `db:"uid" json:"uid"`
		ToUid      int64     `db:"to_uid" json:"to_uid"`
	}
)

func NewChatHistoriesModel(conn sqlx.SqlConn, c cache.CacheConf) ChatHistoriesModel {
	return &defaultChatHistoriesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`chat_histories`",
	}
}

func (m *defaultChatHistoriesModel) Insert(ctx context.Context, data *ChatHistories) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, chatHistoriesRowsExpectAutoSet)
	ret, err := m.ExecNoCacheCtx(ctx, query, data.Id, data.Uid, data.ToUid)

	return ret, err
}

func (m *defaultChatHistoriesModel) FindOne(ctx context.Context, id int64) (*ChatHistories, error) {
	chatHistoriesIdKey := fmt.Sprintf("%s%v", cacheChatHistoriesIdPrefix, id)
	var resp ChatHistories
	err := m.QueryRowCtx(ctx, &resp, chatHistoriesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatHistoriesRows, m.table)
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

func (m *defaultChatHistoriesModel) Update(ctx context.Context, data *ChatHistories) error {
	chatHistoriesIdKey := fmt.Sprintf("%s%v", cacheChatHistoriesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, chatHistoriesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.ToUid, data.Id)
	}, chatHistoriesIdKey)
	return err
}

func (m *defaultChatHistoriesModel) Delete(ctx context.Context, id int64) error {
	chatHistoriesIdKey := fmt.Sprintf("%s%v", cacheChatHistoriesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, chatHistoriesIdKey)
	return err
}

func (m *defaultChatHistoriesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheChatHistoriesIdPrefix, primary)
}

func (m *defaultChatHistoriesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatHistoriesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
