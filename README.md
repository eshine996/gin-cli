# gin-cli

gin工程化助手

### 安装

```shell
go install github.com/codeHauler-1/gin-cli@latest
```

### 使用
* 初始化项目
```shell
gin-cli init {project name} 
```
* 生成dao层代码</br>
备注：执行命令前前先在config.yaml中配置mysql的数据库连接串
```shell
gin-cli gen dao
```

### 启动项目

```shell
cd {project name}
go mod tidy
go run main.go
```