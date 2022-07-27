package main

import (
	"encoding/json"
	"fileManager2/pkg/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (wa *WebApplication) Action(w http.ResponseWriter, r *http.Request) {
	s := make(chan *models.Stack)
	if r.Method == "GET" {
		fmt.Println("the server is ok")
	} else {

		data, err := ioutil.ReadAll(r.Body)
		wa.utils.CheckErr(err)
		param := models.DataTemplate{}
		wa.utils.CheckErr(json.Unmarshal(data, &param))
		wa.fileManager.StartProcessing(&param, s)
		fmt.Println(param)
		w.Write([]byte("data sent"))
	}

}
