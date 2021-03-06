FROM golang:1.13 as builder

WORKDIR /go/src/github.com/ayubmalik/nhsfinder

COPY . /go/src/github.com/ayubmalik/nhsfinder

RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux go build -o dist/linux/finder ./cmd/finder

RUN GO111MODULE=off CGO_ENABLED=0 GOOS=darwin go build -o dist/darwin/finder ./cmd/finder

VOLUME /go/src/github.com/ayubmalik/nhsfinder
