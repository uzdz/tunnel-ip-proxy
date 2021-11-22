package auth

import (
	"ip-proxy/pkg/config"
)

func UserAuth(credential string) string {

	if config.AuthListWithUNameMap == nil {
		return ""
	}

	return config.AuthListWithUNameMap[credential]
}
