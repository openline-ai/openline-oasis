
build:
	go build -v -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" -o bin/channels-api

test:
	go test -v

clean:
	find ent -mindepth 1 -maxdepth 1 -not \( -name '.placeholder' \) -exec rm -r {} \;
	rm -rf bin

all: clean build
