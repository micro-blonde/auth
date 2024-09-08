# IMPORTANT
if you get error about google proto files import, you'll need to clone protofiles and put them into ./proto directory:
```sh
git clone https://github.com/protocolbuffers/protobuf.git
cp -r ./protobuf/src/google ./proto
```

---

- To generate/update go files run:
    ```shell script
    ./generate.sh
    ```

- To generate/update ts/js files run:
    ```sh
    npm install -g grpc-tools
    npm install -g @grpc/proto-loader grpc
    ```
    ```sh
    ./generate-ts.sh
    ```
