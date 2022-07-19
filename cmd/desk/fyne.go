package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var exts = []string{"pdf", "png", "txt", "jpeg", "csv", "doc", "docs"}

func main() {
	app := app.New()
	w := app.NewWindow("Files Manager")
	w.SetFixedSize(true)
	w.CenterOnScreen()

	inDir := widget.NewEntry()
	inDir.SetPlaceHolder("select the input folder")
	inDir.Resize(fyne.NewSize(300, 40))
	inDir.Move(fyne.NewPos(2, 2))

	outDir := widget.NewEntry()
	outDir.SetPlaceHolder("select the output folder")
	outDir.Resize(fyne.NewSize(300, 40))
	outDir.Move(fyne.NewPos(2, 2))

	otherExt := widget.NewEntry()
	otherExt.SetPlaceHolder("Other extentions for exmeples : msi;rar...")
	otherExt.Resize(fyne.NewSize(396, 40))
	otherExt.Move(fyne.NewPos(2, 2))

	logs := widget.NewMultiLineEntry()

	oif := widget.NewButton("...", func() {
		d := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			inDir.SetText(uri.Path())
		}, w)
		d.Show()
	}) //oif => open input folder
	oif.Resize(fyne.NewSize(90, 40))
	oif.Move(fyne.NewPos(306, 3))

	oof := widget.NewButton("...", func() {

		d := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			outDir.SetText(uri.Path())
		}, w)
		d.Show()
	}) //oof => open out folder
	oof.Resize(fyne.NewSize(90, 40))
	oof.Move(fyne.NewPos(306, 2))

	btnCancel := widget.NewButton("Quit", func() {})
	btnGo := widget.NewButton("Go", func() {})

	btnContainer := container.NewHBox(layout.NewSpacer(), btnCancel, btnGo)
	//extContent := container.NewGridWrap(fyne.NewSize(60, 40))
	extContent := container.NewGridWrap(fyne.NewSize(60, 40))

	for _, e := range exts {
		cb := widget.NewCheck(e, func(b bool) {})
		extContent.Add(cb)
	}

	w.SetContent(container.NewVBox(
		container.NewGridWrap(
			fyne.NewSize(400, 40), container.NewWithoutLayout(inDir, oif)),
		container.NewGridWrap(
			fyne.NewSize(400, 40), container.NewWithoutLayout(outDir, oof)),
		container.NewGridWrap(
			fyne.NewSize(400, 80), extContent),
		container.NewGridWrap(
			fyne.NewSize(400, 40), otherExt),
		container.NewGridWrap(
			fyne.NewSize(400, 200), logs), btnContainer))
	w.ShowAndRun()

}
