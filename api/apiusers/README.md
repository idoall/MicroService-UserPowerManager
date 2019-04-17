# Template Service

This is the Template service

Generated with

```
micro new github.com/idoall/MicroService-UserPowerManager/api/CUsers --namespace=go.micro --alias=template --type=api
```

## Getting Started

- [Template Service](#template-service)
  - [Getting Started](#getting-started)
  - [Configuration](#configuration)
  - [Dependencies](#dependencies)
  - [Usage](#usage)
  - [测试运行](#)
- [swagger](#swagger)

## Configuration

- FQDN: go.micro.api.template
- Type: api
- Alias: template

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./template-api
```

Build a docker image
```
make docker
```

## 测试运行

运行api
```bash
go run main.go init.go --registry=mdns
```

启动micro api
```bash
micro --registry=mdns api --address=:8080  --handler=api 
# micro --registry=mdns api --address=:8080  --handler=api --namespace=mshk.top.api
```


测试方法
```
curl -XPOST -H 'Content-Type: application/json' -d '{
      "service": "go.micro.api.mshk.api.v1",
      "method": "ApiUsers.Add",
      "request": {
        "username": "This is a test"
      }
    }' --url http://localhost:8080/rpc
```

```shell
# 添加用户
curl "http://localhost:8080/mshk/api/v1/ApiUsers/add?UserName=John&password=123"
```

```shell
# 获取用户列表
curl "http://localhost:8080/mshk/api/v1/ApiUsers/getList?PageSize=2&CurrentPageIndex=1&OrderBy=-id"
```

curl "http://localhost:8080/mshk/api/v1/cusers/add?name=John"

docker run -p 8080:8080 -e MICRO_REGISTRY=mdns microhq/micro api --handler=rpc --address=:8080 --namespace=shippy

# swagger

官方文档：https://goswagger.io/install.html

中文参考：https://studygolang.com/articles/12354

安装方法：
```
brew tap go-swagger/go-swagger
brew install go-swagger
```

生成 `swagger`
```
swagger generate spec -o ./swagger.json
```


浏览 `swagger`
```
swagger serve -F=swagger swagger.json
```