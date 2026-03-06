//go:build windows

package log

import (
	"syscall"
)

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	getCurrentThreadId = kernel32.NewProc("GetCurrentThreadId")
)

// getTID 获取 Windows 下的线程 ID
func getTID() int {
	r, _, _ := getCurrentThreadId.Call()
	return int(r)
}
