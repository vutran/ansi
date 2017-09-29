package styles

import (
	"github.com/vutran/fansi"
)

func Bold(value string) string {
	return fansi.Bold() + value + fansi.BoldOff()
}

func Faint(value string) string {
	return fansi.Faint() + value + fansi.FaintOff()
}

func Italic(value string) string {
	return fansi.Italic() + value + fansi.ItalicOff()
}

func Underline(value string) string {
	return fansi.Underline() + value + fansi.UnderlineOff()
}

func CrossedOut(value string) string {
	return fansi.CrossedOut() + value + fansi.CrossedOutOff()
}
