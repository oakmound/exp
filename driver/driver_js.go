// +build !android
// +build !darwin
// +build !linux
// +build !windows
// +build !dragonfly
// +build !openbsd
// +build js

package driver

import "github.com/oakmound/shiny/screen"

func main(f func(screen.Screen)) {
	jsdriver.Main(f)
}
