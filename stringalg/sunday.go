package stringalg

import (
	"bytes"
)

func strStr(origin, aim string) int {
	lo := len(origin)
	la := len(aim)
	if lo == 0 || la == 0 {
		return 0
	}
	if lo < la {
		return -1
	}

	oIndex, aIndex := 0, 0

	for aIndex < la {
		if origin[oIndex] == aim[aIndex] {
			oIndex++
			aIndex++
		} else {
			nextCharIndex := oIndex - aIndex + la
			if nextCharIndex >= lo {
				return -1
			}
			if i := bytes.LastIndexByte([]byte(aim), origin[nextCharIndex]); i == -1 {
				oIndex = nextCharIndex + 1
			} else {
				oIndex = nextCharIndex - i
			}
			aIndex = 0
		}
	}
	return oIndex - aIndex
}
