package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)	

type AppData struct {
	eventsChannel chan string
}

func (appData *AppData) App() *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Add a new App", theme.ContentAddIcon() , appData.addAppComponent()),
		container.NewTabItemWithIcon("Apps", theme.ListIcon(), appData.applicationsListTab()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	go func() {
		for range appData.eventsChannel {
			tabs.Items[1].Content = appData.applicationsListTab()
		}
	}()

	return tabs
}

