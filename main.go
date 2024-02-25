package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	err := SetupStorage(true)

	if err != nil {
		log.Fatal(err)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	myWindow.Resize(fyne.Size{Width: 1080, Height: 980})

	myWindow.SetContent(App())
	myWindow.ShowAndRun()
}