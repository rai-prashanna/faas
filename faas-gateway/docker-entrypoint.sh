#!/bin/bash

echo "setting environment variables in faas-gateway-service"
export GOPATH=`pwd`
export PATH=$GOPATH/bin:$PATH
cd src/app
govendor sync
cd ..
echo "running go app in faas-gateway-service"

go run main.go
echo "executed go app in faas-gateway-service"
#fswatch -config /fsw.yml

