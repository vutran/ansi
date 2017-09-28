package fansi

import (
	"strconv"
)

const EraseDisplay = Esc + "2J"

const EraseLine = Esc + "K"

func SelectGraphicsRendition(code int) string {
	return Esc + strconv.Itoa(code) + "m"
}
