// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package colornames

import (
	"image/color"
	"testing"
)

func TestColornames(t *testing.T) {
	if len(Map) != len(Names) {
		t.Fatalf("Map and Names have different length: %d vs %d", len(Map), len(Names))
	}

	for name, want := range testCases {
		got, ok := Map[name]
		if !ok {
			t.Errorf("Did not find %s", name)
			continue
		}
		if got != want {
			t.Errorf("%s:\ngot  %x\nwant %x", name, got, want)
		}
	}
}

var testCases = map[string]color.RGBA{
	"Red500":         {0xf4, 0x43, 0x36, 0xff},
	"Red50":          {0xff, 0xeb, 0xee, 0xff},
	"Red900":         {0xb7, 0x1c, 0x1c, 0xff},
	"RedA700":        {0xd5, 0x00, 0x00, 0xff},
	"Pink300":        {0xf0, 0x62, 0x92, 0xff},
	"Purple100":      {0xe1, 0xbe, 0xe7, 0xff},
	"Cyan400":        {0x26, 0xc6, 0xda, 0xff},
	"LightGreen800":  {0x55, 0x8b, 0x2f, 0xff},
	"DeepOrangeA200": {0xff, 0x6e, 0x40, 0xff},
	"Brown50":        {0xef, 0xeb, 0xe9, 0xff},
	"Grey500":        {0x9e, 0x9e, 0x9e, 0xff},
	"Grey600":        {0x75, 0x75, 0x75, 0xff},
	"Grey700":        {0x61, 0x61, 0x61, 0xff},
	"BlueGrey400":    {0x78, 0x90, 0x9c, 0xff},
	"Black":          {0x00, 0x00, 0x00, 0xff},
	"White":          {0xff, 0xff, 0xff, 0xff},
}
