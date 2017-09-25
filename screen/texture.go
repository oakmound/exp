package screen

import "image"

// Texture is a pixel buffer, but not one that is directly accessible as a
// []byte. Conceptually, it could live on a GPU, in another process or even be
// across a network, instead of on a CPU in this process.
//
// Images can be uploaded to Textures, and Textures can be drawn on Windows.
//
// When specifying a sub-Texture via Draw, a Texture's top-left pixel is always
// (0, 0) in its own coordinate space.
type Texture interface {
	// Release releases the Texture's resources, after all pending uploads and
	// draws resolve.
	//
	// The behavior of the Texture after Release, whether calling its methods
	// or passing it as an argument, is undefined.
	Release()

	// Size returns the size of the Texture's image.
	Size() image.Point

	// Bounds returns the bounds of the Texture's image. It is equal to
	// image.Rectangle{Max: t.Size()}.
	Bounds() image.Rectangle

	Uploader

	// TODO: also implement Drawer? If so, merge the Uploader and Drawer
	// interfaces??
}
