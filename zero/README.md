- api中新增接口

```shell
goctl api go -api zero.api  -style go_zero  -dir  .
`````

- api生成swagger文件
```shell
goctl api plugin -plugin goctl-swagger="swagger -filename zero.json" -api zero.api
```
- api生成swagger文件
```shell
goctl api swagger --api zero.api --dir ./
```

- dockerfile生成
```shell
goctl docker -go zero.go
```