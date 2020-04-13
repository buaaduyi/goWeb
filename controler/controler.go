package controler

import (
	"fmt"
	"goweb/db"
	"goweb/util"
	"html/template"
	"net/http"
)

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
	c.HandlerMap["/test/"] = test
}

func (c *Controler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if c.HandlerMap[r.URL.Path] != nil {
		handle := c.HandlerMap[r.URL.Path]
		handle(w, r)
	} else {
		notFound(w, r)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/home.html")
		util.CheckErr(err)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		// cookie := r.Header["Cookie"]
		// jump to myHomePage
		w.Header().Set("Location", "http://localhost:8080/myhome/")
		w.WriteHeader(302)
	}
}

func myHomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/myhome.html")
		util.CheckErr(err)
		posts := db.GetPostByAuthor("duyi")
		t.Execute(w, posts)
	} else if r.Method == "POST" {
		w.Header().Set("Location", "http://localhost:8080/post/")
		w.WriteHeader(302)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "404 not found")
}

func postsFmt(posts []db.Post) string {
	var reply string
	for _, post := range posts {
		reply += fmt.Sprintf("时间: %s 作者: %s\n%s\n", post.Time, post.Author, post.Content)
		for i, comment := range post.Comments {
			if i == 0 {
				reply += "评论:\n"
			}
			reply += fmt.Sprintf("%s: %s\n", comment.Author, comment.Content)
		}
		reply += "\n"
	}
	return reply
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
		r.ParseForm()
		content := r.PostForm.Get("content")
		post := db.Post{}
		post.Author = "duyi"
		post.Content = content
		post.Create()
		w.Header().Set("Location", "http://localhost:8080/myhome/")
		w.WriteHeader(302)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/test.tmpl")
	// posts := db.GetPostByAuthor("duyi")
	// reply := postsFmt(posts)
	reply := []string{"hello", "world"}
	t.Execute(w, reply)
}
