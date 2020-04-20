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
		Addr:    controler.HostIP + ":" + controler.HostPort,
		Handler: mux,
	}

	util.ColorPrintf("Service at: ", util.Green)
	util.ColorPrintf(controler.HostAddr+"\n", util.Yellow)
	server.ListenAndServe()
}
