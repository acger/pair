package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"pair/common/aerror"
	"pair/service/chat/api/internal/logic"

	"pair/service/chat/api/internal/config"
	"pair/service/chat/api/internal/handler"
	"pair/service/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chat-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	manager := logic.NewManager(ctx)
	go manager.Run()

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/chat/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			logic.ServeWs(manager, w, r)
		},
	})

	httpx.SetErrorHandler(aerror.ErrorHandler)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
