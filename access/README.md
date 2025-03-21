- 新增接口

```shell
goctl api go -api access.api  -style go_zero  -dir  .
`````

- 生成swagger文件
```shell
goctl api plugin -plugin goctl-swagger="swagger -filename access.json" -api access.api
```

- dockerfile生成
```shell
goctl docker -go access.go
```