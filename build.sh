#!/bin/sh

# Build for Linux, statically linked
(cd load; GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build)

docker-compose build
