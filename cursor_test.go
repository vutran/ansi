package ansi_test

import (
	"github.com/vutran/ansi"
	"testing"
)

func TestCursorUp(t *testing.T) {
	if ansi.CursorUp(1) != "\u001B[1A" {
		t.Error("CursorUp(1) failed")
	}
	if ansi.CursorUp(10) != "\u001B[10A" {
		t.Error("CursorUp(10) failed")
	}
}

func TestCursorDown(t *testing.T) {
	if ansi.CursorDown(1) != "\u001B[1B" {
		t.Error("CursorDown(1) failed")
	}
	if ansi.CursorDown(10) != "\u001B[10B" {
		t.Error("CursorDown(10) failed")
	}
}

func TestCursorForward(t *testing.T) {
	if ansi.CursorForward(1) != "\u001B[1C" {
		t.Error("CursorForward(1) failed")
	}
	if ansi.CursorForward(10) != "\u001B[10C" {
		t.Error("CursorForward(10) failed")
	}
}

func TestCursorBackward(t *testing.T) {
	if ansi.CursorBackward(1) != "\u001B[1D" {
		t.Error("CursorBackward(1) failed")
	}
	if ansi.CursorBackward(10) != "\u001B[10D" {
		t.Error("CursorBackward(10) failed")
	}
}
