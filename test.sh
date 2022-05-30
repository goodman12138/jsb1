#!/bin/bash
rm -rf jsb/
git clone git@github.com:goodman12138/jsb.git

chmod 777 jsb

cd jsb
go build main.go

docker stop 82web
docker rm 82web
docker rmi 82webimages
docker build -t 82webimages .
docker run -dit -v /root/shtest/wlog:/go/wlog --name=82web -p 8082:8082 82webimages



your_name="start test ..."
echo $your_name