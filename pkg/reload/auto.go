package reload

import (
	"ip-proxy/pkg/utils"
	"log"
	"os"
	"os/exec"
	"time"
)

// 添加注册表（用于windows重启自动登陆账户），目前由运营人员手动维护
//Windows Registry Editor Version 5.00
//
//[HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Winlogon]
//"DefaultUserName"="administrator"
//"DefaultPassword"="haxiwa"
//"AutoAdminLogon"="1"

var windowsAutoStartFilePath = "C:\\Users\\Administrator\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\proxy_auto.ink"

func winDisconnect(originPath, startupPath string) {
	cmd := exec.Command("cmd", "/C", "mklink", startupPath, originPath)
	utils.Output(cmd, "GB18030")
}

func AutoStart() {
	// 等待窗口初始化，然后可视化日志。
	time.Sleep(time.Second * 3)

	defer func() {
		if p := recover(); p != nil {
			log.Println("开机自启动设置失败！请联系开发人员排查...")
		}
	}()

	exist := utils.FileExist(windowsAutoStartFilePath)

	if exist {
		log.Println("删除源绑定开机自启动文件...")
		os.Remove(windowsAutoStartFilePath)
	}

	execPath, _ := exec.LookPath(os.Args[0])
	winDisconnect(execPath, windowsAutoStartFilePath)

	log.Println("开机自启动文件已成功设置...")
}
