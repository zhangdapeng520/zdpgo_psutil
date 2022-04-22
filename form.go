package zdpgo_psutil

// ProcessCpuInfo CPU 信息
type ProcessCpuInfo struct {
	AllCpuPercent        float64 `json:"all_cpu_percent"`        // 占用所有CPU的百分比
	SingleCpuPercent     float64 `json:"single_cpu_percent"`     // 占用所有单个CPU的百分比
	ProcessMemoryPercent float32 `json:"process_memory_percent"` // 进程占用内存百分比
	ThredsNumber         int     `json:"threads_number"`         // 创建的线程数量
	GoroutineNumber      int     `json:"goroutine_number"`       // 创建的goroutine数量
}
