
generate:
	export PATH="${PATH}:$(shell go env GOPATH)/bin";go generate ./...
	export PATH="${PATH}:$(shell go env GOPATH)/bin";protoc --go_out=ent --go_opt=paths=source_relative \
        --go-grpc_out=ent --go-grpc_opt=paths=source_relative \
        proto/messageevent.proto

build: generate
	go build -v -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" -o bin/channels-api

test:
	go test -v

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install gotest.tools/gotestsum@latest

clean:
	find ent -mindepth 1 -maxdepth 1 -not \( -name '.placeholder' \) -exec rm -r {} \;
	rm -rf bin

all: install build
