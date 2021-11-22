package report

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/config/po"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func Report(rq po.Rq) error {

	err, rp := ReHttpProtoPost(rq, config.RemoteUrl)

	if err != nil {
		return fmt.Errorf("[上报] 服务器响应错误，错误信息：" + err.Error())
	}

	if rp.Code == 200 {
		// log.Println("[上报] VPS信息已同步服务器！")
		config.Fill(rp)
		return nil
	}

	return fmt.Errorf("[上报] 服务器响应码错误，：%d \n", rp.Code)
}

func ReHttpProtoPost(rq po.Rq, url string) (err error, v po.Rp) {

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
	rp := &po.Rp{}
	proto.Unmarshal(respBytes, rp)

	return nil, *rp
}
