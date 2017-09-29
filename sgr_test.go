package ansi

import (
	"fmt"
	"testing"
)

func TestBold(t *testing.T) {
	fmt.Print(Bold())
	fmt.Println("This should be bold")
	fmt.Print(BoldOff())
}

func TestFaint(t *testing.T) {
	fmt.Print(Faint())
	fmt.Println("This should be faint")
	fmt.Print(FaintOff())
}

func TestItalic(t *testing.T) {
	fmt.Print(Italic())
	fmt.Println("This should be italic")
	fmt.Print(ItalicOff())
}

func TestUnderline(t *testing.T) {
	fmt.Print(Underline())
	fmt.Println("This should be underline")
	fmt.Print(UnderlineOff())
}

func TestCrossedOut(t *testing.T) {
	fmt.Print(CrossedOut())
	fmt.Println("This should be crossed-out")
	fmt.Print(CrossedOutOff())
}

func TestBlackWithWhiteBg(t *testing.T) {
	fmt.Print(Black() + WhiteBg())
	fmt.Println("This should be black with a white background")
	fmt.Print(DefaultColor() + DefaultBg())
}

func TestWhiteWithBlackBg(t *testing.T) {
	fmt.Print(White() + BlackBg())
	fmt.Println("This should be white with a black background")
	fmt.Print(DefaultColor() + DefaultBg())
}

func TestRedWithYellowBg(t *testing.T) {
	fmt.Print(Red() + YellowBg())
	fmt.Println("This should be red with a yellow background")
	fmt.Print(DefaultColor() + DefaultBg())
}

func TestYellowWithRedBg(t *testing.T) {
	fmt.Print(Yellow() + RedBg())
	fmt.Println("This should be yellow with a red background")
	fmt.Print(DefaultColor() + DefaultBg())
}

func TestBlueWithGreenBg(t *testing.T) {
	fmt.Print(Blue() + GreenBg())
	fmt.Println("This should be blue with a green background")
	fmt.Print(DefaultColor() + DefaultBg())
}

func TestGreenWithBlueBg(t *testing.T) {
	fmt.Print(Green() + BlueBg())
	fmt.Println("This should be green with a blue background")
	fmt.Print(DefaultColor() + DefaultBg())
}

func TestCyanWithMagentaBg(t *testing.T) {
	fmt.Print(Cyan() + MagentaBg())
	fmt.Println("This should be cyan with a magenta background")
	fmt.Print(DefaultColor() + DefaultBg())
}

func TestMagentaWithCyanBg(t *testing.T) {
	fmt.Print(Magenta() + CyanBg())
	fmt.Println("This should be magenta with a cyan background")
	fmt.Print(DefaultColor() + DefaultBg())
}
