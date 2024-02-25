package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)	

func App() *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Add a new App", theme.ContentAddIcon() , addAppComponent()),
		container.NewTabItemWithIcon("Apps", theme.ListIcon(), applicationsListTab()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	return tabs
}

