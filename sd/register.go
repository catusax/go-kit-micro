// from https://github.com/wencan/kit-demo

package sd

import (
	"context"
	"fmt"
	"github.com/coolrc136/go-kit-micro/config"
	"github.com/coolrc136/go-kit-micro/sd/mdns"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/log"
	consulApi "github.com/hashicorp/consul/api"
	syslog "log"
	"time"
)

var (
	ttl = 5 * time.Second
)

func NewRegistrar(ctx context.Context, host string, port int, service string, logger log.Logger) (sd.Registrar, error) {
	instance := fmt.Sprintf("%s:%d", host, port)

	if len(config.C.Etcd) > 0 {
		// etcd
		etcdClient, err := etcdv3.NewClient(ctx, config.C.Etcd, etcdv3.ClientOptions{
			DialTimeout:   ttl,
			DialKeepAlive: ttl,
		})
		if err != nil {
			syslog.Println(err)
			return nil, err
		}
		registrar := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
			Key:   service + "/" + instance,
			Value: instance,
		}, logger)
		return registrar, nil
	} else if config.C.Consul != "" {
		// consul
		consulConfig := consulApi.DefaultConfig()
		consulConfig.Address = config.C.Consul
		c, err := consulApi.NewClient(consulConfig)
		if err != nil {
			syslog.Println(err)
			return nil, err
		}
		client := consul.NewClient(c)
		registration := &consulApi.AgentServiceRegistration{
			Name:    service,
			ID:      service + "/" + instance,
			Address: host,
			Port:    port,
			Check: &consulApi.AgentServiceCheck{
				GRPC:     instance + "/" + service, // gRPC地址 + 健康检查service参数
				Interval: "10s",                    // 必须
			},
		}
		registrar := consul.NewRegistrar(client, registration, logger)
		return registrar, nil
	} else {
		// mDNS
		service := mdns.Service{
			Instance: instance, // unique
			Service:  service,
			Port:     port,
		}
		registrar, err := mdns.NewRegistrar(service, logger)
		if err != nil {
			return nil, err
		}
		return registrar, nil
	}
}
