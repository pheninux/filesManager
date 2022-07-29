package main

//
//import (
//	"fileManager2/cmd/common"
//	"fileManager2/pkg/models"
//	"fmt"
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/app"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/dialog"
//	"fyne.io/fyne/v2/layout"
//	"fyne.io/fyne/v2/widget"
//	"os"
//	"strings"
//)
//
//type FyneGui struct {
//	in        *widget.Entry
//	out       *widget.Entry
//	action    *widget.Select
//	exts      []*widget.Check
//	otherExts *widget.Entry
//	log       *widget.TextGrid
//	oif       *widget.Button
//	oof       *widget.Button
//	btnGo     *widget.Button
//	btnCnl    *widget.Button
//	Pb        *widget.ProgressBarInfinite
//}
//
//var exts = []string{"pdf", "png", "txt", "jpeg", "csv", "doc", "docs"}
//var selectedExts []string
//var cbs []*widget.Check
//
//func startGuiFyne(da *DeskApplication, s chan *models.Stack) {
//	//func main() {
//
//	app := app.New()
//	// for setting new theme
//	//app.Settings().SetTheme(&myTheme{})
//	w := app.NewWindow("Files Manager")
//	w.SetFixedSize(true)
//	w.CenterOnScreen()
//
//	/*********          ********************/
//	f := FyneGui{}
//
//	/************             *****************/
//	inDir := widget.NewEntry()
//	inDir.SetPlaceHolder("select the input folder")
//	inDir.Resize(fyne.NewSize(270, 40))
//	inDir.Move(fyne.NewPos(0, 2))
//
//	/*************        ****************/
//	oif := widget.NewButton("...", func() {
//		d := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
//			if uri == nil {
//				return
//			}
//			if da.os == "windows" {
//				inDir.SetText(uri.Path())
//			} else {
//				inDir.SetText(uri.String())
//			}
//		}, w)
//		d.Show()
//	}) // button ... select input folder
//	oif.Resize(fyne.NewSize(60, 40))
//	oif.Move(fyne.NewPos(275, 3))
//
//	/****************        ******************/
//	outDir := widget.NewEntry()
//	outDir.SetPlaceHolder("select the output folder")
//	outDir.Resize(fyne.NewSize(270, 40))
//	outDir.Move(fyne.NewPos(0, 2))
//
//	/******************       **********************/
//	oof := widget.NewButton("...", func() {
//		d := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
//			if uri == nil {
//				return
//			}
//			if da.os == "windows" {
//				outDir.SetText(uri.Path())
//			} else {
//				outDir.SetText(uri.String())
//			}
//		}, w)
//		d.Show()
//	}) // button ... select out folder
//	oof.Resize(fyne.NewSize(60, 40))
//	oof.Move(fyne.NewPos(275, 2))
//
//	/***********             ********************/
//
//	action := widget.NewSelect([]string{"Move", "Copy"}, func(s string) {})
//	action.PlaceHolder = "Select Action"
//
//	/*************          *********************/
//	extContent := container.NewGridWrap(fyne.NewSize(60, 40))
//
//	/***************        ********************/
//	for _, e := range exts {
//		cb := widget.NewCheck(e, func(b bool) {})
//		cbs = append(cbs, cb)
//		extContent.Add(cb)
//	}
//
//	/**************           ***************/
//	otherExt := widget.NewEntry()
//	otherExt.SetPlaceHolder("Other extentions for exmeples : msi;rar...")
//	otherExt.Resize(fyne.NewSize(350, 40))
//	otherExt.Move(fyne.NewPos(2, 2))
//
//	//************* logs widget **********************/
//	logs := widget.NewTextGrid()
//
//	/**************  progress bar*****************/
//
//	pb := widget.NewProgressBarInfinite()
//	pb.Hide()
//	/******* ******************/
//	btnCancel := widget.NewButton("Quit", func() {
//		if da.os == "windows" {
//			app.Quit()
//		} else {
//			os.Exit(1)
//		}
//
//	})
//	btnGo := widget.NewButton("Go", func() {
//
//		// when press de button "go" first we collect de selected exts
//		selectedExts = []string{}
//		for _, cb := range cbs {
//			if cb.Checked {
//				selectedExts = append(selectedExts, cb.Text)
//			}
//		}
//
//		// write logs
//		go func(da *DeskApplication) {
//			for {
//				select {
//				case r := <-s:
//					fmt.Println(r)
//					if r.Err != "" {
//						d := dialog.NewInformation("Error", fmt.Sprintf("%s", r.Err), w)
//						d.Show()
//					} else if r.Ffound != nil {
//						sb := strings.Builder{}
//						sb.WriteString("Files found \n\n")
//						for k, v := range r.Ffound {
//							sb.WriteString(fmt.Sprintf("[%s] : %v file(s) \n", k, v))
//						}
//						sb.WriteString("\nwant to confirm ?")
//						dc := dialog.NewConfirm("INFO", sb.String(), func(b bool) {
//							if b {
//								// start processing
//								///***** make progress bar visible ***************/
//								f.Pb.Show()
//								f.btnGo.Disable()
//								go da.fileManager.StartProcessing(wrapFyneFormEntry(f), s)
//
//							}
//						}, w)
//						dc.Show()
//
//					} else if r.Done {
//
//						d := dialog.NewInformation("INFO","DONE", w)
//						d.Show()
//						f.btnGo.Enable()
//						f.Pb.Hide()
//					}
//				}
//			}
//		}(da)
//
//		// check validate form entry
//		if !validateFyneEntry(f, w) {
//			return
//		}
//
//		if err := da.utils.ValidateArgs(wrapFyneFormEntry(f), s); err != nil {
//			return
//		}
//
//		common.CountFileWithExtention(wrapFyneFormEntry(f), s)
//
//	})
//	btnGo.Importance = widget.HighImportance
//
//	f = FyneGui{
//		in:        inDir,
//		out:       outDir,
//		action:    action,
//		exts:      cbs,
//		otherExts: otherExt,
//		log:       logs,
//		oif:       oif,
//		oof:       oof,
//		btnGo:     btnGo,
//		btnCnl:    btnCancel,
//		Pb:        pb,
//	}
//
//	/*************** *************************/
//	btnContainer := container.NewHBox(layout.NewSpacer(), btnCancel, btnGo)
//
//	/***************            **********************/
//	w.SetContent(container.NewVBox(
//		container.NewGridWrap(
//			fyne.NewSize(335, 40), container.NewWithoutLayout(inDir, oif)),
//		container.NewGridWrap(
//			fyne.NewSize(335, 40), container.NewWithoutLayout(outDir, oof)),
//		container.NewGridWrap(
//			fyne.NewSize(335, 40), action),
//		container.NewGridWrap(
//			fyne.NewSize(335, 80), extContent),
//		container.NewGridWrap(
//			fyne.NewSize(335, 40), otherExt),
//		container.NewGridWrap(
//			fyne.NewSize(335, 200), logs), btnContainer,
//		container.NewGridWrap(fyne.NewSize(335, 10), pb)))
//	w.ShowAndRun()
//
//}
//
//func removeFromSlice(s []string, r string) []string {
//	for i, v := range s {
//		if v == r {
//			return append(s[:i], s[i+1:]...)
//		}
//	}
//	return s
//}
//
//// wrap data from from  fyne entry to the templateData model
//func wrapFyneFormEntry(frm FyneGui) *models.DataTemplate {
//	dt := &models.DataTemplate{
//		DirIn:  frm.in.Text,
//		DirOut: frm.out.Text,
//		Action: parseFyneSelectedCombo(frm.action),
//		Exts:   manageExts(frm),
//		Stack:  new(models.Stack),
//	}
//
//	//fmt.Println(da.dataTemplate)
//	return dt
//}
//
//// parse the selected option in combo widget
//func parseFyneSelectedCombo(cd *widget.Select) string {
//	return cd.Selected
//}
//
//// manage and clean the slice of exts
//func manageExts(frm FyneGui) []string {
//	if frm.otherExts.Text == "" {
//		return selectedExts
//	}
//	// else i return the sum of 2 slice ext and delete empty value
//	return deleteEmptySliceValue(append(strings.Split(frm.otherExts.Text, ";"), selectedExts...))
//}
//
//// delete de empty value from slice
//func deleteEmptySliceValue(s []string) []string {
//	var r []string
//	for _, str := range s {
//		if str != "" {
//			r = append(r, str)
//		}
//	}
//	return r
//}
//
//// function to validate the fyne form entry
//func validateFyneEntry(frm FyneGui, w fyne.Window) bool {
//	if frm.in.Text == "" {
//		d := dialog.NewInformation("Information", "Input directory is empty !", w)
//		d.Show()
//		w.Canvas().Focus(frm.in)
//		return false
//	} else if frm.out.Text == "" {
//		d := dialog.NewInformation("Information", "Output directory is empty !", w)
//		d.Show()
//		w.Canvas().Focus(frm.out)
//		return false
//	} else if frm.action.Selected == "" {
//		d := dialog.NewInformation("Information", "Please select the action to do", w)
//		d.Show()
//		w.Canvas().Focus(frm.action)
//		return false
//	} else if manageExts(frm) == nil || len(manageExts(frm)) == 0 {
//		d := dialog.NewInformation("Information", "Please select at less one extension", w)
//		d.Show()
//		w.Canvas().Focus(frm.otherExts)
//		return false
//	}
//
//	return true
//}
