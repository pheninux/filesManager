package main

import (
	"encoding/json"
	"fileManager2/pkg/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (wa *WebApplication) Action(w http.ResponseWriter, r *http.Request) {
	ccount := make(chan int)
	if r.Method == "GET" {
		fmt.Println("the server is ok")
	} else {

		time.Sleep(time.Second * 7000)
		data, err := ioutil.ReadAll(r.Body)
		wa.utils.CheckErr(err)
		param := models.DataTemplate{}
		wa.utils.CheckErr(json.Unmarshal(data, &param))
		wa.fileManager.StartProcessing(&param, ccount)
		fmt.Println(param)
		w.Write([]byte("data sent"))
	}

}
