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

demo: [test_client.pb.go](cmd/protoc-gen-go-kit-grpc/proto/test_client.pb.go)

### usage

```shell
## dep
go get -u google.golang.org/protobuf/proto
go install github.com/golang/protobuf/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go install github.com/catusax/go-kit-micro/cmd/protoc-gen-go-kit-grpc@latest
protoc --proto_path=. --go-grpc_out=. --go_out=:. --go-kit-grpc_out=. proto/<file>.proto

```

to generate testing code ,use:

```shell
protoc --proto_path=. --go-grpc_out=. --go_out=:. --go-kit-grpc_out=. --go-kit-grpc_opt=test=true proto/<file>.proto
```

## thanks

<https://github.com/rotemtam/protoc-gen-go-ascii>

<https://github.com/wencan/kit-demo>

<https://github.com/go-kit/kit>

<https://github.com/asim/go-micro>

## License

Copyright (c) 2022 catusax

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.