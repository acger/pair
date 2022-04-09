package svc

import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"pair/common/database"
	"pair/service/pair/model"
	"pair/service/pair/rpc/internal/config"
	"pair/service/user/rpc/user"
)

type ServiceContext struct {
	Config       config.Config
	DB           *gorm.DB
	Cache        *redis.Redis
	UserRPC      user.User
	ES           *es.Client
	ElementModel model.ElementsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		DB:           database.NewMysql(c.DB.DataSource),
		Cache:        redis.New(c.Cache[0].Host, redis.WithPass(c.Cache[0].Pass)),
		UserRPC:      user.NewUser(zrpc.MustNewClient(c.UserRPC)),
		ES:           database.NewElasticsearch(&c.Elasticsearch),
		ElementModel: model.NewElementsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
