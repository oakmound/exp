package win32

import (
	"syscall"
	"unsafe"

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
	moduser32  = windows.NewLazySystemDLL("user32.dll")
	modshell32 = syscall.NewLazyDLL("shell32.dll")
)

var (
	procShell_NotifyIconW = modshell32.NewProc("Shell_NotifyIconW")

	procRegisterClass     = moduser32.NewProc("RegisterClassW")
	procIsZoomed          = moduser32.NewProc("IsZoomed")
	procLoadIcon          = moduser32.NewProc("LoadIconW")
	procLoadImageW        = moduser32.NewProc("LoadImageW")
	procLoadCursor        = moduser32.NewProc("LoadCursorW")
	procShowWindow        = moduser32.NewProc("ShowWindow")
	procCreateWindowEx    = moduser32.NewProc("CreateWindowExW")
	procDestroyWindow     = moduser32.NewProc("DestroyWindow")
	procDefWindowProc     = moduser32.NewProc("DefWindowProcW")
	procPostQuitMessage   = moduser32.NewProc("PostQuitMessage")
	procGetMessage        = moduser32.NewProc("GetMessageW")
	procTranslateMessage  = moduser32.NewProc("TranslateMessage")
	procDispatchMessage   = moduser32.NewProc("DispatchMessageW")
	procSendMessage       = moduser32.NewProc("SendMessageW")
	procPostMessage       = moduser32.NewProc("PostMessageW")
	procSetWindowText     = moduser32.NewProc("SetWindowTextW")
	procGetWindowRect     = moduser32.NewProc("GetWindowRect")
	procMoveWindow        = moduser32.NewProc("MoveWindow")
	procScreenToClient    = moduser32.NewProc("ScreenToClient")
	procSetWindowLong     = moduser32.NewProc("SetWindowLongW")
	procGetClientRect     = moduser32.NewProc("GetClientRect")
	procGetDC             = moduser32.NewProc("GetDC")
	procReleaseDC         = moduser32.NewProc("ReleaseDC")
	procSetWindowPos      = moduser32.NewProc("SetWindowPos")
	procGetKeyboardLayout = moduser32.NewProc("GetKeyboardLayout")
	procGetKeyboardState  = moduser32.NewProc("GetKeyboardState")
	procMonitorFromWindow = moduser32.NewProc("MonitorFromWindow")
	procGetMonitorInfo    = moduser32.NewProc("GetMonitorInfoW")
	procGetKeyState       = moduser32.NewProc("GetKeyState")
	procToUnicodeEx       = moduser32.NewProc("ToUnicodeEx")
)

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

func _PostQuitMessage(exitCode int32) {
	syscall.Syscall(procPostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
	return
}

func _ToUnicodeEx(wVirtKey uint32, wScanCode uint32, lpKeyState *byte, pwszBuff *uint16, cchBuff int32, wFlags uint32, dwhkl syscall.Handle) (ret int32) {
	r0, _, _ := syscall.Syscall9(procToUnicodeEx.Addr(), 7, uintptr(wVirtKey), uintptr(wScanCode), uintptr(unsafe.Pointer(lpKeyState)), uintptr(unsafe.Pointer(pwszBuff)), uintptr(cchBuff), uintptr(wFlags), uintptr(dwhkl), 0, 0)
	ret = int32(r0)
	return
}

func IsZoomed(hwnd HWND) bool {
	ret, _, _ := procIsZoomed.Call(uintptr(hwnd))
	return ret != 0
}

func RegisterClass(wc *_WNDCLASS) (atom uint16, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterClass.Addr(), 1, uintptr(unsafe.Pointer(wc)), 0, 0)
	atom = uint16(r0)
	if atom == 0 {
		err = errnoErr(e1)
	}
	return
}

func LoadIcon(instance HINSTANCE, iconName uintptr) (HICON, error) {
	var err error
	r0, _, e1 := syscall.Syscall(procLoadIcon.Addr(), 2, uintptr(instance), iconName, 0)
	icon := HICON(r0)
	if icon == 0 {
		err = errnoErr(e1)
	}
	return icon, err
}

func LoadCursor(instance HINSTANCE, cursorName uintptr) (HCURSOR, error) {
	var err error
	r0, _, e1 := syscall.Syscall(procLoadCursor.Addr(), 2, uintptr(instance), cursorName, 0)
	cursor := HCURSOR(r0)
	if cursor == 0 {
		err = errnoErr(e1)
	}
	return cursor, err
}

func ShowWindow(hwnd HWND, cmdshow int) bool {
	ret, _, _ := procShowWindow.Call(
		uintptr(hwnd),
		uintptr(cmdshow))

	return ret != 0
}

func CreateWindowEx(exStyle uint32, className, windowName *uint16,
	style uint32, x, y, width, height int, parent HWND, menu HMENU,
	instance HINSTANCE, param uintptr) (HWND, error) {
	ret, _, err := procCreateWindowEx.Call(
		uintptr(exStyle),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		param)
	if ret == 0 {
		return 0, err
	}
	return HWND(ret), nil
}

func DestroyWindow(hwnd HWND) bool {
	ret, _, _ := procDestroyWindow.Call(
		uintptr(hwnd))

	return ret != 0
}

func DefWindowProc(hwnd HWND, msg uint32, wParam, lParam uintptr) (uintptr, error) {
	ret, _, err := procDefWindowProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)
	if ret == 0 {
		return 0, err
	}
	return ret, nil
}

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) (int, error) {
	ret, _, err := procGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))
	if ret == 0 {
		return 0, err
	}
	return int(ret), nil
}

func TranslateMessage(msg *MSG) bool {
	ret, _, _ := procTranslateMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret != 0

}

func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := procDispatchMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret

}

func SendMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procSendMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func PostMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) bool {
	ret, _, _ := procPostMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret != 0
}

func SetWindowText(hwnd HWND, text string) {
	procSetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
}

func GetWindowRect(hwnd HWND) (*RECT, error) {
	var rect RECT
	ret, _, err := procGetWindowRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))
	if ret == 0 {
		return nil, err
	}
	return &rect, nil
}

func MoveWindow(hwnd HWND, x, y, width, height int32, repaint bool) error {
	ret, _, err := procMoveWindow.Call(
		uintptr(hwnd),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(boolToBOOL(repaint)))
	if ret == 0 {
		return err
	}
	return nil

}

func ScreenToClient(hwnd HWND, x, y int) (X, Y int, ok bool) {
	pt := POINT{X: int32(x), Y: int32(y)}
	ret, _, _ := procScreenToClient.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)))

	return int(pt.X), int(pt.Y), ret != 0
}

func SetWindowLong(hwnd HWND, index int, value int32) int32 {
	ret, _, _ := procSetWindowLong.Call(
		uintptr(hwnd),
		uintptr(index),
		uintptr(value))

	return int32(ret)
}

func GetClientRect(hwnd HWND) (*RECT, error) {
	var rect RECT
	ret, _, err := procGetClientRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))

	if ret == 0 {
		return nil, err
	}

	return &rect, nil
}

func GetDC(hwnd HWND) (HDC, error) {
	ret, _, err := procGetDC.Call(
		uintptr(hwnd))
	if ret == 0 {
		return 0, err
	}
	return HDC(ret), nil
}

func ReleaseDC(hwnd HWND, hDC HDC) bool {
	ret, _, _ := procReleaseDC.Call(
		uintptr(hwnd),
		uintptr(hDC))

	return ret != 0
}

func SetWindowPos(hwnd, hWndInsertAfter HWND, x, y, cx, cy int32, uFlags uint) bool {
	ret, _, _ := procSetWindowPos.Call(
		uintptr(hwnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags))

	return ret != 0
}

func MonitorFromWindow(hwnd HWND, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromWindow.Call(
		uintptr(hwnd),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

func LoadImage(hInst uintptr, name *uint16, type_ uint32, cx, cy int32, fuLoad uint32) (uintptr, error) {
	r1, _, err := procLoadImageW.Call(hInst, uintptr(unsafe.Pointer(name)), uintptr(type_), uintptr(cx), uintptr(cy), uintptr(fuLoad))
	if r1 != 0 {
		err = nil
	}
	return r1, err
}

func Shell_NotifyIcon(action Shell_NotifyAction, data *NOTIFYICONDATA) bool {
	r1, _, _ := procShell_NotifyIconW.Call(
		uintptr(action),
		uintptr(unsafe.Pointer(data)),
	)
	return r1 == 1
}

func GetMonitorInfo(hMonitor HMONITOR, lmpi *MONITORINFO) bool {
	ret, _, _ := procGetMonitorInfo.Call(
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(lmpi)),
	)
	return ret != 0
}
