#! /bin/bash

for GOOS in darwin linux windows; do
  for GOARCH in amd64 arm64; do
    echo "Building $GOOS-$GOARCH"
    export GOOS=$GOOS
    export GOARCH=$GOARCH
    go build -ldflags="-s -w" -o bin/uxmal-$GOOS-$GOARCH
  done
done
