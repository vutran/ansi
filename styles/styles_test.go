package styles

import (
	"fmt"
	"testing"
)

func TestBold(t *testing.T) {
	fmt.Println(Bold("Hello, world!"))
}

func TestFaint(t *testing.T) {
	fmt.Println(Faint("Hello, world!"))
}

func TestItalic(t *testing.T) {
	fmt.Println(Italic("Hello, world!"))
}

func TestUnderline(t *testing.T) {
	fmt.Println(Underline("Hello, world!"))
}

func TestCrossedOut(t *testing.T) {
	fmt.Println(CrossedOut("Hello, world!"))
}
