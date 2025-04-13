package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MySQL struct {
		DataSource string
	}

	RedisCache cache.CacheConf

	JWT struct {
		Secret string
		Expire    int64
	}
}
