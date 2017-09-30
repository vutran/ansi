package colors_test

import (
	"fmt"
	"github.com/vutran/ansi/colors"
	"testing"
)

const emptySpace = "               "

func TestBlack(t *testing.T) {
	fmt.Println(colors.Black("This is black"))
}

func TestRed(t *testing.T) {
	fmt.Println(colors.Red("This is red"))
}

func TestGreen(t *testing.T) {
	fmt.Println(colors.Green("This is green"))
}

func TestYellow(t *testing.T) {
	fmt.Println(colors.Yellow("This is yellow"))
}

func TestBlue(t *testing.T) {
	fmt.Println(colors.Blue("This is blue"))
}

func TestMagenta(t *testing.T) {
	fmt.Println(colors.Magenta("This is magenta"))
}

func TestCyan(t *testing.T) {
	fmt.Println(colors.Cyan("This is cyan"))
}

func TestWhite(t *testing.T) {
	fmt.Println(colors.White("This is white"))
}

func TestBlackBg(t *testing.T) {
	fmt.Println(colors.BlackBg(emptySpace))
}

func TestRedBg(t *testing.T) {
	fmt.Println(colors.RedBg(emptySpace))
}

func TestGreenBg(t *testing.T) {
	fmt.Println(colors.GreenBg(emptySpace))
}

func TestYellowBg(t *testing.T) {
	fmt.Println(colors.YellowBg(emptySpace))
}

func TestBlueBg(t *testing.T) {
	fmt.Println(colors.BlueBg(emptySpace))
}

func TestMagentaBg(t *testing.T) {
	fmt.Println(colors.MagentaBg(emptySpace))
}

func TestCyanBg(t *testing.T) {
	fmt.Println(colors.CyanBg(emptySpace))
}

func TestWhiteBg(t *testing.T) {
	fmt.Println(colors.WhiteBg(emptySpace))
}
