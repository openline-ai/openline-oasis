# syntax=docker/dockerfile:1

FROM golang:1.19-alpine as builder


WORKDIR /app

COPY channels-api ./channels-api
COPY oasis-api ./oasis-api

WORKDIR /app/oasis-api
RUN go build -o /go/bin/app github.com/openline-ai/openline-oasis/packages/server/oasis-api

FROM alpine:3.14
COPY --chown=65534:65534 --from=builder /go/bin/app .
USER 65534
EXPOSE 3006

ENTRYPOINT [ "./app" ]
