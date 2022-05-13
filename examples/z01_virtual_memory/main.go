package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/mem"
)

func main() {
	// 创建虚拟内存对象
	v, _ := mem.VirtualMemory()

	// 查看信息
	fmt.Printf("总内存: %v, 可用内存:%v, 使用百分比:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// 查看json信息
	fmt.Println(v)
}
