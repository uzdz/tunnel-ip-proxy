package ip

import (
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/utils"
	"os/exec"
)

func winDisconnect() {
	cmd := exec.Command("cmd", "/C", "rasdial", config.NetDeviceName, "/disconnect")
	utils.Output(cmd, "GB18030")
}

func winConnect() {
	cmd := exec.Command("cmd", "/C", "rasdial", config.NetDeviceName, config.IpAC, config.IpPW)
	utils.Output(cmd, "GB18030")
}
