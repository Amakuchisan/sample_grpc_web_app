#!/bin/bash

set -eux

cd $(pwd)/pb/

docker build protoc-go -t streaming-protoc-go
docker build protoc-python -t streaming-protoc-python


mkdir -p go/picture

docker run -v "$(pwd):/pb" -w /pb --rm streaming-protoc-go \
  protoc \
    --go_out=plugins=grpc:go/picture\
    --go_opt=paths=source_relative \
    picture.proto

mkdir -p python/picture

docker run -v "$(pwd):/pb" -w /pb --rm streaming-protoc-python \
  python -m grpc_tools.protoc -I. \
    --python_out=python/picture \
    --grpc_python_out=python/picture \
    picture.proto

sed -i -e "s/import picture_pb2 as picture__pb2/from . import picture_pb2 as picture__pb2/g" python/picture/picture_pb2_grpc.py # インポートに失敗するため、置換している

mkdir -p ../services/server-go/pb
mkdir -p ../services/server-python/pb
mkdir -p ../services/db-manager/pb

cp -r ./go/* ../services/server-go/pb/
cp -r ./go/* ../services/db-manager/pb/
cp -r ./python/* ../services/server-python/pb/

docker build protoc-web -t streaming-protoc-web
mkdir -p js/picture
docker run -v "$(pwd):/pb" -w /pb --rm streaming-protoc-web \
  protoc --proto_path=. picture.proto \
    --js_out=import_style=commonjs:js/picture \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:js/picture
mkdir -p ../services/client/src/pb
cp -r ./js/* ../services/client/src/pb/
