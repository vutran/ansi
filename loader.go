package ansi

import (
	"fmt"
	"time"
)

func cleanup() {
	fmt.Println(EraseLine(2))
	fmt.Print(ShowCursor())
}

type LoaderControl struct {
	anim   []string
	value  string
	cursor int
	max    int
	ticker *time.Ticker
	done   chan bool
}

func (loader *LoaderControl) Start() {
	go func() {
		for {
			select {
			case <-loader.done:
				return
			case <-loader.ticker.C:
				fmt.Print(HideCursor())
				fmt.Print(EraseLine(2))
				fmt.Print(CursorStart(1))

				if loader.cursor >= loader.max {
					loader.cursor = 0
				}

				fmt.Print(loader.anim[loader.cursor] + " " + loader.value)

				loader.cursor++
			}
		}
	}()
}

func (loader *LoaderControl) Stop() {
	loader.done <- true
	cleanup()
}

func (loader *LoaderControl) SetValue(value string) {
	loader.value = value
}

func Loader(anim []string, value string) *LoaderControl {
	return &LoaderControl{
		cursor: 0,
		anim:   anim,
		value:  value,
		max:    len(anim),
		ticker: time.NewTicker(time.Millisecond * 100),
		done:   make(chan bool),
	}
}
