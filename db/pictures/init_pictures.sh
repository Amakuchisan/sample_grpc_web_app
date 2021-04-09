#!/bin/bash

cd db/pictures

docker run -w /work -v $PWD:/work --rm golang:1.16-alpine go run main.go

cp 2_insert.sql ../init/
rm -f 2_insert.sql
