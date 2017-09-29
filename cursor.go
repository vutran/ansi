package fansi

import (
	"strconv"
)

const HideCursor = Esc + "[?25l"

const ShowCursor = Esc + "[?25h"

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

const SaveCursorPosition = Esc + "[s"

const RestoreCursorPosition = Esc + "[u"

const GetCursorPosition = Esc + "[6n"
