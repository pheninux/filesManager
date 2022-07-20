package boImpl

import (
	"fileManager2/cmd/common"
	"fileManager2/pkg/models"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type FileModel struct {
	*common.Utils
}

func (fm *FileModel) StartProcessing(param *models.DataTemplate) {

	if err := fm.ValidateArgs(param); err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	fi, err := ioutil.ReadDir(param.DirIn) //read the content of dir files and folder
	fm.CheckErr(err)
	go fm.MakeOutDirs(param.Exts, param.DirOut, &wg) // create directories for different extentions of files
	go fm.Process(fi, param, &wg)

	wg.Wait()
}

func IsEmpty(path string) (bool, error) {

	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// OR f.Readdir(1)
	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func (fm *FileModel) Process(fi []os.FileInfo, param *models.DataTemplate, wg *sync.WaitGroup) {
	var wgp sync.WaitGroup
	defer wg.Done()

	for _, entry := range fi {
		if entry.IsDir() {
			continue
		}
		wgp.Add(1)
		go fm.CheckExtAndCopy(entry, param, &wgp)
		fmt.Println(" ", entry.Name(), entry.IsDir(), filepath.Ext(entry.Name()))
	}
	wgp.Wait()
}

/** check every extentions and copy file in the appropriate folder**/
func (fm *FileModel) CheckExtAndCopy(entry os.FileInfo, param *models.DataTemplate, wgp *sync.WaitGroup) {
	defer wgp.Done()
	var wgc sync.WaitGroup
	for _, ext := range param.Exts {
		if filepath.Ext(entry.Name()) == "."+ext { // if the entry extention is equal to the given extention args
			src := param.DirIn + fm.Slash + entry.Name()                               // create a source path
			dest := param.DirOut + fm.Slash + ext + "-files" + fm.Slash + entry.Name() // create a distination path
			wgc.Add(1)
			go fm.CopyOrMove(src, dest, param, &wgc)
		}
	}
	// delete the folder if is empty
	for _, ext := range param.Exts {
		path := param.DirOut + fm.Slash + ext + "-files"
		empty, _ := IsEmpty(path)
		if !empty {
			os.Remove(path)
		}
	}
	wgc.Wait()

}

func (fm *FileModel) CopyOrMove(src, dst string, param *models.DataTemplate, wgc *sync.WaitGroup) {
	defer wgc.Done()
	switch strings.ToLower(param.Action) {
	case "copy":
		src_file, err := os.Open(src)
		fm.CheckErr(err)
		defer src_file.Close()

		src_file_stat, err := src_file.Stat()
		fm.CheckErr(err)

		if !src_file_stat.Mode().IsRegular() {
			fmt.Errorf("%s is not a regular file", src)
		}
		dst_file, err := os.Create(dst)
		fm.CheckErr(err)
		defer dst_file.Close()

		_, err = io.Copy(dst_file, src_file)
		fm.CheckErr(err)
	case "move":
		fm.CheckErr(os.Rename(src, dst))
	}

	// delete the folder if is empty
	for _, ext := range param.Exts {
		path := param.DirOut + fm.Slash + ext + "-files"
		empty, _ := IsEmpty(path)
		if !empty {
			os.Remove(path)
		}
	}
}

func (fm *FileModel) MakeOutDirs(exts []string, dirOut string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range exts {
		os.Mkdir(dirOut+fm.Slash+v+"-files", 0755)
	}
}

func (fm *FileModel) detectFileType(name string, dir string) string {

	f, err := os.Open(dir + name)
	fm.CheckErr(err)
	defer f.Close()
	// Only the first 512 bytes are used to sniff the content type.
	buff := make([]byte, 512)
	_, err = f.Read(buff)
	fm.CheckErr(err)
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	cType := http.DetectContentType(buff)
	return cType
}
