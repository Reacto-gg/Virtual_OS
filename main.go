package main

import (

	// t "Fyne-app/test/best"

	"fmt"
	"log"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var a fyne.App = app.New()
var w fyne.Window

func main() {

	w := a.NewWindow("	Pepcoding Contest 2021 - By Kalyan Santra")
	w.Resize(fyne.NewSize(1280, 720))

	//Logo Icondssss....................
	textEditorLogo, _ := fyne.LoadResourceFromPath("./images/textEditorLogo.png")
	weatherLogo, _ := fyne.LoadResourceFromPath("./images/weatherLogo.png")
	musicLogo, _ := fyne.LoadResourceFromPath("./images/musicLogo.png")
	calculatorLogo, _ := fyne.LoadResourceFromPath("./images/calcLogo.png")
	clockLogo, _ := fyne.LoadResourceFromPath("./images/clockLogo.png")
	backgroundChangeLogo, _ := fyne.LoadResourceFromPath("./images/backgroundChangeLogo.png")
	galleryLogo, _ := fyne.LoadResourceFromPath("./images/galleryLogo.png")
	rpsGameLogo, _ := fyne.LoadResourceFromPath("./images/rpcLogo.png")

	// label := widget.NewLabel("There")
	//  Windows Startup Sound
	go StartupSound()

	a.Settings().SetTheme(theme.LightTheme())
	//BackGround Image
	lightBackground := canvas.NewImageFromFile("./images/light.png")
	darkBackground := canvas.NewImageFromFile("./images/dark.jpg")
	//Default case
	darkBackground.Hide()
	imgContainer := container.NewGridWithColumns(1, lightBackground, darkBackground)

	//Clock Text
	clockText := canvas.NewText("", color.White)
	clockText.TextSize = 50
	clockText.Alignment = fyne.TextAlignCenter

	go Clock(clockText)
	clockTextShow := true

	darkTheme := false
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		// widget.NewToolbarSeparator(),
		widget.NewToolbarAction(textEditorLogo, func() {
			TextEditor()
		}),
		widget.NewToolbarAction(weatherLogo, func() {
			WeatherApp()
		}),
		widget.NewToolbarAction(galleryLogo, func() {
			Gallery(lightBackground, darkBackground, darkTheme)
		}),

		widget.NewToolbarAction(musicLogo, func() {
			Player()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			if darkTheme {
				a.Settings().SetTheme(theme.LightTheme())
				// lightBackground.File = fmt.Sprintf("./light.png")
				darkBackground.Hide()
				lightBackground.Show()
				darkTheme = !darkTheme

			} else {
				a.Settings().SetTheme(theme.DarkTheme())
				// lightBackground.File = fmt.Sprintf("./dark.jpg")
				lightBackground.Hide()
				darkBackground.Show()
				darkTheme = !darkTheme
			}
		}),
		widget.NewToolbarAction(calculatorLogo, func() {
			log.Println("Display help")
			Calculator()
		}),
		widget.NewToolbarAction(clockLogo, func() {
			if clockTextShow {
				clockText.Hide()
				clockTextShow = !clockTextShow
			} else {
				clockText.Show()
				clockTextShow = !clockTextShow
			}
		}),

		widget.NewToolbarAction(backgroundChangeLogo, func() {

			openImageDialog := dialog.NewFileOpen(func(r fyne.URIReadCloser, _ error) {

				selectImage := r.URI().Name()
				// fmt.Printf("value  = %v", r.URI().Path())
				source := r.URI().Path()
				source = source[:len(source)-len(selectImage)]

				if darkTheme {
					darkBackground.File = fmt.Sprintf(source + selectImage)
					darkBackground.Refresh()
				} else {
					lightBackground.File = fmt.Sprintf(source + selectImage)
					lightBackground.Refresh()
				}
			}, w)
			openImageDialog.SetFilter(storage.NewExtensionFileFilter([]string{".jpg", ".png", ".jpeg", ".PNG"}))
			openImageDialog.Show()

		}),
		widget.NewToolbarAction(rpsGameLogo, func() {
			RPSGame()
		}),
		widget.NewToolbarSpacer(),
	)

	textWithImgContainer := fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		imgContainer,
		clockText)

	content := container.NewBorder(nil, toolbar, nil, nil, textWithImgContainer)

	iconRes, _ := fyne.LoadResourceFromPath("./images/osLogo.png")
	w.SetIcon(iconRes)

	w.CenterOnScreen()
	w.SetPadded(false)
	// Full Screen Funtionality
	// w.SetFullScreen(true)
	fullScreen := true
	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == fyne.KeyF11 {
			if fullScreen {
				w.SetFullScreen(false)
				fullScreen = !fullScreen
			} else {
				w.SetFullScreen(true)
				fullScreen = !fullScreen
			}
		}
	})

	w.SetContent(content)

	w.ShowAndRun()

}
