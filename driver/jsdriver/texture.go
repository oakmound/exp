// +build js

package jsdriver

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/oakmound/shiny/screen"
)

type JSTexture struct {
	JSBuffer
}

func (jst *JSTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	// We only operate on the expected use case of upload, where sr = src.Bounds
	// and dp = zeroPoint
	jst.rect = sr
	jst.rgba = image.NewRGBA(sr)
	*jst.rgba = *src.RGBA()
}

func (jst *JSTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	// Todo
}
