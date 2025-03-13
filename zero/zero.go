package main

import (
	"flag"
	"fmt"

	"github/zjzjzjzj1874/go-zero-temp/zero/internal/config"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/handler"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/zero.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.PrintRoutes() // prints routes

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
