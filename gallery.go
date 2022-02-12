package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Gallery(lightBackground *canvas.Image, darkBackground *canvas.Image, darkTheme bool) {

	w := a.NewWindow("GF Gallery")
	w.Resize(fyne.NewSize(1280, 720))

	slice := make([]string, 0)

	source := "image Path Here updated"

	imgContainer := container.NewGridWithColumns(1)

	var selectImage string
	var b int = -1
	checkImageSelect := 0
	var img *canvas.Image
	//icon for set as background inside the
	galSetBackground, _ := fyne.LoadResourceFromPath("./images/gallerySetBackgroundIcon.jpg")

	// Create Menu Items with Open Image Functionality
	galleryMenuItem1 := fyne.NewMenuItem("Open", func() {

		openImageDialog := dialog.NewFileOpen(func(r fyne.URIReadCloser, _ error) {

			selectImage = r.URI().Name()
			// fmt.Printf("value  = %v", r.URI().Path())
			source = r.URI().Path()
			source = source[:len(source)-len(selectImage)]

			files, _ := ioutil.ReadDir(source)
			for _, file := range files {
				// fmt.Println(file.Name(), file.IsDir(), file)
				if file.IsDir() == false {
					extension := strings.Split(file.Name(), ".")[1]
					if extension == "png" || extension == "jpeg" || extension == "jpg" {
						slice = append(slice, file.Name())
					}
				}
			}
			b = imgFinder(slice, selectImage)
			img = canvas.NewImageFromFile(source + slice[b])

			//Checking Image container
			if checkImageSelect == 0 {
				checkImageSelect++
			} else {
				imgContainer.Remove(img)
			}

			imgContainer.Add(img)

		}, w)
		openImageDialog.SetFilter(storage.NewExtensionFileFilter([]string{".jpg", ".png", ".jpeg", ".PNG"}))
		openImageDialog.Show()
	})
	galleryMenuItem2 := fyne.NewMenuItem("Set as Background", func() {
		if b >= 0 {
			imagePathVal := source + slice[b]
			setAsBackground(lightBackground, darkBackground, darkTheme, imagePathVal)
		}
	})

	newMenu := fyne.NewMenu("File", galleryMenuItem1, galleryMenuItem2)
	mainMenu := fyne.NewMainMenu(newMenu)

	// fmt.Printf(" %v", b)

	//Tool bar for navigate
	imageToolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		// widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {
			if b > 0 {
				b--
				img.File = fmt.Sprint(source + slice[b])
				img.Refresh()

			}
		}),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() {
			if b < len(slice)-1 {
				b++
				img.File = fmt.Sprint(source + slice[b])
				img.Refresh()
			}
		}),
		widget.NewToolbarSpacer(),
		// Change BackGround Image
		widget.NewToolbarAction(galSetBackground, func() {
			if b >= 0 {
				imagePathVal := source + slice[b]
				setAsBackground(lightBackground, darkBackground, darkTheme, imagePathVal)
			}

		}),
	)

	// sliderImage := container.New(layout.NewBorderLayout(nil, nil, btnLeft, btnRight),imgContainer)
	sliderImage := container.NewBorder(imageToolbar, nil, nil, nil, imgContainer)

	r, _ := fyne.LoadResourceFromPath("./images/galleryLogo.png")
	w.SetIcon(r)

	w.SetMainMenu(mainMenu)

	w.SetPadded(false)
	w.CenterOnScreen()
	fullScreen := false
	// KeyBoard Key Pressing
	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == fyne.KeyF11 {
			if fullScreen {
				w.SetFullScreen(false)
				fullScreen = !fullScreen
			} else {
				w.SetFullScreen(true)
				fullScreen = !fullScreen
			}
		} else if k.Name == fyne.KeyLeft {
			fmt.Println("Gallery left")
			if b > 0 {
				b--
				img.File = fmt.Sprint(source + slice[b])
				img.Refresh()
			}

		} else if b < len(slice)-1 {
			fmt.Println("Gallery Right")
			b++
			img.File = fmt.Sprint(source + slice[b])
			img.Refresh()
		}
	})
	// w.SetFullScreen(true)
	w.SetContent(sliderImage)
	w.Show()
}

//String Finder in slice or array
func imgFinder(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a) - 1
}

//Changing Background Image functionality
func setAsBackground(lightBackground *canvas.Image, darkBackground *canvas.Image, darkTheme bool, imagePathVal string) {
	if darkTheme {
		darkBackground.File = fmt.Sprintf(imagePathVal)
		darkBackground.Refresh()
	} else {
		lightBackground.File = fmt.Sprintf(imagePathVal)
		lightBackground.Refresh()
	}
}
