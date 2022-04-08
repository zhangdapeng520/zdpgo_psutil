//go:build darwin && !cgo
// +build darwin,!cgo

package disk

import (
	"context"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/internal/common"
)

func IOCountersWithContext(ctx context.Context, names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}
