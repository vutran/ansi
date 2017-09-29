package colors

import (
	"fmt"
	"testing"
)

const emptySpace = "               "

func TestBlack(t *testing.T) {
	fmt.Println(Black("This is black"))
}

func TestRed(t *testing.T) {
	fmt.Println(Red("This is red"))
}

func TestGreen(t *testing.T) {
	fmt.Println(Green("This is green"))
}

func TestYellow(t *testing.T) {
	fmt.Println(Yellow("This is yellow"))
}

func TestBlue(t *testing.T) {
	fmt.Println(Blue("This is blue"))
}

func TestMagenta(t *testing.T) {
	fmt.Println(Magenta("This is magenta"))
}

func TestCyan(t *testing.T) {
	fmt.Println(Cyan("This is cyan"))
}

func TestWhite(t *testing.T) {
	fmt.Println(White("This is white"))
}

func TestBlackBg(t *testing.T) {
	fmt.Println(BlackBg(emptySpace))
}

func TestRedBg(t *testing.T) {
	fmt.Println(RedBg(emptySpace))
}

func TestGreenBg(t *testing.T) {
	fmt.Println(GreenBg(emptySpace))
}

func TestYellowBg(t *testing.T) {
	fmt.Println(YellowBg(emptySpace))
}

func TestBlueBg(t *testing.T) {
	fmt.Println(BlueBg(emptySpace))
}

func TestMagentaBg(t *testing.T) {
	fmt.Println(MagentaBg(emptySpace))
}

func TestCyanBg(t *testing.T) {
	fmt.Println(CyanBg(emptySpace))
}

func TestWhiteBg(t *testing.T) {
	fmt.Println(WhiteBg(emptySpace))
}
