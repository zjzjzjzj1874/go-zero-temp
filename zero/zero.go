package main

import (
	_ "embed"
	"flag"
	"fmt"
	"net/http"

	"github/zjzjzjzj1874/go-zero-temp/constants/errors"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/config"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/handler"
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/zero.yaml", "the config file")

func main() {
	flag.Parse()
	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.Swagger = swagger

	server := rest.MustNewServer(c.RestConf, Opt)
	defer server.Stop()

	httpx.SetErrorHandler(errors.ErrorRestfulHandler) // 自定义错误返回,restful类型

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.PrintRoutes() // prints routes

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//go:embed zero.json
var swagger []byte

// Opt 设置自定义跨域请求,或者返回的header
var Opt = rest.WithCustomCors(func(header http.Header) {
	header.Set("Access-Control-Allow-Headers", "*")
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Expose-Headers", "*")
	header.Set("Access-Control-Allow-Methods", "*")
	header.Set("Access-Control-Allow-Credentials", "false")
}, nil, "*")
