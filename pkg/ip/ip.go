package ip

import (
	"fmt"
	"ip-proxy/pkg/config"
	"net"
)

func ExternalIP() (ip string, err error) {
	ips, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, ip := range ips {

		if ip.Name == config.NetDeviceName {
			adds, err := ip.Addrs()
			if err == nil {
				for _, addr := range adds {

					netIp := getIpFromAddr(addr)
					if netIp == nil {
						continue
					}
					return netIp.String(), nil
				}
			}
		}
	}
	return "", fmt.Errorf("[%s] 未找到，当前可能已断开网络！", config.NetDeviceName)
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
