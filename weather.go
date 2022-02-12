package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func WeatherApp() {
	// a := app.New()
	w := a.NewWindow("GF Weather app")

	w.Resize(fyne.NewSize(600, 400))

	byteContent := getApiValue("kolkata") //By  Default Kolkata

	// fmt.Println(weather.Weather[0].Main)
	label := widget.NewLabel("Weather Details")
	label1 := widget.NewLabel("")
	label2 := widget.NewLabel("")
	label3 := widget.NewLabel("")
	label4 := widget.NewLabel("")
	label5 := widget.NewLabel("")

	label.Alignment = fyne.TextAlignCenter
	label1.Alignment = fyne.TextAlignCenter
	label2.Alignment = fyne.TextAlignCenter
	label3.Alignment = fyne.TextAlignCenter
	label4.Alignment = fyne.TextAlignCenter
	label5.Alignment = fyne.TextAlignCenter

	image := labelTextChange(label1, label2, label3, label4, label5, byteContent)
	imageContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(350, 250)), image)
	image.FillMode = canvas.ImageFillOriginal

	boxLabels := container.NewGridWithColumns(1, label1,
		label2,
		label3,
		label4,
		label5)

	boxContainer := container.NewVBox(
		label,
	)
	boxContainer.Add(imageContainer)

	//ComboBox setup
	combo := widget.NewSelect([]string{"kolkata", "mumbai", "delhi", "noida", "Miami"}, func(value string) {
		fmt.Println("Select set to", value)
		byteContent = getApiValue(value)

		imageContainer.Remove(image)
		image = labelTextChange(label1, label2, label3, label4, label5, byteContent)
		image.FillMode = canvas.ImageFillOriginal
		imageContainer.Add(image)

	})
	combo.Selected = "kolkata"
	boxContainer.Add(combo)
	boxContainer.Add(boxLabels)

	r, _ := fyne.LoadResourceFromPath("./images/weatherLogo.png")
	w.SetIcon(r)

	w.CenterOnScreen()
	w.SetContent(boxContainer)
	w.Show()
}

//Api
func getApiValue(variablePlace string) []uint8 {
	// Api functon
	linkSource := "https://api.openweathermap.org/data/2.5/weather?q=" + variablePlace + "&appid=6e6d83b7865d0917805fcf3caebf827c"

	res, err := http.Get(linkSource)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	byteContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return byteContent
}

//Canvas Text Change
func labelTextChange(label1 *widget.Label, label2 *widget.Label, label3 *widget.Label, label4 *widget.Label, label5 *widget.Label, byteContent []uint8) *canvas.Image {

	weather, _ := UnmarshalWeather(byteContent)

	imageUrl := "https://openweathermap.org/img/wn/" + weather.Weather[0].Icon + "@4x.png"
	r, _ := fyne.LoadResourceFromURLString(imageUrl)

	image2 := canvas.NewImageFromResource(r)

	temperature := (weather.Main.Temp - 273.15)
	// fmt.Println(weather.Weather[0].Main)
	label1.SetText(fmt.Sprintf("Country : %v", weather.Sys.Country))
	label2.SetText(fmt.Sprintf("Weather : %v", weather.Weather[0].Main))
	label3.SetText(fmt.Sprintf("Temperature : %.2f °C", temperature))
	label4.SetText(fmt.Sprintf("Wind Speed : %.2f m/s", weather.Wind.Speed))
	label5.SetText(fmt.Sprintf("Humidity : %v %%", weather.Main.Humidity))

	return image2

}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
