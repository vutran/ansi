package ansi_test

import (
	"github.com/vutran/ansi"
	"testing"
)

func TestEsc(t *testing.T) {
	if ansi.Esc != "\u001b" {
		t.Error("Esc code incorrect")
	}
}
