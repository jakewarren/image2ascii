//+build darwin dragonfly freebsd linux netbsd openbsd

package ascii

import (
	"syscall"
	"unsafe"
)

func termWidth(fd uintptr) int {
	var sz struct {
		rows    uint16
		cols    uint16
		xpixels uint16
		ypixels uint16
	}
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
		fd, uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))
	return int(sz.cols)
}
