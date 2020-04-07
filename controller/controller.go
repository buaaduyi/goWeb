package controller

import (
	"fmt"
	"goweb/utils"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

// Controller struct
type Controller struct {
	HTTPAddr  string
	RedisAddr string
	r         redis.Conn
	message   []byte
}

// Init controller
func (c *Controller) Init() {
	http.HandleFunc("/db", c.Get)
	var err error
	c.r, err = redis.Dial("tcp", c.RedisAddr)
	if utils.ErrOccur(err) == true {
		c.r.Close()
		return
	}
	utils.ColorPrintf("Server is ready\n", utils.Green)
	err = http.ListenAndServe(c.HTTPAddr, nil)
	if utils.ErrOccur(err) == true {
		return
	}

}

// Finish controller
func (c *Controller) Finish() {
	c.r.Close()
	utils.ColorPrintf("bye bye ...\n", utils.Green)
}

// Get method
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		method := r.Form.Get("method")
		key := r.Form.Get("key")
		if method == "get" || method == "GET" {
			reply, err := redis.String(c.r.Do("GET", key))
			if utils.ErrOccur(err) == true {
				fmt.Fprintf(w, "%s\n", err.Error())
			} else {
				fmt.Fprintf(w, "%s\n", reply)
				utils.ColorPrintf("GET SUCCESS\n", utils.Blue)
			}
		} else if method == "set" || method == "SET" {
			val := r.Form.Get("val")
			var err error
			reply, err := redis.String(c.r.Do("SET", key, val))
			if utils.ErrOccur(err) == true {
				fmt.Fprintf(w, "%s\n", err.Error())
			} else {
				fmt.Fprintf(w, "%s\n", reply)
				utils.ColorPrintf("SET SUCCESS\n", utils.Blue)
			}
		}
	} else {
		utils.ColorPrintf("METHOD ERROR\n", utils.Red)
	}
}
