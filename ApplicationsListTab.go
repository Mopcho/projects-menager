package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func applicationsListTab() *fyne.Container {
	applications, err := GetApplications()

	if err != nil {
		log.Fatal(err)
	}

	container := container.NewVBox()

	for _, application := range applications {
		container.Add(signeleAppCard(application))
	}
	
	return container
}

func signeleAppCard(application Application) *widget.Card {
	playBtn := widget.NewButton("Execute Start Command", func() {})
	card := widget.NewCard(application.Name, "", playBtn)
	return card
}