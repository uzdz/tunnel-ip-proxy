package csv

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/utils"
	"log"
	"os"
	"strings"
	"time"
)

func InitRequestUpload() {
	utils.PathExists(config.RequestLogFolder)

	config.OutFile, config.OutCsvFile = CreateTodayFile()

	// 每天创建新文件
	go TimeToNewFile()

	// 删除指定天前的数据
	go TimeToClear()
}

func TimeToNewFile() {
	for {
		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))

		<-t.C

		config.OutCsvFile.Flush()
		config.OutFile.Close()
		config.OutFile, config.OutCsvFile = CreateTodayFile()
	}
}

func CreateTodayFile() (*os.File, *csv.Writer) {
	now := time.Now()
	todayFile := now.Format("2006-01-02") + ".csv"
	f, err := os.OpenFile(config.RequestLogFolder+todayFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	return f, csv.NewWriter(f)
}

func TimeToClear() {
	for {
		now := time.Now()
		next := now.Add(time.Hour * 24)
		t := time.NewTimer(next.Sub(now))

		<-t.C

		ClearFile()
	}
}

func ClearFile() {

	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	clearEndTime := newTime.AddDate(0, 0, config.ClearFileInterval)

	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(config.RequestLogFolder)
	if err != nil {
		log.Fatal(err)
	}

	var shouldClearFile = make([]string, 0, 10)

	for i := range fileInfoList {
		name := fileInfoList[i].Name()

		config := strings.Split(strings.Trim(name, " "), ".")
		if len(config) != 2 {
			continue
		}

		old, err := time.ParseInLocation("2006-01-02", config[0], time.Local)

		if err != nil {
			continue
		}

		if old.Before(clearEndTime) {
			shouldClearFile = append(shouldClearFile, name)
		}
	}

	for _, name := range shouldClearFile {
		_ = os.Remove(config.OsDesktop + name)
	}
}
