package main

import (
	"crypto/md5"
	"fmt"
	"goweb/utils"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/websocket"
)

type rocket struct {
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func saybye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bye bye")
}

var temp string

type MD5Code struct {
	MD5 string
}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method: ", r.Method)

	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		temp = token
		t := template.New("")
		t, _ = t.Parse("ID: {{.MD5}}\n")
		var md5 MD5Code
		md5.MD5 = token
		t.Execute(w, md5)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		utils.ColorPrintf(temp+"\n", utils.Yellow)
		utils.ColorPrintf(token+"\n", utils.Blue)

		if temp == token {
			utils.ColorPrintf("pass\n", utils.Green)
		} else {
			utils.ColorPrintf("not pass\n", utils.Red)
		}

	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {

	}
}

func (roc *rocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hel" {
		sayhello(w, r)
		return
	}
	if r.URL.Path == "/bye" {
		saybye(w, r)
		return
	}
	if r.URL.Path == "/login" {
		login(w, r)
		return
	}
	if r.URL.Path == "/upload" {

	}
	http.NotFound(w, r)
	return
}

func echo(ws *websocket.Conn) {
	for {
		var reply string
		websocket.Message.Receive(ws, &reply)
		fmt.Println("Received: " + reply)
		for i := 0; i < 10; i++ {
			msg := "Hi :" + strconv.Itoa(i)
			err := websocket.Message.Send(ws, msg)
			if err != nil {
				fmt.Println("cant send")
				return
			}
			time.Sleep(time.Second)
		}

	}
}

var name string

func main() {

	fmt.Println("begin")
	var count int

	for {
		count++
		fmt.Println(count)
		if count == 10 {
			break
		}
	}

	fmt.Println("end")
	// var roc rocket
	// utils.ColorPrintf(">>>", utils.Red)
	// http.ListenAndServe("localhost:8080", &roc)

}
