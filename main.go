package main

import (
	"fmt"
	"goweb/controler"
	"goweb/util"
	"net/http"
)

func main() {

	mux := controler.Controler{}
	controler.Init(&mux)
	server := http.Server{
		Addr:    controler.HostIP + ":" + controler.HostPort,
		Handler: &mux,
	}
	message := fmt.Sprintf("Service at %s\n> > > > > > >\n", controler.HostAddr)
	util.ColorPrintf(message, util.Green)
	server.ListenAndServe()
}
