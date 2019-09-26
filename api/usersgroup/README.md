# 测试运行

## 服务发现

所有服务都需要服务发现。默认为多播DNS，简单不需要任何配置。

如果您需要多主机或更具弹性可以使用 Consul。

```
# install
$ brew install consul

# run http://localhost:8500
$ consul agent -dev
```

## 运行 srv usersgroup
```
# run srv usersgroup
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/srv/usersgroup
$ go run main.go init.go --registry=consul
```

## 运行 api usersgroup
```bash
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/api/usersgroup
$ go run main.go init.go --registry=consul
```

## 运行 micro api
```bash
$ micro --registry=consul api --address=:8080  --handler=api 
```


# 测试方法
```http://localhost:8080/mshk/v1/columns/Columns/add
# Post 方法，添加
$ curl -XPOST -H 'Content-Type: application/x-www-form-urlencoded' \
      -d 'Name=test Columns&URL=This is a URL&ParentID=123abc' \
      --url http://localhost:8080/mshk/v1/usersgroup/UsersGroup/add
{"id":"go.micro.api.mshk.v1.usersgroup","code":500,"detail":"Sorts 不能为空","status":"Internal Server Error"}

# Get 方法，获取列表
$ curl "http://localhost:8080/mshk/v1/usersgroup/UsersGroup/getList?PageSize=2&CurrentPageIndex=1&OrderBy=id%20desc"
```
