package zdpgo_psutil

import (
	"fmt"
	"testing"
	"time"
)

/*
@Time : 2022/5/13 14:26
@Author : 张大鹏
@File : get_test.go
@Software: Goland2021.3.1
@Description: get类型方法的测试
*/

// 测试获取进程占用的内存信息
func TestPsutil_GetThreadMemoryUsage(t *testing.T) {
	p := getPsutil()
	for i := 0; i < 100; i++ {
		fmt.Println(p.GetThreadMemoryUsage().M)
	}
}

// 测试获取进程占用的CPU信息
func TestPsutil_GetThreadCpuInfo(t *testing.T) {
	p := getPsutil()

	// 执行多个耗时很长的计算
	for i := 0; i < 100; i++ {
		go func() {
			sum := 0
			for j := 0; j < 10000000000; j++ {
				sum += j
			}
			fmt.Println(sum)
		}()
	}

	// 监测CPU占用
	for i := 0; i < 10; i++ {
		info, err := p.GetThreadCpuInfo()
		if err != nil {
			panic(err)
		}
		t.Log(info.CpuNumber, info.AllCpuPercent, info.SingleCpuPercent)
		time.Sleep(time.Second)
	}
}

// 测试获取本机IP
func TestIP_LocalIP(t *testing.T) {
	pu := getPsutil()

	// 获取ip
	ip, err := pu.GetLocalIP()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip)

	// 遍历10次
	for j := 0; j < 10; j++ {
		ip1, err := pu.GetLocalIP()
		if err != nil {
			t.Error(err)
		}
		fmt.Println(ip1)
		if ip != ip1 {
			t.Error("not equal")
		}
	}
}
