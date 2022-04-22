package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_psutil"
)

func main() {
	p := zdpgo_psutil.New()
	info, err := p.GetProcessCpuInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println("进程1s内占用所有CPU比例：", info.AllCpuPercent)
	fmt.Println("进程1s内占用单个CPU比例：", info.SingleCpuPercent)
	fmt.Println("进程占用内存比例：", info.ProcessMemoryPercent)
	fmt.Println("创建的线程数：", info.ThredsNumber)
	fmt.Println("创建的goroutine数：", info.GoroutineNumber)
}
