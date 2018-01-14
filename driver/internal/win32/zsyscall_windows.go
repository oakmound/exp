package win32

import (
	"syscall"
	"unsafe"

	"github.com/oakmound/w32"
	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return syscall.EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	moduser32 = windows.NewLazySystemDLL("user32.dll")

	procGetDC             = moduser32.NewProc("GetDC")
	procReleaseDC         = moduser32.NewProc("ReleaseDC")
	procCreateWindowExW   = moduser32.NewProc("CreateWindowExW")
	procDefWindowProcW    = moduser32.NewProc("DefWindowProcW")
	procDispatchMessageW  = moduser32.NewProc("DispatchMessageW")
	procGetClientRect     = moduser32.NewProc("GetClientRect")
	procGetWindowRect     = moduser32.NewProc("GetWindowRect")
	procGetKeyboardLayout = moduser32.NewProc("GetKeyboardLayout")
	procGetKeyboardState  = moduser32.NewProc("GetKeyboardState")
	procGetKeyState       = moduser32.NewProc("GetKeyState")
	procGetMessageW       = moduser32.NewProc("GetMessageW")
	procLoadCursorW       = moduser32.NewProc("LoadCursorW")
	procLoadIconW         = moduser32.NewProc("LoadIconW")
	procMoveWindow        = moduser32.NewProc("MoveWindow")
	procPostQuitMessage   = moduser32.NewProc("PostQuitMessage")
	procRegisterClassW    = moduser32.NewProc("RegisterClassW")
	procShowWindow        = moduser32.NewProc("ShowWindow")
	procScreenToClient    = moduser32.NewProc("ScreenToClient")
	procToUnicodeEx       = moduser32.NewProc("ToUnicodeEx")
	procTranslateMessage  = moduser32.NewProc("TranslateMessage")
	procSetWindowPos      = moduser32.NewProc("SetWindowPos")
)

func GetDC(hwnd w32.HWND) (dc syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetDC.Addr(), 1, uintptr(hwnd), 0, 0)
	dc = syscall.Handle(r0)
	if dc == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func ReleaseDC(hwnd w32.HWND, dc syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procReleaseDC.Addr(), 2, uintptr(hwnd), uintptr(dc), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _CreateWindowEx(exstyle uint32, className *uint16, windowText *uint16, style uint32, x int32, y int32, width int32, height int32, parent w32.HWND, menu syscall.Handle, hInstance w32.HINSTANCE, lpParam uintptr) (hwnd w32.HWND, err error) {
	r0, _, e1 := syscall.Syscall12(procCreateWindowExW.Addr(), 12, uintptr(exstyle), uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(windowText)), uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(parent), uintptr(menu), uintptr(hInstance), uintptr(lpParam))
	hwnd = w32.HWND(r0)
	if hwnd == 0 {
		err = errnoErr(e1)
	}
	return
}

func _DefWindowProc(hwnd w32.HWND, uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	r0, _, _ := syscall.Syscall6(procDefWindowProcW.Addr(), 4, uintptr(hwnd), uintptr(uMsg), uintptr(wParam), uintptr(lParam), 0, 0)
	lResult = uintptr(r0)
	return
}

func _GetClientRect(hwnd w32.HWND, rect *_RECT) (err error) {
	r1, _, e1 := syscall.Syscall(procGetClientRect.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(rect)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _GetWindowRect(hwnd w32.HWND, rect *_RECT) (err error) {
	r1, _, e1 := syscall.Syscall(procGetWindowRect.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(rect)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _GetKeyboardLayout(threadID uint32) (locale syscall.Handle) {
	r0, _, _ := syscall.Syscall(procGetKeyboardLayout.Addr(), 1, uintptr(threadID), 0, 0)
	locale = syscall.Handle(r0)
	return
}

func _GetKeyboardState(lpKeyState *byte) (err error) {
	r1, _, e1 := syscall.Syscall(procGetKeyboardState.Addr(), 1, uintptr(unsafe.Pointer(lpKeyState)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _GetKeyState(virtkey int32) (keystatus int16) {
	r0, _, _ := syscall.Syscall(procGetKeyState.Addr(), 1, uintptr(virtkey), 0, 0)
	keystatus = int16(r0)
	return
}

func _GetMessage(msg *w32.MSG, hwnd syscall.Handle, msgfiltermin uint32, msgfiltermax uint32) (ret int32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetMessageW.Addr(), 4, uintptr(unsafe.Pointer(msg)), uintptr(hwnd), uintptr(msgfiltermin), uintptr(msgfiltermax), 0, 0)
	ret = int32(r0)
	if ret == -1 {
		err = errnoErr(e1)
	}
	return
}

func _LoadCursor(hInstance syscall.Handle, cursorName uintptr) (cursor w32.HCURSOR, err error) {
	r0, _, e1 := syscall.Syscall(procLoadCursorW.Addr(), 2, uintptr(hInstance), uintptr(cursorName), 0)
	cursor = w32.HCURSOR(r0)
	if cursor == 0 {
		err = errnoErr(e1)
	}
	return
}

func _LoadIcon(hInstance w32.HINSTANCE, iconName uintptr) (icon w32.HICON, err error) {
	r0, _, e1 := syscall.Syscall(procLoadIconW.Addr(), 2, uintptr(hInstance), uintptr(iconName), 0)
	icon = w32.HICON(r0)
	if icon == 0 {
		err = errnoErr(e1)
	}
	return
}

func MoveWindow(hwnd w32.HWND, x int32, y int32, w int32, h int32, repaint bool) (err error) {
	var _p0 uint32
	if repaint {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r1, _, e1 := syscall.Syscall6(procMoveWindow.Addr(), 6, uintptr(hwnd), uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(_p0))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _PostQuitMessage(exitCode int32) {
	syscall.Syscall(procPostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
	return
}

func _RegisterClass(wc *_WNDCLASS) (atom uint16, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterClassW.Addr(), 1, uintptr(unsafe.Pointer(wc)), 0, 0)
	atom = uint16(r0)
	if atom == 0 {
		err = errnoErr(e1)
	}
	return
}

func _ToUnicodeEx(wVirtKey uint32, wScanCode uint32, lpKeyState *byte, pwszBuff *uint16, cchBuff int32, wFlags uint32, dwhkl syscall.Handle) (ret int32) {
	r0, _, _ := syscall.Syscall9(procToUnicodeEx.Addr(), 7, uintptr(wVirtKey), uintptr(wScanCode), uintptr(unsafe.Pointer(lpKeyState)), uintptr(unsafe.Pointer(pwszBuff)), uintptr(cchBuff), uintptr(wFlags), uintptr(dwhkl), 0, 0)
	ret = int32(r0)
	return
}
