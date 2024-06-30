#!/bin/bash

mkdir ./out
rm -r auth/* || mkdir ./auth

protoc --go_out=./out --go_opt=paths=source_relative --go-grpc_out=./out\
    --go-grpc_opt=paths=source_relative -I . \
    proto/auth.proto \
    proto/account/account.proto \
    proto/account/profile/profile.proto \
    proto/account/request.proto \
    proto/account/profile/request.proto

mv out/proto/* ./auth
rm -r ./out
