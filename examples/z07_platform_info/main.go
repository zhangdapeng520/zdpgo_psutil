package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/host"
)

func main() {
	// 获取内核版本
	version, _ := host.KernelVersion()
	fmt.Println(version)

	// 获取平台信息
	platform, family, version, _ := host.PlatformInformation()
	fmt.Println("platform:", platform)
	fmt.Println("family:", family)
	fmt.Println("version:", version)
}
