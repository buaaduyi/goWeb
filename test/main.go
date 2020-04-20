package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func test(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/test.html")
		t.Execute(w, nil)
	}
}

func main() {
	images := http.FileServer(http.Dir("images"))
	http.Handle("/static/", http.StripPrefix("/static/", images))
	http.HandleFunc("/", test)
	fmt.Println("ready")
	http.ListenAndServe("localhost:8080", nil)
}
