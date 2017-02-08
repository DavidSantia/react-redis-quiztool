#!/bin/sh

PROJECT=github.com/DavidSantia/react-redis-quiztool/load

# Build for Linux, statically linked
echo Building load app $GOPATH/src/$PROJECT
docker run --rm --name golang -v $GOPATH/src:/go/src golang:alpine /bin/sh -l -c \
    "cd /go/src/$PROJECT; CGO_ENABLED=0 /usr/local/go/bin/go build -v -i"

docker-compose build
