package auth

import (
	"encoding/base64"
	"strings"
)

func HttpAuthorization(userPw, ip string) string {

	var uid string

	if ip != "" {
		uid = IpAuth(ip)
	}

	if uid == "" && userPw != "" {

		basic := strings.Fields(userPw)
		if len(basic) == 2 {
			user, err := base64.StdEncoding.DecodeString(basic[1])
			if err == nil {
				uid = UserAuth(string(user))
			}
		}
	}

	return uid
}
