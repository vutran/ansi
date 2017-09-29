package ansi

import (
	"testing"
)

func TestEsc(t *testing.T) {
	if Esc != "\u001b" {
		t.Error("Esc code incorrect")
	}
}
