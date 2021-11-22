package auth

import (
	"ip-proxy/pkg/config"
	"strings"
)

func IpAuth(ip string) string {
	u := strings.Split(ip, ":")

	if config.AuthListWithIpMap != nil && len(u) == 2 {
		return config.AuthListWithIpMap[u[0]]
	}
	return ""
}
