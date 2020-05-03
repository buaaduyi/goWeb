package main

import (
	"html/template"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/test.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe("localhost:8080", nil)
}
