package svc

import (
	"kkim/app/user/api/internal/config"
	"kkim/app/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// UserRpc is a rpc client for user service
	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// 创建一个 UserRpc 客户端
		User: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
