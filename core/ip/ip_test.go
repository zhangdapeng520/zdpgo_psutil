package ip

import "testing"

func getIP() *IP {
	return NewIP()
}

// 测试获取本机IP
func TestIP_LocalIP(t *testing.T) {
	i := getIP()

	// 获取ip
	ip, err := i.LocalIP()
	if err != nil {
		t.Error(err)
	}

	// 遍历10次
	for j := 0; j < 10; j++ {
		ip1, err := i.LocalIP()
		if err != nil {
			t.Error(err)
		}
		if ip != ip1 {
			t.Error("not equal")
		}
	}
}
