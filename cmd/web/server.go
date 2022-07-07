package main

import (
	"fileManager2/cmd/common"
	"fmt"
	"log"
	"net/http"
	"time"
)

type WebApplication struct {
	fileManager common.IFileManager
	utils       common.Utils
}

func main() {
	srv := &http.Server{ReadTimeout: time.Second * 10000, WriteTimeout: time.Second * 10000, Addr: ":4000"}
	fmt.Println("starting server port :4000")
	log.Fatal(srv.ListenAndServe())
}
