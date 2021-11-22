package config

import (
	"ip-proxy/pkg/config/po"
	"ip-proxy/pkg/limit"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mutex sync.Mutex

func Fill(rp po.Rp) {
	mutex.Lock()
	defer mutex.Unlock()

	IpInterval = rp.IpInterval
	IpRepeat = rp.IpRepeat
	NoAuth = rp.NoAuth
	UserConnectLimit = rp.ConfigData.UserConnectLimit
	AuthListWithIpMap = rp.ConfigData.AuthListWithIpMap
	AuthListWithUNameMap = rp.ConfigData.AuthListWithUNameMap

	// 限流刷新
	limitReset()
}

func limitReset() {
	for uid, _ := range limit.UserLimiter {

		if UserConnectLimit == nil || UserConnectLimit[uid] == "" {
			limit.UserLimiter[uid] = nil
		} else {
			comp(uid, UserConnectLimit[uid])
		}
	}

	if UserConnectLimit != nil {
		for uid, userLimit := range UserConnectLimit {
			comp(uid, userLimit)
		}
	}
}

func comp(uid, userLimit string) {
	config := strings.Split(strings.Trim(userLimit, " "), ":")
	if len(config) != 2 {
		return
	}

	second, err := strconv.ParseInt(config[0], 10, 64)
	if err != nil {
		return
	}

	number, ne := strconv.Atoi(config[1])
	if ne != nil {
		return
	}

	limiter := limit.UserLimiter[uid]

	if limiter == nil ||
		limiter.Number != number ||
		limiter.Duration != second {
		limit.UserLimiter[uid] = limit.New(number, time.Second*time.Duration(second), second)
	}
}
