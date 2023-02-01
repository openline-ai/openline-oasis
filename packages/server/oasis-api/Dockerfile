# syntax=docker/dockerfile:1

FROM golang:1.19-alpine as builder
RUN apk add --update make protobuf

WORKDIR /app

COPY channels-api ./channels-api
COPY oasis-api ./oasis-api

WORKDIR /app/channels-api
RUN make clean install generate

WORKDIR /app/oasis-api
RUN make clean all

FROM alpine:3.14
COPY --chown=65534:65534 --from=builder /app/bin/oasis-api .
USER 65534
EXPOSE 3006

ENTRYPOINT [ "./oasis-api" ]
