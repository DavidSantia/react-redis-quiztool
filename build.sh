#!/bin/sh

export PROJECT=github.com/DavidSantia/react-redis-quiztool

# Build for Linux, statically linked

export NAME=load

echo "## Building $NAME app"
docker run --rm --name golang -v $GOPATH/src:/go/src golang:alpine /bin/sh -l -c \
    "cd /go/src/$PROJECT/$NAME; CGO_ENABLED=0 /usr/local/go/bin/go build -v -i; tar -cf $NAME.tar $NAME"

export NAME=redis-ws

echo "## Building $NAME app"
docker run --rm --name golang -v $GOPATH/src:/go/src golang:alpine /bin/sh -l -c \
    "cd /go/src/$PROJECT/$NAME; CGO_ENABLED=0 /usr/local/go/bin/go build -v -i; tar -cf $NAME.tar $NAME"

echo "## Running docker-compose build"
docker-compose build
