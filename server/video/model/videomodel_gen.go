// Code generated by goctl. DO NOT EDIT.

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
	videoFieldNames          = builder.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheVideoIdPrefix = "cache:video:id:"
)

type (
	videoModel interface {
		Insert(ctx context.Context, data *Video) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Video, error)
		FindOneById(ctx context.Context, id int64) (*Video, error)
		Update(ctx context.Context, data *Video) error
		Delete(ctx context.Context, id int64) error
	}

	defaultVideoModel struct {
		sqlc.CachedConn
		table string
	}

	Video struct {
		Id         int64     `db:"id" json:"id"`
		UserId     int64     `db:"user_id" json:"user_id"`
		Title      string    `db:"title" json:"title"`
		Partition  int64     `db:"partition" json:"partition"`
		PlayUrl    string    `db:"play_url" json:"play_url"`
		CoverUrl   string    `db:"cover_url" json:"cover_url"`
		CreateTime time.Time `db:"create_time" json:"create_time"`
		UpdateTime time.Time `db:"update_time" json:"update_time"`
		DeleteTime time.Time `db:"delete_time" json:"delete_time"`
	}
)

func newVideoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultVideoModel {
	return &defaultVideoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`video`",
	}
}

func (m *defaultVideoModel) withSession(session sqlx.Session) *defaultVideoModel {
	return &defaultVideoModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`video`",
	}
}

func (m *defaultVideoModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, data.Id)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, videoIdKey)
	return err
}

func (m *defaultVideoModel) FindOne(ctx context.Context, id int64) (*Video, error) {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, id)
	var resp Video
	err := m.QueryRowCtx(ctx, &resp, videoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
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

func (m *defaultVideoModel) FindOneById(ctx context.Context, id int64) (*Video, error) {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, id)
	var resp Video
	err := m.QueryRowIndexCtx(ctx, &resp, videoIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, id); err != nil {
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

func (m *defaultVideoModel) Insert(ctx context.Context, data *Video) (sql.Result, error) {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.Title, data.Partition, data.PlayUrl, data.CoverUrl, data.DeleteTime)
	}, videoIdKey)
	return ret, err
}

func (m *defaultVideoModel) Update(ctx context.Context, newData *Video) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.Title, newData.Partition, newData.PlayUrl, newData.CoverUrl, newData.DeleteTime, newData.Id)
	}, videoIdKey)
	return err
}

func (m *defaultVideoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheVideoIdPrefix, primary)
}

func (m *defaultVideoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideoModel) tableName() string {
	return m.table
}
