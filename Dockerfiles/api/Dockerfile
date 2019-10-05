# https://qiita.com/kkeisuke/items/7cd4d5834386666faab3
FROM golang:1.13-alpine

WORKDIR /go/src/api

COPY . .

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

EXPOSE 8080

CMD go run ./cmd/api/main.go
