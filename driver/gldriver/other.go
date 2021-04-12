// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (!darwin || ios || !cgo) && (!linux || android || !cgo) && (!openbsd || !cgo) && !windows
// +build !darwin ios !cgo
// +build !linux android !cgo
// +build !openbsd !cgo
// +build !windows
// +build !openbsd

package gldriver

import (
	"fmt"
	"runtime"

	"github.com/oakmound/shiny/screen"
)

func newWindow(opts screen.WindowGenerator) (uintptr, error) { return 0, nil }

func moveWindow(w *windowImpl, opts screen.WindowGenerator) error { return nil }

const useLifecycler = true
const handleSizeEventsAtChannelReceive = true

func initWindow(id *windowImpl) {}
func showWindow(id *windowImpl) {}
func closeWindow(id uintptr)    {}
func drawLoop(w *windowImpl)    {}

func surfaceCreate() error             { return errUnsupported }
func main(f func(screen.Screen)) error { return errUnsupported }
