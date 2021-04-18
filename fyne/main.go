package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("fyne")

	// w.Resize(fyne.NewSize(300, 300))
	// w.SetFixedSize(true)
	w.SetFullScreen(true)

	hello := widget.NewLabel("Hello Fyne!")

	img := canvas.NewImageFromFile("logo.png")
	img.FillMode = canvas.ImageFillOriginal

	w.SetContent(container.NewVBox(
		img,
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
