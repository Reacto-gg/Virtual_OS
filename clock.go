package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2/canvas"
)

func Clock(clockText *canvas.Text) {
	cc := true

	ShowTime(clockText, cc)
	go func() {
		t := time.NewTicker(time.Second)

		for range t.C {
			// Color changing option
			if cc {
				cc = !cc
			} else {
				cc = !cc
			}
			// call ShowTime Function
			ShowTime(clockText, cc)

		}
	}()
}

func ShowTime(clockText *canvas.Text, cc bool) {

	times := time.Now().Format(time.Stamp)
	clockText.Text = times
	clockText.Refresh()
	if cc {
		clockText.Color = color.RGBA{0, 236, 249, 255}
	} else {
		clockText.Color = color.RGBA{255, 255, 1, 255}
	}

}
