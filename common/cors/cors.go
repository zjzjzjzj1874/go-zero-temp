package cors

import (
	"net/http"
	
	"github.com/zeromicro/go-zero/rest"
)

// Opt 设置自定义跨域请求,或者返回的header
var Opt = rest.WithCustomCors(func(header http.Header) {
	header.Set("Access-Control-Allow-Headers", "*")
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Expose-Headers", "*")
	header.Set("Access-Control-Allow-Methods", "*")
	header.Set("Access-Control-Allow-Credentials", "false")
}, nil, "*")
