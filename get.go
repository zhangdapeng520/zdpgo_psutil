package zdpgo_psutil

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/cpu"
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/host"
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/mem"
	psnet "github.com/zhangdapeng520/zdpgo_psutil/gopsutil/net"
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/process"
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
		p.Log.Error("获取所有IP地址失败", "error", err, "address", address)
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

// GetNetworkInfo 获取网卡信息
func (p *Psutil) GetNetworkInfo() (networkInfo NetworkInfo) {
	// 创建网卡状态列表
	ifStats, err := psnet.Interfaces()
	if err != nil {
		p.Log.Error("创建网卡状态列表失败", "error", err)
		return
	}

	// 网卡信息
	networkInfo = NetworkInfo{}

	// 获取网卡速率
	ifRates, err := p.GetNetworkRate()
	if err != nil {
		p.Log.Error("获取网卡上行速率和下行速率失败", "error", err)
		return
	}

	// 遍历网卡速率
	for _, ifState := range ifStats {
		// 对网卡进行分组
		var ipv4 []string
		var ipv6 []string
		for _, addr := range ifState.Addrs {
			// 获取网络地址的IP地址和IP类型
			ipAddr, ipType := p.GetNetworkIpType(addr.Addr)

			// 进行分类
			if ipType == "ipv4" {
				ipv4 = append(ipv4, ipAddr)
			} else if ipType == "ipv6" {
				ipv6 = append(ipv6, ipAddr)
			} else {
				p.Log.Debug("未知的ip地址类型", "addr", addr.Addr, "ipAddr", ipAddr, "ipType", ipType)
			}
		}

		// 封装分组数据
		networkInfo[ifState.Name] = RateInfo{
			Ip:       ipv4,
			Ipv6:     ipv6,
			UpRate:   ifRates[ifState.Name].UpRate,   // 上行速率
			DownRate: ifRates[ifState.Name].DownRate, // 下行速率
		}
	}

	p.Log.Debug("获取网卡信息成功", "networkInfo", networkInfo)
	return
}

// GetNetworkRate 获取网卡速率信息
func (p *Psutil) GetNetworkRate() (rates NetworkInfo, err error) {
	// 获取网卡网速信息
	getIfIO := func() (ioInfo NetworkIOInfo, err error) {
		// {网卡名称：{xx：xx}}
		ioInfo = NetworkIOInfo{}

		// 获取网卡个数
		ifs, err := psnet.IOCounters(true)
		if err != nil {
			p.Log.Error("获取网卡个数失败", "error", err, "ifs", ifs)
			return nil, err
		}

		// 遍历网卡
		for _, if_ := range ifs {
			// 每个网卡的发送字节数和接收字节数
			ioInfo[if_.Name] = IOInfo{
				SendBytes:    if_.BytesSent,
				ReceiveBytes: if_.BytesRecv,
			}
		}

		// 返回网卡信息
		return
	}

	var (
		IO1 NetworkIOInfo
		IO2 NetworkIOInfo
	)

	// 第一次获取网卡信息
	IO1, err = getIfIO()
	if err != nil {
		p.Log.Error("获取网卡信息失败", "error", err)
		return
	}
	time.Sleep(time.Second)

	// 第二次获取网卡信息
	IO2, err = getIfIO()
	if err != nil {
		p.Log.Error("获取网卡信息失败", "error", err)
		return
	}

	// 计算网卡速率信息
	rates = NetworkInfo{}
	for ifName := range IO1 {
		// 下行速率：指的是本机的接收速率
		// 网卡下行速率 = （第二次发送字节数 - 第一次发送字节数） / 1024
		downRate := float32(IO2[ifName].ReceiveBytes-IO1[ifName].ReceiveBytes) / 1024

		// 上行速率：指的是本机的发送速率
		// 网卡上行速率 = （第二次接收字节数 - 第一次接收字节数） / 1024
		upRate := float32(IO2[ifName].SendBytes-IO1[ifName].SendBytes) / 1024

		// 网卡速率信息
		rates[ifName] = RateInfo{
			UpRate:   upRate,
			DownRate: downRate,
			Ip:       nil,
			Ipv6:     nil,
		}
	}

	// 返回网卡速率信息
	return
}

// GetBaseInfo 获取基本信息
func (p *Psutil) GetBaseInfo() (info BaseInfo, err error) {
	memoryInfo, err := mem.VirtualMemory()
	if err != nil {
		p.Log.Error("获取内存信息失败", "error", err)
		return
	}

	// 获取内存使用率
	info.MemoryTotal = memoryInfo.Total
	info.MemoryFree = memoryInfo.Free
	info.MemoryUsedPercent = memoryInfo.UsedPercent

	// cpu使用率
	cpuPercentArr, err := cpu.Percent(time.Second, false)
	if err != nil {
		p.Log.Error("获取cpu使用率失败", "error", err)
		return
	}
	var (
		cpuPercent float64
		total      float64
	)
	for _, cpuP := range cpuPercentArr {
		total += cpuP
	}
	cpuPercent = total / float64(len(cpuPercentArr))
	info.CpuUsedPercent = cpuPercent

	// 平台
	platform, family, version, err := host.PlatformInformation()
	if err != nil {
		p.Log.Error("获取平台信息失败", "error", err)
		return
	}
	info.Platform = platform
	info.Family = family
	info.Version = version

	// CPU数量
	info.CpuNum = runtime.NumCPU()

	// 返回
	p.Log.Debug("获取基本信息成功", "info", info)
	return
}
