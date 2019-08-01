#!/bin/bash

echo "setting environment variables in dig-service"
export GOPATH=`pwd`
export PATH=$GOPATH/bin:$PATH
cd src/app
govendor sync
cd ..
echo "running go app in dig-service"

go run main.go
echo "executed go app in dig-service"
#fswatch -config /fsw.yml
