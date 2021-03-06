//go:build darwin || freebsd || openbsd
// +build darwin freebsd openbsd

package process

import (
	"bytes"
	"context"
	"encoding/binary"
	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/cpu"
	common2 "github.com/zhangdapeng520/zdpgo_psutil/gopsutil/internal/common"
)

type MemoryInfoExStat struct{}

type MemoryMapsStat struct{}

func (p *Process) TgidWithContext(ctx context.Context) (int32, error) {
	return 0, common2.ErrNotImplementedError
}

func (p *Process) IOniceWithContext(ctx context.Context) (int32, error) {
	return 0, common2.ErrNotImplementedError
}

func (p *Process) RlimitWithContext(ctx context.Context) ([]RlimitStat, error) {
	return nil, common2.ErrNotImplementedError
}

func (p *Process) RlimitUsageWithContext(ctx context.Context, gatherUsed bool) ([]RlimitStat, error) {
	return nil, common2.ErrNotImplementedError
}

func (p *Process) NumCtxSwitchesWithContext(ctx context.Context) (*NumCtxSwitchesStat, error) {
	return nil, common2.ErrNotImplementedError
}

func (p *Process) NumFDsWithContext(ctx context.Context) (int32, error) {
	return 0, common2.ErrNotImplementedError
}

func (p *Process) CPUAffinityWithContext(ctx context.Context) ([]int32, error) {
	return nil, common2.ErrNotImplementedError
}

func (p *Process) MemoryInfoExWithContext(ctx context.Context) (*MemoryInfoExStat, error) {
	return nil, common2.ErrNotImplementedError
}

func (p *Process) PageFaultsWithContext(ctx context.Context) (*PageFaultsStat, error) {
	return nil, common2.ErrNotImplementedError
}

func (p *Process) OpenFilesWithContext(ctx context.Context) ([]OpenFilesStat, error) {
	return nil, common2.ErrNotImplementedError
}

func (p *Process) MemoryMapsWithContext(ctx context.Context, grouped bool) (*[]MemoryMapsStat, error) {
	return nil, common2.ErrNotImplementedError
}

func (p *Process) ThreadsWithContext(ctx context.Context) (map[int32]*cpu.TimesStat, error) {
	return nil, common2.ErrNotImplementedError
}

func (p *Process) EnvironWithContext(ctx context.Context) ([]string, error) {
	return nil, common2.ErrNotImplementedError
}

func parseKinfoProc(buf []byte) (KinfoProc, error) {
	var k KinfoProc
	br := bytes.NewReader(buf)
	err := common2.Read(br, binary.LittleEndian, &k)
	return k, err
}
