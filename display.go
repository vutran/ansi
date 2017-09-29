package fansi

import (
	"strconv"
)

func EraseDisplay(code int) string {
	return Esc + "[" + strconv.Itoa(code) + "2J"
}

func EraseLine(code int) string {
	return Esc + "[" + strconv.Itoa(code) + "K"
}

func SelectGraphicsRendition(code int) string {
	return Esc + "[" + strconv.Itoa(code) + "m"
}
