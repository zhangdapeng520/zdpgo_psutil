package zdpgo_psutil

// BaseInfo 基础信息
type BaseInfo struct {
	MemoryUsedPercent float64 `json:"memory_used_percent"` // 内存使用百分比
	MemoryTotal       uint64  `json:"memory_total"`        // 总内存
	MemoryFree        uint64  `json:"memory_free"`         // 可用内存
	CpuUsedPercent    float64 `json:"cpu_used_percent"`    // CPU使用百分比
	Platform          string  `json:"platform"`            // 平台
	Family            string  `json:"family"`              // 类别
	Version           string  `json:"version"`             // 版本
	CpuNum            int     `json:"cpu_num"`             // CPU数量
}

// ProcessCpuInfo CPU 信息
type ProcessCpuInfo struct {
	AllCpuPercent        float64 `json:"all_cpu_percent"`        // 占用所有CPU的百分比
	SingleCpuPercent     float64 `json:"single_cpu_percent"`     // 占用所有单个CPU的百分比
	ProcessMemoryPercent float32 `json:"process_memory_percent"` // 进程占用内存百分比
	ThreadNumber         int     `json:"thread_number"`          // 创建的线程数量
	CpuNumber            int     `json:"cpu_number"`             // CPU的数量
	GoroutineNumber      int     `json:"goroutine_number"`       // 创建的goroutine数量
}

// UsageInfo 使用信息
type UsageInfo struct {
	G uint64 `json:"gb"` // 多少gb
	M uint64 `json:"mb"` // 多少mb
	K uint64 `json:"kb"` // 多少kb
	B uint64 `json:"b"`  // 多少b
}

// NetworkInfo 网卡信息
type NetworkInfo map[string]RateInfo

// RateInfo 网卡速率信息
type RateInfo struct {
	UpRate   float32  `json:"up_rate"`   // 上行速率
	DownRate float32  `json:"down_rate"` // 下行速率
	Ip       []string `json:"ip"`
	Ipv6     []string `json:"ipv6"`
}

// IOInfo 读写信息
type IOInfo struct {
	SendBytes    uint64 `json:"send_bytes"`    // 发送字节数
	ReceiveBytes uint64 `json:"receive_bytes"` // 接收字节数
}

// NetworkIOInfo 网卡的IO信息
type NetworkIOInfo map[string]IOInfo
