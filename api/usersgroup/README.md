# 测试运行


# 测试运行

运行 srv usersgroup
```
# run srv usersgroup
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/srv/usersgroup
$ go run main.go init.go --registry=mdns
```

运行 api usersgroup
```bash
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/api/usersgroup
$ go run main.go init.go --registry=mdns
```

运行 micro api
```bash
$ micro --registry=mdns api --address=:8080  --handler=api 
```


# 测试方法
```
# Post 方法，添加
$ curl -XPOST -H 'Content-Type: application/x-www-form-urlencoded' \
      -d 'Name=test Columns&URL=This is a URL&ParentID=123abc' \
      --url http://localhost:8080/mshk/api/v1/usersgroup/UsersGroup/add
{"id":"go.micro.api.mshk.api.v1","code":500,"detail":"ParentID 的格式不正确:unable to parse as int: string","status":"Internal Server Error"}

# Get 方法，获取列表
$ curl "http://localhost:8080/mshk/api/v1/usersgroup/UsersGroup/getList?PageSize=2&CurrentPageIndex=1&OrderBy=-id"
```
