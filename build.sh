#!/usr/bin/env bash
repo=nhsfinder
command=finder

rm -rf dist

docker build -t ayubmalik/${command} .

docker create -it --name ${command} ayubmalik/${command}

docker cp ${command}:/go/src/github.com/ayubmalik/${repo}/dist/ .

docker rm ${command}
