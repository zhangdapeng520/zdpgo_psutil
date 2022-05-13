package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/mem"
)

func main() {
	// 获取交换内存
	swapMemory, _ := mem.SwapMemory()

	// 打印交换内存的信息
	data, _ := json.MarshalIndent(swapMemory, "", " ")
	fmt.Println(string(data))
}
