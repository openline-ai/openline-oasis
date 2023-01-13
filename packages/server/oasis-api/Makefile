
clean:
	find ent -mindepth 1 -maxdepth 1 -not \( -name '.placeholder' \) -exec rm -r {} \;
	rm -rf bin

build:
	go build -v -o bin/oasis-api

all: clean build
