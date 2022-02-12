package main

// import fyne
import (
	"fmt"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func TextEditor() {

	// a := app.New()
	w := a.NewWindow("Text Editor")

	w.Resize(fyne.NewSize(500, 500))

	multilineInput := widget.NewMultiLineEntry()
	multilineInput.SetPlaceHolder("Enter Your content Here......")

	//Menu
	// Menu Items
	menuItem1 := fyne.NewMenuItem("Open", func() { openFile(w) })
	menuItem2 := fyne.NewMenuItem("Save", func() { fileSave(multilineInput.Text, w) })

	// New Menu
	newMenu := fyne.NewMenu("File", menuItem1, menuItem2)
	// creating new main menu
	menu := fyne.NewMainMenu(newMenu)

	btnSave := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		fileSave(multilineInput.Text, w)

	})
	btnOpen := widget.NewButtonWithIcon("Open", theme.FolderOpenIcon(), func() {
		openFile(w)
	})

	textEditorBtn := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btnSave, btnOpen, layout.NewSpacer())

	r, _ := fyne.LoadResourceFromPath("./images/textEditorLogo.png")
	w.SetIcon(r)

	w.SetMainMenu(menu)
	w.SetPadded(false)
	w.CenterOnScreen()
	w.SetContent(
		container.NewBorder(nil, textEditorBtn, nil, nil, multilineInput),
	)
	w.Show()
}

//Save File Function
func fileSave(input string, w fyne.Window) {
	newTextFileSave := dialog.NewFileSave(func(uc fyne.URIWriteCloser, _ error) {

		textData := []byte(input)
		uc.Write(textData)

		fmt.Println("File Not selected")

	}, w)
	newTextFileSave.SetFileName("New file.txt")
	newTextFileSave.Show()
}

// Open File function
func openFile(w fyne.Window) {
	newTextFileOpen := dialog.NewFileOpen(func(r fyne.URIReadCloser, _ error) {
		// if e != nil {
		readData, _ := ioutil.ReadAll(r)

		data := fyne.NewStaticResource("New File", readData)

		viweData := widget.NewMultiLineEntry()
		viweData.SetText(string(data.StaticContent))

		// fmt.Println(r.URI().Name())
		newWindow := fyne.CurrentApp().NewWindow(r.URI().Name())

		// ***** Menu Create  ******
		newWindowMenuItem := fyne.NewMenuItem("Save", func() {
			fmt.Println("Save pressed")
			fileSave(viweData.Text, newWindow)
		})
		newWindowMenu := fyne.NewMenu("File", newWindowMenuItem)
		menu := fyne.NewMainMenu(newWindowMenu)
		newWindow.SetMainMenu(menu)

		newWindow.SetPadded(false)
		newWindow.CenterOnScreen()
		newWindow.Resize(fyne.NewSize(500, 500))
		newWindow.SetContent(container.NewHScroll(viweData))

		newWindow.Show()
		// } else {
		// 	fmt.Println(e)
		// }
	}, w)
	newTextFileOpen.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
	newTextFileOpen.Show()
}
