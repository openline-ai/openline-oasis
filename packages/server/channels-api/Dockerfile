# syntax=docker/dockerfile:1

FROM golang:1.19-alpine as builder

RUN apk add --update make protobuf

WORKDIR /app

COPY oasis-api ./oasis-api
COPY channels-api ./channels-api

WORKDIR /app/oasis-api
RUN make install
RUN make generate

WORKDIR /app/channels-api
RUN make install
RUN make generate
RUN go build -o /go/bin/app openline-ai/channels-api

FROM alpine:3.14
COPY --chown=65534:65534 --from=builder /go/bin/app .
USER 65534

ENTRYPOINT [ "./app" ]
