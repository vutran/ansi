package fansi

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
