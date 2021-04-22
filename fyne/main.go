package main

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("fyne")

	w.Resize(fyne.NewSize(300, 300))
	w.SetFixedSize(true)
	//w.SetFullScreen(true)

	counter := 0

	hello := widget.NewLabel("0")

	incCounter := func(c *int) {
		*c++
		hello.SetText(strconv.Itoa(*c))
	}

	img := canvas.NewImageFromFile("logo.png")
	img.FillMode = canvas.ImageFillOriginal

	w.SetContent(container.NewVBox(
		img,
		hello,
		widget.NewButton("inc", func() {
			incCounter(&counter)
		}),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))

	altF4 := desktop.CustomShortcut{KeyName: fyne.KeyF4, Modifier: desktop.AltModifier}
	w.Canvas().AddShortcut(&altF4, func(s fyne.Shortcut) {
		incCounter(&counter)
		w.Hide()
		time.Sleep(3e9)
		w.Show()
	})

	w.CenterOnScreen()

	// w.RequestFocus() // TODO(dvrkps): panic on macos.

	w.ShowAndRun()
}
