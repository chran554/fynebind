package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	app2 "fynebind/internal/app"
)

func main() {
	application := app.New()
	application.SetIcon(theme.ComputerIcon())

	window := application.NewWindow("Untyped binding test")

	mainLayout := app2.ApplicationContent()

	window.Resize(fyne.NewSize(800, 600))
	window.SetContent(mainLayout)

	window.ShowAndRun()
}
