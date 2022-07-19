package main

import (
	"fileManager2/cmd/common"
	"fileManager2/pkg/models"
	"fileManager2/pkg/models/boImpl"
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"runtime"
	"time"
)

var win *ui.Window
var count int

type DeskApplication struct {
	fileManager  common.IFileManager
	dataTemplate *models.DataTemplate
	utils        common.Utils
}

func main() {

	count = 0
	//params := models.DataTemplate{Action: "copy", DirOut: "C:\\Users\\a706836\\go\\src\\filesManager2", DirIn: "C:\\Users\\a706836\\Downloads", Exts: []string{"pdf"}}
	u := common.Utils{}
	if runtime.GOOS == "windows" {
		u.Slash = "\\"
	} else {
		u.Slash = "//"
	}

	da := DeskApplication{fileManager: &boImpl.FileModel{Utils: &u}, utils: common.Utils{}}

	//currPath, err := os.Getwd() // get current path
	//app.utils.CheckErr(err)
	//dir := flag.String("dir", "C:\\Utilisateurs\\a706836\\go\\src\\fileManager\\files", "dir for all directory")
	//flag.StringVar(&params.DirIn, "dirIn", "", "dir to scan")
	//ext := flag.String("exts", "", "extentions to find")
	//flag.StringVar(&params.DirOut, "dirOut", currPath, "dir out ") // the dir out pet default will be the path of execution app (currPath)
	//flag.StringVar(&params.Action, "action", "copy", "action to make for files (copy or move) ")
	//flag.Parse()
	//e := strings.Split(*ext, ",")
	//params.Exts = e

	//da.fileManager.StartProcessing(&params) // start processing traitements
	//go counter()
	startGui(da)

}

func counter() {
	for {
		time.Sleep(1 * time.Second)
		count++

		// Update the UI using the QueueMain function
		ui.QueueMain(func() {
			labelcount.SetText(fmt.Sprintf("%d", count))
		})
	}
}
