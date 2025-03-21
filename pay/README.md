- pay服务中新增接口

```shell
goctl api go -api pay.api  -style go_zero  -dir  .
```

- 生成swagger文件
```shell
goctl api plugin -plugin goctl-swagger="swagger -filename pay.json" -api pay.api
```

- dockerfile生成
```shell
goctl docker -go pay.go
```