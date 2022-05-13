package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/host"
	"time"
)

func main() {
	// 获取主机的开机时间
	timestamp, _ := host.BootTime()

	// 将时间转换为数字
	t := time.Unix(int64(timestamp), 0)

	// 格式化时间
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))
}
