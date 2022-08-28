package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		widget.NewVBox(
			widget.NewLabel("Hello Fyne!"),
			widget.NewLabel("This is sampleapplication!"),
		),
	)
	w.ShowAndRun()
}
