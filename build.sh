#!/bin/bash
echo "build go program"
go mod tidy
go build -o main .
echo "build docker image"
docker build -t api-bridge .
