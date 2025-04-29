FROM golang:1.24-alpine AS builder

WORKDIR /app/

COPY . .

RUN go mod download


RUN go build -o bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=0 /app/bin . 
COPY --from=0 /app/services/inventry-service/config.yaml ./services/inventry-service/config.yaml
COPY --from=0 /app/services/order-service/config.yaml ./services/order-service/config.yaml
COPY --from=0 /app/services/payment-service/config.yaml ./services/payment-service/config.yaml

ENTRYPOINT ["./app"]