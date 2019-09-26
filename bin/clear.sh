#!/bin/bash

dir=`pwd`
buildprefix="mshk"
build() {
	for d in $(ls ./$1); do
		echo "delete $1/$d"
		# pushd命令常用于将目录加入到栈中，加入记录到目录栈顶部，并切换到该目录；若pushd命令不加任何参数，则会将位于记录栈最上面的2个目录对换位置
		pushd $dir/$1/$d >/dev/null
		# -s 忽略符号表和调试信息，-w忽略DWARF符号表，通过这两个参数，可以进一步减少编译的程序的尺寸，更多的参数可以参考go link,
		#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w'
		rm -rf $buildprefix.*
		# opd用于删除目录栈中的记录；如果popd命令不加任何参数，则会先删除目录栈最上面的记录，然后切换到删除过后的目录栈中的最上面的目录
		popd >/dev/null
	done
}

build api
build srv

# 单独编译 strategy
echo "delete api/strategy/strategy.tar.gz"
rm -rf api/strategy/strategy.tar.gz

# 单独编译 web
echo "delete web/$buildprefix.web"
cd web
rm -rf $buildprefix.web
rm -rf web
cd ..


# 批量删除镜像
docker rmi --force $(docker images | grep microservice-userpowermanager | awk '{print $3}')
docker rmi --force $(docker images | grep none | awk '{print $3}')


# docker stop $(docker ps -a | grep Exited | awk '{print $1}')