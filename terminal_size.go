//go:build !windows
// +build !windows

package uilive

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

type windowSize struct {
	rows   uint16
	cols   uint16
	xPixel uint16
	yPixel uint16
}

var out *os.File
var err error
var sz windowSize

func getTermSize() (int, int) {
	if runtime.GOOS == "openbsd" {
		out, err = os.OpenFile("/dev/tty", os.O_RDWR, 0)
		if err != nil {
			return 0, 0
		}
		defer out.Close()
	} else {
		out, err = os.OpenFile("/dev/tty", os.O_WRONLY, 0)
		if err != nil {
			return 0, 0
		}
		defer out.Close()
	}

	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
		out.Fd(), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))
	return int(sz.cols), int(sz.rows)
}
