package reload

import (
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/utils"
	"os"
	"os/exec"
	"time"
)

// vbs重启脚本
var vbsFileName = "rp.bat"

// vbs内容
var content = "@echo off \n ping -n 3 127.0.0.1>nul" +
	" \n taskkill /f /t /im " + fileLikeName +
	" \n ping -n 2 127.0.0.1>nul " +
	" \n del " + config.OsDesktop + fileLikeName +
	" \n ren " + config.OsDesktop + newFileName + " " + programmerFileName +
	" \n start " + config.OsDesktop + programmerFileName +
	" \n exit"

// 最新版本文件名
var newFileName = "new.exe"

// 模糊名称
var fileLikeName = "ip-proxy*"

// 运行版本文件名
var programmerFileName = "ip-proxy.exe"

var vbsRealPath = config.OsDesktop + vbsFileName

func UpdateVbsInit() {

	exist := utils.FileExist(vbsRealPath)

	if exist {
		os.Remove(vbsRealPath)
	}

	// 写入版本更新文件bat批处理
	utils.WriteToFile(vbsRealPath, content)
}

// 关闭当前旧客户端，并启动新客户端
func RunProxy() {
	exist := utils.FileExist(vbsRealPath)

	if exist {
		config.CommandChan <- config.ShutdownCommand

		for {
			if config.ProxyRun == false {
				cmd := exec.Command("cmd.exe", "/c", "start "+vbsRealPath)
				utils.Output(cmd, "GB18030")
				return
			}

			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}
