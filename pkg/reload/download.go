package reload

import (
	"bytes"
	"io/ioutil"
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/config/po"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
)

// 定时检查，是否有新版本客户端
func ReportVAndDownload() {

	for {
		if config.ProxyRun && config.DownloadNow == false {

			rq := po.FRq{}
			rq.DeviceId = config.Number
			rq.Version = config.Version
			rq.Model = config.ClientModel
			err, v := FileHttpProtoPost(rq, config.FileRequestUrl)

			if err != nil {
				log.Println("[版本检查] 服务器响应错误，错误信息：" + err.Error())
			} else {
				if v.Code == 200 && v.Need {
					config.DownloadNow = true
					err := DownloadExec(v.FilePath)
					config.DownloadNow = false

					if err == nil {
						RunProxy()
					} else {
						log.Printf("[版本检查] 客户端下载错误，错误信息：" + err.Error())
					}
				}
			}
		}

		time.Sleep(time.Second * time.Duration(10))
	}
}

func FileHttpProtoPost(rq po.FRq, url string) (err error, v po.FRp) {

	bytesData, err := proto.Marshal(&rq)
	if err != nil {
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/x-protobuf")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	rp := &po.FRp{}
	proto.Unmarshal(respBytes, rp)

	return nil, *rp
}
