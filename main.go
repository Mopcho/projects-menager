package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	err := SetupStorage(false)

	if err != nil {
		log.Fatal(err)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	myWindow.Resize(fyne.Size{Width: 1080, Height: 980})

	eventsChannel := make(chan string)

	appData := AppData{
		eventsChannel: eventsChannel,
	} 

	myWindow.SetContent(appData.App())
	myWindow.ShowAndRun()
}