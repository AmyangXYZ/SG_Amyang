package model

import (
	"strconv"
	"time"
)

// Post Structure.
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Cat   string `json:"cat"`
	HTML  string `json:"html"`
	MD    string `json:"md"`
	Time  string `json:"time"`
}

// GetPosts returns 5 posts per page.
func GetPosts(page string) ([]Post, error) {
	var post Post
	posts := make([]Post, 0)
	p, _ := strconv.Atoi(page)
	rows, err := db.Query("SELECT * from posts ORDER BY id DESC LIMIT ?,5", (p-1)*5)
	defer rows.Close()
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		rows.Scan(&post.ID, &post.Title, &post.Cat, &post.HTML, &post.MD, &post.Time)
		posts = append(posts, post)
	}
	return posts, nil
}

// GetPostByTitle returns the post.
func GetPostByTitle(title string) (Post, error) {
	var post Post
	rows, err := db.Query("SELECT * from posts where title=? limit 1", title)
	defer rows.Close()
	if err != nil {
		return post, err
	}
	for rows.Next() {
		rows.Scan(&post.ID, &post.Title, &post.Cat, &post.HTML, &post.MD, &post.Time)
	}
	return post, nil
}

// GetPostsByCat returns the posts, 5 per page.
func GetPostsByCat(cat string, page string) ([]Post, error) {
	var post Post
	posts := make([]Post, 0)
	p, _ := strconv.Atoi(page)
	rows, err := db.Query("SELECT * from posts where cat=? ORDER BY id DESC LIMIT ?,5", cat, (p-1)*5)
	defer rows.Close()
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		rows.Scan(&post.ID, &post.Title, &post.Cat, &post.HTML, &post.MD, &post.Time)
		posts = append(posts, post)
	}
	return posts, nil
}

// NewPost inserts a new post into database.
func NewPost(title, cat, html, md string) error {
	time := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.Prepare("INSERT posts SET title=?, cat=?, html=?, md=?, time=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(title, cat, html, md, time)
	return err
}

// UpdatePost updates an existed post.
func UpdatePost(newTitle, cat, html, md, oldTitle string) error {
	stmt, err := db.Prepare("UPDATE posts SET title=?, cat=?, html=?, md=? WHERE title=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newTitle, cat, html, md, oldTitle)
	return err
}
