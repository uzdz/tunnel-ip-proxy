package ip

import (
	"fmt"
	"runtime"
)

func Dial() string {
	// log.Println("\n开始进行PPPOE拨号：")
	cIp, err := ModIp("dct")
	// log.Println("\nPPPOE拨号结束：" + cIp)

	if err != nil {
		return ""
	}

	return cIp
}

func ModIp(do string) (string, error) {

	var connect func()
	var disconnect func()

	switch runtime.GOOS {
	case "windows":
		connect = winConnect
		disconnect = winDisconnect
	case "linux":
		connect = linuxConnect
		disconnect = linuxDisconnect
	default:
		return "", fmt.Errorf("the platform does not support dialing! %s", runtime.GOOS)
	}

	switch do {
	case "ct":
		connect()
	case "dc":
		disconnect()
	case "dct":
		disconnect()
		connect()
	default:
		return "", fmt.Errorf("I don't know what I want? %s", do)
	}

	return ExternalIP()
}
