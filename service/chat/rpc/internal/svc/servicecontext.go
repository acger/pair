package svc

import (
	"pair/common/database"
	"pair/service/chat/rpc/internal/config"
	"pair/service/user/rpc/user"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config       config.Config
	DB           *gorm.DB
	UserRPC      user.User
	ChatKqPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		DB:           database.NewMysql(c.DB.DataSource),
		UserRPC:      user.NewUser(zrpc.MustNewClient(c.UserRPC)),
		ChatKqPusher: kq.NewPusher(c.ChatKq.Brokers, c.ChatKq.Topic),
	}
}
