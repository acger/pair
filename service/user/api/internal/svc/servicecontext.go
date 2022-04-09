package svc

import (
	"pair/service/user/api/internal/config"
	"pair/service/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserRPC     user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserRPC:     user.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}
