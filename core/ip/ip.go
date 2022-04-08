package ip

import "net"

type IP struct {
}

func NewIP() *IP {
	i := IP{}

	return &i
}

func (i *IP) LocalIP() (ip string, err error) {
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

// GetNetworkType 获取在网络地址中的IP的分类，返回IP地址和IP分类类型
func (i *IP) GetNetworkType(netWorkAddr string) (string, string) {
	return GetIpInNetworkType(netWorkAddr)
}
