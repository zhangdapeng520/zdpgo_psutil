//go:build solaris
// +build solaris

package disk

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"

	"github.com/zhangdapeng520/zdpgo_psutil/gopsutil/internal/common"
	"golang.org/x/sys/unix"
)

const (
	// _DEFAULT_NUM_MOUNTS is set to `cat /etc/mnttab | wc -l` rounded up to the
	// nearest power of two.
	_DEFAULT_NUM_MOUNTS = 32

	// _MNTTAB default place to read mount information
	_MNTTAB = "/etc/mnttab"
)

// A blacklist of read-only virtual filesystems.  Writable filesystems are of
// operational concern and must not be included in this list.
var fsTypeBlacklist = map[string]struct{}{
	"ctfs":   {},
	"dev":    {},
	"fd":     {},
	"lofs":   {},
	"lxproc": {},
	"mntfs":  {},
	"objfs":  {},
	"proc":   {},
}

func PartitionsWithContext(ctx context.Context, all bool) ([]PartitionStat, error) {
	ret := make([]PartitionStat, 0, _DEFAULT_NUM_MOUNTS)

	// Scan mnttab(4)
	f, err := os.Open(_MNTTAB)
	if err != nil {
	}
	defer func() {
		if err == nil {
			err = f.Close()
		} else {
			f.Close()
		}
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "\t")

		if _, found := fsTypeBlacklist[fields[2]]; found {
			continue
		}

		ret = append(ret, PartitionStat{
			// NOTE(seanc@): Device isn't exactly accurate: from mnttab(4): "The name
			// of the resource that has been mounted."  Ideally this value would come
			// from Statvfs_t.Fsid but I'm leaving it to the caller to traverse
			// unix.Statvfs().
			Device:     fields[0],
			Mountpoint: fields[1],
			Fstype:     fields[2],
			Opts:       strings.Split(fields[3], ","),
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to scan %q: %v", _MNTTAB, err)
	}

	return ret, err
}

func IOCountersWithContext(ctx context.Context, names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}

func UsageWithContext(ctx context.Context, path string) (*UsageStat, error) {
	statvfs := unix.Statvfs_t{}
	if err := unix.Statvfs(path, &statvfs); err != nil {
		return nil, fmt.Errorf("unable to call statvfs(2) on %q: %v", path, err)
	}

	usageStat := &UsageStat{
		Path:   path,
		Fstype: common.IntToString(statvfs.Basetype[:]),
		Total:  statvfs.Blocks * statvfs.Frsize,
		Free:   statvfs.Bfree * statvfs.Frsize,
		Used:   (statvfs.Blocks - statvfs.Bfree) * statvfs.Frsize,

		// NOTE: ZFS (and FreeBZSD's UFS2) use dynamic inode/dnode allocation.
		// Explicitly return a near-zero value for InodesUsedPercent so that nothing
		// attempts to garbage collect based on a lack of available inodes/dnodes.
		// Similarly, don't use the zero value to prevent divide-by-zero situations
		// and inject a faux near-zero value.  Filesystems evolve.  Has your
		// filesystem evolved?  Probably not if you care about the number of
		// available inodes.
		InodesTotal:       1024.0 * 1024.0,
		InodesUsed:        1024.0,
		InodesFree:        math.MaxUint64,
		InodesUsedPercent: (1024.0 / (1024.0 * 1024.0)) * 100.0,
	}

	usageStat.UsedPercent = (float64(usageStat.Used) / float64(usageStat.Total)) * 100.0

	return usageStat, nil
}

func SerialNumberWithContext(ctx context.Context, name string) (string, error) {
	cfgadm, err := exec.LookPath("cfgadm")
	if err != nil {
		return "", fmt.Errorf("find cfgadm: %w", err)
	}
	out, err := invoke.CommandWithContext(ctx, cfgadm, "-ls", "select=type(disk),cols=ap_id:info,cols2=,noheadings")
	if err != nil {
		return "", fmt.Errorf("exec cfgadm: %w", err)
	}

	suf := "::" + strings.TrimPrefix(name, "/dev/")
	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		flds := strings.Fields(s.Text())
		if strings.HasSuffix(flds[0], suf) {
			flen := len(flds)
			if flen >= 3 {
				for i, f := range flds {
					if i > 0 && i < flen-1 && f == "SN:" {
						return flds[i+1], nil
					}
				}
			}
			return "", nil
		}
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	return "", nil
}

func LabelWithContext(ctx context.Context, name string) (string, error) {
	return "", common.ErrNotImplementedError
}
