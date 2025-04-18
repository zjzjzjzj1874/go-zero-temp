- 新增接口

```shell
goctl api go -api user.api  -style go_zero  -dir  .
`````

- 生成swagger文件
```shell
goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api
```
- api生成swagger文件
```shell
goctl api swagger --api user.api --dir ./
```

- dockerfile生成
```shell
goctl docker -go user.go
```