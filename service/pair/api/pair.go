package main

import (
	"pair/common/aerror"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"pair/service/pair/api/internal/config"
	"pair/service/pair/api/internal/handler"
	"pair/service/pair/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/pair-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	httpx.SetErrorHandler(aerror.ErrorHandler)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
