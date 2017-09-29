package colors

import (
	"github.com/vutran/fansi"
)

func Black(value string) string {
	return fansi.Black() + value + fansi.DefaultColor()
}

func Red(value string) string {
	return fansi.Red() + value + fansi.DefaultColor()
}

func Green(value string) string {
	return fansi.Green() + value + fansi.DefaultColor()
}

func Yellow(value string) string {
	return fansi.Yellow() + value + fansi.DefaultColor()
}

func Blue(value string) string {
	return fansi.Blue() + value + fansi.DefaultColor()
}

func Magenta(value string) string {
	return fansi.Magenta() + value + fansi.DefaultColor()
}

func Cyan(value string) string {
	return fansi.Cyan() + value + fansi.DefaultColor()
}

func White(value string) string {
	return fansi.White() + value + fansi.DefaultColor()
}

func BlackBg(value string) string {
	return fansi.BlackBg() + value + fansi.DefaultBg()
}

func RedBg(value string) string {
	return fansi.RedBg() + value + fansi.DefaultBg()
}

func GreenBg(value string) string {
	return fansi.GreenBg() + value + fansi.DefaultBg()
}

func YellowBg(value string) string {
	return fansi.YellowBg() + value + fansi.DefaultBg()
}

func BlueBg(value string) string {
	return fansi.BlueBg() + value + fansi.DefaultBg()
}

func MagentaBg(value string) string {
	return fansi.MagentaBg() + value + fansi.DefaultBg()
}

func CyanBg(value string) string {
	return fansi.CyanBg() + value + fansi.DefaultBg()
}

func WhiteBg(value string) string {
	return fansi.WhiteBg() + value + fansi.DefaultBg()
}
