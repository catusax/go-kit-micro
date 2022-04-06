package generator

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

// GenerateFile generates a _ascii.pb.go file containing gRPC service definitions.
func GenerateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_client.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-ascii. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	//for _, msg := range file.Messages {
	//	g.P("func (x *", msg.GoIdent, ") Ascii() string {")
	//	g.P("return `", " fig.String()", "`")
	//	g.P("}")
	//}

	g.P(`import(
	"context"
	"github.com/catusax/go-kit-micro/sd"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)
`)

	//create client
	for _, srv := range file.Services {
		g.P("type ClientImpl struct {")
		for _, method := range srv.Methods {
			g.P(FirstLower(method.GoName), "  endpoint.Endpoint")
		}
		g.P("}")

		// create NewClient func
		g.P("func New", srv.GoName, "ClientImpl(logger log.Logger) *ClientImpl {")
		g.P("  instancer, err := sd.NewInstancer(\"" + strings.ToLower(srv.GoName) + ".service\", logger)")
		g.P(`if err != nil {
		panic(err)
	}

    return &ClientImpl{
`)
		for _, method := range srv.Methods {

			if method.Desc.IsStreamingClient() {
				g.P(FirstLower(method.GoName), fmt.Sprintf(`: sd.GetEndPoint(instancer, func(conn *grpc.ClientConn, ctx context.Context, request interface{}) (interface{}, error) {
			client := New%sClient(conn)

			return client.%s(ctx)
		}, logger),`, srv.GoName, method.GoName))
			} else {
				g.P(FirstLower(method.GoName), fmt.Sprintf(`: sd.GetEndPoint(instancer, func(conn *grpc.ClientConn, ctx context.Context, request interface{}) (interface{}, error) {
			client := New%sClient(conn)
			req := request.(*%s)

			return client.%s(ctx, req)
		}, logger),`, srv.GoName, method.Input.GoIdent.GoName, method.GoName))
			}

		}

		g.P("}}")

		//create Client Methods
		for _, method := range srv.Methods {
			g.P("// ", method.GoName)
			//
			//for _, comm := range method.Comments.LeadingDetached {
			//	g.P(comm.String())
			//}

			leadcomm := method.Comments.Leading.String()
			if len(leadcomm) > 0 {
				g.P(leadcomm[:len(leadcomm)-1])
			}

			//g.P(method.Comments.Trailing.String())

			//bidi
			if method.Desc.IsStreamingClient() && method.Desc.IsStreamingClient() {
				g.P(fmt.Sprintf(`func (n *ClientImpl) %s(ctx context.Context) (%s_BidiStreamClient, error) {
	rsp, err := n.%s(ctx, nil)
	if err != nil {
		return nil, err
	}
	res := rsp.(%s_BidiStreamClient)
	return res, err
}`, method.GoName, srv.GoName, FirstLower(method.GoName), srv.GoName))
			} else if method.Desc.IsStreamingServer() {
				g.P(fmt.Sprintf(`func (n *ClientImpl) %s(ctx context.Context, req *%s) (%s_ServerStreamClient, error) {
	rsp, err := n.%s(ctx, req)
	if err != nil {
		return nil, err
	}
	res := rsp.(%s_ServerStreamClient)
	return res, err
}`, method.GoName, method.Input.GoIdent.GoName, srv.GoName, FirstLower(method.GoName), srv.GoName))

			} else if method.Desc.IsStreamingClient() {
				g.P(fmt.Sprintf(`func (n *ClientImpl) %s(ctx context.Context) (%s_ClientStreamClient, error) {
	rsp, err := n.%s(ctx, nil)
	if err != nil {
		return nil, err
	}
	res := rsp.(%s_ClientStreamClient)
	return res, err
}`, method.GoName, srv.GoName, FirstLower(method.GoName), srv.GoName))

			} else {
				g.P(fmt.Sprintf(`func (n *ClientImpl) %s(ctx context.Context, req *%s) (*%s, error) {
	rsp, err := n.%s(ctx, req)
	if err != nil {
		return nil, err
	}
	res := rsp.(*%s)
	return res, err
}`, method.GoName, method.Input.GoIdent.GoName, method.Output.GoIdent.GoName, FirstLower(method.GoName), method.Output.GoIdent.GoName))

			}
		}
	}

	return g
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
