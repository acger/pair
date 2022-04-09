package svc

import (
	"pair/service/chat/api/internal/config"
	"pair/service/chat/rpc/chat"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	ChatRPC  chat.Chat
	KqPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		ChatRPC:  chat.NewChat(zrpc.MustNewClient(c.ChatRPC)),
		KqPusher: kq.NewPusher(c.ChatKq.Brokers, c.ChatKq.Topic),
	}
}
