#!/bin/bash

echo "setting environment variables"
export GOPATH=`pwd`
export PATH=$GOPATH/bin:$PATH
cd src/app
govendor sync
cd ..
echo "running go app"

go run main.go
echo "executed go app"
