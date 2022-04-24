package main

import (
	"flag"
	"github.com/catusax/go-kit-micro/cmd/protoc-gen-go-kit-grpc/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var test *bool
var handler *bool

func main() {

	var flags flag.FlagSet
	test = flags.Bool("test", false, "generate testing code")
	handler = flags.Bool("handler", false, "generate handler code")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generator.GenerateClientFile(gen, f)

			if *test {
				generator.GenerateTestFile(gen, f)
			}

			if *handler {
				generator.GenerateHandlerFile(gen, f)
			}
		}
		return nil
	})
}
