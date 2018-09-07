#!/bin/bash

printf "\033[1;33m"
echo "Begin to build ..."
printf "\033[m"

cd ..

export GOPATH=$GOPATH:`pwd`
echo $GOPATH

cd -

# creational
# go build 04_builder/builder.go
go test -v creational/01_factory/ 
go test -v creational/02_abstractfactory/
go test -v creational/03_singleton/
go test -v creational/04_builder/
go test -v creational/05_prototype/

# structral
go test -v structural/06_adapter/
go test -v structural/07_bridge/
go test -v structural/08_filter/
go test -v structural/09_composite/
go test -v structural/10_decorator/
go test -v structural/11_facade/
go test -v structural/12_flyweight/
go test -v structural/13_proxy/

# behavioral
go test -v behavioral/14_chain/
go test -v behavioral/15_command/

