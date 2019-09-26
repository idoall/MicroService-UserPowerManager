
#  package Comms

测试编译

```
$ cd $GOPATH/src/github.com/idoall/MicroService-UserPowerManager/api/dbproxy
$ docker build -t idoall/dbproxy .

$ docker run -it \
--rm \
--name shazam \
--hostname shazam \
-e MYSQL_HOST="192.168.8.41" \
-e MYSQL_PORT=20081 \
-e MYSQL_DBNAME="UserPowerManager" \
-e MYSQL_USERNAME="lion" \
-e MYSQL_PASSWORD="123456" \
-e DBPROXY_USERNAME="UserPowerManager" \
-e DBPROXY_PASSWORD="UserPowerManager123456" \
-p 30080:13306 \
idoall/dbproxy


$ docker run -it \
--rm \
--name shazam \
--hostname shazam \
-e MYSQL_HOST="10.0.0.30" \
-e MYSQL_PORT=20081 \
-e MYSQL_DBNAME="UserPowerManager" \
-e MYSQL_USERNAME="lion" \
-e MYSQL_PASSWORD="123456" \
-e DBPROXY_USERNAME="UserPowerManager" \
-e DBPROXY_PASSWORD="UserPowerManager123456" \
-p 30080:13306 \
idoall/dbproxy
```