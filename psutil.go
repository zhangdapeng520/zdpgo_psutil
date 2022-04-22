package zdpgo_psutil

import (
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/zhangdapeng520/zdpgo_psutil/core/ip"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/process"
)

type Psutil struct {
	IP *ip.IP // 操作IP的核心对象
}

func New() *Psutil {
	p := Psutil{}

	// 实例化对象
	p.IP = ip.NewIP()

	return &p
}

func (pu *Psutil) GetProcessCpuInfo() (info ProcessCpuInfo, err error) {
	// 创建一个进程
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return
	}

	// 进程的CPU使用率需要通过计算指定时间内的进程的CPU使用时间变化计算出来
	info.AllCpuPercent, err = p.Percent(time.Second)
	if err != nil {
		return
	}

	// 上面返回的是占所有CPU时间的比例，如果想更直观的看占比，可以算一下占单个核心的比例。
	info.SingleCpuPercent = info.AllCpuPercent / float64(runtime.NumCPU())

	//获取进程占用内存的比例
	info.ProcessMemoryPercent, err = p.MemoryPercent()
	if err != nil {
		return
	}

	// 创建的线程数
	info.ThredsNumber = pprof.Lookup("threadcreate").Count()

	// Goroutine数
	info.GoroutineNumber = runtime.NumGoroutine()
	return
}
