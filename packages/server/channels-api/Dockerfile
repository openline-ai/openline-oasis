# syntax=docker/dockerfile:1

FROM golang:1.19-alpine as builder

WORKDIR /app

COPY oasis-api/proto ./oasis-api/proto
COPY oasis-api/go.mod ./oasis-api/go.mod
COPY channels-api ./channels-api

WORKDIR /app/channels-api

RUN go build -o /go/bin/app openline-ai/channels-api

FROM alpine:3.14
COPY --chown=65534:65534 --from=builder /go/bin/app .
USER 65534

ENTRYPOINT [ "./app" ]
