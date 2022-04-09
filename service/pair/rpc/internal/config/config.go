package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"pair/common/database"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache         cache.CacheConf
	UserRPC       zrpc.RpcClientConf
	Elasticsearch database.ElasticsearchConf
}
