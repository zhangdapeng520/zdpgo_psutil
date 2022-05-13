package zdpgo_psutil

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
