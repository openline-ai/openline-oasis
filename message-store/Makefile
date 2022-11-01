generate:
	go generate ./...
	cd ent/proto/entpb; go generate ./...

build:
	go build -v -o bin/message-store

test:
	go test -v

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"

clean:
	find ent -mindepth 1 -maxdepth 1 -not \( -name 'proto' -or -name 'schema' -or -name generate.go \) -exec rm -r {} \;
	find ent/proto/entpb -mindepth 1 -maxdepth 1 -not \( -name 'entpb.proto' -or -name generate.go \) -exec rm -r {} \;
	rm -rf bin
