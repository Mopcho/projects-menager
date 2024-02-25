package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (appData *AppData) applicationsListTab() *fyne.Container {
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
	playBtn := widget.NewButton("Execute Start Command", func() {
		args := strings.Split(application.StartCommand, " ") 
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = application.Location
		out, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(out))
	})
	card := widget.NewCard(application.Name, "", playBtn)
	return card
}