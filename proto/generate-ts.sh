#!/bin/bash

mkdir ./out

protoc-gen-grpc-ts \
    --ts_out=paths=source_relative,grpc_js:./out \
    -I ./proto \
    auth.proto \
    account/account.proto \
    account/profile/profile.proto \
    account/request.proto \
    account/profile/request.proto

protoc-gen-grpc \
    --js_out=import_style=commonjs,binary:./out \
    --grpc_out=grpc_js:./out \
    -I ./proto \
    auth.proto \
    account/account.proto \
    account/profile/profile.proto \
    account/request.proto \
    account/profile/request.proto

rm -r ts/* || mkdir ./ts
mv out/* ./ts
rm -r ./out
