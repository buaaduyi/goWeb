package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:     "duyi",
		Value:    "12345678",
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header["Cookie"]
	fmt.Println(cookie)
}

func main() {

}
