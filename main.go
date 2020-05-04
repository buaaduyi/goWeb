package main

import (
	"goweb/controler"
	"goweb/util"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	controler.Init(mux)
	server := http.Server{
		Addr:   "172.21.0.8:8080",
		Handler: mux,
	}
	util.ColorPrintf("Listening and serving HTTP on ", util.Green)
	util.ColorPrintf(controler.HostAddr+"\n", util.Yellow)
	server.ListenAndServe()
}
