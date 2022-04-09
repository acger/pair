package config

import (
	"pair/common/database"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB      database.MysqlConf
	Cache   cache.CacheConf
	UserRPC zrpc.RpcClientConf
	ChatKq  kq.KqConf
}
