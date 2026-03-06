//go:build linux

package log

import (
	"syscall"
)

// getTID 获取 Linux 下的线程 ID
func getTID() int {
	return syscall.Gettid()
}
