package utils

import (
	"github.com/coolrc136/go-kit-micro/config"
	"net"
)

func GetOutboundIP() (net.IP, error) {
	var ip = "8.8.8.8:80"
	if len(config.C.Etcd) > 0 {
		ip = config.C.Etcd[0]
	}
	if config.C.Consul != "" {
		ip = config.C.Consul
	}
	conn, err := net.Dial("udp", ip)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, err
}
