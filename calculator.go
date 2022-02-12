package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func Calculator() {

	// a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(320, 380))

	output := ""
	input := widget.NewLabel("Calculator")
	input.Alignment = fyne.TextAlignTrailing
	input.TextStyle = fyne.TextStyle{Bold: true}
	historySlice := make([]string, 0) //Dynamic array or slice
	history := ""
	historyInput := widget.NewLabel("")

	flag := true

	historyBtn := widget.NewButton("History", func() {
		if flag {
			for _, val := range historySlice {
				history += val + "\n"
			}
		} else {
			history = ""
		}
		flag = !flag
		historyInput.SetText(history)
	})
	backBtn := widget.NewButton("Back", func() {
		if output != "" {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})
	clearBtn := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
	})
	openBtn := widget.NewButton("(", func() {
		output += "("
		input.SetText(output)
	})

	closeBtn := widget.NewButton(")", func() {
		output += ")"
		input.SetText(output)
	})
	divideBtn := widget.NewButton("/", func() {
		output += "/"
		input.SetText(output)
	})
	sevenBtn := widget.NewButton("7", func() {
		output += "7"
		input.SetText(output)
	})
	eightBtn := widget.NewButton("8", func() {
		output += "8"
		input.SetText(output)
	})
	nineBtn := widget.NewButton("9", func() {
		output += "9"
		input.SetText(output)
	})
	multiplicationBtn := widget.NewButton("*", func() {
		output += "*"
		input.SetText(output)
	})
	fourBtn := widget.NewButton("4", func() {
		output += "4"
		input.SetText(output)
	})
	fiveBtn := widget.NewButton("5", func() {
		output += "5"
		input.SetText(output)
	})
	sixBtn := widget.NewButton("6", func() {
		output += "6"
		input.SetText(output)
	})
	minusBtn := widget.NewButton("-", func() {
		output += "-"
		input.SetText(output)
	})
	oneBtn := widget.NewButton("1", func() {
		output += "1"
		input.SetText(output)
	})
	twoBtn := widget.NewButton("2", func() {
		output += "2"
		input.SetText(output)
	})
	threeBtn := widget.NewButton("3", func() {
		output += "3"
		input.SetText(output)
	})
	plusBtn := widget.NewButton("+", func() {
		output += "+"
		input.SetText(output)
	})
	zeroBtn := widget.NewButton("0", func() {
		output += "0"
		input.SetText(output)
	})
	dotBtn := widget.NewButton(".", func() {
		output += "."
		input.SetText(output)
	})
	equalBtn := widget.NewButton("=", func() {

		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, _ := expression.Evaluate(nil)
			fmt.Printf("value = %v  Type = %T", result, result)
			if err == nil {
				//Output added to the history variable
				history = output

				//Float value to string (result is interface type float)
				output = strconv.FormatFloat(result.(float64), 'f', -1, 64) //Type assertion of result

				//Added to the History
				historySlice = append(historySlice, (history + " = " + output))
				history = ""
			} else {
				output = "Error"
			}
		}
		input.SetText(output)
		output = ""
	})

	r, _ := fyne.LoadResourceFromPath("./images/calcLogo.png")
	w.SetIcon(r)

	w.CenterOnScreen()
	w.SetContent(container.NewVBox(
		input,
		historyInput,
		container.NewGridWithColumns(1,
			// First Row
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),

			// Second Row
			container.NewGridWithColumns(4,
				clearBtn,
				openBtn,
				closeBtn,
				divideBtn,
			),

			//Third Row
			container.NewGridWithColumns(4,
				sevenBtn,
				eightBtn,
				nineBtn,
				multiplicationBtn,
			),

			//Fourth Row
			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				minusBtn,
			),

			//Fifth Row
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				plusBtn,
			),

			//Sixth Row
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					dotBtn,
				),
				equalBtn,
			),
		),
	))

	w.Show()
}
