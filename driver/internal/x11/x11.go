package x11

import (
	"fmt"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/xevent"
)

func MoveWindow(xc *xgb.Conn, xw xproto.Window, x, y, width, height int32) (int32, int32, int32, int32) {
	vals := []uint32{}

	flags := xproto.ConfigWindowHeight |
		xproto.ConfigWindowWidth |
		xproto.ConfigWindowX |
		xproto.ConfigWindowY

	vals = append(vals, uint32(x))
	vals = append(vals, uint32(y))

	if int16(width) <= 0 {
		width = 1
	}
	vals = append(vals, uint32(width))

	if int16(height) <= 0 {
		height = 1
	}
	vals = append(vals, uint32(height))

	cook := xproto.ConfigureWindowChecked(xc, xw, uint16(flags), vals)
	if err := cook.Check(); err != nil {
		fmt.Println("X11 configure window failed: ", err)
	}
	return x, y, width, height
}

// todo: upgrade all of our xgb.Conns to xgbutil.XUtil types
// this code is stitched together from xgbutil examples
func SetFullScreen(xc *xgb.Conn, xw xproto.Window, fullscreen bool) error {

	fmt.Println("Entering x11.SetFullScreen")
	window := xw
	action := ewmh.StateToggle

	first := "_NET_WM_STATE_FULLSCREEN"
	messageType := "_NET_WM_STATE"
	source := 2
	var err error

	var atom1, atom2 xproto.Atom
	atom2 = 0

	reply, err := xproto.InternAtom(xc, false, uint16(len(first)), first).Reply()
	if err != nil {
		return err
	}
	atom1 = reply.Atom

	reply, err = xproto.InternAtom(xc, false, uint16(len(messageType)), messageType).Reply()
	if err != nil {
		return err
	}
	mstype := reply.Atom

	evMask := uint32(xproto.EventMaskSubstructureNotify |
		xproto.EventMaskSubstructureRedirect)

	ev, err := xevent.NewClientMessage(32, window, mstype, action, int(atom1), int(atom2), source)
	if err != nil {
		return err
	}

	err = xproto.SendEventChecked(xc, false, xw, evMask,
		string(ev.Bytes())).Check()

	if err == nil {
		fmt.Println("x11.SetFullScreen: no errors")
	}

	return err
}
