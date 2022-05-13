package disk

import (
	"context"
	"encoding/json"
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/internal/common"
)

var invoke common.Invoker = common.Invoke{}

// UsageStat 使用状态
type UsageStat struct {
	Path              string  `json:"path"`        // 路径，传入的参数
	Fstype            string  `json:"fstype"`      // 文件系统类型
	Total             uint64  `json:"total"`       // 该分区总容量
	Free              uint64  `json:"free"`        // 空闲容量
	Used              uint64  `json:"used"`        // 已使用的容量
	UsedPercent       float64 `json:"usedPercent"` // 使用百分比
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}

// PartitionStat 分区信息
type PartitionStat struct {
	Device     string   `json:"device"`     // 分区标识，在 Windows 上即为C:这类格式
	Mountpoint string   `json:"mountpoint"` // 挂载点，即该分区的文件路径起始位置
	Fstype     string   `json:"fstype"`     // 文件系统类型，Windows 常用的有 FAT、NTFS 等，Linux 有 ext、ext2、ext3等
	Opts       []string `json:"opts"`       // 选项，与系统相关
}

// IOCountersStat 磁盘信息
type IOCountersStat struct {
	ReadCount        uint64 `json:"readCount"`
	MergedReadCount  uint64 `json:"mergedReadCount"`
	WriteCount       uint64 `json:"writeCount"`
	MergedWriteCount uint64 `json:"mergedWriteCount"`
	ReadBytes        uint64 `json:"readBytes"`
	WriteBytes       uint64 `json:"writeBytes"`
	ReadTime         uint64 `json:"readTime"`
	WriteTime        uint64 `json:"writeTime"`
	IopsInProgress   uint64 `json:"iopsInProgress"`
	IoTime           uint64 `json:"ioTime"`
	WeightedIO       uint64 `json:"weightedIO"`
	Name             string `json:"name"`
	SerialNumber     string `json:"serialNumber"`
	Label            string `json:"label"`
}

func (d UsageStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

func (d PartitionStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

func (d IOCountersStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

// Usage 获得路径path所在磁盘的使用情况，返回一个UsageStat结构
func Usage(path string) (*UsageStat, error) {
	return UsageWithContext(context.Background(), path)
}

// Partitions 返回分区信息。
// 如果all = false，只返回实际的物理分区（包括硬盘、CD-ROM、USB），忽略其它的虚拟分区。
// 如果all = true则返回所有的分区。返回类型为[]PartitionStat，每个分区对应一个PartitionStat结构
func Partitions(all bool) ([]PartitionStat, error) {
	return PartitionsWithContext(context.Background(), all)
}

// IOCounters 获取磁盘个数，返回磁盘的信息映射字典
func IOCounters(names ...string) (map[string]IOCountersStat, error) {
	return IOCountersWithContext(context.Background(), names...)
}

// SerialNumber returns Serial Number of given device or empty string
// on error. Name of device is expected, eg. /dev/sda
func SerialNumber(name string) (string, error) {
	return SerialNumberWithContext(context.Background(), name)
}

// Label returns label of given device or empty string on error.
// Name of device is expected, eg. /dev/sda
// Supports label based on devicemapper name
// See https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-block-dm
func Label(name string) (string, error) {
	return LabelWithContext(context.Background(), name)
}
