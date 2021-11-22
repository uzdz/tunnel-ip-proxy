package reload

import (
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/utils"
	"log"
)

var reloadName = config.OsDesktop + "reload.csv"

// 查找是否存在自动启动文件，存在则启动代理
// 规定隧道文件格式如下，csv格式，逗号分割
// 第一列：设备号
// 第二列：拨号账号
// 第三列：拨号密码
// 第四列：网络设备名
func FindReloadFileAndLoad() {

	defer func() {
		if p := recover(); p != nil {
			log.Printf("load#FindReloadFileAndLoad internal error: %v", p)
		}
	}()

	exist := utils.FileExist(reloadName)

	if exist {
		data := ReadOneLineWithCSV(reloadName)

		if data == nil || len(data) != 4 {
			return
		}

		config.Number = data[0]
		config.IpAC = data[1]
		config.IpPW = data[2]
		config.NetDeviceName = data[3]

		// 首次启动
		config.CommandChan <- config.StartCommand
	}
}
