package ip

import (
	"strings"
)

// GetIpInNetworkType 获取在网络地址中的IP的分类。
// 网络地址示例1：fe80::4d90:6461:8e7e:2870/64
// 网络地址示例2：10.1.3.12/24
// @return 返回(IP地址，IP类型)
func GetIpInNetworkType(netWorkAddr string) (string, string) {
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
