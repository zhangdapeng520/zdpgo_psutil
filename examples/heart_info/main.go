package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/cpu"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/mem"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/net"
	"strings"
	"time"
)

// getIpInNetworkType 获取在网络地址中的IP的分类。
// 网络地址示例1：fe80::4d90:6461:8e7e:2870/64
// 网络地址示例2：10.1.3.12/24
// @return 返回(IP地址，IP类型)
func getIpInNetworkType(netWorkAddr string) (string, string) {
	// 默认是未知类型
	defaultType := "unknown"

	// 校验参数的合法性
	if netWorkAddr == "" || !strings.Contains(netWorkAddr, "/") {
		return "", defaultType
	}

	// 切割地址
	splitAddr := strings.Split(netWorkAddr, "/")
	if len(splitAddr) == 0 {
		return "", defaultType
	}

	// 取第一个参数
	ipAddr := splitAddr[0]

	// 判断类型
	if strings.Contains(ipAddr, ":") && strings.Count(ipAddr, ":") >= 2 {
		return ipAddr, "ipv6"
	} else if strings.Contains(ipAddr, ".") && strings.Count(ipAddr, ".") == 3 {
		return ipAddr, "ipv4"
	}

	// 既不是ipv4，也不是ipv6
	return ipAddr, defaultType
}

// HeartData 获取要发送的心跳数据
// @param version 版本号
// @param system 操作系统
// @param uuid 唯一编号
func HeartData(version string, system string, uuid, hostName string) (data map[string]interface{}, err error) {
	// 心跳数据
	heartbeatData := map[string]interface{}{
		"version":   version,                         // 版本号
		"ifs":       make(map[string]interface{}, 5), // 网卡速率
		"mem_usage": 0,                               // 内存使用率
		"cpu_usage": 0,                               // cpu使用率
		"system":    system,
		"uuid":      uuid,
		"host_name": hostName,
	}

	// 创建网卡状态列表
	ifStats, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// 网卡信息
	ifs := make(map[string]interface{}, 5)

	// 获取网卡速率
	ifRates, err := GetNetworkCardRate()
	if err != nil {
		return nil, err
	}

	// 遍历网卡速率
	for _, if_ := range ifStats {
		// 对网卡进行分组
		ipv4 := make([]string, 0, 1)
		ipv6 := make([]string, 0, 1)
		for _, addr := range if_.Addrs {
			// 获取网络地址的IP地址和IP类型
			ipAddr, ipType := getIpInNetworkType(addr.Addr)

			// 进行分类
			if ipType == "ipv4" {
				ipv4 = append(ipv4, ipAddr)
			} else if ipType == "ipv6" {
				ipv6 = append(ipv6, ipAddr)
			} else {
				fmt.Println("未知的ip地址类型：", addr.Addr, ipAddr, ipType)
			}
		}

		// 封装分组数据
		ifs[if_.Name] = map[string]interface{}{
			"ip":        ipv4,
			"ipv6":      ipv6,
			"up_rate":   ifRates[if_.Name]["up_rate"],   // 上行速率
			"down_rate": ifRates[if_.Name]["down_rate"], // 下行速率
		}
	}

	// 封装心跳中的网卡速率
	heartbeatData["ifs"] = ifs

	// 内存使用率
	memoryInfo, err := mem.VirtualMemory()
	heartbeatData["mem_usage"] = memoryInfo.UsedPercent
	if err != nil {
		return nil, err
	}

	// cpu使用率
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}
	heartbeatData["cpu_usage"] = fmt.Sprintf("%f", cpuPercent[0])

	// 返回心跳数据
	return heartbeatData, nil
}

// GetNetworkCardRate 获取网卡的上行速率和下行速率
func GetNetworkCardRate() (map[string]map[string]float32, error) {
	// 获取网卡网速信息
	getIfIO := func() (map[string]map[string]uint64, error) {
		IOs := make(map[string]map[string]uint64, 5)

		// 获取网卡个数
		ifs, err := net.IOCounters(true)
		if err != nil {
			return nil, err
		}

		// 遍历网卡
		for _, if_ := range ifs {
			// 每个网卡的发送字节数和接收字节数
			IOs[if_.Name] = map[string]uint64{
				"bytesSent": if_.BytesSent,
				"bytesRecv": if_.BytesRecv,
			}
		}

		// 返回网卡信息
		return IOs, nil
	}

	// 第一次获取网卡信息
	IO1, err := getIfIO()
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Second)

	// 第二次获取网卡信息
	IO2, err := getIfIO()
	if err != nil {
		return nil, err
	}

	// 计算网卡速率信息
	rates := make(map[string]map[string]float32)
	for ifName := range IO1 {
		// 下行速率：指的是本机的接收速率
		// 网卡下行速率 = （第二次发送字节数 - 第一次发送字节数） / 1024
		upRate := float32(IO2[ifName]["bytesSent"]-IO1[ifName]["bytesSent"]) / 1024

		// 上行速率：指的是本机的发送速率
		// 网卡上行速率 = （第二次接收字节数 - 第一次接收字节数） / 1024
		downRate := float32(IO2[ifName]["bytesRecv"]-IO1[ifName]["bytesRecv"]) / 1024

		// 网卡速率信息
		rates[ifName] = map[string]float32{
			"down_rate": downRate,
			"up_rate":   upRate,
		}
	}

	// 返回网卡速率信息
	return rates, nil
}

func main() {
	rate, err := GetNetworkCardRate()
	fmt.Println("获取网卡速率：", rate, err)
	heartbeatData, err := HeartData("1", "systemInfo", "1", "hostName")
	fmt.Println("心跳数据：", heartbeatData, err)

	// 打印心跳数据
	for k, v := range heartbeatData {
		fmt.Println(k, v)
	}
}
