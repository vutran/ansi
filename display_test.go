package ansi_test

import (
	"github.com/vutran/ansi"
	"testing"
)

func TestSelectGraphicsRendition(t *testing.T) {
	if ansi.SelectGraphicsRendition(0) != "\u001b[0m" {
		t.Error("Should be SGR 0")
	}

	if ansi.SelectGraphicsRendition(10) != "\u001b[10m" {
		t.Error("Should be SGR 10")
	}
}
