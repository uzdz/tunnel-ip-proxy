package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"unsafe"
)

func SamplePost(u map[string]interface{}, url string) (err error, v string) {

	bytesData, err := json.Marshal(u)
	if err != nil {
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))

	return nil, *str
}

func SampleGet(path string) (err error, value string) {
	resp, err := http.Get(path)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return nil, string(b)
}
