# go-kit-micro

tools to help building micro services with go-kit

## cli

command line tool to create micro service scaffold

### usage

```shell
go install github.com/catusax/go-kit-micro/cmd/micro@latest
micro new service test
cd test
make init proto tidy
```

### config

`github.com/catusax/go-kit-micro/sd` reads registry config from ENV:

- `$ETCD` : etcd addresses,multiple address can be split by `,`,do not insert blank between
  addresses,eg: `127.0.0.1:2379` or `127.0.0.1:2379,127.0.0.1:2380`
- `$CONSUL` : Consul address,eg: `127.0.0.1:8300`

if you don't set both `$ETCD` and `$CONSUL`,by default,your micro service will use mDNS as registry.

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