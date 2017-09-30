package ansi_test

import (
	"fmt"
	"github.com/vutran/ansi"
	"github.com/vutran/ansi/loaders"
	"testing"
	"time"
)

var styles = [][]string{
	loaders.Circle,
	loaders.Clock,
	loaders.Dots,
	loaders.Lines,
	loaders.Moon,
	loaders.Triangle,
}

func run(loader []string) {
	s := ansi.Loader(loader)
	s.SetValue("Loading")
	s.Start()
	time.Sleep(1 * time.Second)
	s.SetValue("Finalizing...")
	time.Sleep(1 * time.Second)
	s.Stop()

	// reset cursors
	fmt.Print(ansi.EraseLine(2))
	fmt.Print(ansi.CursorStart(1))
}

func TestLoader(t *testing.T) {
	for _, l := range styles {
		t.Run("test", func(t *testing.T) {
			run(l)
		})
	}
}
