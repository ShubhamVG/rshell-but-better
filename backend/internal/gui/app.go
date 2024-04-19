package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func App() fyne.Window {
	a := app.New()
	w := a.NewWindow("RShell")
	w.Resize(fyne.NewSize(1000, 1000))

	// topBar := getTopBarArea()
	// middleArea := getMiddleArea()
	entryArea := getEntryArea()

	screenContainer := container.NewBorder(nil, entryArea, nil, nil)
	// screenContainer := container.NewBorder(topBar, entryArea, nil, nil, middleArea)
	w.SetContent(screenContainer)

	return w
}
