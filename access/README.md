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


## 需求
现在我要设计一个权限角色的模块，其作用是控制用户访问api的权限；
所以api的接口就是一个一个的权限；然后角色是一个一个权限的集合；
角色和权限是一对多的关系；但是对于用户来讲，只需要给用户分配对应的角色，用户和角色是一对多的关系；
按照这个设计，给我设计一个PlantUML的设计关系图。