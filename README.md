# MicroService-UserPowerManager

# 使用 go-micro 构建的一个项目
> 实例中默认您已经安装好 go-micro 环境，关于 go-micro 的环境安装请移步 [这里](https://micro.mu/docs/go-micro.html)

[升级记录](README_Feature.MD)

主要使用以下第三方工具：
* `jaeger` 来收集分布式服务追踪
* `beego` 做 web 可访问的服务
* `XORM` 做为数据库 ORM 处理
* `JWT` 实现了跨域登录
* `Casbin` 结合实际应用做权限管理
* `Shazam` shazam ([ʃə'zæm], 沙赞)是一款兼容MySQL协议的数据库中间件, 支持分库分表, 其前身是Gaea.
* `logrus` 做日志记录，未来也可以通过 hook 方式记录到其他第三方组件，在输出控制台的同时也会在 `./logs` 目录下，根据日期每天创建一个日志文件。

## 目录说明
```
|____bin                                # 存放 shell 脚本，目前主要用于批量 build
|____swagger                            # swagger 格式的 go-micro 对外提供 API 文档，可以通过 Docker 直接运行
|____Makefile                           # 编译规则文件
|____web                                # Beego 运行的 Web 程序，用于演示如何通过 Web 调用 go-micro 提供的 API
| |____conf                             # Web 程序的配置文件目录
| |____models                           # Web 程序使用的 models 定义
| |____routers                      
| | |____router.go                      # 路由，在管理后台 `/admin` 的路由加入 Filter 做登录的 token 判断
...
| |____static                           # Beego Web 模板用到的静态文件目录
| |____controllers                      # Beego controllers 目录
...
| |____views                            # Beege Web 模板目录
... 
|____utils                              # 工具库
| |____log4                             # 对日志组件封装
...
| |____inner                            # 内部使用的公共变量和方法
...
| |____jaeger                           # jaeger组件封装
...
| |____TConfig.go                       # 操作配置文件
| |____request                          # 统一封装 http 请求
...
| |____orm                              # ORM封装
...
| |____encrypt                          # 加密、解密相关    
...
|____common                             # 公共函数库
...
|____docker-compose.dbproxy.env         # docker-compose 文件用到的环境变量，MySql代理DbProxy用的环境变量
|____docker-compose.jaeger.env          # docker-compose 文件用到的环境变量，Jaeger的环境变量
|____docker-compose.srv.common.env      # docker-compose 文件用到的环境变量，微服务 Service 通用环境变量
|____docker-compose.yml                 # docker-compose 编排文件
|____data                               # 用到的数据库相关文件夹
| |____insert.sql
| |____UserPowerManager.mwb             # 数据库建模文件
| |____UserPowerManager.sql
| |____UserPowerManager.mwb.bak
|____main.go
|____api                                # 微服务 API 目录
| |____role                             # 用户和角色（组）的权限管理
| |____users                            # 用户增、删、改、查, 登录、验证Token相关操作的API
| |____columns                          # 栏目（导航菜单）增、删、改、查
| |____usersgroup                       # 用户组增、删、改、查
| |____dbproxy                          # MySql 数据库代理，基于 Docker 的封装
|____srv                                # 微服务 Service  目录
| |____role                             # 用户和角色（组）的权限管理 Service，包括和数据库交互
| |____historyuserlogin                 # 用户登录历史记录的 Service
| |____users                            # 用户增、删、改、查相关操作的 Service，包括和数据库交互
| |____columns                          # 栏目（导航菜单）增、删、改、查 Service，包括和数据库交互
| |____usersgroup                       # 用户组增、删、改、查 Service，包括和数据库交互
```

## 数据库代理配置介绍

### MySQL 数据库配置

`data/UserPowerManager.mwb` 文件，是使用 [MySQL Workbench](https://www.mysql.com/cn/products/workbench/) 创建的数据库建模。

首先创建数据库：`UserPowerManager`，在库中运行 `data/UserPowerManager.sql` 文件来创建初始需要的表

然后导入 `data/insert.sql` ,会在 `users_0001` 表，创建用户 `admin` 密码是 `admin`，程序启动以后，可以在管理后台登录。

### MySQL 数据库代理配置说明

使用 `MySql` 代理 `Shazam` 是考虑到未来数据量大，可以做分库分表优化。示例中对用户表做了分表，分别是 `users_0000`、`users_0001`，使用数据库代理后可以通过 `select count(*) from users where id>1` 来获取需要的数据结果集。
> 示例中使用的 `MySql` 和 代理 `Shazam` 配置信息，如需要修改，请修改 `docker-compose.dbproxy.env` 和 `docker-compose.srv.common.env` 文件

启动 `MySql` 和数据库代理后，使用以下命令可以通过代理连接 `MySql`:
```
$ mysql -h10.0.0.10 -P 30080 -uUserPowerManager -pUserPowerManager123456
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 10001
Server version: 5.6.20-gaea MySQL Community Server (GPL)

Copyright (c) 2000, 2016, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> use UserPowerManager;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> select count(*) from users;
+----------+
| COUNT(1) |
+----------+
|        2 |
+----------+
1 row in set (0.08 sec)

mysql>
```


## go-micro 环境安装


### 编译安装 Protoc 3.9.1

本示例中，使用 Protocol Buffers 3 做为微服务的通讯。

```
$ wget https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protobuf-all-3.9.1.tar.gz
$ unzip protobuf-all-3.8.0.zip
$ cd protobuf-3.8.0

# Mac 编译安装
$ brew install automake libtool autoconf
$ ./autogen.sh
$ ./configure CPPFLAGS=-DGTEST_USE_OWN_TR1_TUPLE=1
$ make && make install

# 查看版本
$ protoc --version
```


### 安装 Toolkit

```
$ go get github.com/micro/micro
```

## 快速运行

```
# （这个步骤可忽略）验证代码风格,`Golang` 的开发团队制定了统一的官方代码风格，并且推出了 `gofmt` 工具（`gofmt` 或 `go fmt`）来帮助开发者格式化他们的代码到统一的风格。
$ make fmt
```
![mshk.top](https://img.mshk.top/MicroService-UserPowerManager1.gif-2)

```
# 下载项目依赖包,需要 >golang 1.12+
$ make update_deps
```
![mshk.top](https://img.mshk.top/MicroService-UserPowerManager2.gif-2)

```
# （这个步骤可忽略）使用 golangci-lint 检查代码，`GolangCI-Lint` 是一个 `lint` 聚合器，它的速度很快，平均速度是 `gometalinter` 的5倍。它易于集成和使用，具有良好的输出并且具有最小数量的误报。
$ curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin vX.Y.Z
$ make linter
```
![mshk.top](https://img.mshk.top/MicroService-UserPowerManager3.gif-2)


```
# 生成 protobuf 文件
# mac 下安装有些库依赖的 bazaar
$ brew install bazaar
$ make proto
```
![mshk.top](https://img.mshk.top/MicroService-UserPowerManager4.gif-2)

```
# 编译程序
$ make build
```

```
# 启动所有的微服务
# 启动前，请先编辑 `.env` 文件，**修改数据库的配置**，`docker-compose` 运行时会优先读取 `.env` 中的变量进行替换
$ make run
```
![mshk.top](https://img.mshk.top/MicroService-UserPowerManager5.gif-2)

启动所有微服务后，浏览 http://localhost:8500 ，能够看到注册到 `consul` 的所有微服务。

![mshk.top](https://img.mshk.top/MicroService-UserPowerManager7.png-2)

```
## 关闭所有微服务
$ make down


## 清除编译文件
$ make clear
```


## 浏览微服务


浏览 http://localhost:18080/v1/admin 用户列表页面，默认没有登录，会 302 跳转到登录页面，使用用户名 `admin` 密码 `admin` 登录。

登录成功以后，可以看到下图中的页面

![mshk.top](https://img.mshk.top/MicroService-UserPowerManager6.gif-2)

同时浏览 http://localhost:16686 可以看到登录过程中，记录到 `jaeger` 的微服务之间调用关系，如下图：

![mshk.top](https://img.mshk.top/MicroService-UserPowerManager8.png-2)

## swagger 文档查看

官方文档：https://swagger.io/docs/specification/2-0/what-is-swagger/

编译：
```
$ docker-compose build swagger
```

运行以下命令，然后浏览 http://localhost:18081 查看 API 文档
```bash
$ docker-compose up swagger
```
> 可以看接口，不能进行测试，因为go-micro默认是不支持跨域的，如果需要跨域，需要重新编译micro，参考文章：https://github.com/micro/go-plugins/tree/master/micro/cors


运行 `swagger` 编辑器，浏览 http://localhost，可以看到效果
```
$ docker run -it --rm -p 80:8080 --name swagger-editor swaggerapi/swagger-editor
```
