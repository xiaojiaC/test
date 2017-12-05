```
cd ${GOPATH}/src/git.augmentum.com.cn/test/proto
```

## Generate gRPC stub

```
protoc -I. \
  -I$GOPATH/src/git.augmentum.com.cn/test/proto \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  example/helloworld.proto
```

## Generate reverse-proxy

```
protoc -I. \
  -I$GOPATH/src/git.augmentum.com.cn/test/proto \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  example/helloworld.proto
```

## Generate swagger definitions

```
protoc -I. \
  -I$GOPATH/src/git.augmentum.com.cn/test/proto \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:. \
  example/helloworld.proto
```