package log

import (
	"unsafe"
)

// getg 由汇编实现
func getg() uintptr

// getGID 极速获取当前 Goroutine ID
// 性能比 runtime.Stack 快约 1000 倍，且零内存分配
func getGID() uint64 {
	g := getg()
	if g == 0 {
		return 0
	}
	// 在 Go 1.10 - 1.24+ 的 amd64 架构中，goid 的偏移量通常是 152
	// 如果未来 Go 版本变更，只需调整此偏移量
	return *(*uint64)(unsafe.Pointer(g + 152))
}
