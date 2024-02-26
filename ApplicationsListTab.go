package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
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
		vbox.Add(appData.signeleAppCard(application))
		separator := NewCustomSeparator(colornames.Aliceblue, 1)
		vbox.Add(separator)
	}

	content := container.NewVScroll(vbox)
	
	return content
}

func (appData *AppData) signeleAppCard(application Application) *fyne.Container {
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

	deletBtn := widget.NewButton("Delete This App", func() {
		err := DeleteApplication(application.ID)

		if err != nil {
			log.Fatal(err)
		}

		appData.eventsChannel <- "Refresh"
	})

	hBox := container.NewHBox(playBtn, deletBtn)

	card := widget.NewCard(application.Name, fmt.Sprintf("Location: %v", application.Location), widget.NewLabel("ID: " + application.ID))

	vBox := container.NewVBox(card, hBox)

	return vBox
}