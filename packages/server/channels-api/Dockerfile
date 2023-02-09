# syntax=docker/dockerfile:1

FROM golang:1.19-alpine as builder
RUN apk add --update make protobuf protobuf-dev bash

WORKDIR /app

COPY oasis-api ./oasis-api
COPY channels-api ./channels-api

WORKDIR /app/oasis-api
RUN make clean install generate

WORKDIR /app/channels-api
RUN make clean all

FROM alpine:3.14
COPY --chown=65534:65534 --from=builder /app/channels-api/bin/channels-api .
USER 65534

ENTRYPOINT [ "./channels-api" ]
