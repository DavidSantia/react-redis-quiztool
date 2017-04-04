#!/bin/sh

PROJECT=github.com/DavidSantia/react-redis-quiztool

# Build for Linux, statically linked

NAME=load

echo "## Building $NAME app"
docker run --rm --name golang -v $GOPATH/src:/go/src golang:alpine /bin/sh -l -c \
    "cd /go/src/$PROJECT/$NAME; CGO_ENABLED=0 /usr/local/go/bin/go build -v -i"

NAME=redis-ws

echo "## Building $NAME app"
docker run --rm --name golang -v $GOPATH/src:/go/src golang:alpine /bin/sh -l -c \
    "cd /go/src/$PROJECT/$NAME; CGO_ENABLED=0 /usr/local/go/bin/go build -v -i"

echo "## Running docker-compose build"
docker-compose build
