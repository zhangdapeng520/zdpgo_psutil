package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/process"
)

func main() {
	// process可用于获取系统当前运行的进程信息，创建新进程，对进程进行一些操作等
	var rootProcess *process.Process

	// 获取所有进程
	processes, _ := process.Processes()
	for _, p := range processes {
		if p.Pid == 0 {
			rootProcess = p
			break
		}
	}

	// 查看主进程
	fmt.Println(rootProcess)

	// 查看子进程
	fmt.Println("children:")
	children, _ := rootProcess.Children()
	for _, p := range children {
		fmt.Println(p)
	}
}
