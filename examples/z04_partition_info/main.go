package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/disk"
)

func main() {
	// 获取分区信息
	infos, _ := disk.Partitions(false)

	// 遍历分区信息
	for _, info := range infos {
		// json格式化分区信息
		data, _ := json.MarshalIndent(info, "", "  ")

		// 打印分区信息
		fmt.Println(string(data))
	}
}
