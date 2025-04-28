FROM golang:1.23-alpine AS builder

ARG TAG=NO_TAG
ARG DEBUG=true

WORKDIR /app/
COPY . .

RUN go mod download

FROM alpine:latest

ARG MODE

WORKDIR /app

COPY --from=0 /app/bin . 
COPY --from=0 /app/config.yml config.yml

ENTRYPOINT ["./app"]