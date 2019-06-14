
# 测试运行

运行 srvusers 和 srvhistoryuserlogin
```
# run srvusers
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/srv/srvusers
$ go run main.go init.go --registry=mdns

# run srvhistoryuserlogin
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/srv/srvhistoryuserlogin/
$ go run main.go init.go --registry=mdns
```

运行 apiusers
```bash
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/api/apiusers
$ go run main.go init.go --registry=mdns
```

运行 micro api
```bash
$ micro --registry=mdns api --address=:8080  --handler=api 
```


# 测试方法
```
# Post 方法，添加用户
$ curl -XPOST -H 'Content-Type: application/x-www-form-urlencoded' \
      -d 'UserName=This is a Name&PassWord=123password' \
      --url http://localhost:8080/mshk/v1/Users/add
{"id":"go.micro.api.mshk.api.v1","code":500,"detail":"Email 不能为空","status":"Internal Server Error"}

# Get 方法，获取用户列表
$ curl "http://localhost:8080/mshk/v1/users/Users/getList?PageSize=2&CurrentPageIndex=1&OrderBy=-id"
```
