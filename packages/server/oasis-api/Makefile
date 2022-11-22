
generate:
	export PATH="${PATH}:$(shell go env GOPATH)/bin";go generate ./...
	export PATH="${PATH}:$(shell go env GOPATH)/bin";protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/oasisapi.proto

build: generate
	go build -v -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" -o bin/oasis-api

test:
	go test -v ./... -cover

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

clean:
	find proto -mindepth 1 -maxdepth 1 -not \( -name 'oasisapi.proto' \) -exec rm -r {} \;
	rm -rf bin

all: install build
