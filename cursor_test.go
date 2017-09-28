package fansi

import (
	"testing"
)

func TestCursorUp(t *testing.T) {
	if CursorUp(1) != "\u001B[1A" {
		t.Error("CursorUp(1) failed")
	}
	if CursorUp(10) != "\u001B[10A" {
		t.Error("CursorUp(10) failed")
	}
}

func TestCursorDown(t *testing.T) {
	if CursorDown(1) != "\u001B[1B" {
		t.Error("CursorDown(1) failed")
	}
	if CursorDown(10) != "\u001B[10B" {
		t.Error("CursorDown(10) failed")
	}
}

func TestCursorForward(t *testing.T) {
	if CursorForward(1) != "\u001B[1C" {
		t.Error("CursorForward(1) failed")
	}
	if CursorForward(10) != "\u001B[10C" {
		t.Error("CursorForward(10) failed")
	}
}

func TestCursorBackward(t *testing.T) {
	if CursorBackward(1) != "\u001B[1D" {
		t.Error("CursorBackward(1) failed")
	}
	if CursorBackward(10) != "\u001B[10D" {
		t.Error("CursorBackward(10) failed")
	}
}
