package ansi

import (
	"strconv"
)

func HideCursor() string {
	return Esc + "[?25l"
}

func ShowCursor() string {
	return Esc + "[?25h"
}

func CursorUp(value int) string {
	return Esc + "[" + strconv.Itoa(value) + "A"
}

func CursorDown(value int) string {
	return Esc + "[" + strconv.Itoa(value) + "B"
}

func CursorForward(value int) string {
	return Esc + "[" + strconv.Itoa(value) + "C"
}

func CursorBackward(value int) string {
	return Esc + "[" + strconv.Itoa(value) + "D"
}

func SaveCursorPosition() string {
	return Esc + "[s"
}

func RestoreCursorPosition() string {
	return Esc + "[u"
}

func GetCursorPosition() string {
	return Esc + "[6n"
}
