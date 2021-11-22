package main

import (
	"ip-proxy/pkg"
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/gui"
	"ip-proxy/pkg/ticker"
)

func main() {

	// 系统时间设置
	//ticker.TimeLoad()
	//go ticker.TimeResetTicker()

	// 设置开机自启动脚本
	//go reload.AutoStart()

	// 先关闭拨号任务
	config.DialTicker.Stop()

	// 开启拨号任务监听
	go ticker.TimingDialTask()

	// 记录日志文件，并且每天24点定时清空文件
	//go ticker.CleanYesterdayLog()

	// 监控堆使用大小
	go ticker.MemoryLimit()

	// 写入更新VPS脚本（覆盖）
	//reload.UpdateVbsInit()

	// 定时检查是否有新版本更新
	//go reload.ReportVAndDownload()

	// 任务执行池
	go pkg.CommandServerThread()

	// 清除每天统计的不重复IP
	//go ticker.CleanYesterdayIP()

	// 每天生成请求记录文件，并且清除老文件
	// go csv.InitRequestUpload()

	// 查找文件，存在则自启动
	//reload.FindReloadFileAndLoad()

	// GUI窗口启动
	gui.Start()
}
