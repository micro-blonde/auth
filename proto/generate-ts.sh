#!/bin/bash

mkdir ./out

protoc-gen-grpc-ts \
    --ts_out=paths=source_relative,grpc_js:./out \
    -I ./template \
    auth.proto \
    request.proto \
    account/account.proto \
    account/profile/profile.proto

protoc-gen-grpc \
    --js_out=import_style=commonjs,binary:./out \
    --grpc_out=grpc_js:./out \
    -I ./template \
    auth.proto \
    request.proto \
    account/account.proto \
    account/profile/profile.proto

rm -r ts/* || mkdir ./ts
mv out/* ./ts
rm -r ./out
