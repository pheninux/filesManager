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

func (fm *FileModel) StartProcessing(dt *models.DataTemplate, s chan *models.Stack) {

	fi, err := ioutil.ReadDir(dt.DirIn) //read the content of dir files and folder
	if err != nil {
		dt.Stack.Err = err.Error()
		s <- dt.Stack
	}

	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	wg.Add(1)

	go fm.MakeOutDirs(dt.Exts, dt.DirOut, &wg) // create directories for different extentions of files
	wg.Wait()
	fm.Process(fi, dt, &wg2, s)

	wg2.Wait()

	fmt.Println("done")
	dt.Stack.Done = true

	fmt.Println(dt.Stack)
	s <- dt.Stack

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

func (fm *FileModel) Process(fi []os.FileInfo, dt *models.DataTemplate, wg *sync.WaitGroup, s chan *models.Stack) {
	//var wgp sync.WaitGroup
	//defer wg.Done()

	for _, entry := range fi {
		if entry.IsDir() {
			continue
		}
		//wgp.Add(1)
		fm.CheckExtAndCopy(entry, dt, wg, s)
		//fmt.Println(" ", entry.Name(), entry.IsDir(), filepath.Ext(entry.Name()))
	}

	//wgp.Wait()
}

/** check every extentions and copy file in the appropriate folder**/
func (fm *FileModel) CheckExtAndCopy(entry os.FileInfo, dt *models.DataTemplate, wgp *sync.WaitGroup, s chan *models.Stack) {

	//defer wgp.Done()
	//var wgc sync.WaitGroup
	// counter for nbr of file ara copied

	for _, ext := range dt.Exts {
		if filepath.Ext(entry.Name()) == "."+ext { // if the entry extention is equal to the given extention args
			src := dt.DirIn + fm.Slash + entry.Name()                               // create a source path
			dest := dt.DirOut + fm.Slash + ext + "-files" + fm.Slash + entry.Name() // create a distination path
			//wgc.Add(1)
			wgp.Add(1)
			go fm.CopyOrMove(src, dest, dt, wgp, s)
		}
	}

	//wgc.Wait()

}

func (fm *FileModel) CopyOrMove(src, dst string, dt *models.DataTemplate, wgp *sync.WaitGroup, s chan *models.Stack) {

	defer wgp.Done()
	switch strings.ToLower(dt.Action) {
	case "copy":

		/********** 1 er solution *********/

		src_file, err := os.Open(src)

		if err != nil {
			dt.Stack.Err = err.Error()
			s <- dt.Stack
		}
		defer src_file.Close()

		src_file_stat, err := src_file.Stat()

		if err != nil {
			dt.Stack.Err = err.Error()
			s <- dt.Stack
		}

		if !src_file_stat.Mode().IsRegular() {
			dt.Stack.Err = fmt.Errorf("%s is not a regular file", src).Error()
			s <- dt.Stack
		}

		dst_file, err := os.Create(dst)

		if err != nil {
			dt.Stack.Err = err.Error()
			s <- dt.Stack
		}
		defer dst_file.Close()

		_, err = io.Copy(dst_file, src_file)
		if err != nil {
			dt.Stack.Err = err.Error()
			s <- dt.Stack
		}

		/********** 2 eme solution mais les benchmarck c'est pareil avec la premiere******/
		//input, err := ioutil.ReadFile(src)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//
		//err = ioutil.WriteFile(dst, input, 0644)
		//if err != nil {
		//	fmt.Println("Error creating", dst)
		//	fmt.Println(err)
		//	return
		//}

	case "move":

		if err := os.Rename(src, dst); err != nil {
			dt.Stack.Err = err.Error()
			s <- dt.Stack
		}

	}

	// delete the folder if is empty
	for _, ext := range dt.Exts {
		path := dt.DirOut + fm.Slash + ext + "-files"
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
