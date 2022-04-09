package svc

import (
	"pair/service/pair/api/internal/config"
	"pair/service/pair/rpc/pair"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	PairRPC pair.Pair
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		PairRPC: pair.NewPair(zrpc.MustNewClient(c.PairRPC)),
	}
}
