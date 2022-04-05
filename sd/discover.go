package sd

import (
	"context"
	"github.com/catusax/go-kit-micro/config"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/log"
	consul_api "github.com/hashicorp/consul/api"
	"github.com/wencan/kit-plugins/sd/mdns"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	syslog "log"
	"time"
)

func NewInstancer(service string, log log.Logger) (sd.Instancer, error) {
	// etcd参数
	etcdServers := config.C.Etcd
	// consul参数
	consulServer := config.C.Consul

	if len(etcdServers) > 0 {
		// 如果提供etcd服务器地址参数
		// 使用etcd发现服务
		// etcd客户端
		etcdClient, err := etcdv3.NewClient(context.TODO(), etcdServers, etcdv3.ClientOptions{})
		if err != nil {
			syslog.Println(err)
			return nil, err
		}
		instancer, err := etcdv3.NewInstancer(etcdClient, service, log)
		if err != nil {
			syslog.Println(err)
			return nil, err
		}
		return instancer, nil
	} else if consulServer != "" {
		consulConfig := consul_api.DefaultConfig()
		consulConfig.Address = consulServer
		c, err := consul_api.NewClient(consulConfig)
		if err != nil {
			syslog.Println(err)
			return nil, err
		}
		client := consul.NewClient(c)
		instancer := consul.NewInstancer(client, log, service, nil, false)
		return instancer, nil
	} else {
		// 如果没提供etcd服务器地址参数
		// 使用mDNS发现服务
		instancer, err := mdns.NewInstancer(service, mdns.InstancerOptions{}, log)
		if err != nil {
			syslog.Println(err)
			return nil, err
		}
		return instancer, nil
	}
}

func GetEndPoint(instancer sd.Instancer, clientEndpoint func(conn *grpc.ClientConn, ctx context.Context, request interface{}) (interface{}, error), logger log.Logger) endpoint.Endpoint {
	end := sd.NewEndpointer(instancer, func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, nil, err
		}
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			return clientEndpoint(conn, ctx, request)

		}, conn, nil

	}, logger)
	blancer := lb.NewRandom(end, time.Now().UnixNano())
	return lb.Retry(3, 500*time.Millisecond, blancer)
}
