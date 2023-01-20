#!/bin/bash

if [ ! -f ./tmp.go ]; then
echo "package main" > ./tmp.go
cat "../../structs.txt" >> ./tmp.go
fi
go fmt ./tmp.go
