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
	elementsFieldNames          = builder.RawFieldNames(&Elements{})
	elementsRows                = strings.Join(elementsFieldNames, ",")
	elementsRowsExpectAutoSet   = strings.Join(stringx.Remove(elementsFieldNames, "`create_time`", "`update_time`"), ",")
	elementsRowsWithPlaceHolder = strings.Join(stringx.Remove(elementsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheElementsIdPrefix  = "cache:elements:id:"
	cacheElementsUidPrefix = "cache:elements:uid:"
)

type (
	ElementsModel interface {
		Insert(ctx context.Context, data *Elements) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Elements, error)
		FindOneByUid(ctx context.Context, uid int64) (*Elements, error)
		Update(ctx context.Context, data *Elements) error
		Delete(ctx context.Context, id int64) error
	}

	defaultElementsModel struct {
		sqlc.CachedConn
		table string
	}

	Elements struct {
		Id         int64     `db:"id" json:"id"`
		CreateTime time.Time `db:"create_time" json:"create_time" gorm:"-"`
		UpdateTime time.Time `db:"update_time" json:"update_time" gorm:"-"`
		Uid        int64     `db:"uid" json:"uid"`
		Skill      string    `db:"skill" json:"skill"`
		SkillNeed  string    `db:"skill_need" json:"skill_need"`
		Star       int64     `db:"star" json:"star"`
		Boost      int64     `db:"boost" json:"boost"`
	}
)

func NewElementsModel(conn sqlx.SqlConn, c cache.CacheConf) ElementsModel {
	return &defaultElementsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`elements`",
	}
}

func (m *defaultElementsModel) Insert(ctx context.Context, data *Elements) (sql.Result, error) {
	elementsIdKey := fmt.Sprintf("%s%v", cacheElementsIdPrefix, data.Id)
	elementsUidKey := fmt.Sprintf("%s%v", cacheElementsUidPrefix, data.Uid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, elementsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Uid, data.Skill, data.SkillNeed, data.Star, data.Boost)
	}, elementsIdKey, elementsUidKey)
	return ret, err
}

func (m *defaultElementsModel) FindOne(ctx context.Context, id int64) (*Elements, error) {
	elementsIdKey := fmt.Sprintf("%s%v", cacheElementsIdPrefix, id)
	var resp Elements
	err := m.QueryRowCtx(ctx, &resp, elementsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", elementsRows, m.table)
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

func (m *defaultElementsModel) FindOneByUid(ctx context.Context, uid int64) (*Elements, error) {
	elementsUidKey := fmt.Sprintf("%s%v", cacheElementsUidPrefix, uid)
	var resp Elements
	err := m.QueryRowIndexCtx(ctx, &resp, elementsUidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", elementsRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, uid); err != nil {
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

func (m *defaultElementsModel) Update(ctx context.Context, data *Elements) error {
	elementsIdKey := fmt.Sprintf("%s%v", cacheElementsIdPrefix, data.Id)
	elementsUidKey := fmt.Sprintf("%s%v", cacheElementsUidPrefix, data.Uid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, elementsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Uid, data.Skill, data.SkillNeed, data.Star, data.Boost, data.Id)
	}, elementsIdKey, elementsUidKey)
	return err
}

func (m *defaultElementsModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	elementsUidKey := fmt.Sprintf("%s%v", cacheElementsUidPrefix, data.Uid)
	elementsIdKey := fmt.Sprintf("%s%v", cacheElementsIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, elementsIdKey, elementsUidKey)
	return err
}

func (m *defaultElementsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheElementsIdPrefix, primary)
}

func (m *defaultElementsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", elementsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
