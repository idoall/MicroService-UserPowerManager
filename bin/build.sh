#!/bin/bash

BUILD_DATE=$(date '+%Y-%m-%d %H:%M:%S')
PACKAGENAME=$(go list -e)
LDFLAGS="-extldflags -static"
BUILD_USER=$(whoami)
BUILD_HOST=$(hostname -f)

if BUILD_GIT_REVISION=$(git rev-parse HEAD 2> /dev/null); then if ! git diff-index --quiet HEAD; then BUILD_GIT_REVISION=${BUILD_GIT_REVISION}"-dirty"
    fi
else
    BUILD_GIT_REVISION=unknown
fi

# Check for local changes
if git diff-index --quiet HEAD --; then
  tree_status="Clean"
else
  tree_status="Modified"
fi

# XXX This needs to be updated to accomodate tags added after building, rather than prior to builds
RELEASE_TAG=$(git describe --match '[0-9]*\.[0-9]*\.[0-9]*' --exact-match --tags 2> /dev/null || echo "")

# security wanted VERSION='unknown'
VERSION=$(git rev-parse --short HEAD || echo "GitNotFound")
if [[ -n "${RELEASE_TAG}" ]]; then
  VERSION="${RELEASE_TAG}"
elif [[ -n ${MY_VERSION} ]]; then
  VERSION="${MY_VERSION}"
fi

dir=`pwd`
buildprefix="mshk"
build() {
	for d in $(ls ./$1); do
		if [ "$d" == "strategy" ]; then
			echo "当前迭代 strategy ，跳出。"
			continue
		elif [ "$d" == "dbproxy" ]; then
			echo "当前迭代 dbproxy ，跳出。"
			continue
		fi
		echo "building $1/$d filename:$buildprefix.$d"
		# pushd命令常用于将目录加入到栈中，加入记录到目录栈顶部，并切换到该目录；若pushd命令不加任何参数，则会将位于记录栈最上面的2个目录对换位置
		pushd $dir/$1/$d >/dev/null
		# -s 忽略符号表和调试信息，-w忽略DWARF符号表，通过这两个参数，可以进一步减少编译的程序的尺寸，更多的参数可以参考go link,
		#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w'
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $buildprefix.$d -a -installsuffix cgo -ldflags "-s -w $LDFLAGS -X '$PACKAGENAME/common.buildTime=$BUILD_DATE' -X '$PACKAGENAME/common.buildGitCommit=$VERSION' -X '$PACKAGENAME/common.buildGitRevision=$BUILD_GIT_REVISION' -X '$PACKAGENAME/common.buildGolangVersion=`go version`' -X '$PACKAGENAME/common.buildUser=$BUILD_USER' -X '$PACKAGENAME/common.buildHost=$BUILD_HOST' -X '$PACKAGENAME/common.buildStatus=$tree_status'"
		# opd用于删除目录栈中的记录；如果popd命令不加任何参数，则会先删除目录栈最上面的记录，然后切换到删除过后的目录栈中的最上面的目录
		popd >/dev/null
	done
}

build api
build srv

# 单独编译 web
echo "building web filename:$buildprefix.web"
cd web
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $buildprefix.web  -a -installsuffix cgo -ldflags '-s -w'
bee pack -be GOOS=linux -exs=.git:.go:.DS_Store:Dockerfile:.tmp:.sql:.pdm:.bak:.gz:.md:.bak -exp=logs:docker/deploy:build:controllers:vendor:.git:docker:temp
cd ..
