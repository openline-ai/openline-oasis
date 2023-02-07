generated = proto/generated/oasisapi.pb.go

clean:
	rm -rf bin
	rm -rf proto/generated/

generate: $(generated)

$(generated): proto/generated/%.pb.go: proto/%.proto
	export PATH="${PATH}:$(shell go env GOPATH)/bin";go generate ./...

build: generate bin/oasis-api

bin/oasis-api:
	go build -v -o bin/oasis-api

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install gotest.tools/gotestsum@latest

all: clean install build
