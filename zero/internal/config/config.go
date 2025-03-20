package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Swagger     []byte `json:",optional"`
	SwaggerPath string `json:",default=/app/swagger.json"`
}
