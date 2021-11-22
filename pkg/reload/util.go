package reload

import (
	"encoding/csv"
	"fmt"
	"io"
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/utils"
	"log"
	"net/http"
	"os"
)

func DownloadExec(path string) error {

	var newFilePath = config.OsDesktop + newFileName

	exist := utils.FileExist(newFilePath)
	if exist {
		os.Remove(newFilePath)
	}

	// Get the data
	resp, err := http.Get(path)
	if err != nil {
		return fmt.Errorf("download failed")
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(newFilePath)
	if err != nil {
		return fmt.Errorf("create file failed")
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("save file failed")
	}

	return nil
}

func ReadOneLineWithCSV(fileName string) []string {
	//准备读取文件
	fs, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
		return nil
	}

	defer fs.Close()

	//针对大文件，一行一行的读取文件
	r := csv.NewReader(fs)
	row, err := r.Read()

	if err != nil && err != io.EOF {
		log.Fatalf("can not read, err is %+v", err)
		return nil
	}

	if err == io.EOF {
		return nil
	}

	return row
}
