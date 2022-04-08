//go:build !darwin && !linux && !freebsd && !openbsd && !solaris && !windows && !dragonfly && !plan9 && !aix
// +build !darwin,!linux,!freebsd,!openbsd,!solaris,!windows,!dragonfly,!plan9,!aix

package cpu

import (
	"context"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/internal/common"
	"runtime"
)

func Times(percpu bool) ([]TimesStat, error) {
	return TimesWithContext(context.Background(), percpu)
}

func TimesWithContext(ctx context.Context, percpu bool) ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}

func Info() ([]InfoStat, error) {
	return InfoWithContext(context.Background())
}

func InfoWithContext(ctx context.Context) ([]InfoStat, error) {
	return []InfoStat{}, common.ErrNotImplementedError
}

func CountsWithContext(ctx context.Context, logical bool) (int, error) {
	return runtime.NumCPU(), nil
}
