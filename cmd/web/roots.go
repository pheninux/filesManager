package main

import "net/http"

func (wa *WebApplication) root() {
	http.HandleFunc("/action", http.HandlerFunc(wa.Action))
}
