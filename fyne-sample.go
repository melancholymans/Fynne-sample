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
			widget.NewTabContainer(
				widget.NewTabItem("First",
					widget.NewLabel("This is First tab item"),
				),
				widget.NewTabItem("Second",
					widget.NewLabel("This is Second tab item"),
				),
			),
		),
	)
	w.ShowAndRun()
}
