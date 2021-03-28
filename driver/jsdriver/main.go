// +build js

package jsdriver

import "github.com/oakmound/shiny/screen"

func Main(f func(screen.Screen)) {
	f(new(JSScreen))
}
