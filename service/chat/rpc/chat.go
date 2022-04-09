package main

import (
	"flag"
	"fmt"
	"pair/common/queue"
	"pair/service/chat/mq"

	"pair/service/chat/rpc/chat/pb"
	"pair/service/chat/rpc/internal/config"
	"pair/service/chat/rpc/internal/server"
	"pair/service/chat/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewChatServer(ctx)

	queue.ListenKq(mq.GetKqueueList(ctx.DB), c.ChatKq)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterChatServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("starting rpc at at %s...\n", c.ListenOn)
	s.Start()
}
