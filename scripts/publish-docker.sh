#!/bin/bash

TAG=`git describe --abbrev=0 --tags`

rm til-system
make build-linux
docker login
docker build -t robinthrift/til-system:$TAG .
docker tag robinthrift/til-system:$TAG robinthrift/til-system:latest
docker push robinthrift/til-system
