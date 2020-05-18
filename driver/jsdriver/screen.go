package jsdriver

import (
	"image"

	"syscall/js"

	"github.com/oakmound/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/mouse"
)

type JSScreen struct{}

func (jss *JSScreen) NewImage(p image.Point) (screen.Image, error) {
	rect := image.Rect(0, 0, p.X, p.Y)
	rgba := image.NewRGBA(rect)
	buffer := &JSImage{
		rect,
		rgba,
	}
	return buffer, nil
}

func (jss *JSScreen) NewWindow(opts screen.WindowGenerator) (screen.Window, error) {
	jsc := new(JSWindow)

	document := js.Global().Get("document")
	canvas := document.Call("createElement", "canvas")
	canvas.Get("style").Set("display", "block")
	canvas.Set("width", opts.Width)
	canvas.Set("height", opts.Height)
	jsc.ctx = canvas.Call("getContext", "2d")
	bdy := document.Get("body")
	bdy.Call("appendChild", canvas)

	// These bindings are modified from the bindings engi uses for its js support.

	canvas.Call("addEventListener", "mousemove", func(ev js.Value) {
		rect := canvas.Call("getBoundingClientRect")
		x := float32((ev.Get("clientX").Int() - rect.Get("left").Int()))
		y := float32((ev.Get("clientY").Int() - rect.Get("top").Int()))
		jsc.Send(mouse.Event{X: x, Y: y, Button: mouse.ButtonNone, Direction: mouse.DirNone})
	}, false)

	canvas.Call("addEventListener", "mousedown", func(ev js.Value) {
		rect := canvas.Call("getBoundingClientRect")
		x := float32((ev.Get("clientX").Int() - rect.Get("left").Int()))
		y := float32((ev.Get("clientY").Int() - rect.Get("top").Int()))
		button := jsMouseButton(ev.Get("button").Int())
		jsc.Send(mouse.Event{X: x, Y: y, Button: button, Direction: mouse.DirPress})
	}, false)

	canvas.Call("addEventListener", "mouseup", func(ev js.Value) {
		rect := canvas.Call("getBoundingClientRect")
		x := float32((ev.Get("clientX").Int() - rect.Get("left").Int()))
		y := float32((ev.Get("clientY").Int() - rect.Get("top").Int()))
		button := jsMouseButton(ev.Get("button").Int())
		jsc.Send(mouse.Event{X: x, Y: y, Button: button, Direction: mouse.DirRelease})
	}, false)

	js.Global().Call("addEventListener", "keydown", func(ev js.Value) {
		k := ev.Get("keyCode").Int()
		jsc.Send(key.Event{Code: jsKey(k), Direction: key.DirPress})
	}, false)

	js.Global().Call("addEventListener", "keyup", func(ev js.Value) {
		k := ev.Get("keyCode").Int()
		jsc.Send(key.Event{Code: jsKey(k), Direction: key.DirRelease})
	}, false)

	return jsc, nil
}

func (jss *JSScreen) NewTexture(p image.Point) (screen.Texture, error) {
	txt := new(JSTexture)
	return txt, nil
}
