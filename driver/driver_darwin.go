// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package driver

import (
	"github.com/oakmound/shiny/driver/gldriver"
	"github.com/oakmound/shiny/screen"
)

func main(f func(screen.Screen)) {
	gldriver.Main(f)
}
