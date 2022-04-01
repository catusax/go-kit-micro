package main

import (
	"github.com/coolrc136/go-kit-micro/cmd/micro/cli"
	// register commands
	_ "github.com/coolrc136/go-kit-micro/cmd/micro/cli/generate"
	_ "github.com/coolrc136/go-kit-micro/cmd/micro/cli/new"
)

func main() {
	cli.Run()
}
