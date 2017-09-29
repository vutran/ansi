package fansi

func Reset() string {
	return SelectGraphicsRendition(0)
}

func Bold() string {
	return SelectGraphicsRendition(1)
}

func BoldOff() string {
	// 21 not widely supported, so we're using 22 here instead
	return SelectGraphicsRendition(22)
}

func Faint() string {
	return SelectGraphicsRendition(2)
}

func FaintOff() string {
	return SelectGraphicsRendition(22)
}

func Italic() string {
	return SelectGraphicsRendition(3)
}

func ItalicOff() string {
	return SelectGraphicsRendition(23)
}

func Underline() string {
	return SelectGraphicsRendition(4)
}

func UnderlineOff() string {
	return SelectGraphicsRendition(24)
}

func BlinkSlow() string {
	return SelectGraphicsRendition(5)
}

func BlinkFast() string {
	return SelectGraphicsRendition(6)
}

func BlinkOff() string {
	return SelectGraphicsRendition(25)
}

func ImageNegative() string {
	return SelectGraphicsRendition(7)
}

func ImagePositive() string {
	return SelectGraphicsRendition(27)
}

func Conceal() string {
	return SelectGraphicsRendition(8)
}

func Reveal() string {
	return SelectGraphicsRendition(28)
}

func CrossedOut() string {
	return SelectGraphicsRendition(9)
}

func CrossedOutOff() string {
	return SelectGraphicsRendition(29)
}

func DefaultColor() string {
	return SelectGraphicsRendition(39)
}

func Black() string {
	return SelectGraphicsRendition(30)
}

func Red() string {
	return SelectGraphicsRendition(31)
}

func Green() string {
	return SelectGraphicsRendition(32)
}

func Yellow() string {
	return SelectGraphicsRendition(33)
}

func Blue() string {
	return SelectGraphicsRendition(34)
}

func Magenta() string {
	return SelectGraphicsRendition(35)
}

func Cyan() string {
	return SelectGraphicsRendition(36)
}

func White() string {
	return SelectGraphicsRendition(37)
}

func DefaultBg() string {
	return SelectGraphicsRendition(49)
}

func BlackBg() string {
	return SelectGraphicsRendition(40)
}

func RedBg() string {
	return SelectGraphicsRendition(41)
}

func GreenBg() string {
	return SelectGraphicsRendition(42)
}

func YellowBg() string {
	return SelectGraphicsRendition(43)
}

func BlueBg() string {
	return SelectGraphicsRendition(44)
}

func MagentaBg() string {
	return SelectGraphicsRendition(45)
}

func CyanBg() string {
	return SelectGraphicsRendition(46)
}

func WhiteBg() string {
	return SelectGraphicsRendition(47)
}
