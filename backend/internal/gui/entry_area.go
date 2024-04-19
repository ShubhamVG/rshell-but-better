package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// func getEntryArea(winPtr *fyne.Window) fyne.CanvasObject {
// 	window := *winPtr

// 	entry := widget.NewMultiLineEntry()
// 	button := widget.NewButton("Send", func() {})

// 	buttonSize := fyne.NewSize(500, 500)
// 	button.Resize(buttonSize)

// 	entrySize := fyne.NewSize(
// 		window.Canvas().Size().Width-buttonSize.Width,
// 		entry.Size().Height,
// 	)
// 	entry.Resize(entrySize)

// 	hboxContainer := container.NewHBox(entry, button)
// 	hboxContainer.Resize(window.Canvas().Size())

// 	return hboxContainer
// }

func getEntryArea() fyne.CanvasObject {
	entry := widget.NewMultiLineEntry()
	button := widget.NewButton("Send", func() {})
	cri := container.NewVSplit(button, entry)

	return container.NewBorder(nil, nil, nil, button, cri)
}
