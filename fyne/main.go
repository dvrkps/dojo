package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("fyne")

	w.Resize(fyne.NewSize(300, 300))

	hello := widget.NewLabel("Hello Fyne!")

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
			w.Hide()
			time.Sleep(5 * time.Second)
			w.Show()
		}),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	w.ShowAndRun()
}
