package zdpgo_psutil

import "github.com/zhangdapeng520/zdpgo_psutil/core/ip"

type Psutil struct {
	IP *ip.IP // 操作IP的核心对象
}

func New() *Psutil {
	p := Psutil{}

	// 实例化对象
	p.IP = ip.NewIP()

	return &p
}
