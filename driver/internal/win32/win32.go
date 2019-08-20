// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

// Package win32 implements a partial shiny screen driver using the Win32 API.
// It provides window, lifecycle, key, and mouse management, but no drawing.
// That is left to windriver (using GDI) or gldriver (using DirectX via ANGLE).
package win32

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"syscall"
	"unsafe"

	"github.com/oakmound/shiny/screen"
	"github.com/oakmound/w32"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/geom"
)

// screenHWND is the handle to the "Screen window".
// The Screen window encapsulates all screen.Screen operations
// in an actual Windows window so they all run on the main thread.
// Since any messages sent to a window will be executed on the
// main thread, we can safely use the messages below.
var screenHWND w32.HWND

const (
	msgCreateWindow = _WM_USER + iota
	msgMainCallback
	msgShow
	msgQuit
	msgLast
)

// userWM is used to generate private (WM_USER and above) window message IDs
// for use by screenWindowWndProc and windowWndProc.
type userWM struct {
	sync.Mutex
	id uint32
}

func (m *userWM) next() uint32 {
	m.Lock()
	if m.id == 0 {
		m.id = msgLast
	}
	r := m.id
	m.id++
	m.Unlock()
	return r
}

var currentUserWM userWM

func newWindow(opts screen.WindowGenerator) (w32.HWND, error) {
	// TODO(brainman): convert windowClass to *uint16 once (in initWindowClass)
	wcname, err := syscall.UTF16PtrFromString(windowClass)
	if err != nil {
		return 0, err
	}
	title, err := syscall.UTF16PtrFromString(opts.Title)
	if err != nil {
		return 0, err
	}
	style, exStyle := WindowsStyle(opts)
	// This should be a feature, putting windows on the top layer
	if opts.TopMost {
		exStyle = exStyle | WS_EX_TOPMOST
	}
	hwnd, err := _CreateWindowEx(exStyle,
		wcname, title,
		style,
		_CW_USEDEFAULT, _CW_USEDEFAULT,
		_CW_USEDEFAULT, _CW_USEDEFAULT,
		0, 0, hThisInstance, 0)
	if err != nil {
		return 0, err
	}

	// This is interesting and we'll use it eventually
	//w32.SetWindowLongPtr(hwnd, w32.GWL_STYLE, 0)
	// TODO(andlabs): use proper nCmdShow
	// TODO(andlabs): call UpdateWindow()

	return hwnd, nil
}

// WindowsStyle converts a screen.BorderStyle into a style and
// exStyle for a Windows window
func WindowsStyle(gen screen.WindowGenerator) (uint32, uint32) {
	return WS_OVERLAPPEDWINDOW, 0
}

// ResizeClientRect makes hwnd client rectangle opts.Width by opts.Height in size.
func ResizeClientRect(hwnd w32.HWND, opts screen.WindowGenerator) error {
	if opts.Width <= 0 || opts.Height <= 0 {
		return errors.New("Invalid inputs to ResizeClientRect")
	}
	var cr, wr _RECT
	err := _GetClientRect(hwnd, &cr)
	if err != nil {
		return err
	}
	err = _GetWindowRect(hwnd, &wr)
	if err != nil {
		return err
	}
	w := (wr.Right - wr.Left) - (cr.Right - int32(opts.Width))
	h := (wr.Bottom - wr.Top) - (cr.Bottom - int32(opts.Height))
	x := wr.Left
	if opts.X != 0 {
		x = opts.X
	}
	y := wr.Top
	if opts.Y != 0 {
		y = opts.Y
	}
	return MoveWindow(hwnd, x, y, w, h, false)
}

// Show shows a newly created window.
// It sends the appropriate lifecycle events, makes the window appear
// on the screen, and sends an initial size event.
//
// This is a separate step from NewWindow to give the driver a chance
// to setup its internal state for a window before events start being
// delivered.
func Show(hwnd w32.HWND) {
	w32.SendMessage(hwnd, msgShow, 0, 0)
}

func Release(hwnd w32.HWND) {
	w32.SendMessage(hwnd, w32.WM_CLOSE, 0, 0)
}

func sendFocus(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	switch uMsg {
	case _WM_SETFOCUS:
		LifecycleEvent(hwnd, lifecycle.StageFocused)
	case _WM_KILLFOCUS:
		LifecycleEvent(hwnd, lifecycle.StageVisible)
	default:
		panic(fmt.Sprintf("unexpected focus message: %d", uMsg))
	}
	return _DefWindowProc(hwnd, uMsg, wParam, lParam)
}

func sendShow(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	LifecycleEvent(hwnd, lifecycle.StageVisible)
	w32.ShowWindow(hwnd, _SW_SHOWDEFAULT)
	sendSize(hwnd)
	return 0
}

func sendSizeEvent(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	wp := (*_WINDOWPOS)(unsafe.Pointer(lParam))
	if wp.Flags&_SWP_NOSIZE != 0 {
		return 0
	}
	sendSize(hwnd)
	return 0
}

func sendSize(hwnd w32.HWND) {
	var r _RECT
	if err := _GetClientRect(hwnd, &r); err != nil {
		panic(err) // TODO(andlabs)
	}

	width := int(r.Right - r.Left)
	height := int(r.Bottom - r.Top)

	// TODO(andlabs): don't assume that PixelsPerPt == 1
	SizeEvent(hwnd, size.Event{
		WidthPx:     width,
		HeightPx:    height,
		WidthPt:     geom.Pt(width),
		HeightPt:    geom.Pt(height),
		PixelsPerPt: 1,
	})
}

func sendClose(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	LifecycleEvent(hwnd, lifecycle.StageDead)
	ptr, _ := w32.DefWindowProc(hwnd, uMsg, wParam, lParam)
	return ptr
}

func sendMouseEvent(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	e := mouse.Event{
		X:         float32(_GET_X_LPARAM(lParam)),
		Y:         float32(_GET_Y_LPARAM(lParam)),
		Modifiers: keyModifiers(),
	}

	switch uMsg {
	case _WM_MOUSEMOVE:
		e.Direction = mouse.DirNone
	case _WM_LBUTTONDOWN, _WM_MBUTTONDOWN, _WM_RBUTTONDOWN:
		e.Direction = mouse.DirPress
	case _WM_LBUTTONUP, _WM_MBUTTONUP, _WM_RBUTTONUP:
		e.Direction = mouse.DirRelease
	case _WM_MOUSEWHEEL:
		// TODO: On a trackpad, a scroll can be a drawn-out affair with a
		// distinct beginning and end. Should the intermediate events be
		// DirNone?
		e.Direction = mouse.DirStep

		x, y, _ := w32.ScreenToClient(hwnd, int(e.X), int(e.Y))
		e.X = float32(x)
		e.Y = float32(y)
	default:
		panic("sendMouseEvent() called on non-mouse message")
	}

	switch uMsg {
	case _WM_MOUSEMOVE:
		// No-op.
	case _WM_LBUTTONDOWN, _WM_LBUTTONUP:
		e.Button = mouse.ButtonLeft
	case _WM_MBUTTONDOWN, _WM_MBUTTONUP:
		e.Button = mouse.ButtonMiddle
	case _WM_RBUTTONDOWN, _WM_RBUTTONUP:
		e.Button = mouse.ButtonRight
	case _WM_MOUSEWHEEL:
		// TODO: handle horizontal scrolling
		delta := _GET_WHEEL_DELTA_WPARAM(wParam) / _WHEEL_DELTA
		switch {
		case delta > 0:
			e.Button = mouse.ButtonWheelUp
		case delta < 0:
			e.Button = mouse.ButtonWheelDown
			delta = -delta
		default:
			return
		}
		for delta > 0 {
			MouseEvent(hwnd, e)
			delta--
		}
		return
	}

	MouseEvent(hwnd, e)

	return 0
}

// Precondition: this is called in immediate response to the message that triggered the event (so not after w.Send).
func keyModifiers() (m key.Modifiers) {
	down := func(x int32) bool {
		// GetKeyState gets the key state at the time of the message, so this is what we want.
		return _GetKeyState(x)&0x80 != 0
	}

	if down(_VK_CONTROL) {
		m |= key.ModControl
	}
	if down(_VK_MENU) {
		m |= key.ModAlt
	}
	if down(_VK_SHIFT) {
		m |= key.ModShift
	}
	if down(_VK_LWIN) || down(_VK_RWIN) {
		m |= key.ModMeta
	}
	return m
}

var (
	MouseEvent     func(hwnd w32.HWND, e mouse.Event)
	PaintEvent     func(hwnd w32.HWND, e paint.Event)
	SizeEvent      func(hwnd w32.HWND, e size.Event)
	KeyEvent       func(hwnd w32.HWND, e key.Event)
	LifecycleEvent func(hwnd w32.HWND, e lifecycle.Stage)

	// TODO: use the golang.org/x/exp/shiny/driver/internal/lifecycler package
	// instead of or together with the LifecycleEvent callback?
)

func sendPaint(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr) {
	PaintEvent(hwnd, paint.Event{})
	return _DefWindowProc(hwnd, uMsg, wParam, lParam)
}

var screenMsgs = map[uint32]func(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr){}

func AddScreenMsg(fn func(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr)) uint32 {
	uMsg := currentUserWM.next()
	screenMsgs[uMsg] = func(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) uintptr {
		fn(hwnd, uMsg, wParam, lParam)
		return 0
	}
	return uMsg
}

var (
	windowInit = sync.Once{}
)

func screenWindowWndProc(hwnd w32.HWND, uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	switch uMsg {
	case msgCreateWindow:
		p := (*newWindowParams)(unsafe.Pointer(lParam))
		p.w, p.err = newWindow(p.opts)
	case msgMainCallback:
		windowInit.Do(func() {
			go func() {
				mainCallback()
				SendScreenMessage(msgQuit, 0, 0)
			}()
		})
	case msgQuit:
		_PostQuitMessage(0)
	}
	fn := screenMsgs[uMsg]
	if fn != nil {
		return fn(hwnd, uMsg, wParam, lParam)
	}
	return _DefWindowProc(hwnd, uMsg, wParam, lParam)
}

//go:uintptrescapes

func SendScreenMessage(uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	return w32.SendMessage(screenHWND, uMsg, wParam, lParam)
}

var windowMsgs = map[uint32]func(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) (lResult uintptr){
	_WM_SETFOCUS:         sendFocus,
	_WM_KILLFOCUS:        sendFocus,
	_WM_PAINT:            sendPaint,
	msgShow:              sendShow,
	_WM_WINDOWPOSCHANGED: sendSizeEvent,
	_WM_CLOSE:            sendClose,

	_WM_LBUTTONDOWN: sendMouseEvent,
	_WM_LBUTTONUP:   sendMouseEvent,
	_WM_MBUTTONDOWN: sendMouseEvent,
	_WM_MBUTTONUP:   sendMouseEvent,
	_WM_RBUTTONDOWN: sendMouseEvent,
	_WM_RBUTTONUP:   sendMouseEvent,
	_WM_MOUSEMOVE:   sendMouseEvent,
	_WM_MOUSEWHEEL:  sendMouseEvent,

	_WM_KEYDOWN: sendKeyEvent,
	_WM_KEYUP:   sendKeyEvent,
	// TODO case _WM_SYSKEYDOWN, _WM_SYSKEYUP:
}

func AddWindowMsg(fn func(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr)) uint32 {
	uMsg := currentUserWM.next()
	windowMsgs[uMsg] = func(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) uintptr {
		fn(hwnd, uMsg, wParam, lParam)
		return 0
	}
	return uMsg
}

func windowWndProc(hwnd w32.HWND, uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	fn := windowMsgs[uMsg]
	if fn != nil {
		return fn(hwnd, uMsg, wParam, lParam)
	}
	return _DefWindowProc(hwnd, uMsg, wParam, lParam)
}

type newWindowParams struct {
	opts screen.WindowGenerator
	w    w32.HWND
	err  error
}

func NewWindow(opts screen.WindowGenerator) (w32.HWND, error) {
	var p newWindowParams
	p.opts = opts
	SendScreenMessage(msgCreateWindow, 0, uintptr(unsafe.Pointer(&p)))
	return p.w, p.err
}

const windowClass = "shiny_Window"

func initWindowClass() (err error) {
	wcname, err := syscall.UTF16PtrFromString(windowClass)
	if err != nil {
		return err
	}
	_, err = _RegisterClass(&_WNDCLASS{
		LpszClassName: wcname,
		LpfnWndProc:   syscall.NewCallback(windowWndProc),
		HIcon:         hDefaultIcon,
		HCursor:       hDefaultCursor,
		HInstance:     hThisInstance,
		HbrBackground: w32.COLOR_BTNSHADOW,
	})
	return err
}

func initScreenWindow() (err error) {
	const screenWindowClass = "shiny_ScreenWindow"
	swc, err := syscall.UTF16PtrFromString(screenWindowClass)
	if err != nil {
		return err
	}
	emptyString, err := syscall.UTF16PtrFromString("")
	if err != nil {
		return err
	}
	wc := _WNDCLASS{
		LpszClassName: swc,
		LpfnWndProc:   syscall.NewCallback(screenWindowWndProc),
		HIcon:         hDefaultIcon,
		HCursor:       hDefaultCursor,
		HInstance:     hThisInstance,
		HbrBackground: w32.HWND(w32.COLOR_BTNSHADOW),
	}
	_, err = _RegisterClass(&wc)
	if err != nil {
		return err
	}
	screenHWND, err = _CreateWindowEx(0,
		swc, emptyString,
		windowStyle,
		_CW_USEDEFAULT, _CW_USEDEFAULT,
		_CW_USEDEFAULT, _CW_USEDEFAULT,
		w32.HWND_MESSAGE, 0, hThisInstance, 0)
	if err != nil {
		return err
	}
	return nil
}

var (
	windowStyle uint32 = WS_OVERLAPPEDWINDOW
)

var (
	hDefaultIcon   w32.HICON
	hDefaultCursor w32.HCURSOR
	hThisInstance  w32.HINSTANCE
)

func initCommon() (err error) {
	hDefaultIcon, err = _LoadIcon(0, w32.IDI_APPLICATION)
	if err != nil {
		return err
	}
	hDefaultCursor, err = _LoadCursor(0, w32.IDC_ARROW)
	if err != nil {
		return err
	}
	// TODO(andlabs) hThisInstance
	return nil
}

// Todo: this (and other globals) forces this package to only be able to run one window.
// Can we change this?

var mainCallback func()

func Main(f func()) (retErr error) {
	// It does not matter which OS thread we are on.
	// All that matters is that we confine all UI operations
	// to the thread that created the respective window.
	runtime.LockOSThread()

	if err := initCommon(); err != nil {
		return err
	}

	if err := initScreenWindow(); err != nil {
		return err
	}
	defer func() {
		// TODO(andlabs): log an error if this fails?
		w32.DestroyWindow(screenHWND)
		// TODO(andlabs): unregister window class
	}()

	if err := initWindowClass(); err != nil {
		return err
	}

	// Prime the pump.
	mainCallback = f
	w32.PostMessage(screenHWND, msgMainCallback, 0, 0)

	// Main message pump.
	var m w32.MSG
	for {
		done, err := _GetMessage(&m, 0, 0, 0)
		if err != nil {
			return fmt.Errorf("win32 GetMessage failed: %v", err)
		}
		if done == 0 { // WM_QUIT
			break
		}
		w32.TranslateMessage(&m)
		w32.DispatchMessage(&m)
	}

	return nil
}
