package db

import (
	"database/sql"
	"fmt"
	"goweb/util"
	"os"
	"time"

	//
	_ "github.com/go-sql-driver/mysql"
)

// User content the user information
type User struct {
	ID    string
	Name  string
	Pwd   string
	Email string
}

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
	Image    bool
}

// Comment store the comment
type Comment struct {
	ID      string
	Content string
	Author  string
	PostID  string
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

// Create user
func (user *User) Create() bool {
	exist := "0"
	query := fmt.Sprintf("select 1 from users where name = '%s' limit 1", user.Name)
	db.QueryRow(query).Scan(&exist)
	if exist == "0" {
		query = fmt.Sprintf("INSERT INTO users (id, name, pwd, email) VALUES ('%s', '%s','%s', '%s')", user.ID, user.Name, user.Pwd, user.Email)
		_, err := db.Exec(query)
		if util.CheckErr(err) == true {
			util.InfoLog("New user: " + user.Name)
			return true
		}
	}
	return false
}

// Create a new post
func (post *Post) Create() {
	post.Time = time.Now().Format("2006-01-02 15:04:05")
	post.ID = util.MD5Code(post.Content + post.Author + post.Time)
	image := 0
	if post.Image == true {
		image = 1
	}
	query := fmt.Sprintf("INSERT INTO posts (post_id, time, content, image, author) VALUES ('%s', '%s', '%s', '%d', '%s')", post.ID, post.Time, post.Content, image, post.Author)
	_, err := db.Exec(query)
	if util.CheckErr(err) == true {
		util.InfoLog("New post: " + post.ID)
	}
}

// DeletePost by id
func DeletePost(id string) {
	var imageFlag bool
	query := fmt.Sprintf("SELECT image from posts WHERE post_id='%s'", id)
	db.QueryRow(query).Scan(&imageFlag)
	if imageFlag == true {
		path := "images/" + id + ".jpeg"
		err := os.Remove(path)
		util.ErrorLog(err)
	}
	query = fmt.Sprintf("DELETE FROM posts WHERE post_id='%s'", id)
	_, err := db.Exec(query)
	util.CheckErr(err)
	DeleteCommentByPostID(id)
	util.InfoLog("Delete post: " + id)

}

// Create a new comment
func (comment *Comment) Create() {
	comment.Time = time.Now().Format("2006-01-02 15:04:05")
	comment.ID = util.MD5Code(comment.Content + comment.Author + comment.Time)
	query := fmt.Sprintf("INSERT INTO comments (comment_id, time, content, author, post_id) VALUES ('%s', '%s','%s', '%s', '%s')", comment.ID, comment.Time, comment.Content, comment.Author, comment.PostID)
	_, err := db.Exec(query)
	util.CheckErr(err)
}

//DeleteComment by id
func DeleteComment(id string) {
	query := fmt.Sprintf("DELETE FROM comments WHERE comment_id='%s'", id)
	_, err := db.Exec(query)
	util.CheckErr(err)
}

//DeleteCommentByPostID by post_id
func DeleteCommentByPostID(postID string) {
	query := fmt.Sprintf("DELETE FROM comments WHERE post_id='%s'", postID)
	_, err := db.Exec(query)
	util.CheckErr(err)
}

// GetAllPosts get all posts
func GetAllPosts() []Post {
	posts := []Post{}
	query := fmt.Sprintf("SELECT post_id, time, content, image, author from posts ORDER BY time DESC")
	postRows, err := db.Query(query)
	defer postRows.Close()
	util.CheckErr(err)
	for postRows.Next() {
		post := Post{}
		post.Comments = []Comment{}
		postRows.Scan(&post.ID, &post.Time, &post.Content, &post.Image, &post.Author)
		query = fmt.Sprintf("SELECT comment_id, time, content, author from comments where post_id='%s' ORDER BY time", post.ID)
		commentRows, err := db.Query(query)
		defer commentRows.Close()
		util.CheckErr(err)
		for commentRows.Next() {
			comment := Comment{PostID: post.ID}
			commentRows.Scan(&comment.ID, &comment.Time, &comment.Content, &comment.Author)
			post.Comments = append(post.Comments, comment)
		}
		posts = append(posts, post)
	}
	return posts
}

// GetPostByID get post from database
func GetPostByID(id string) Post {
	post := Post{}
	post.Comments = []Comment{}
	query := fmt.Sprintf("SELECT post_id, time, content, image,  author from posts where id='%s'", id)
	db.QueryRow(query).Scan(&post.ID, &post.Time, &post.Content, &post.Image, &post.Author)
	query = fmt.Sprintf("SELECT comment_id, time, content, author from comments where post_id='%s'", id)
	rows, err := db.Query(query)
	util.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		comment := Comment{PostID: post.ID}
		rows.Scan(&comment.ID, &comment.Time, &comment.Content, &comment.Author)
		post.Comments = append(post.Comments, comment)
	}
	return post
}

// GetPostByAuthor get post from database
func GetPostByAuthor(author string) []Post {
	posts := []Post{}
	query := fmt.Sprintf("SELECT post_id, time, content, image, author from posts where author='%s' ORDER BY time DESC", author)
	postRows, err := db.Query(query)
	defer postRows.Close()
	util.CheckErr(err)
	for postRows.Next() {
		post := Post{}
		post.Comments = []Comment{}
		postRows.Scan(&post.ID, &post.Time, &post.Content, &post.Image, &post.Author)
		query = fmt.Sprintf("SELECT comment_id, time, content, author from comments where post_id='%s' ORDER BY time", post.ID)
		commentRows, err := db.Query(query)
		defer commentRows.Close()
		util.CheckErr(err)
		for commentRows.Next() {
			comment := Comment{PostID: post.ID}
			commentRows.Scan(&comment.ID, &comment.Time, &comment.Content, &comment.Author)
			post.Comments = append(post.Comments, comment)
		}
		posts = append(posts, post)
	}
	return posts
}

// GetUserByID by id
func GetUserByID(id string) User {
	user := User{}
	query := fmt.Sprintf("SELECT name, pwd, email from users where id='%s'", id)
	err := db.QueryRow(query).Scan(&user.Name, &user.Pwd, &user.Email)
	if util.CheckErr(err) == true {
		user.ID = id
		return user
	}
	user.ID = ""
	return user
}

// GetUserByName is correct?
func GetUserByName(name string) User {
	user := User{}
	query := fmt.Sprintf("SELECT id, pwd, email from users where name='%s'", name)
	err := db.QueryRow(query).Scan(&user.ID, &user.Pwd, &user.Email)
	if util.CheckErr(err) == true {
		return user
	}
	user.ID = ""
	return user
}
