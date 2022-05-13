//go:build darwin && !cgo
// +build darwin,!cgo

package cpu

import (
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/internal/common"
)

func perCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}

func allCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}
