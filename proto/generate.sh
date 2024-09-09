#!/bin/bash

mkdir ./out

protoc --go_out=./out --go_opt=paths=source_relative --go-grpc_out=./out \
    --go-grpc_opt=paths=source_relative -I ./template \
    auth.proto \
    request.proto \
    account/account.proto \
    account/profile/profile.proto

rm -r auth/* || mkdir ./auth
mv out/* ./auth
rm -r ./out
