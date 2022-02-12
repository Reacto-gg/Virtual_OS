package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func Player() {
	// a := app.New()
	w := a.NewWindow("Audio Player")
	w.Resize(fyne.NewSize(500, 300))

	//Mp3 Player Background Image
	playerBackgroundImage := canvas.NewImageFromFile("./images/playerBackgroundImage.jpg")
	playerBackgroundImage.Translucency = 0.4

	var format beep.Format
	var streamer beep.StreamSeekCloser
	var file fyne.URIReadCloser

	lock := true
	var start int
	label := widget.NewLabel("")
	label2 := widget.NewLabel("")
	label.TextStyle = fyne.TextStyle{Bold: true}
	label2.TextStyle = fyne.TextStyle{Bold: true}

	toolbarPlayer := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			if start == 0 {
				// fmt.Println(file)
				streamer, format, _ = mp3.Decode(file)
				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
				// speaker.Play(streamer)
				go func() {
					done := make(chan bool)
					speaker.Play(beep.Seq(streamer, beep.Callback(func() {
						done <- true
					})))

					for {
						select {
						case <-done:
							label.SetText(shortDur(format.SampleRate.D(streamer.Len()).Round(time.Second))) //Only Use for matching the time displaying
							return
						case <-time.After(time.Second):
							speaker.Lock()
							// fmt.Println(format.SampleRate.D(streamer.Position()).Round(time.Second))
							label.SetText(shortDur(format.SampleRate.D(streamer.Position()).Round(time.Second)))
							speaker.Unlock()
						}
					}
				}()
				label2.SetText("/  " + shortDur(format.SampleRate.D(streamer.Len()).Round(time.Second)))
				start = 1
			}

			if lock == false {
				speaker.Unlock()
				lock = true
			}
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if lock == true {
				speaker.Lock()
				lock = false
			}
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()
		}),
		widget.NewToolbarSpacer(),
	)

	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), label, label2, layout.NewSpacer())

	songName := widget.NewLabel("Your PlayList Empty")
	songName.Alignment = fyne.TextAlignCenter

	//Browse the file from local storage
	PlayFile := widget.NewButton("Browse...", func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {
			songName.Text = uc.URI().Name()
			songName.Refresh()
			file = uc
			start = 0
		}, w)
		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))

	})

	funContainer := container.NewVBox(
		songName,
		centered,
		toolbarPlayer,
		PlayFile,
	)
	//Set the container using combind
	content := fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		playerBackgroundImage,
		funContainer,
	)

	r, _ := fyne.LoadResourceFromPath("./images/musicLogo.png")
	w.SetIcon(r)
	w.SetPadded(false)
	w.CenterOnScreen()
	w.SetOnClosed(func() {
		speaker.Clear()
		fmt.Println("Player Closed")
	})

	w.SetContent(content)

	w.Show()
}

//type time.Duration  to string Conversion function
func shortDur(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}
