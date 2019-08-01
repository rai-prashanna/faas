#!/bin/bash

echo "setting environment variables in factorial-service"
export GOPATH=`pwd`
export PATH=$GOPATH/bin:$PATH
cd src/app
govendor sync
cd ..
echo "running go app in factorial-service"

go run main.go
echo "executed go app in factorial-service"
#fswatch -config /fsw.yml

