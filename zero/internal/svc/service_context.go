package svc

import (
	"github/zjzjzjzj1874/go-zero-temp/zero/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
