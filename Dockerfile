FROM golang:alpine

WORKDIR /app

COPY ./ /app

RUN go mod download

ENTRYPOINT go run cmd/server.go