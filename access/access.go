package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github/zjzjzjzj1874/go-zero-temp/access/internal/config"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/handler"
	"github/zjzjzjzj1874/go-zero-temp/access/internal/svc"
	"github/zjzjzjzj1874/go-zero-temp/common/cors"
	"github/zjzjzjzj1874/go-zero-temp/constants/errors"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/access.yaml", "the config file")

func main() {
	flag.Parse()
	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.Swagger = swagger

	server := rest.MustNewServer(c.RestConf, cors.Opt)
	defer server.Stop()

	httpx.SetErrorHandler(errors.ErrorRestfulHandler) // 自定义错误返回,restful类型

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.PrintRoutes() // prints routes

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//go:embed access.json
var swagger []byte
