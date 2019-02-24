#!/bin/sh

fuser -k 80/tcp
fuser -k 443/tcp

cd app/frontend

yarn
yarn build

cd ..

go get -t
nohup go run . > log.log &

tail log.log
