
# 测试运行

运行 srvusers 和 srvhistoryuserlogin
```
# run srv users
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/srv/users
$ go run main.go init.go --registry=mdns

# run srv historyuserlogin
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/srv/historyuserlogin/
$ go run main.go init.go --registry=mdns
```

运行 api users
```bash
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/api/users
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
      --url http://localhost:8080/mshk/v1/users/Users/add
{"id":"go.micro.api.mshk.api.v1","code":500,"detail":"Email 不能为空","status":"Internal Server Error"}

# Get 方法，获取用户列表
$ curl "http://localhost:8080/mshk/v1/users/Users/getList?PageSize=2&CurrentPageIndex=1&OrderBy=-id"
{"rows":[{"ID":1,"UserName":"admin","RealyName":"admin","Password":"93a57b286d7f77fdce1c8e17f5c2dfb6459c739b058c85b168cdd1df599e1f35","AuthKey":"1118447772383584256","Email":"admin@mshk.top","Note":"adminnote","CreateTime":1555048377,"LastUpdateTime":1555051238}],"total":1}

# Get 方法，获取单个用户列表
$ curl "http://localhost:8080/mshk/v1/users/Users/getUser?ID=1"
```
