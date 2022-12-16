# syntax=docker/dockerfile:1

FROM golang:1.19-alpine as builder

RUN apk add --update make wget unzip libstdc++ gcompat

WORKDIR /tmp
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v21.9/protoc-21.9-linux-x86_64.zip
RUN unzip protoc-21.9-linux-x86_64.zip
RUN mv bin/protoc /usr/local/bin
RUN mv include/* /usr/local/include/
RUN ldd /usr/local/bin/protoc && false
WORKDIR /app



COPY channels-api ./channels-api
COPY oasis-api ./oasis-api

WORKDIR /app/channels-api
RUN make install
RUN make generate

WORKDIR /app/oasis-api
RUN make install
RUN make generate
RUN go build -o /go/bin/app openline-ai/oasis-api

FROM alpine:3.14
COPY --chown=65534:65534 --from=builder /go/bin/app .
USER 65534
EXPOSE 3006

ENTRYPOINT [ "./app" ]
