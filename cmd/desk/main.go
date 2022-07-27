package main

import (
	"fileManager2/cmd/common"
	"fileManager2/pkg/models"
	"fileManager2/pkg/models/boImpl"
	"runtime"
)

type DeskApplication struct {
	fileManager  common.IFileManager
	dataTemplate *models.DataTemplate
	utils        common.Utils
	os           string
}

func main() {

	s := make(chan *models.Stack)
	//count = 0
	//params := models.DataTemplate{Action: "copy", DirOut: "C:\\Users\\a706836\\go\\src\\filesManager2", DirIn: "C:\\Users\\a706836\\Downloads", Exts: []string{"pdf"}}
	u := common.Utils{}
	da := &DeskApplication{fileManager: &boImpl.FileModel{Utils: &u}, utils: common.Utils{}}
	switch runtime.GOOS {
	case "windows":
		u.Slash = "\\"
		da.os = "windows"
	case "android":
		u.Slash = "//"
		da.os = "android"
	}

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
	//startGui(da)
	startGuiFyne(da, s)

	// TODO REGLER LE SOUCIE SI UN DOSSIER EST VIDE FAUT LE SUPPRIMER
	// todo logs ( numbre file , taille file , ect ...)
	// todo theming
	// todo shwo error in msgbox when something turn wrong

}
