package main

import (
	"goweb/controler"
	"goweb/db"
	"goweb/util"
	"net/http"
)

func main() {

	var d db.DSN
	d.User = "root"
	d.Pwd = "19960115"
	d.Hostname = "localhost"
	d.Port = "3306"
	d.Schema = "ChitChat"
	// //////////////////////

	mux := controler.Controler{}
	controler.Init(&mux, d)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mux,
	}
	util.ColorPrintf("server is ready\n> > > > > > >\n", util.Green)
	server.ListenAndServe()
}
