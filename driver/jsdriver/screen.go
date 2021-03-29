package jsdriver

import (
	"image"

	//"syscall/js"

	"github.com/oakmound/shiny/screen"
	//"golang.org/x/mobile/event/key"
	//"golang.org/x/mobile/event/mouse"
	"honnef.co/go/js/dom/v2"
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
	jsc.width = opts.Width
	jsc.height = opts.Height

	//document := js.Global().Get("document")
	document := dom.GetWindow().Document()
	//canvas := document.Call("createElement", "canvas")
	canvas := document.CreateElement("canvas")
	//canvas.Get("style").Set("display", "block")
	//canvas.Set("width", opts.Width)
	//canvas.Set("height", opts.Height)
	canvasHTML := canvas.(*dom.HTMLCanvasElement)
	canvasHTML.SetHeight(opts.Height)
	canvasHTML.SetWidth(opts.Width)
	canvasHTML.Style().SetProperty("display", "block", "")

	//jsc.ctx = canvas.Call("getContext", "2d")
	jsc.ctx = canvasHTML.GetContext2d()

	//bdy := document.Get("body")
	//bdy.Call("appendChild", canvas)
	bdy := document.GetElementsByTagName("body")
	bdy[0].AppendChild(canvas)

	jsc.imageData = jsc.ctx.CreateImageData(jsc.width, jsc.height)

	// These bindings are modified from the bindings engi uses for its js support.

	// canvas.Call("addEventListener", "mousemove", js.FuncOf(func(ev js.Value, args []js.Value) interface{} {
	// 	rect := canvas.Call("getBoundingClientRect")
	// 	x := float32((ev.Get("clientX").Int() - rect.Get("left").Int()))
	// 	y := float32((ev.Get("clientY").Int() - rect.Get("top").Int()))
	// 	jsc.Send(mouse.Event{X: x, Y: y, Button: mouse.ButtonNone, Direction: mouse.DirNone})
	// 	return nil
	// }), false)

	// canvas.Call("addEventListener", "mousedown", js.FuncOf(func(ev js.Value, args []js.Value) interface{} {
	// 	rect := canvas.Call("getBoundingClientRect")
	// 	x := float32((ev.Get("clientX").Int() - rect.Get("left").Int()))
	// 	y := float32((ev.Get("clientY").Int() - rect.Get("top").Int()))
	// 	button := jsMouseButton(ev.Get("button").Int())
	// 	jsc.Send(mouse.Event{X: x, Y: y, Button: button, Direction: mouse.DirPress})
	// 	return nil
	// }), false)

	// canvas.Call("addEventListener", "mouseup", js.FuncOf(func(ev js.Value, args []js.Value) interface{} {
	// 	rect := canvas.Call("getBoundingClientRect")
	// 	x := float32((ev.Get("clientX").Int() - rect.Get("left").Int()))
	// 	y := float32((ev.Get("clientY").Int() - rect.Get("top").Int()))
	// 	button := jsMouseButton(ev.Get("button").Int())
	// 	jsc.Send(mouse.Event{X: x, Y: y, Button: button, Direction: mouse.DirRelease})
	// 	return nil
	// }), false)

	// js.Global().Call("addEventListener", "keydown", js.FuncOf(func(ev js.Value, args []js.Value) interface{} {
	// 	k := ev.Get("keyCode").Int()
	// 	jsc.Send(key.Event{Code: jsKey(k), Direction: key.DirPress})
	// 	return nil
	// }), false)

	// js.Global().Call("addEventListener", "keyup", js.FuncOf(func(ev js.Value, args []js.Value) interface{} {
	// 	k := ev.Get("keyCode").Int()
	// 	jsc.Send(key.Event{Code: jsKey(k), Direction: key.DirRelease})
	// 	return nil
	// }), false)

	return jsc, nil
}

func (jss *JSScreen) NewTexture(p image.Point) (screen.Texture, error) {
	txt := new(JSTexture)
	return txt, nil
}
