package report

import (
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/utils"
)

func Load(w config.WRequest) {

	maps := make(map[string]interface{})
	maps["deviceNum"] = w.DeviceNum
	maps["uId"] = w.Uid
	maps["originIp"] = w.OriginIp
	maps["proxyIp"] = w.ProxyIp
	maps["remoteIp"] = w.RemoteIp
	maps["eventTime"] = w.EventTime

	_, _ = utils.SamplePost(maps, config.ReportRequestUrl)
	//if err != nil {
	//	log.Println("[上报] 用户请求数据上报服务器失败：" + err.Error())
	//} else {
	//	log.Println("[上报] 用户请求数据上报服务器结果：" + resp)
	//}
}
