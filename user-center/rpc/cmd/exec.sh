#/bin/bash
cd ../ && goctl rpc protoc *.proto --go_out=. --go-grpc_out=. --zrpc_out=.

