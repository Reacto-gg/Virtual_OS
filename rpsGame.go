package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func RPSGame() {
	// a := app.New()
	w := a.NewWindow("Rock Paper Scissor Game")
	w.Resize(fyne.NewSize(400, 620))

	choiceImage := []string{"Rock", "Paper", "Scissor"}

	rpcBackground := canvas.NewImageFromFile("./images/back.jpg")
	rpcBackground.Translucency = 0.9

	playerImg := canvas.NewImageFromFile("./images/Rock.png")
	computerImg := canvas.NewImageFromFile("./images/Paper.png")
	playerImg.FillMode = canvas.ImageFillOriginal
	computerImg.FillMode = canvas.ImageFillOriginal

	scoreLabel := widget.NewLabel("Round Result")
	scoreLabel.Alignment = fyne.TextAlignCenter
	scoreLabel.TextStyle = fyne.TextStyle{Bold: true}

	computerImg.Hide()
	//store result
	playerChoice := "Rock"
	var computerChoice string
	playerWon := 0
	computerWon := 0

	//Win Record
	playerRecord := widget.NewLabel("Player Won : 0")
	computerRecord := widget.NewLabel("Computer Won : 0")

	playerRecord.Alignment = fyne.TextAlignCenter
	computerRecord.Alignment = fyne.TextAlignCenter

	// Select from List
	playerCombo := widget.NewSelect([]string{"Rock", "Paper", "Scissor"}, func(value string) {
		playerChoice = value

		playerImg.File = fmt.Sprintf("./images/" + value + ".png")
		playerImg.Refresh()
		playerImg.Show()

	})
	playerCombo.Alignment = fyne.TextAlignCenter
	playerCombo.Selected = "Rock"

	// button
	btn1 := widget.NewButton("Play", func() {
		// UI is finished.. Now Logic part
		rand := rand.Intn(3)
		computerChoice = choiceImage[rand]

		computerImg.File = fmt.Sprintf("./images/" + choiceImage[rand] + ".png")
		computerImg.Refresh()
		computerImg.Show()
		result(playerChoice, computerChoice, scoreLabel, &playerWon, &computerWon)
		playerRecord.SetText("Player Won : " + strconv.Itoa(playerWon))
		computerRecord.SetText("Computer Won : " + strconv.Itoa(computerWon))
	})

	//Container Record
	recordContainer := container.NewGridWithColumns(1,
		container.NewGridWithColumns(2,
			playerRecord,
			computerRecord,
		),
	)

	otherContain := container.NewVBox(
		fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), playerImg, layout.NewSpacer()), // Center Align the Image
		playerCombo,
		fyne.NewContainerWithLayout(layout.NewHBoxLayout(), layout.NewSpacer(), computerImg, layout.NewSpacer()), // Center Align the Image
		scoreLabel,
		recordContainer,
		btn1,
	)
	content := fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		rpcBackground,
		otherContain,
	)

	//Set Window Icon
	r, _ := fyne.LoadResourceFromPath("./images/rpcLogo.png")
	w.SetIcon(r)

	w.CenterOnScreen()
	w.SetPadded(false)
	w.SetContent(
		// NewVBox.. More than on Widgets
		content,
	)
	// show
	w.Show()
}

func result(p string, c string, scoreLabel *widget.Label, playerWon *int, computerWon *int) {
	if p == "Rock" {
		switch c {
		case "Rock":
			scoreLabel.SetText("Draw")
		case "Paper":
			scoreLabel.SetText("You Lost")
			*computerWon++
		case "Scissor":
			scoreLabel.SetText("You Won")
			*playerWon++

		}
	} else if p == "Paper" {
		switch c {
		case "Rock":
			scoreLabel.SetText("You Won")
			*playerWon++
		case "Paper":
			scoreLabel.SetText("Draw")
		case "Scissor":
			scoreLabel.SetText("You Lost")
			*computerWon++
		}
	} else if p == "Scissor" {
		switch c {
		case "Rock":
			scoreLabel.SetText("You Lost")
			*computerWon++
		case "Paper":
			scoreLabel.SetText("You Won")
			*playerWon++
		case "Scissor":
			scoreLabel.SetText("Draw")

		}
	}
}
