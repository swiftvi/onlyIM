package main

import (
	"flag"
	"fmt"

	"kkim/app/user/api/internal/config"
	"kkim/app/user/api/internal/handler"
	"kkim/app/user/api/internal/svc"
	"kkim/pkg/xcode"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//在此处设置错误处理函数
	httpx.SetErrorHandlerCtx(xcode.ErrHandler(c.Name))
	httpx.SetOkHandler(xcode.OkHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
