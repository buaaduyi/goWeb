package controler

import (
	"fmt"
	"goweb/db"
	"goweb/util"
	"html/template"
	"net/http"
)

const host = "http://192.168.0.104:8080/"

// Controler controle every thing
type Controler struct {
	HandlerMap map[string]func(w http.ResponseWriter, r *http.Request)
}

// Init all
func Init(mux *Controler, d db.DSN) {
	db.InitDB(d)
	mux.InitControler()
}

// InitControler init the controler
func (c *Controler) InitControler() {
	c.HandlerMap = map[string]func(w http.ResponseWriter, r *http.Request){}
	c.HandlerMap["/"] = homePage
	c.HandlerMap["/myhome/"] = myHomePage
	c.HandlerMap["/post/"] = createPost
	c.HandlerMap["/singup/"] = singUP
	c.HandlerMap["/login/"] = login
}

func (c *Controler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if c.HandlerMap[r.URL.Path] != nil {
		handle := c.HandlerMap[r.URL.Path]
		handle(w, r)
	} else {
		notFound(w, r)
	}
}

func singUP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/singup.html")
		util.CheckErr(err)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		choice := r.PostForm.Get("button")
		if choice == "singup" {
			name := r.PostForm.Get("username")
			pwd := r.PostForm.Get("password")
			email := r.PostForm.Get("email")
			if name != "" && pwd != "" && email != "" {
				user := db.User{}
				user.Name = name
				user.Pwd = pwd
				user.Email = email
				user.ID = util.MD5Code(name + pwd)
				if user.Create() == true {
					w.Header().Set("Location", host+"login/")
					w.WriteHeader(302)
				} else {
					t, err := template.ParseFiles("template/singup.html")
					util.CheckErr(err)
					t.Execute(w, "注册失败")
				}
			} else {
				t, err := template.ParseFiles("template/singup.html")
				util.CheckErr(err)
				t.Execute(w, "信息缺失")
			}
		} else if choice == "cancel" {
			w.Header().Set("Location", host)
			w.WriteHeader(302)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/login.html")
		util.CheckErr(err)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		choice := r.PostForm.Get("button")
		if choice == "login" {
			name := r.PostForm.Get("username")
			pwd := r.PostForm.Get("password")
			user := db.GetUserByName(name)
			if user.Pwd != "" && user.Pwd == pwd {
				cookie := http.Cookie{
					Name:     name,
					Value:    user.ID,
					HttpOnly: true,
					Path:     "/",
				}
				http.SetCookie(w, &cookie)
				w.Header().Set("Location", host+"myhome/")
				w.WriteHeader(302)
			} else {
				t, err := template.ParseFiles("template/login.html")
				util.CheckErr(err)
				t.Execute(w, "密码错误")
			}
		} else if choice == "singup" {
			w.Header().Set("Location", host+"singup/")
			w.WriteHeader(302)
		}
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/home.html")
		util.CheckErr(err)
		posts := db.GetAllPosts()
		t.Execute(w, posts)
	} else if r.Method == "POST" {
		r.ParseForm()
		choice := r.PostForm.Get("button")
		if choice == "myhome" {
			w.Header().Set("Location", host+"myhome/")
			w.WriteHeader(302)
		} else {
			cookie := r.Cookies()
			if len(cookie) != 0 {
				content := r.PostForm.Get("comment" + choice)
				user := db.GetUserByID(cookie[0].Value)
				comment := db.Comment{
					Author:  user.Name,
					PostID:  choice,
					Content: content,
				}
				comment.Create()
				w.Header().Set("Location", host)
				w.WriteHeader(302)
			} else {
				w.Header().Set("Location", host+"login/")
				w.WriteHeader(302)
			}
		}
	}
}

func myHomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cookie := r.Cookies()
		if len(cookie) != 0 {
			user := db.GetUserByID(cookie[0].Value)
			if user.ID != "" {
				t, err := template.ParseFiles("template/myhome.html")
				util.CheckErr(err)
				posts := db.GetPostByAuthor(user.Name)
				t.Execute(w, posts)
			} else {
				w.Header().Set("Location", host+"login/")
				w.WriteHeader(302)
			}
		} else {
			w.Header().Set("Location", host+"login/")
			w.WriteHeader(302)
		}

	} else if r.Method == "POST" {
		r.ParseForm()
		choice := r.PostForm.Get("button")
		if choice == "create" {
			w.Header().Set("Location", host+"post/")
			w.WriteHeader(302)
		} else if choice == "homepage" {
			w.Header().Set("Location", host)
			w.WriteHeader(302)
		} else {
			db.DeletePost(choice)
			w.Header().Set("Location", host+"myhome/")
			w.WriteHeader(302)
		}
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "404 not found")
}

func getPostByID(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	post := db.GetPostByID(id)
	fmt.Println(post.Content)
	reply := fmt.Sprintf("时间: %s 作者: %s\n%s\n", post.Time, post.Author, post.Content)
	for i, comment := range post.Comments {
		if i == 0 {
			reply += "评论:\n"
		}
		reply += fmt.Sprintf("%s: %s\n", comment.Author, comment.Content)
	}
	fmt.Fprint(w, reply)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/post.html")
		util.CheckErr(err)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		cookie := r.Cookies()
		r.ParseForm()
		choice := r.PostForm.Get("button")
		if choice == "post" {
			if len(cookie) != 0 {
				user := db.GetUserByID(cookie[0].Value)
				content := r.PostForm.Get("content")
				post := db.Post{}
				post.Author = user.Name
				post.Content = content
				post.Create()
				w.Header().Set("Location", host+"myhome/")
				w.WriteHeader(302)
			} else {
				w.Header().Set("Location", host+"login/")
				w.WriteHeader(302)
			}
		} else if choice == "back" {
			w.Header().Set("Location", host+"myhome/")
			w.WriteHeader(302)
		}
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/test.tmpl")
	// posts := db.GetPostByAuthor("duyi")
	// reply := postsFmt(posts)
	reply := []string{"hello", "world"}
	t.Execute(w, reply)
}
