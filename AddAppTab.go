package main

import "fyne.io/fyne/v2/widget"

func addAppComponent() *widget.Card {
	nameEntry := widget.NewEntry()
	nameFormItem := widget.NewFormItem("Name of project:", nameEntry)

	location := widget.NewEntry()
	locationFormItem := widget.NewFormItem("Location on disk:", location)

	form := widget.NewForm(nameFormItem, locationFormItem)

	heading := widget.NewCard("Add a new project to your list here", "", form)

	return heading
}