FROM golang:1.11

WORKDIR /go/src/github.com/ayubmalik/pharmacyfinder/

COPY . ./

WORKDIR /go/src/github.com/ayubmalik/pharmacyfinder/cmd/server

RUN go get goji.io goji.io/pat && go build .

RUN ls -lrt

# $ docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.8 bash
#$ for GOOS in darwin linux; do
#>   for GOARCH in 386 amd64; do
#>     export GOOS GOARCH
#>     go build -v -o myapp-$GOOS-$GOARCH
#>   done
#> done