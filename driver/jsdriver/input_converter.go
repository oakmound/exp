// +build js

package jsdriver

import (
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/mouse"
)

var (
	jsMouseButtons = map[int]mouse.Button{
		0: mouse.ButtonLeft,
		1: mouse.ButtonMiddle,
		2: mouse.ButtonRight,
	}
	jsKeys = map[int]key.Code{
		9:  key.CodeTab,
		13: key.CodeReturnEnter,
		32: key.CodeSpacebar,
		33: key.CodePageUp,
		34: key.CodePageDown,
		35: key.CodeEnd,
		36: key.CodeHome,
		37: key.CodeLeftArrow,
		38: key.CodeUpArrow,
		39: key.CodeRightArrow,
		40: key.CodeDownArrow,
		45: key.CodeInsert,
		46: key.CodeDeleteBackspace,
		49: key.Code0,
		50: key.Code1,
		51: key.Code2,
		52: key.Code3,
		53: key.Code4,
		54: key.Code5,
		55: key.Code6,
		56: key.Code7,
		57: key.Code8,
		58: key.Code9,
		// 59: colon
		// 60: <
		61: key.CodeEqualSign,
		// 62: >
		// 63: ?
		// 64: @
		65: key.CodeA,
		66: key.CodeB,
		67: key.CodeC,
		68: key.CodeD,
		69: key.CodeE,
		70: key.CodeF,
		71: key.CodeG,
		72: key.CodeH,
		73: key.CodeI,
		74: key.CodeJ,
		75: key.CodeK,
		76: key.CodeL,
		77: key.CodeM,
		78: key.CodeN,
		79: key.CodeO,
		80: key.CodeP,
		81: key.CodeQ,
		82: key.CodeR,
		83: key.CodeS,
		84: key.CodeT,
		85: key.CodeU,
		86: key.CodeV,
		87: key.CodeW,
		88: key.CodeX,
		89: key.CodeY,
		90: key.CodeZ,
	}
)

func jsMouseButton(i int) mouse.Button {
	if v, ok := jsMouseButtons[i]; ok {
		return v
	}
	return mouse.ButtonNone
}

func jsKey(i int) key.Code {
	if v, ok := jsKeys[i]; ok {
		return v
	}
	return key.CodeUnknown
}
