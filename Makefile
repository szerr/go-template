GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
CONF_PATH=configs/config.yaml

.PHONY: init
# 初始化依赖
init:
	go install github.com/swaggo/swag/cmd/swag@v1.16.3
	go install github.com/google/wire/cmd/wire@latest
	go install go.uber.org/mock/mockgen@latest

.PHONY: gen-mock
# 生成 mock
gen-mock:
	@mkdir -p ./test/mock/
	@for path in $(shell find ./internal/biz -name "*.go" | grep -v mock | grep -v biz.go); do \
  		echo $$path ;\
  		mockgen -source=$$path -destination=./internal/biz/mock_$$(basename $$path) -package=biz; \
	done
	@mockgen -source=./internal/pkg/db/db.go -destination=./internal/pkg/db/mock_db.go -package=db

.PHONY: generate
# generate
generate:
	@go generate ./...
	make gen-mock

.PHONY: swagger
# 生成 swagger 文档
swagger:
	@swag init -g handle.go -d api/handle -o docs --parseDependency

.PHONY: server
# 运行 server
server:
	@mkdir -p bin/ && go build -o bin/server cmd/server/main.go cmd/server/wire_gen.go && ./bin/server -c $(CONF_PATH)

.PHONY: init-db
# 初始化数据库
init-db:
	@go run cmd/init-db/main.go cmd/init-db/wire_gen.go -c $(CONF_PATH)

#.PHONY: gen-dao
# 生成 dao 和 model
#gen-dao:
#	@go run cmd/gen-dao/main.go cmd/gen-dao/wire_gen.go -c $(CONF_PATH)

.PHONY: add-user
# 添加新用户
add-user:
	@go run cmd/add-user/main.go cmd/add-user/wire_gen.go -c $(CONF_PATH)

.PHONY: gen-config
# 生成配置文件，自动整理旧的配置文件
gen-config:
	@go run cmd/gen-config/main.go -c $(CONF_PATH)

.PHONY: syncdb
# 同步 model 的结构到 db
syncdb:
	@go run cmd/syncdb/main.go cmd/syncdb/wire_gen.go -c $(CONF_PATH)

.PHONY: run
# 运行服务
run:
	@make generate
	@make server

.PHONY: doc
# 运行 doc 服务
doc:
	@mkdir -p bin/ && go build -o bin/doc cmd/doc/main.go cmd/doc/wire_gen.go && ./bin/doc -c $(CONF_PATH)

.PHONY: te
# 实验入口命令
te:
#	@go run cmd/te/main.go -c $(CONF_PATH)
	@go run cmd/te/main.go cmd/te/wire_gen.go -c $(CONF_PATH)

.PHONY: test
# 运行单元测试
test:
	@go test --count=1 ./...

.PHONY: build
# 编译
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: all
# generate all
all:
	make init;
	make swagger;
	make gen-config;
	make generate;
	make test;
	make build;

.PHONY: help
# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help