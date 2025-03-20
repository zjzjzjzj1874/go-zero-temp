- api中新增接口

```shell
goctl api go -api zero.api  -style go_zero  -dir  .
`````

- api生成swagger文件
```shell
goctl api plugin -plugin goctl-swagger="swagger -filename zero.json" -api zero.api
```

- dockerfile生成
```shell
goctl docker -go zero.go
```