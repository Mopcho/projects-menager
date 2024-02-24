package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)	

func App() *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Add a new App", theme.ContentAddIcon() , addAppComponent()),
		container.NewTabItemWithIcon("Apps", theme.ListIcon(), widget.NewLabel("Not implemented...")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	return tabs
}

