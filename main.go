package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AppPage() (*widget.Label, *widget.Entry, *widget.Button) {
	label := widget.NewLabel("Username:")
	input := widget.NewEntry()
	submitButton := widget.NewButton("Login", func() {
		fmt.Printf("%v", input.Text)
	})

	return label, input, submitButton
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	w.Resize(fyne.NewSize(100, 100))

	w.SetContent(container.NewVBox(AppPage()))

	// clock := widget.NewLabel("")
	// updateTime(clock)

	// w.SetContent(clock)
	// go func() {
	// 	for range time.Tick(time.Second) {
	// 		updateTime(clock)
	// 	}
	// }()
	w.ShowAndRun()
}