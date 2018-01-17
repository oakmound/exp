// +build js

package jsdriver

import (
	"image"
)

type JSBuffer struct {
	rect image.Rectangle
	rgba *image.RGBA
}

func (jsb *JSBuffer) Release() {
}

func (jsb *JSBuffer) Size() image.Point {
	return jsb.rect.Max
}

func (jsb *JSBuffer) Bounds() image.Rectangle {
	return jsb.rect
}

func (jsb *JSBuffer) RGBA() *image.RGBA {
	return jsb.rgba
}
