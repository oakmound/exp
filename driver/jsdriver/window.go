// +build js

package jsdriver

import (
	"image"
	"image/color"
	"image/draw"
	"runtime"
	"sync"

	"syscall/js"

	"github.com/oakmound/shiny/screen"
	"golang.org/x/image/math/f64"
)

type JSWindow struct {
	ctx       js.Value
	jsUint8   js.Value
	uint8Set  bool
	imgData   js.Value
	events    []interface{}
	eventLock sync.Mutex
}

func (jsc *JSWindow) Release() {
}

func (jsc *JSWindow) Publish() screen.PublishResult {
	// Publish doesn't do anything on JS
	// (Publish doesn't do anything on windows either)
	return screen.PublishResult{}
}

/////////////
// EventDeque
func (jsc *JSWindow) Send(event interface{}) {
	jsc.eventLock.Lock()
	jsc.events = append(jsc.events, event)
	jsc.eventLock.Unlock()
}

func (jsc *JSWindow) SendFirst(event interface{}) {
}

func (jsc *JSWindow) NextEvent() interface{} {
	if len(jsc.events) > 0 {
		jsc.eventLock.Lock()
		ev := jsc.events[0]
		jsc.events = jsc.events[1:]
		jsc.eventLock.Unlock()
		return ev
	}
	return nil
}

//////////////
// Uploader

func (jsc *JSWindow) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {
	if !jsc.uint8Set {
		jsc.jsUint8 = js.Global().Get("Uint8ClampedArray").New(src.RGBA().Pix, sr.Max.X, sr.Max.Y)
		jsc.uint8Set = true
	} else {
		jsc.jsUint8.Call("set", src.RGBA().Pix)
	}
	// This uses a heck of a lot of memory. It'd be wonderful if we didn't need to call New here
	// but could just refill the old variable
	jsc.imgData = js.Global().Get("ImageData").New(jsc.jsUint8, sr.Max.X, sr.Max.Y)
	jsc.ctx.Call("putImageData", jsc.imgData, dp.X, dp.Y)
	runtime.GC()
}

func (jsc *JSWindow) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	//Todo
}

///////////////
// Drawer

func (jsc *JSWindow) Draw(src2dst f64.Aff3, src screen.Texture, sr image.Rectangle, op draw.Op) {
	//Todo
}

// DrawUniform is like Draw except that the src is a uniform color instead
// of a Texture.
func (jsc *JSWindow) DrawUniform(src2dst f64.Aff3, src color.Color, sr image.Rectangle, op draw.Op) {
	//Todo
}

// Copy copies the sub-Texture defined by src and sr to the destination
// (the method receiver), such that sr.Min in src-space aligns with dp in
// dst-space.
func (jsc *JSWindow) Copy(dp image.Point, src screen.Texture, sr image.Rectangle, op draw.Op) {
	//Todo
}

// Scale scales the sub-Texture defined by src and sr to the destination
// (the method receiver), such that sr in src-space is mapped to dr in
// dst-space.
func (jsc *JSWindow) Scale(dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op) {
	//Todo
}
