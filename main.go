package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	myWindow.Resize(fyne.Size{Width: 1080, Height: 980})

	myWindow.SetContent(App())
	myWindow.ShowAndRun()
}