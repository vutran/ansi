package styles_test

import (
	"fmt"
	"github.com/vutran/ansi/styles"
	"testing"
)

func TestBold(t *testing.T) {
	fmt.Println(styles.Bold("Bold"))
}

func TestFaint(t *testing.T) {
	fmt.Println(styles.Faint("Faint"))
}

func TestItalic(t *testing.T) {
	fmt.Println(styles.Italic("Italic"))
}

func TestUnderline(t *testing.T) {
	fmt.Println(styles.Underline("Underline"))
}

func TestCrossedOut(t *testing.T) {
	fmt.Println(styles.CrossedOut("CrossedOut"))
}
