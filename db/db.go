package db

import (
	"database/sql"
	"fmt"
	"goweb/util"
	"time"

	//
	_ "github.com/go-sql-driver/mysql"
)

// DSN data source name
type DSN struct {
	User     string
	Pwd      string
	Hostname string
	Port     string
	Schema   string
}

// Post store the post
type Post struct {
	ID       string
	Content  string
	Author   string
	Comments []Comment
	Time     string
}

// Comment store the comment
type Comment struct {
	ID      string
	Content string
	Author  string
	Post    *Post
	Time    string
}

var db *sql.DB

// InitDB open mysql database
func InitDB(dsn DSN) {
	var err error
	dsnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dsn.User, dsn.Pwd, dsn.Hostname, dsn.Port, dsn.Schema)
	db, err = sql.Open("mysql", dsnStr)
	util.CheckErr(err)
}

// Create a new post
func (post *Post) Create() {
	post.Time = time.Now().Format("2006-01-02 15:04:05")
	post.ID = util.MD5Code(post.Content + post.Author + post.Time)
	query := fmt.Sprintf("INSERT INTO posts (id, time, content, author) VALUES ('%s', '%s','%s', '%s')", post.ID, post.Time, post.Content, post.Author)
	_, err := db.Exec(query)
	if util.CheckErr(err) == true {
		util.ColorPrintf(post.Author+" post new blog\n", util.Blue)
	}
}

// Create a new comment
func (comment *Comment) Create() {
	if comment.Post == nil {
		fmt.Println("Post not found")
		return
	}
	comment.Time = time.Now().Format("2006-01-02 15:04:05")
	comment.ID = util.MD5Code(comment.Content + comment.Author + comment.Time)
	query := fmt.Sprintf("INSERT INTO comments (id, time, content, author, post_id) VALUES ('%s', '%s','%s', '%s', '%s')", comment.ID, comment.Time, comment.Content, comment.Author, comment.Post.ID)
	_, err := db.Exec(query)
	util.CheckErr(err)
}

// GetPostByID get post from database
func GetPostByID(id string) Post {
	post := Post{}
	post.Comments = []Comment{}
	query := fmt.Sprintf("SELECT id, time, content, author from posts where id='%s'", id)
	db.QueryRow(query).Scan(&post.ID, &post.Time, &post.Content, &post.Author)
	query = fmt.Sprintf("SELECT id, time, content, author from comments where post_id='%s'", id)
	rows, err := db.Query(query)
	util.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		comment := Comment{Post: &post}
		rows.Scan(&comment.ID, &comment.Time, &comment.Content, &comment.Author)
		post.Comments = append(post.Comments, comment)
	}
	return post
}

// GetPostByAuthor get post from database
func GetPostByAuthor(author string) []Post {
	posts := []Post{}
	query := fmt.Sprintf("SELECT id, time, content, author from posts where author='%s'", author)
	postRows, err := db.Query(query)
	defer postRows.Close()
	util.CheckErr(err)
	for postRows.Next() {
		post := Post{}
		post.Comments = []Comment{}
		postRows.Scan(&post.ID, &post.Time, &post.Content, &post.Author)
		query = fmt.Sprintf("SELECT id, time, content, author from comments where post_id='%s'", post.ID)
		commentRows, err := db.Query(query)
		defer commentRows.Close()
		util.CheckErr(err)
		for commentRows.Next() {
			comment := Comment{Post: &post}
			commentRows.Scan(&comment.ID, &comment.Time, &comment.Content, &comment.Author)
			post.Comments = append(post.Comments, comment)
		}
		posts = append(posts, post)
	}
	return posts
}
