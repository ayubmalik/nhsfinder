#!/usr/bin/env bash
cname=trams

rm -rf dist

docker build -t ayubmalik/${cname} .

docker create -it --name ${cname} ayubmalik/${cname}

docker cp ${cname}:/go/src/github.com/ayubmalik/trams/dist/ .

docker rm ${cname}
