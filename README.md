# faas-scaffold

Clone the scaffold and run `docker-compose up` in dir to get started.

For support in your IDEA, a sugestion is to put it into your `GOPATH` and work form there, eg. `$GOPATH/src/github.com/<YOUR_HANDLE>/faas`

 export GOPATH=`pwd`
 export PATH=$GOPATH/bin:$PATH
cd src/app
govendor sync
cd ..
go run main.go