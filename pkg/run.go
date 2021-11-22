package pkg

import (
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/config/po"
	"ip-proxy/pkg/reload"
	"ip-proxy/pkg/server"
	"ip-proxy/pkg/utils"
	"log"
	"strconv"
	"time"
)

func start(command string) {
	config.StartMutex.Lock()
	defer config.StartMutex.Unlock()

	if config.ProxyRun {
		// 已经启动
		return
	}

	//if config.RestartState >= 3 {
	//	// 关机并重启
	//	log.Println("触发重启限制，正在重启VPS...")
	//	cmd := exec.Command("cmd", "/C", "shutdown", "-r", "-t", "0")
	//	utils.Output(cmd, "GB18030")
	//	return
	//}

	//var ipStatus int
	//switch command {
	//case config.StartCommand:
	//	ipStatus = 0
	//case config.StartWithIpChangeCommand:
	//	ipStatus = 1
	//}

	//currentIp, err := ip.ExternalIP()
	//if ipStatus == 1 || currentIp == "" || err != nil {
	//
	//	if ipStatus == 1 {
	//		time.Sleep(time.Duration(utils.RangeRand(5, 15)) * time.Second)
	//	}
	//
	//	//currentIp = ip.Dial()
	//}

	//if currentIp != "" {
	// 重制重启次数
	config.RestartState = 0

	//config.Ip = currentIp
	config.Port = utils.RangeRand(30000, 40000)
	config.DialTime = time.Now().Unix()

	rq := po.Rq{}
	rq.DeviceId = config.Number
	rq.Ip = config.Ip
	rq.Port = config.Port
	rq.DialTime = config.DialTime

	//err := report.Report(rq)
	//if err != nil {
	//	log.Printf("WARING 代理服务器上报失败！\n" + err.Error())
	//
	//	time.Sleep(time.Second * time.Duration(5))
	//	config.CommandChan <- command
	//	return
	//}

	if config.TodayIpMap[config.Ip] {
		if config.IpRepeat == 1 {
			config.CommandChan <- config.StartWithIpChangeCommand
			return
		}
	} else {
		config.TodayIpMap[config.Ip] = true
	}

	port := ":" + strconv.Itoa(int(config.Port))
	server.P = &server.S{}
	go server.P.RunServer(port)

	// 拨号任务重新计时
	var waitTime int64
	if config.IpInterval == 0 {
		waitTime = 30
	} else {
		waitTime = config.IpInterval
	}

	config.DialTicker.Reset(time.Second * time.Duration(waitTime))
	config.ProxyRun = true
	// 重新写入热加载文件
	go reload.WriteReloadCsv()
	//} else {
	//	config.RestartState = config.RestartState + 1
	//	config.CommandChan <- config.StartWithIpChangeCommand
	//}
}

func shutdown() {
	config.ShutdownMutex.Lock()
	defer config.ShutdownMutex.Unlock()

	if server.P != nil {
		b := server.P.Shutdown()

		if b {
			server.P = nil
		}
	}

	config.DialTicker.Stop()
	config.ProxyRun = false
	log.Println("代理服务器已成功关闭...")
}

func CommandServerThread() {

	for {
		command := <-config.CommandChan

		for {
			// 如果下载状态下，阻塞等待下载完成
			if config.DownloadNow {
				time.Sleep(time.Second * time.Duration(5))
			} else {
				break
			}
		}

		switch command {
		case config.StartCommand:
			start(command)
		case config.StartWithIpChangeCommand:
			start(command)
		case config.ShutdownCommand:
			shutdown()
		case config.ShutdownThenStartCommand:
			shutdown()
			start(config.StartWithIpChangeCommand)
		}
	}
}
