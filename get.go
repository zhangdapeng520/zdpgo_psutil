package zdpgo_psutil

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/process"
	"net"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

/*
@Time : 2022/5/13 14:20
@Author : 张大鹏
@File : get.go
@Software: Goland2021.3.1
@Description: get类型的获取方法
*/

//GetUsageInfo 获取使用信息
func (p *Psutil) GetUsageInfo(bNum uint64) UsageInfo {
	return UsageInfo{
		G: bNum / 1024 / 1024 / 1024,
		M: bNum / 1024 / 1024,
		K: bNum / 1024,
		B: bNum,
	}
}

// GetThreadMemoryUsage 获取当前进程的内存使用信息
func (p *Psutil) GetThreadMemoryUsage() UsageInfo {
	var m runtime.MemStats

	// 监测内存
	runtime.ReadMemStats(&m)

	// 将内存转换为找
	usageInfo := p.GetUsageInfo(m.Sys)
	return usageInfo
}

func (p *Psutil) GetThreadCpuInfo() (info ProcessCpuInfo, err error) {
	// 创建一个进程
	pr, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		p.Log.Error("创建进程失败", "error", err)
		return
	}

	// 进程的CPU使用率需要通过计算指定时间内的进程的CPU使用时间变化计算出来
	info.AllCpuPercent, err = pr.Percent(time.Second)
	if err != nil {
		p.Log.Error("计算所有CPU的占用百分比失败", "error", err)
		return
	}

	// 上面返回的是占所有CPU时间的比例，如果想更直观的看占比，可以算一下占单个核心的比例。
	info.CpuNumber = runtime.NumCPU()
	info.SingleCpuPercent = info.AllCpuPercent / float64(runtime.NumCPU())

	// 获取进程占用内存的比例
	info.ProcessMemoryPercent, err = pr.MemoryPercent()
	if err != nil {
		fmt.Println("获取进程占用的内存比例失败", "error", err)
		return
	}

	// 创建的线程数
	info.ThreadNumber = pprof.Lookup("threadcreate").Count()

	// Goroutine数
	info.GoroutineNumber = runtime.NumGoroutine()
	return
}

// GetLocalIP 获取本地IP地址
func (p *Psutil) GetLocalIP() (ip string, err error) {
	// 获取所有的地址
	address, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	// 遍历地址
	for _, addr := range address {
		// 转换为IP地址
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}

		// 返回IP地址的字符串表示
		return ipAddr.IP.String(), nil
	}
	return
}

// GetNetworkIpType 获取在网络地址中的IP的分类，返回IP地址和IP分类类型
func (p *Psutil) GetNetworkIpType(netWorkAddr string) (string, string) {
	return GetIpInNetworkType(netWorkAddr)
}
