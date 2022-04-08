package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/disk"
)

func main() {
	// 获取文件夹的使用信息
	info, _ := disk.Usage("C:/projects")

	// 将信息json格式化
	data, _ := json.MarshalIndent(info, "", "  ")

	// 打印信息
	fmt.Println(string(data))
}
