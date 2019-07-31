# faas-scaffold

Clone the scaffold and run `docker-compose up` in dir to get started.

For support in your IDEA, a sugestion is to put it into your `GOPATH` and work form there, eg. `$GOPATH/src/github.com/<YOUR_HANDLE>/faas`

 export GOPATH=`pwd`
 export PATH=$GOPATH/bin:$PATH
cd src/app
govendor sync
cd ..
go run main.go


edit /etc/default/docker
uncomment export http_proxy="http://127.0.0.1:3128/"



[{0.0.0.0 %!s(uint16=7070) %!s(uint16=7070) tcp}]


'{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'
docker inspect --format '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' 22c790841c5a