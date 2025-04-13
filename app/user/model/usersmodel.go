package model

import (
	"context"
	"errors"
	"fmt"
	"strings"


	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		FindOneByPhone(ctx context.Context, phone string) (*Users, error)
		ListByIds(ctx context.Context, ids []string) ([]*Users, error)
		ListByName(ctx context.Context, name string) ([]*Users, error)	
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}

func (m *customUsersModel) FindOneByPhone(ctx context.Context, phone string) (*Users, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, phone)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, usersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, phone)
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

func (m *customUsersModel) ListByIds(ctx context.Context, ids []string) ([]*Users, error) {
	if(len(ids) == 0){
		return nil, errors.New("Invalid name")
	}
	var resp []*Users
	quotedIds := make([]string, len(ids))
	for i, id := range ids {
    	quotedIds[i] = "'" + id + "'"
	}
	query := fmt.Sprintf("select %s from %s where `id` in (%s)", usersRows, m.table, strings.Join(quotedIds, ","))
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil :
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customUsersModel) ListByName(ctx context.Context, name string) ([]*Users, error) {
	if(len(name) == 0){
		return nil, errors.New("Invalid name")
	}
	var resp []*Users
	query := fmt.Sprintf("select %s from %s where nickname like ?", usersRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, "%"+name+"%")
	switch err {
	case nil :
		return resp, nil
	default:
		return nil, err
	}
}