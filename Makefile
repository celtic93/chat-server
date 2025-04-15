include .env

LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

lint-fix:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml --fix

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-api

generate-api:
	mkdir -p pkg/v1/chat
	protoc --proto_path api/v1/chat \
	--go_out=pkg/v1/chat --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/v1/chat --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/v1/chat/chat.proto

migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} status -v

db-migrate:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} up -v

db-rollback:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} down -v

run:
	go run cmd/main.go
