package svc

import (
	"kkim/app/user/model"
	"kkim/app/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserModels model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config: c,
		UserModels: model.NewUsersModel(sqlConn, c.RedisCache),
	}
}
