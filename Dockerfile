FROM golang:1.13 as builder

WORKDIR /go/src/github.com/ayubmalik/trams

COPY . /go/src/github.com/ayubmalik/trams

RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux go build -o dist/linux/trams ./cmd/trams

RUN GO111MODULE=off CGO_ENABLED=0 GOOS=darwin go build -o dist/darwin/trams ./cmd/trams

VOLUME /go/src/github.com/ayubmalik/trams
