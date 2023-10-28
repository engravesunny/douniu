package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FavoriteModel = (*customFavoriteModel)(nil)

type (
	// FavoriteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFavoriteModel.
	FavoriteModel interface {
		favoriteModel
		FindPage(ctx context.Context, pageNum int, pageSize int) ([]*Favorite, error)
	}

	customFavoriteModel struct {
		*defaultFavoriteModel
	}
)

func (m *customFavoriteModel) FindPage(ctx context.Context, pageNum int, pageSize int) ([]*Favorite, error) {
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", favoriteRows, m.table)
	var resp []*Favorite
	err := m.conn.QueryRowsCtx(ctx, &resp, query, (pageNum-1)*pageSize, pageSize)

	return resp, err
}

// NewFavoriteModel returns a model for the database table.
func NewFavoriteModel(conn sqlx.SqlConn) FavoriteModel {
	return &customFavoriteModel{
		defaultFavoriteModel: newFavoriteModel(conn),
	}
}
