package main

import (
	"fileManager2/pkg/models"
	"fmt"
	"github.com/andlabs/ui"
	"os"
	"strings"
)

var labelcount *ui.Label

var extentions []string = []string{"pdf", "zip", "xls", "doc", "png", "jpeg", "exe"}
var selectedExt map[string]struct{} = make(map[string]struct{})
var keys []interface{}

type FormGui struct {
	inEntry       *ui.Entry
	outEntry      *ui.Entry
	otherExtEntry *ui.Entry
	log           *ui.MultilineEntry
	action        *ui.Combobox
	submit        *ui.Button
}

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

	labelcount = ui.NewLabel(fmt.Sprintf("%d", count))

	frm := FormGui{
		inEntry:       ui.NewEntry(),
		outEntry:      ui.NewEntry(),
		otherExtEntry: ui.NewEntry(),
		log:           ui.NewMultilineEntry(),
		action:        actionCombo(),
		submit:        ui.NewButton("GO"),
	}

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	grp := ui.NewGroup("")
	grp.SetMargined(true)
	uiForm := ui.NewForm()
	uiForm.SetPadded(true)
	uiForm.Append("Répertoire d'entrée", frm.inEntry, false)
	uiForm.Append("Répertoire de sortie", frm.outEntry, false)

	uiForm.Append("Action", frm.action, false)
	uiForm.Append("", makeExtentionsUi(), false)
	uiForm.Append("Autres extentions", frm.otherExtEntry, false)
	uiForm.Append("Logs", frm.log, false)
	uiForm.Append("", frm.submit, false)
	uiForm.Append("", labelcount, false)
	grp.SetChild(uiForm)
	vbox.Append(grp, true)

	frm.submit.OnClicked(func(button *ui.Button) {

		da.fileManager.StartProcessing(wrapFormEntryValue(frm))
	})

	frm.inEntry.OnChanged(func(entry *ui.Entry) {

	})
	return vbox
}

func extractMapkeys(m map[string]struct{}) []string {

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func wrapFormEntryValue(frm FormGui) (dt *models.DataTemplate) {
	dt = &models.DataTemplate{
		DirIn:  frm.inEntry.Text(),
		DirOut: frm.outEntry.Text(),
		Action: parseSelectedCombo(frm.action),
		Exts:   append(strings.Split(frm.otherExtEntry.Text(), ";"), extractMapkeys(selectedExt)...),
	}
	return dt
}

func parseSelectedCombo(cd *ui.Combobox) string {
	if cd.Selected() == 0 {
		return "Copy"
	} else {
		return "Move"
	}
}

func actionCombo() (cmb *ui.Combobox) {
	cmb = ui.NewCombobox()
	cmb.Append("Copy")
	cmb.Append("Move")
	return cmb
}

func makeExtentionsUi() ui.Control {
	vbox := ui.NewVerticalBox()
	//l := ui.NewLabel("Extentions :")
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	for _, v := range extentions {
		checkB := ui.NewCheckbox(v)
		checkB.OnToggled(func(checkbox *ui.Checkbox) {
			_, ok := selectedExt[checkbox.Text()]
			if !ok {
				selectedExt[checkbox.Text()] = struct{}{}
			} else {
				delete(selectedExt, checkbox.Text())
			}

			fmt.Println(checkbox.Text())
			fmt.Println(selectedExt)
			fmt.Println(keys)
		})
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
