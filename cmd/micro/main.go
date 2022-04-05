package main

import (
	"github.com/catusax/go-kit-micro/cmd/micro/cli"
	// register commands
	_ "github.com/catusax/go-kit-micro/cmd/micro/cli/generate"
	_ "github.com/catusax/go-kit-micro/cmd/micro/cli/new"
)

func main() {
	cli.Run()
}
