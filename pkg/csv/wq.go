package csv

import (
	"fmt"
	"ip-proxy/pkg/config"
)

func WriteRequestCsv(w config.WRequest) {
	var data = []string{w.DeviceNum, w.OriginIp, w.ProxyIp, w.RemoteIp, fmt.Sprint(w.EventTime)}
	config.OutCsvFile.Write(data)
	config.OutCsvFile.Flush()
}
