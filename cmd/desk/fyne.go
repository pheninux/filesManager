package main

import (
	"fileManager2/pkg/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
)

type FyneGui struct {
	in        *widget.Entry
	out       *widget.Entry
	action    *widget.Select
	exts      []*widget.Check
	otherExts *widget.Entry
	log       *widget.Entry
	oif       *widget.Button
	oof       *widget.Button
	btnGo     *widget.Button
	btnCnl    *widget.Button
}

var exts = []string{"pdf", "png", "txt", "jpeg", "csv", "doc", "docs"}
var selectedExts []string
var cbs []*widget.Check

func startGuiFyne(da DeskApplication) {
	//func main() {

	app := app.New()
	w := app.NewWindow("Files Manager")
	//w.SetFixedSize(true)
	w.CenterOnScreen()

	/*********          ********************/
	f := FyneGui{}

	/************             *****************/
	inDir := widget.NewEntry()
	inDir.SetPlaceHolder("select the input folder")
	inDir.Resize(fyne.NewSize(302, 40))
	inDir.Move(fyne.NewPos(0, 2))

	/*************        ****************/
	oif := widget.NewButton("...", func() {
		d := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			inDir.SetText(uri.Path())
		}, w)
		d.Show()
	}) // button ... select input folder
	oif.Resize(fyne.NewSize(90, 40))
	oif.Move(fyne.NewPos(310, 3))

	/****************        ******************/
	outDir := widget.NewEntry()
	outDir.SetPlaceHolder("select the output folder")
	outDir.Resize(fyne.NewSize(302, 40))
	outDir.Move(fyne.NewPos(0, 2))

	/******************       **********************/
	oof := widget.NewButton("...", func() {
		d := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			outDir.SetText(uri.Path())
		}, w)
		d.Show()
	}) // button ... select out folder
	oof.Resize(fyne.NewSize(90, 40))
	oof.Move(fyne.NewPos(310, 2))

	/***********             ********************/

	action := widget.NewSelect([]string{"Move", "Copy"}, func(s string) {})
	action.PlaceHolder = "Select Action"

	/*************          *********************/
	extContent := container.NewGridWrap(fyne.NewSize(60, 40))

	/***************        ********************/
	for _, e := range exts {
		cb := widget.NewCheck(e, func(b bool) {})
		cbs = append(cbs, cb)
		extContent.Add(cb)
	}

	/**************           ***************/
	otherExt := widget.NewEntry()
	otherExt.SetPlaceHolder("Other extentions for exmeples : msi;rar...")
	otherExt.Resize(fyne.NewSize(396, 40))
	otherExt.Move(fyne.NewPos(2, 2))

	logs := widget.NewMultiLineEntry()

	/******* ******************/
	btnCancel := widget.NewButton("Quit", func() {})
	btnGo := widget.NewButton("Go", func() {

		selectedExts = []string{}
		for _, cb := range cbs {
			if cb.Checked {
				selectedExts = append(selectedExts, cb.Text)
			}
		}
		da.fileManager.StartProcessing(wrapFyneFormEntry(f))
	})

	f = FyneGui{
		in:        inDir,
		out:       outDir,
		action:    action,
		exts:      cbs,
		otherExts: otherExt,
		log:       logs,
		oif:       oif,
		oof:       oof,
		btnGo:     btnGo,
		btnCnl:    btnCancel,
	}

	/*************** *************************/
	btnContainer := container.NewHBox(layout.NewSpacer(), btnCancel, btnGo)

	/***************            **********************/
	w.SetContent(container.NewVBox(
		container.NewGridWrap(
			fyne.NewSize(400, 40), container.NewWithoutLayout(inDir, oif)),
		container.NewGridWrap(
			fyne.NewSize(400, 40), container.NewWithoutLayout(outDir, oof)),
		container.NewGridWrap(
			fyne.NewSize(400, 40), action),
		container.NewGridWrap(
			fyne.NewSize(400, 80), extContent),
		container.NewGridWrap(
			fyne.NewSize(400, 40), otherExt),
		container.NewGridWrap(
			fyne.NewSize(400, 200), logs), btnContainer))
	w.ShowAndRun()

}

func removeFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func wrapFyneFormEntry(frm FyneGui) (dt *models.DataTemplate) {

	dt = &models.DataTemplate{
		DirIn:  frm.in.Text,
		DirOut: frm.out.Text,
		Action: parseFyneSelectedCombo(frm.action),
		Exts:   manageExts(frm),
	}
	fmt.Println(dt)
	return dt
}

func parseFyneSelectedCombo(cd *widget.Select) string {
	return cd.Selected
}

func manageExts(frm FyneGui) []string {
	if frm.otherExts.Text == "" {
		return selectedExts
	}
	return deleteEmptySliceValue(append(strings.Split(frm.otherExts.Text, ";"), selectedExts...))
}

func deleteEmptySliceValue(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
