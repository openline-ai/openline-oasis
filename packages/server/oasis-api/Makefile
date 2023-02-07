clean:
	rm -rf bin
	rm -rf proto/generated/

generate:
	export PATH="${PATH}:$(shell go env GOPATH)/bin";go generate ./...

build:
	go build -v -o bin/oasis-api

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install gotest.tools/gotestsum@latest

all: clean install generate build
