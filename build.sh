#!/usr/bin/env bash

echo "=========== unit test ============"
mkdir unittest
go test -v -cover -coverprofile=./unittest/cover.out -covermode=count github.com/beanzhou/tandemo/storage
go tool cover -html=./unittest/cover.out -o ./unittest/cover.html
echo "======= unit test Done ==========="

cd cmd/tandemo/ && go build -v && cd ../../
mv cmd/tandemo/tandemo ./
