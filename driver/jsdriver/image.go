// +build js

package jsdriver

import (
	"image"
)

type JSImage struct {
	rect image.Rectangle
	rgba *image.RGBA
}

func (jsb *JSImage) Release() {
}

func (jsb *JSImage) Size() image.Point {
	return jsb.rect.Max
}

func (jsb *JSImage) Bounds() image.Rectangle {
	return jsb.rect
}

func (jsb *JSImage) RGBA() *image.RGBA {
	return jsb.rgba
}
