package main

import (
	"github.com/andlabs/ui"
	"os"
)

var extentions []string = []string{"pdf", "zip", "xls", "doc", "png", "jpeg", "exe"}

func startGui(da DeskApplication) {

	//run gui
	ui.Main(func() {
		win = ui.NewWindow("File manager", 400, 300, true)
		win.SetMargined(true)
		win.OnClosing(func(w *ui.Window) bool {
			ui.Quit()
			os.Exit(1)
			return true
		})

		ui.OnShouldQuit(func() bool {
			win.Destroy()
			os.Exit(1)
			return true
		})
		win.SetChild(makeUiForm(da))
		win.Show()

	})
}

// code gui

func makeUiForm(da DeskApplication) ui.Control {

	btn := ui.NewButton("Go")
	btn.OnClicked(func(button *ui.Button) {
		da.fileManager.StartProcessing(&da.file)
	})

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	grp := ui.NewGroup("")
	grp.SetMargined(true)
	uiForm := ui.NewForm()
	uiForm.SetPadded(true)
	uiForm.Append("Répertoire d'entrée", ui.NewEntry(), false)
	uiForm.Append("Répertoire de sortie", ui.NewEntry(), false)
	cmbAction := ui.NewCombobox()
	cmbAction.Append("copy")
	cmbAction.Append("move")
	uiForm.Append("Action", cmbAction, false)
	uiForm.Append("", makeExtentionsUi(), false)
	uiForm.Append("Autres extentions", ui.NewEntry(), false)
	uiForm.Append("Logs", ui.NewNonWrappingMultilineEntry(), false)
	uiForm.Append("", btn, false)
	grp.SetChild(uiForm)
	vbox.Append(grp, true)
	return vbox
}

func makeExtentionsUi() ui.Control {
	vbox := ui.NewVerticalBox()
	//l := ui.NewLabel("Extentions :")
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	for _, v := range extentions {
		checkB := ui.NewCheckbox(v)
		hbox.Append(checkB, true)
	}
	//vbox.Append(l,true)
	vbox.Append(hbox, true)
	return vbox
}

func hLabelEntry(e *ui.Entry, l string) (*ui.Entry, ui.Control) {
	hbox := ui.NewHorizontalBox()
	lab := ui.NewLabel(l)
	hbox.Append(lab, true)
	hbox.Append(e, true)

	return e, hbox
}

func vLabelEntry(e *ui.Entry, l string) (*ui.Entry, ui.Control) {
	vbox := ui.NewVerticalBox()
	lab := ui.NewLabel(l)
	vbox.Append(lab, true)
	vbox.Append(e, true)

	return e, vbox
}

func hLabelCombo(e *ui.Combobox, l string) (*ui.Combobox, ui.Control) {
	hbox := ui.NewHorizontalBox()
	lab := ui.NewLabel(l)
	hbox.Append(lab, true)
	hbox.Append(e, true)

	return e, hbox
}

func vLabelCombo(e *ui.Combobox, l string) (*ui.Combobox, ui.Control) {
	vbox := ui.NewVerticalBox()
	lab := ui.NewLabel(l)
	vbox.Append(lab, true)
	vbox.Append(e, true)

	return e, vbox
}

func vLabelMultiEntry(e *ui.MultilineEntry, l string) (*ui.MultilineEntry, ui.Control) {
	vbox := ui.NewVerticalBox()
	lab := ui.NewLabel(l)
	vbox.Append(lab, true)
	vbox.Append(e, true)

	return e, vbox
}
