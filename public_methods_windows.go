package figure

import (
	"syscall"
	"unsafe"

	ct "github.com/daviddengcn/go-colortext"
)

func init() {
	if majorVersion, _, _ := RtlGetNtVersionNumbers(); majorVersion < 10 {
		IsOldWindows = true

		ctColors = map[string]ct.Color{
			ColorRed:    ct.Red,
			ColorGreen:  ct.Green,
			ColorYellow: ct.Yellow,
			ColorBlue:   ct.Blue,
			ColorPurple: ct.Magenta,
			ColorCyan:   ct.Cyan,
			ColorGray:   ct.Black,
			ColorWhite:  ct.White,
		}
	}
}

func RtlGetNtVersionNumbers() (majorVersion, minorVersion, buildNumber uint32) {
	ntdll := syscall.NewLazyDLL("ntdll.dll")
	procRtlGetNtVersionNumbers := ntdll.NewProc("RtlGetNtVersionNumbers")
	_, _, _ = procRtlGetNtVersionNumbers.Call(
		uintptr(unsafe.Pointer(&majorVersion)),
		uintptr(unsafe.Pointer(&minorVersion)),
		uintptr(unsafe.Pointer(&buildNumber)),
	)
	buildNumber &= 0xffff
	return
}
