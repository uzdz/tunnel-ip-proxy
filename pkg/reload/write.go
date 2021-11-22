package reload

import (
	"encoding/csv"
	"ip-proxy/pkg/config"
	"ip-proxy/pkg/utils"
	"os"
)

// 写入自动配置文件
func WriteReloadCsv() {

	exist := utils.FileExist(reloadName)

	if exist {
		os.Remove(reloadName)
	}

	//创建文件
	f, err := os.Create(reloadName)
	if err != nil {
		return
	}

	defer f.Close()

	w := csv.NewWriter(f)
	data := [][]string{
		{config.Number, config.IpAC, config.IpPW, config.NetDeviceName},
	}

	w.WriteAll(data)
	w.Flush()
}
