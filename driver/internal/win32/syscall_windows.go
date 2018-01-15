// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package win32

import (
	"syscall"

	"github.com/oakmound/w32"
)

type _COLORREF uint32

func _RGB(r, g, b byte) _COLORREF {
	return _COLORREF(r) | _COLORREF(g)<<8 | _COLORREF(b)<<16
}

type _POINT struct {
	X int32
	Y int32
}

type _RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type _WNDCLASS struct {
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     w32.HINSTANCE
	HIcon         w32.HICON
	HCursor       w32.HCURSOR
	HbrBackground w32.HWND
	LpszMenuName  *uint16
	LpszClassName *uint16
}

type _WINDOWPOS struct {
	HWND            syscall.Handle
	HWNDInsertAfter syscall.Handle
	X               int32
	Y               int32
	Cx              int32
	Cy              int32
	Flags           uint32
}

const (
	_WM_SETFOCUS         = 7
	_WM_KILLFOCUS        = 8
	_WM_PAINT            = 15
	_WM_CLOSE            = 16
	_WM_WINDOWPOSCHANGED = 71
	_WM_KEYDOWN          = 256
	_WM_KEYUP            = 257
	_WM_SYSKEYDOWN       = 260
	_WM_SYSKEYUP         = 261
	_WM_MOUSEMOVE        = 512
	_WM_MOUSEWHEEL       = 522
	_WM_LBUTTONDOWN      = 513
	_WM_LBUTTONUP        = 514
	_WM_RBUTTONDOWN      = 516
	_WM_RBUTTONUP        = 517
	_WM_MBUTTONDOWN      = 519
	_WM_MBUTTONUP        = 520
	_WM_USER             = 0x0400
)

// Docs copied from https://msdn.microsoft.com/en-us/library/windows/desktop/ms632600(v=vs.85).aspx
// These settings have a minimum requirement of windows 2000.
const (
	// The window has a thin-line border.
	WS_BORDER = 0x00800000
	// The window has a title bar (includes the WS_BORDER style).
	WS_CAPTION = 0x00C00000
	// The window is a child window. A window with this style cannot have a menu bar.
	// This style cannot be used with the WS_POPUP style.
	WS_CHILD = 0x40000000
	// Same as the WS_CHILD style.
	WS_CHILDWINDOW = WS_CHILD
	// Excludes the area occupied by child windows when drawing occurs within the parent window.
	// This style is used when creating the parent window.
	WS_CLIPCHILDREN = 0x02000000
	// Clips child windows relative to each other; that is, when a particular child window receives
	// a WM_PAINT message, the WS_CLIPSIBLINGS style clips all other overlapping child windows out
	// of the region of the child window to be updated. If WS_CLIPSIBLINGS is not specified and
	// child windows overlap, it is possible, when drawing within the client area of a child window,
	// to draw within the client area of a neighboring child window.
	WS_CLIPSIBLINGS = 0x04000000
	// The window is initially disabled. A disabled window cannot receive input from the user.
	// To change this after a window has been created, use the EnableWindow function.
	WS_DISABLED = 0x08000000
	// The window has a border of a style typically used with dialog boxes.
	// A window with this style cannot have a title bar.
	WS_DLGFRAME = 0x00400000
	// The window is the first control of a group of controls.
	// The group consists of this first control and all controls defined after it,
	// up to the next control with the WS_GROUP style. The first control in each group usually has
	// the WS_TABSTOP style so that the user can move from group to group. The user can subsequently
	// change the keyboard focus from one control in the group to the next control in the group by
	// using the direction keys.
	//
	// You can turn this style on and off to change dialog box navigation.
	// To change this style after a window has been created, use the SetWindowLong function.
	WS_GROUP = 0x00020000
	// The window has a horizontal scroll bar.
	WS_HSCROLL = 0x00100000
	// The window is initially minimized. Same as the WS_MINIMIZE style.
	WS_ICONIC = 0x20000000
	// The window is initially maximized.
	WS_MAXIMIZE = 0x01000000
	// The window is initially minimized. Same as the WS_ICONIC style.
	WS_MINIMIZE = WS_ICONIC
	// The window is an overlapped window.
	// An overlapped window has a title bar and a border. Same as the WS_TILED style.
	WS_OVERLAPPED = 0x00000000
	// The window has a minimize button. Cannot be combined with the WS_EX_CONTEXTHELP style.
	// The WS_SYSMENU style must also be specified.
	WS_MINIMIZEBOX = 0x00020000
	// The window has a maximize button. Cannot be combined with the WS_EX_CONTEXTHELP style.
	// The WS_SYSMENU style must also be specified.
	WS_MAXIMIZEBOX = 0x00010000
	// The window is an overlapped window. Same as the WS_TILEDWINDOW style.
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	// The windows is a pop-up window. This style cannot be used with the WS_CHILD style.
	WS_POPUP = 0x80000000
	// The window is a pop-up window. The WS_CAPTION and WS_POPUPWINDOW styles must be
	// combined to make the window menu visible.
	WS_POPUPWINDOW = WS_POPUP | WS_BORDER | WS_SYSMENU
	// The window has a sizing border. Same as the WS_THICKFRAME style.
	WS_SIZEBOX = 0x00040000
	// The window has a window menu on its title bar. The WS_CAPTION style must also be specified.
	WS_SYSMENU = 0x00080000
	// The window is a control that can receive the keyboard focus when the user presses the TAB key.
	// Pressing the TAB key changes the keyboard focus to the next control with the WS_TABSTOP style.
	// You can turn this style on and off to change dialog box navigation.
	// To change this style after a window has been created, use the SetWindowLong function.
	// For user-created windows and modeless dialogs to work with tab stops, alter the message loop
	// to call the IsDialogMessage function.
	WS_TABSTOP = 0x00010000
	// The window has a sizing border. Same as the WS_SIZEBOX style.
	WS_THICKFRAME = WS_SIZEBOX
	// The window is an overlapped window. An overlapped window has a title bar and a border.
	// Same as the WS_OVERLAPPED style.
	WS_TILED = WS_OVERLAPPED
	// The window is an overlapped window. Same as the WS_OVERLAPPEDWINDOW style.
	WS_TILEDWINDOW = WS_OVERLAPPEDWINDOW
	// The window is initially visible.
	// This style can be turned on and off by using the ShowWindow or SetWindowPos function.
	WS_VISIBLE = 0x10000000
	// The window has a vertical scroll bar.
	WS_VSCROLL = 0x00200000
)

const (
	// The window accepts drag-drop files.
	WS_EX_ACCEPTFILES = 0x00000010
	// Forces a top-level window onto the taskbar when the window is visible.
	WS_EX_APPWINDOW = 0x00040000
	// The window has a border with a sunken edge.
	WS_EX_CLIENTEDGE = 0x00000200
	// Paints all descendants of a window in bottom-to-top painting order using double-buffering.
	// For more information, see Remarks. This cannot be used if the window has a class style of either CS_OWNDC or CS_CLASSDC.
	// Windows 2000:  This style is not supported.
	WS_EX_COMPOSITED = 0x02000000
	// The title bar of the window includes a question mark. When the user clicks the question mark,
	// the cursor changes to a question mark with a pointer. If the user then clicks a child window,
	// the child receives a WM_HELP message. The child window should pass the message to the parent window procedure,
	// which should call the WinHelp function using the HELP_WM_HELP command. The Help application displays
	// a pop-up window that typically contains help for the child window.
	// WS_EX_CONTEXTHELP cannot be used with the WS_MAXIMIZEBOX or WS_MINIMIZEBOX styles.
	WS_EX_CONTEXTHELP = 0x00000400
	// The window itself contains child windows that should take part in dialog box navigation.
	// If this style is specified, the dialog manager recurses into children of this window when
	// performing navigation operations such as handling the TAB key, an arrow key, or a keyboard mnemonic.
	WS_EX_CONTROLPARENT = 0x00010000
	// The window has a double border; the window can, optionally, be created with a title bar
	// by specifying the WS_CAPTION style in the dwStyle parameter.
	WS_EX_DLGMODALFRAME = 0x00000001
	// The window is a layered window. This style cannot be used if the window has a
	// class style of either CS_OWNDC or CS_CLASSDC.
	// Windows 8:  The WS_EX_LAYERED style is supported for top-level windows and
	// child windows. Previous Windows versions support WS_EX_LAYERED only for top-level windows.
	WS_EX_LAYERED = 0x00080000
	// If the shell language is Hebrew, Arabic, or another language that supports reading
	// order alignment, the horizontal origin of the window is on the right edge.
	// Increasing horizontal values advance to the left.
	WS_EX_LAYOUTRTL = 0x00400000
	// The window has generic left-aligned properties. This is the default.
	WS_EX_LEFT = 0x00000000
	// If the shell language is Hebrew, Arabic, or another language that supports reading order
	// alignment, the vertical scroll bar (if present) is to the left of the client area.
	// For other languages, the style is ignored.
	WS_EX_LEFTSCROLLBAR = 0x00004000
	// The window text is displayed using left-to-right reading-order properties. This is the default.
	WS_EX_LTRREADING = 0x00000000
	// The window is a MDI child window.
	WS_EX_MDICHILD = 0x00000040
	// A top-level window created with this style does not become the foreground window when the
	// user clicks it. The system does not bring this window to the foreground when the user
	// minimizes or closes the foreground window.
	// To activate the window, use the SetActiveWindow or SetForegroundWindow function.
	// The window does not appear on the taskbar by default. To force the window to appear
	// on the taskbar, use the WS_EX_APPWINDOW style.
	WS_EX_NOACTIVATE = 0x08000000
	// The window does not pass its window layout to its child windows.
	WS_EX_NOINHERITLAYOUT = 0x00100000
	// The child window created with this style does not send the WM_PARENTNOTIFY message
	// to its parent window when it is created or destroyed.
	WS_EX_NOPARENTNOTIFY = 0x00000004
	// The window does not render to a redirection surface. This is for windows that do not
	// have visible content or that use mechanisms other than surfaces to provide their visual.
	WS_EX_NOREDIRECTIONBITMAP = 0x00200000
	// The window is an overlapped window.
	WS_EX_OVERLAPPEDWINDOW = (WS_EX_WINDOWEDGE | WS_EX_CLIENTEDGE)
	// The window is palette window, which is a modeless dialog box that presents an array of commands.
	WS_EX_PALETTEWINDOW = (WS_EX_WINDOWEDGE | WS_EX_TOOLWINDOW | WS_EX_TOPMOST)
	// The window has generic "right-aligned" properties. This depends on the window class.
	// This style has an effect only if the shell language is Hebrew, Arabic, or another language
	// that supports reading-order alignment; otherwise, the style is ignored.
	// Using the WS_EX_RIGHT style for static or edit controls has the same effect as using the SS_RIGHT
	// or ES_RIGHT style, respectively. Using this style with button controls has the same effect as
	// using BS_RIGHT and BS_RIGHTBUTTON styles.
	WS_EX_RIGHT = 0x00001000
	//The vertical scroll bar (if present) is to the right of the client area. This is the default.
	WS_EX_RIGHTSCROLLBAR = 0x00000000
	// If the shell language is Hebrew, Arabic, or another language that supports reading-order alignment,
	// the window text is displayed using right-to-left reading-order properties.
	// For other languages, the style is ignored.
	WS_EX_RTLREADING = 0x00002000
	// The window has a three-dimensional border style intended to be used for items that do not accept user input.
	WS_EX_STATICEDGE = 0x00020000
	// The window is intended to be used as a floating toolbar.
	// A tool window has a title bar that is shorter than a normal title bar, and the window title is drawn
	// using a smaller font. A tool window does not appear in the taskbar or in the dialog that
	// appears when the user presses ALT+TAB. If a tool window has a system menu, its icon is not displayed
	// on the title bar. However, you can display the system menu by right-clicking or by typing ALT+SPACE.
	WS_EX_TOOLWINDOW = 0x00000080
	// The window should be placed above all non-topmost windows and should stay above them,
	// even when the window is deactivated. To add or remove this style, use the SetWindowPos function.
	WS_EX_TOPMOST = 0x00000008
	// The window should not be painted until siblings beneath the window (that were created by the same thread)
	// have been painted. The window appears transparent because the bits of underlying sibling windows have
	// already been painted.
	// To achieve transparency without these restrictions, use the SetWindowRgn function.
	WS_EX_TRANSPARENT = 0x00000020
	// The window has a border with a raised edge.
	WS_EX_WINDOWEDGE = 0x00000100
)

// WM_SYSCOMMAND

const (
	_VK_SHIFT   = 16
	_VK_CONTROL = 17
	_VK_MENU    = 18
	_VK_LWIN    = 0x5B
	_VK_RWIN    = 0x5C
)

const (
	_MK_LBUTTON = 0x0001
	_MK_MBUTTON = 0x0010
	_MK_RBUTTON = 0x0002
)

const (
	_COLOR_BTNFACE = 15
)

const (
	_IDI_APPLICATION = 32512
	_IDC_ARROW       = 32512
)

const (
	_CW_USEDEFAULT = 0x80000000 - 0x100000000

	_SW_SHOWDEFAULT = 10

	_HWND_MESSAGE = syscall.Handle(^uintptr(2)) // -3

	_SWP_NOSIZE = 0x0001
)

const (
	_BI_RGB         = 0
	_DIB_RGB_COLORS = 0

	_AC_SRC_OVER  = 0x00
	_AC_SRC_ALPHA = 0x01

	_SRCCOPY = 0x00cc0020

	_WHEEL_DELTA = 120
)

func _GET_X_LPARAM(lp uintptr) int32 {
	return int32(_LOWORD(lp))
}

func _GET_Y_LPARAM(lp uintptr) int32 {
	return int32(_HIWORD(lp))
}

func _GET_WHEEL_DELTA_WPARAM(lp uintptr) int16 {
	return int16(_HIWORD(lp))
}

func _LOWORD(l uintptr) uint16 {
	return uint16(uint32(l))
}

func _HIWORD(l uintptr) uint16 {
	return uint16(uint32(l >> 16))
}

// notes to self
// UINT = uint32
// callbacks = uintptr
// strings = *uint16

//sys	GetDC(hwnd syscall.Handle) (dc syscall.Handle, err error) = user32.GetDC
//sys	ReleaseDC(hwnd syscall.Handle, dc syscall.Handle) (err error) = user32.ReleaseDC

//sys	_CreateWindowEx(exstyle uint32, className *uint16, windowText *uint16, style uint32, x int32, y int32, width int32, height int32, parent syscall.Handle, menu syscall.Handle, hInstance syscall.Handle, lpParam uintptr) (hwnd syscall.Handle, err error) = user32.CreateWindowExW
//sys	_DefWindowProc(hwnd syscall.Handle, uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) = user32.DefWindowProcW
//sys	_GetClientRect(hwnd syscall.Handle, rect *_RECT) (err error) = user32.GetClientRect
//sys	_GetWindowRect(hwnd syscall.Handle, rect *_RECT) (err error) = user32.GetWindowRect
//sys   _GetKeyboardLayout(threadID uint32) (locale syscall.Handle) = user32.GetKeyboardLayout
//sys   _GetKeyboardState(lpKeyState *byte) (err error) = user32.GetKeyboardState
//sys	_GetKeyState(virtkey int32) (keystatus int16) = user32.GetKeyState
//sys	_GetMessage(msg *_MSG, hwnd syscall.Handle, msgfiltermin uint32, msgfiltermax uint32) (ret int32, err error) [failretval==-1] = user32.GetMessageW
//sys	_LoadCursor(hInstance syscall.Handle, cursorName uintptr) (cursor syscall.Handle, err error) = user32.LoadCursorW
//sys	_LoadIcon(hInstance syscall.Handle, iconName uintptr) (icon syscall.Handle, err error) = user32.LoadIconW
//sys	_MoveWindow(hwnd syscall.Handle, x int32, y int32, w int32, h int32, repaint bool) (err error) = user32.MoveWindow
//sys   _PostQuitMessage(exitCode int32) = user32.PostQuitMessage
//sys	_RegisterClass(wc *_WNDCLASS) (atom uint16, err error) = user32.RegisterClassW
//sys   _ToUnicodeEx(wVirtKey uint32, wScanCode uint32, lpKeyState *byte, pwszBuff *uint16, cchBuff int32, wFlags uint32, dwhkl syscall.Handle) (ret int32) = user32.ToUnicodeEx
