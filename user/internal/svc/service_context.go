package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/config"
	"github/zjzjzjzj1874/go-zero-temp/user/internal/middleware"
)

type ServiceContext struct {
	Config     config.Config
	Recover    rest.Middleware
	Log        rest.Middleware
	Permission rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Recover:    middleware.NewRecoverMiddleware().Handle,
		Log:        middleware.NewLogMiddleware().Handle,
		Permission: middleware.NewPermissionMiddleware().Handle,
	}
}
