
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


## 运行 srv columns
```
# run srv columns
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/srv/columns
$ go run main.go init.go --registry=consul
```

## 运行 api columns
```bash
# run api columns
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/api/columns
$ go run main.go init.go --registry=consul
```

## 运行 micro api
```bash
$ micro --registry=consul api --address=:8080  --handler=api 
```


# 测试方法
```
# Post 方法，添加栏目
$ curl -XPOST -H 'Content-Type: application/x-www-form-urlencoded' \
      -d 'Name=test Columns&URL=This is a URL&ParentID=123abc' \
      --url http://localhost:8080/mshk/v1/columns/Columns/add
{"id":"go.micro.api.mshk.api.v1","code":500,"detail":"ParentID 的格式不正确:unable to parse as int: string","status":"Internal Server Error"}

# Get 方法，获取列表
$ curl "http://localhost:8080/mshk/v1/columns/Columns/getList?PageSize=2&CurrentPageIndex=1&OrderBy=sorts%20desc"
```
