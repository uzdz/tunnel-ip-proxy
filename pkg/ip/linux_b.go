package ip

import (
	"ip-proxy/pkg/utils"
	"os/exec"
)

func linuxDisconnect() {
	cmd := exec.Command("bash", "-C", "/sbin/ifdown", "ppp0")
	utils.Output(cmd, "UTF-8")
}

func linuxConnect() {
	cmd := exec.Command("bash", "-C", "/sbin/ifup", "ppp0")
	utils.Output(cmd, "UTF-8")
}
