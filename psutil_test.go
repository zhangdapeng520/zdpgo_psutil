package zdpgo_psutil

import "testing"

func getPsutil() *Psutil {
	return New()
}

func TestPsutil_ip_local(t *testing.T) {
	p := getPsutil()

	// 获取本机IP
	ip, err := p.IP.LocalIP()
	if err != nil {
		t.Error(err)
	}

	// 多次获取，相同
	for j := 0; j < 10; j++ {
		ip1, err := p.IP.LocalIP()
		if err != nil {
			t.Error(err)
		}
		if ip1 != ip {
			t.Error("not equal")
		}
	}
}
