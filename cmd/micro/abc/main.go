package main

import (
	"abc/handler"
	pb "abc/proto"
	"fmt"
	"github.com/coolrc136/go-kit-micro/sd"
	"github.com/coolrc136/go-kit-micro/utils"
	"go.uber.org/zap/zapcore"
	syslog "log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	kitzap "github.com/go-kit/kit/log/zap"
	"go.uber.org/zap"
)

var (
	service = "abc"
    quitChan = make(chan error, 1)
)

func main() {

	logger, err := zap.NewDevelopment(zap.AddStacktrace(zapcore.WarnLevel))
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger.Named(abc))
    rpcLogger = kitzap.NewZapSugarLogger(logger, logLevel)


	grpcListener, err := net.Listen("tcp", ":") //随机端口
	// 服务地址
	host, sport, err := net.SplitHostPort(grpcListener.Addr().String())
	if err != nil {
		syslog.Println(err)
		return
	}
	if host == "::" {
		host = utils.GetOutboundIP().String()
	}
	port, err := strconv.Atoi(sport)
	if err != nil {
		syslog.Println(err)
		return
	}
	registrar, err := sd.NewRegistrar(context.TODO(), host, port, service, rpcLogger)
	if err != nil {
		return
	}
	// Create service

	go func() {
		// Register handler
		srv := NewRpcServer()

		pb.RegisterAbcServer(srv, new(handler.Abc))
		registrar.Register()

		syslog.Println("serving")
		err := srv.Serve(grpcListener)
		// Run service
		if err != nil {
			logger.Error("during listen err:", zap.Error(err))
			quitChan <- err
		}

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		quitChan <- fmt.Errorf("%s", <-c)
	}()
	err = <-quitChan
	syslog.Println("gracefully shutting down...")
	registrar.Deregister()
	syslog.Println(err)
}

func NewRpcServer() *grpc.Server {
	return grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
