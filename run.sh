#!/usr/bin/env bash

if [ -z ${GOPATH+x} ]; then
    echo "GOPATH is unset";
    export GOPATH=$HOME/gowork;
    echo "Setting GOPATH as '$GOPATH'"
else
    echo "GOPATH is set to '$GOPATH'";
fi

if [ ! -d "$GOPATH" ]; then
    mkdir $GOPATH
fi

rm -rf $GOPATH/src/com.crawler
cp -r src/com.crawler $GOPATH/src

go get -u golang.org/x/net/html

go run server.go
