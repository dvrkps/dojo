package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("fyne")

	w.Resize(fyne.NewSize(300, 300))
	w.SetFixedSize(true)

	hello := widget.NewLabel("Hello Fyne!")

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
			w.CenterOnScreen()
		}),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	// w.RequestFocus() // TODO(dvrkps): panic on macos.

	w.ShowAndRun()
}
