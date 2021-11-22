package config

import (
	"encoding/csv"
	"os"
	"sync"
	"time"
)

var ProxyRun bool

var DownloadNow = false

var StartMutex sync.Mutex

var ShutdownMutex sync.Mutex

var OsDesktop = "C:\\Users\\Administrator\\Desktop\\"

var LogPath = OsDesktop + "sys.log"

var RequestLogFolder = OsDesktop + "request\\"

var ServerHost = "www.xxx.com:25435"

// 数据上报服务端地址
var RemoteUrl = "http://" + ServerHost + "/xxx-boot/ip/report/standard"

// 当前服务器时间
var TimeUrl = "http://" + ServerHost + "/xxx-boot/ip/time"

// 检查新版本服务端地址
var FileRequestUrl = "http://" + ServerHost + "/xxx-boot/ip/client/version"

var ReportServerHost = "www.xxx.com:25433"

// 请求数据上报服务端地址
var ReportRequestUrl = "http://" + ReportServerHost + "/e"

// 代理服务器授权Key
var ProxyAuthorizationKey = "Proxy-Authorization"

// 代理隧道随机转发授权KEY
var XTunnelForwardedFor = "X-Tunnel-Forwarded-For"

// 隧道代理服务需添加如下Header
var TunnelXForKey = "X-Tunnel-Key"
var TunnelXForValue = "266aXa2WWe"

// 当前关机重试次数（3次即重启）
var RestartState = 0

// 命令队列
var CommandChan = make(chan string, 10)

// 机器编号
var Number string

// 公网IP
var Ip string

// 端口号
var Port int64

// 最近一次拨号时间戳
var DialTime int64

// 拨号账号
var IpAC string

// 拨号密码
var IpPW string

// 拨号设备名称
var NetDeviceName string

// 访问数据文件清除（天）
var ClearFileInterval = -90

// 当前输出的文件流
var OutFile *os.File

// 当前输出的CSV文件流
var OutCsvFile *csv.Writer

// 日志请求数据体
type WRequest struct {
	DeviceNum string
	Uid       string
	OriginIp  string
	ProxyIp   string
	RemoteIp  string
	EventTime int64
}

// 所有HTTP转发请求忽略的请求头
var IgnoreHeaderMap = []string{
	XTunnelForwardedFor,
	ProxyAuthorizationKey,
	TunnelXForKey,
}

var TodayIpMap = make(map[string]bool)

// 启动代理服务器
var StartCommand = "start"

// 启动代理服务器并进行IP拨号
var StartWithIpChangeCommand = "startWithIpChange"

// 关闭代理服务器
var ShutdownCommand = "shutdown"

// 关闭代理服务器后启动
var ShutdownThenStartCommand = "shutdownThenStart"

// 定义一个任务触发器
var DialTicker = time.NewTicker(time.Second * time.Duration(9999))
