package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/disk"
)

func main() {
	// 获取磁盘信息
	mapStat, _ := disk.IOCounters()

	// 遍历每个磁盘
	for name, stat := range mapStat {
		// 打印磁盘名称
		fmt.Println(name)

		// 磁盘信息json格式化
		data, _ := json.MarshalIndent(stat, "", "  ")

		// 打印磁盘信息
		fmt.Println(string(data))
	}
}
