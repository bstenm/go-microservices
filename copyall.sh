#!/bin/bash
export CGO_ENABLED=0

# build account service
cd accountservice;go get;go build -o accountservice-linux-amd64;echo built `pwd`;cd ..
# build healthchecker
cd healthchecker;go get;go build -o healthchecker-linux-amd64;echo built `pwd`;cd ..
# build quote service
cd quoteservice;go get;go build -o quoteservice-linux-amd64;echo built `pwd`;cd ..

# copy healthchecker binary into the account service folder
cp healthchecker/healthchecker-linux-amd64 accountservice/

docker build -t bstenm/accountservice accountservice/
docker build -t bstenm/quoteservice quoteservice/

docker service rm accountservice
docker service rm quoteservice
docker service create --name=accountservice --replicas=1 --network=my_network -p=6767:6767 bstenm/accountservice
docker service create --name=quoteservice --replicas=1 --network=my_network -p=6768:6768 bstenm/quoteservice