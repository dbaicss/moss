# go module golang包管理解决之道

## 常用命令

1. 初始化
> go mod init [module name]

初始完成后会在目录下生成一个go.mod文件
```
go mod init moss-service
```

2.  自动更新依赖
> go mod tidy

使用go build，go test以及go list时，go会自动得更新go.mod文件。

也可以使用此命令手动更新依赖。

3.  修改版本号
> go mod edit -require

go mod edit -require github.com/bndr/gojenkins@master
修改后直接run，会自动下载对应版本