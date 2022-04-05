# go-kit-micro

tools to help building micro services with go-kit

## protoc-gen-go-kit-grpc

go-kit endpoint generator

demo: [](cmd/protoc-gen-go-kit-grpc/proto/test_client.pb.go)

### usage

```shell
## dep
go get -u google.golang.org/protobuf/proto
go install github.com/golang/protobuf/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go install github.com/catusax/go-kit-micro/cmd/protoc-gen-go-kit-grpc@latest
protoc --proto_path=. --go-grpc_out=. --go_out=:. --go-kit-grpc_out=. proto/<file>.proto

```

## thanks

<https://github.com/wencan/kit-demo>
<https://github.com/go-kit/kit>