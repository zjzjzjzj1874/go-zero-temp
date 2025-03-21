package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/config"
	"github/zjzjzjzj1874/go-zero-temp/pay/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	Log    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Log:    middleware.NewLogMiddleware().Handle,
	}
}
