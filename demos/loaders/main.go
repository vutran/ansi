package main

import (
	"github.com/vutran/ansi"
	"github.com/vutran/ansi/loaders"
	"time"
)

func main() {
	s := ansi.Loader(loaders.Dots)
	s.SetValue("Loading")
	s.Start()
	time.Sleep(2 * time.Second)
	s.SetValue("Finalizing...")
	time.Sleep(2 * time.Second)
	s.Stop()
}
