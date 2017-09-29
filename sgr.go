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
