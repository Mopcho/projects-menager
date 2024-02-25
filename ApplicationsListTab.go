package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
)

func (appData *AppData) applicationsListTab() *container.Scroll {
	applications, err := GetApplications()

	if err != nil {
		log.Fatal(err)
	}

	vbox := container.NewVBox()

	for _, application := range applications {
		vbox.Add(signeleAppCard(application))
		separator := NewCustomSeparator(colornames.Aliceblue, 1)
		vbox.Add(separator)
	}

	content := container.NewVScroll(vbox)
	
	return content
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
	card := widget.NewCard(application.Name, application.Location, playBtn)
	return card
}