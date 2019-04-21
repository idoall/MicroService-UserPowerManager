.PHONY: proto build

proto:
	for d in srv; do \
		for f in $$d/**/proto/*.proto; do \
			protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
	done
	# 去掉omitempty不生效，又不想改原码
	# ls srv/srvcolumns/proto/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

build:
	./bin/build.sh
	
run:
	docker-compose build
	docker-compose up

down:
	docker-compose down
