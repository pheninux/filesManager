package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type parametres struct {
	DirIn  string   `json:"dir_in"`
	DirOut string   `json:"dir_out"`
	Action string   `json:"action"`
	Exts   []string `json:"exts"`
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	param := parametres{}
	http.HandleFunc("/action", http.HandlerFunc(action))

	currPath, err := os.Getwd() // get current path
	checkErr(err)

	//dir := flag.String("dir", "C:\\Utilisateurs\\a706836\\go\\src\\fileManager\\files", "dir for all directory")
	flag.StringVar(&param.DirIn, "dirIn", "", "dir to scan")
	ext := flag.String("exts", "", "extentions to find")
	flag.StringVar(&param.DirOut, "dirOut", currPath, "dir out ") // the dir out pet default will be the path of execution app (currPath)
	flag.StringVar(&param.Action, "action", "copy", "action to make for files (copy or move) ")
	flag.Parse()

	e := strings.Split(*ext, ",")
	param.Exts = e

	startProcessing(&param) // start processing traitements

	fmt.Println("starting server port :4000")
	srv := &http.Server{ReadTimeout: time.Second * 10000, WriteTimeout: time.Second * 10000, Addr: ":4000"}
	log.Fatal(srv.ListenAndServe())

}

func validateArgs(p *parametres) error {

	var regex = "[a-zA-Z]:\\\\(((?![<>:\"/\\\\|?*]).)+((?<![ .])\\\\)?)*$"
	b, err := regexp.MatchString(regex, p.DirIn)
	checkErr(err)
	if b {
		return fmt.Errorf("répertoire cible format incorrect :[%s]", err)

	}
	b, err = regexp.MatchString(regex, p.DirOut)
	checkErr(err)
	if b {
		return fmt.Errorf("répertoire sortie format incorrect :[%s]", err)

	}
	return nil
}

func startProcessing(param *parametres) {
	validateArgs(param) // validate args

	var wg sync.WaitGroup
	wg.Add(2)
	go makeOutDirs(param.Exts, param.DirOut, &wg) // create directories for different extentions of files

	fi, err := ioutil.ReadDir(param.DirIn) //read the content of dir files and folder
	checkErr(err)
	go process(fi, param, &wg)

	wg.Wait()

}

func action(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("the server is ok")
	} else {

		time.Sleep(time.Second * 7000)
		data, err := ioutil.ReadAll(r.Body)
		s := struct {
			DirIn  string   `json:"dir_in"`
			DirOut string   `json:"dir_out"`
			Action string   `json:"action"`
			Exts   []string `json:"exts"`
		}{}
		param := parametres(s)
		checkErr(json.Unmarshal(data, &param))
		checkErr(err)
		startProcessing(&param)
		fmt.Println(param)
		w.Write([]byte("data sent"))
	}

}

func process(fi []os.FileInfo, param *parametres, wg *sync.WaitGroup) {
	var wgp sync.WaitGroup
	defer wg.Done()

	for _, entry := range fi {
		if entry.IsDir() {
			continue
		}
		wgp.Add(1)
		go checkExtAndCopy(entry, param, &wgp)
		fmt.Println(" ", entry.Name(), entry.IsDir(), filepath.Ext(entry.Name()))
	}
	wgp.Wait()
}

// check every extentions and copy file in the appropriate folder
func checkExtAndCopy(entry os.FileInfo, param *parametres, wgp *sync.WaitGroup) {
	defer wgp.Done()
	var wgc sync.WaitGroup
	for _, ext := range param.Exts {
		if filepath.Ext(entry.Name()) == "."+ext { // if the entry extention is equal to the given extention args
			src := param.DirIn + "\\" + entry.Name()                      // create a source path
			dest := param.DirOut + "\\" + ext + "-files\\" + entry.Name() // create a distination path
			wgc.Add(1)
			go copyOrMove(src, dest, param.Action, &wgc)
		}
	}
	wgc.Wait()
}

func copyOrMove(src, dst, action string, wgc *sync.WaitGroup) {
	defer wgc.Done()
	switch strings.ToLower(action) {
	case "copy":
		src_file, err := os.Open(src)
		checkErr(err)
		defer src_file.Close()

		src_file_stat, err := src_file.Stat()
		checkErr(err)

		if !src_file_stat.Mode().IsRegular() {
			fmt.Errorf("%s is not a regular file", src)
		}
		dst_file, err := os.Create(dst)
		checkErr(err)
		defer dst_file.Close()

		_, err = io.Copy(dst_file, src_file)
		checkErr(err)
	case "move":
		checkErr(os.Rename(src, dst))
	}

}

func makeOutDirs(exts []string, dirOut string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range exts {
		os.Mkdir(dirOut+"\\"+v+"-files", 0755)
	}
}

func detectFileType(name string, dir string) string {

	f, err := os.Open(dir + name)
	checkErr(err)
	defer f.Close()
	// Only the first 512 bytes are used to sniff the content type.
	buff := make([]byte, 512)
	_, err = f.Read(buff)
	checkErr(err)
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	cType := http.DetectContentType(buff)
	return cType
}
