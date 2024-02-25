package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (appData *AppData) addAppComponent() *fyne.Container {
	nameEntry := widget.NewEntry()
	nameFormItem := widget.NewFormItem("Name of project:", nameEntry)

	location := widget.NewEntry()
	locationFormItem := widget.NewFormItem("Location on disk:", location)

	commands := widget.NewEntry()
	commandsFormItem := widget.NewFormItem("Commands to execute:", commands)

	submitBtn := widget.NewButton("Create", func() {
		createData := ApplicationCreateData{
			Name: nameEntry.Text,
			Location: location.Text,
			StartCommand: commands.Text,
		}

		err := CreateApplication(createData)

		if err != nil {
			// TODO: Use logger here that saves log to a file
			return
		}

		appData.eventsChannel <- "Refresh"
	})

	form := widget.NewForm(nameFormItem, locationFormItem, commandsFormItem)

	heading := widget.NewCard("Add a new project to your list here", "", form)
	content := container.NewVBox(heading, submitBtn)

	return content
}