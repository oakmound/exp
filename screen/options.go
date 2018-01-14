package screen

import "unicode/utf8"

// Todo: flesh out options as a generator style constructor
// with (optional) variadic arguments

// A WindowGenerator can generate windows based on various new window settings.
type WindowGenerator struct {
	// Width and Height specify the dimensions of the new window. If Width
	// or Height are zero, a driver-dependent default will be used for each
	// zero value dimension.
	Width, Height int

	// Title specifies the window title.
	Title string

	// Fullscreen determines whether the new window will be fullscreen or not.
	Fullscreen bool

	// NoScaling determines whether the new window will have scaling allowed.
	// With a zero value of false, scaling is allowed.
	NoScaling bool

	// X and Y determine the location the new window should be created at. If
	// either are zero, a driver-dependant default will be used for each zero
	// value. If Fullscreen is true, these values will be ignored.
	X, Y int

	// BorderStyle describes the presence of buttons, menus, and thickness of
	// generated windows' borders.
	BorderStyle
}

// A WindowOption is any function that sets up a WindowGenerator.
type WindowOption func(*WindowGenerator)

// BorderStyle describes whether there are maximize, minimize, etc buttons
// the size of the border (thick, thin, stylized)
// todo...
type BorderStyle int

// Title sets a sanitized form of the input string. In particular, its length will
// not exceed 4096, and it may be further truncated so that it is valid UTF-8
// and will not contain the NUL byte.
func Title(s string) WindowOption {
	return func(g *WindowGenerator) {
		g.Title = sanitizeUTF8(s, 4096)
	}
}

// Dimensions sets the width and height of new windows
func Dimensions(w, h int) WindowOption {
	return func(g *WindowGenerator) {
		g.Width = w
		g.Height = h
	}
}

// NewWindowGenerator creates a window generator with zero values,
// then calls all options passed in on it.
func NewWindowGenerator(opts ...WindowOption) WindowGenerator {
	wg := &WindowGenerator{}
	for _, o := range opts {
		o(wg)
	}
	return *wg
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
