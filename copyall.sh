#!/bin/bash
export CGO_ENABLED=0

# build account service
cd accountservice;go get;go build -o accountservice-linux-amd64;echo built `pwd`;cd ..
# build healthchecker
cd healthchecker;go get;go build -o healthchecker-linux-amd64;echo built `pwd`;cd ..

# copy healthchecker binary into the account service folder
cp healthchecker/healthchecker-linux-amd64 accountservice/

docker build -t bstenm/accountservice accountservice/

docker service rm accountservice
docker service create --name=accountservice --replicas=1 --network=my_new_network -p=6767:6767 bstenm/accountservice