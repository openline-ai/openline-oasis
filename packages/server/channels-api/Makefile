generated = proto/generated/common.pb.go proto/generated/messageevent.pb.go

generate: $(generated)

$(generated): proto/generated/%.pb.go: proto/%.proto
	$(shell ./import-protobuf.sh message-store-api)
	export PATH="${PATH}:$(shell go env GOPATH)/bin";go generate ./...

clean:
	rm -rf bin
	rm -rf proto/generated/*
	rm -rf protobuf-import

test:
	go test -v ./... -cover

build: generate bin/channels-api

bin/channels-api:
	go build -v -o bin/channels-api

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install gotest.tools/gotestsum@latest

all: clean install build
