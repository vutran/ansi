package ansi_test

import (
	"fmt"
	"github.com/vutran/ansi"
	"testing"
)

func TestBold(t *testing.T) {
	fmt.Print(ansi.Bold())
	fmt.Println("This should be bold")
	fmt.Print(ansi.BoldOff())
}

func TestFaint(t *testing.T) {
	fmt.Print(ansi.Faint())
	fmt.Println("This should be faint")
	fmt.Print(ansi.FaintOff())
}

func TestItalic(t *testing.T) {
	fmt.Print(ansi.Italic())
	fmt.Println("This should be italic")
	fmt.Print(ansi.ItalicOff())
}

func TestUnderline(t *testing.T) {
	fmt.Print(ansi.Underline())
	fmt.Println("This should be underline")
	fmt.Print(ansi.UnderlineOff())
}

func TestCrossedOut(t *testing.T) {
	fmt.Print(ansi.CrossedOut())
	fmt.Println("This should be crossed-out")
	fmt.Print(ansi.CrossedOutOff())
}

func TestBlackWithWhiteBg(t *testing.T) {
	fmt.Print(ansi.Black() + ansi.WhiteBg())
	fmt.Println("This should be black with a white background")
	fmt.Print(ansi.DefaultColor() + ansi.DefaultBg())
}

func TestWhiteWithBlackBg(t *testing.T) {
	fmt.Print(ansi.White() + ansi.BlackBg())
	fmt.Println("This should be white with a black background")
	fmt.Print(ansi.DefaultColor() + ansi.DefaultBg())
}

func TestRedWithYellowBg(t *testing.T) {
	fmt.Print(ansi.Red() + ansi.YellowBg())
	fmt.Println("This should be red with a yellow background")
	fmt.Print(ansi.DefaultColor() + ansi.DefaultBg())
}

func TestYellowWithRedBg(t *testing.T) {
	fmt.Print(ansi.Yellow() + ansi.RedBg())
	fmt.Println("This should be yellow with a red background")
	fmt.Print(ansi.DefaultColor() + ansi.DefaultBg())
}

func TestBlueWithGreenBg(t *testing.T) {
	fmt.Print(ansi.Blue() + ansi.GreenBg())
	fmt.Println("This should be blue with a green background")
	fmt.Print(ansi.DefaultColor() + ansi.DefaultBg())
}

func TestGreenWithBlueBg(t *testing.T) {
	fmt.Print(ansi.Green() + ansi.BlueBg())
	fmt.Println("This should be green with a blue background")
	fmt.Print(ansi.DefaultColor() + ansi.DefaultBg())
}

func TestCyanWithMagentaBg(t *testing.T) {
	fmt.Print(ansi.Cyan() + ansi.MagentaBg())
	fmt.Println("This should be cyan with a magenta background")
	fmt.Print(ansi.DefaultColor() + ansi.DefaultBg())
}

func TestMagentaWithCyanBg(t *testing.T) {
	fmt.Print(ansi.Magenta() + ansi.CyanBg())
	fmt.Println("This should be magenta with a cyan background")
	fmt.Print(ansi.DefaultColor() + ansi.DefaultBg())
}
