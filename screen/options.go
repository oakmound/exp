package screen

import "unicode/utf8"

// Todo: flesh out options as a generator style constructor
// with (optional) variadic arguments

type WindowGenerator struct {
	Width, Height int
	Title         string
	Fullscreen    bool
	NoScaling     bool
	X, Y          int
	BorderStyle
}

// todo...
// should describe whether there are maximize, minimize, etc buttons
// the size of the border (thick, thin, stylized)
type BorderStyle int

// NewWindowOptions are optional arguments to NewWindow.
type NewWindowOptions struct {
	// Width and Height specify the dimensions of the new window. If Width
	// or Height are zero, a driver-dependent default will be used for each
	// zero value dimension.
	Width, Height int

	// Title specifies the window title.
	Title string

	// TODO: fullscreen, icon, cursorHidden?
}

// GetTitle returns a sanitized form of o.Title. In particular, its length will
// not exceed 4096, and it may be further truncated so that it is valid UTF-8
// and will not contain the NUL byte.
//
// o may be nil, in which case "" is returned.
func (o *NewWindowOptions) GetTitle() string {
	if o == nil {
		return ""
	}
	return sanitizeUTF8(o.Title, 4096)
}

// todo: move this
func sanitizeUTF8(s string, n int) string {
	if n < len(s) {
		s = s[:n]
	}
	i := 0
	for i < len(s) {
		r, n := utf8.DecodeRuneInString(s[i:])
		if r == 0 || (r == utf8.RuneError && n == 1) {
			break
		}
		i += n
	}
	return s[:i]
}

// DrawOptions are optional arguments to Draw.
type DrawOptions struct {
	// TODO: transparency in [0x0000, 0xffff]?
	// TODO: scaler (nearest neighbor vs linear)?
}
