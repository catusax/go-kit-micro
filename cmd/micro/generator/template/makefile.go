package template

// Makefile is the Makefile template used for new projects.
var Makefile = `
GOPATH:=$(shell go env GOPATH)
NAME={{.Service}}{{if .Client}}-client{{end}}
BIN={{.Service}}

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/catusax/go-kit-micro/cmd/protoc-gen-go-kit-grpc@latest

.PHONY: proto
proto:
	@protoc --proto_path=. -I${GOPATH}/src --go-grpc_out=. --go_out=:. --go-kit-grpc_out=. proto/$(NAME).proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o $(BIN) *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t $(BIN):latest .
`
