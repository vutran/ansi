package fansi

import (
	"testing"
)

func TestSelectGraphicsRendition(t *testing.T) {
	if SelectGraphicsRendition(0) != "\u001b[0m" {
		t.Error("Should be SGR 0")
	}

	if SelectGraphicsRendition(10) != "\u001b[10m" {
		t.Error("Should be SGR 10")
	}
}
