module github.com/idoall/MicroService-UserPowerManager

go 1.12

require (
	contrib.go.opencensus.io/exporter/ocagent v0.4.2 // indirect
	github.com/Azure/azure-sdk-for-go v24.1.0+incompatible // indirect
	github.com/Azure/go-autorest v11.3.2+incompatible // indirect
	github.com/SermoDigital/jose v0.9.2-0.20161205224733-f6df55f235c2 // indirect
	github.com/astaxie/beego v1.11.1
	github.com/bwmarrin/snowflake v0.0.0-20180412010544-68117e6bbede
	github.com/casbin/beego-orm-adapter v0.0.0-20180421160615-ba104911ae5b
	github.com/casbin/casbin v1.7.0
	github.com/census-instrumentation/opencensus-proto v0.1.0 // indirect
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.1
	github.com/idoall/TokenExchangeCommon v0.0.0-20190406161816-b77f54f3a4b1
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/lestrrat/go-envload v0.0.0-20180220120943-6ed08b54a570 // indirect
	github.com/lestrrat/go-file-rotatelogs v0.0.0-20180223000712-d3151e2a480f
	github.com/lestrrat/go-strftime v0.0.0-20180220042222-ba3bf9c1d042 // indirect
	github.com/micro/go-api v0.7.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.1.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.3.0
	github.com/tebeka/strftime v0.0.0-20140926081919-3f9c7761e312 // indirect
	github.com/uber-go/atomic v1.3.2 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	sigs.k8s.io/structured-merge-diff v0.0.0-20190416230737-b2ed7e1d99f6 // indirect
)

replace go.etcd.io/etcd v3.3.12+incompatible => github.com/etcd-io/etcd v3.3.12+incompatible
