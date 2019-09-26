.PHONY: proto build

GCTPKG = $(shell go list -e)
LINTPKG = github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1
LINTBIN = $(GOPATH)/bin/golangci-lint
COMMIT_HASH=$(shell git rev-parse --short HEAD || echo "GitNotFound")


get:
	export GOPROXY=https://mirrors.aliyun.com/goproxy/
	GO111MODULE=on go get $(GCTPKG)

linter:
	export GOPROXY=https://mirrors.aliyun.com/goproxy/
	GO111MODULE=on go get $(GCTPKG)
	GO111MODULE=on go get $(LINTPKG)
	golangci-lint run --verbose --skip-dirs=web-exchange --skip-dirs=models | tee /dev/stderr

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic  ./...

update_deps:
	export GOPROXY=https://mirrors.aliyun.com/goproxy/
	GO111MODULE=on go mod verify
	GO111MODULE=on go mod tidy
	rm -rf vendor
	GO111MODULE=on go mod vendor

fmt:
	gofmt -l -w -s $(shell find . -path './vendor' -prune -o -type f -name '*.go' -print)

proto:
	for d in srv; do \
		for f in $$d/**/**/proto/*.proto; do \
			protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
	done

build:
	./bin/build.sh

run:
	docker-compose up

rund:
	docker-compose up -d

down:
	docker-compose down


clear:
	./bin/clear.sh