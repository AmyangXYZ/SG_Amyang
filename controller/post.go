package controller

import (
	"fmt"
	"strings"

	"../model"
	"github.com/AmyangXYZ/sweetygo"
)

// Show Post Page Handler.
func Show(ctx *sweetygo.Context) {
	if title := ctx.Param("title"); title != "" {
		title := strings.Replace(title, "-", " ", -1)
		post, err := model.GetPostByTitle(title)
		if err != nil {
			ctx.Error("get post error", 200)
			return
		}
		if post.ID == 0 {
			ctx.Error("404 page not found", 404)
			return
		}
		ctx.Set("post", post)
		ctx.Set("title", title)
		ctx.Set("show", true)
		ctx.Render(200, "posts/show")
	}
}

// Cat shows posts sorted by category.
func Cat(ctx *sweetygo.Context) {
	if cat := ctx.Param("cat"); cat != "" {
		posts, err := model.GetPostsByCat(cat, "1")
		if err != nil {
			ctx.Error("get posts error", 200)
			return
		}
		ctx.Set("posts", posts)
		b := []byte(cat)
		b[0] -= 32 // uppercase
		cat = string(b)
		ctx.Set("title", cat)
		ctx.Render(200, "posts/cat")
	}
}

// Page shows 5 posts per page.
func Page(ctx *sweetygo.Context) {
	ctx.Text(200, "page is "+ctx.Param("n"))
}

// NewPage is Create Post Page Handler
func NewPage(ctx *sweetygo.Context) {
	ctx.Set("title", "New")
	ctx.Set("editor", true)
	ctx.Render(200, "posts/new")
}

// EditPage is Edit Post Page Handler
func EditPage(ctx *sweetygo.Context) {
	ctx.Set("title", "Edit")
	ctx.Set("editor", true)
	title := ctx.Param("title")
	title = strings.Replace(title, "-", " ", -1)
	post, _ := model.GetPostByTitle(title)
	ctx.Set("post", post)
	ctx.Render(200, "posts/update")
}

// Get Post API Handler.
//
// Usage:
// 	"/api/posts/Hello" 		-> the post.
func Get(ctx *sweetygo.Context) {
	if title := ctx.Param("title"); title != "" {
		post, err := model.GetPostByTitle(title)
		if err != nil {
			ctx.JSON(200, "Get Posts Error", "error")
			return
		}
		// no data
		if post.ID == 0 {
			ctx.JSON(200, "There's nothing you want.", "fail")
			return
		}
		ctx.JSON(200, post, "success")
		return
	}

	if page := ctx.Param("page"); page != "" {
		if cat := ctx.Param("cat"); cat != "" {
			posts, err := model.GetPostsByCat(cat, page)
			if err != nil {
				ctx.JSON(500, "Get Posts Error", "error")
				return
			}
			// no data
			if len(posts) == 0 {
				ctx.JSON(200, "There's nothing u want.", "fail")
				return
			}
			data := map[string][]model.Post{"posts": posts}
			ctx.JSON(200, data, "success")
			return
		}

		posts, err := model.GetPosts(page)
		if err != nil {
			ctx.JSON(500, "Get Posts Error", "error")
			return
		}
		// no data
		if len(posts) == 0 {
			ctx.JSON(200, "There's nothing u want.", "fail")
			return
		}
		data := map[string][]model.Post{"posts": posts}
		ctx.JSON(200, data, "success")
		return
	}
	ctx.JSON(200, "I don't understand what u want.", "fail")
}

// New Post API Handler.
//
// Usage:
//  "/api/posts/new" -X POST -d "title=xx&cat=xx&text=xx"
func New(ctx *sweetygo.Context) {
	title := ctx.Param("title")
	cat := ctx.Param("cat")
	html := ctx.Param("html")
	md := ctx.Param("md")
	if title != "" && cat != "" && html != "" && md != "" {
		err := model.NewPost(title, cat, html, md)
		if err != nil {
			ctx.JSON(200, "create post error", "error")
			return
		}
		ctx.JSON(201, "", "success")
		return
	}
	ctx.JSON(200, "I can't understand what u want", "fail")
}

// Update Post API Handler.
//
// Usage:
// 	"/api/post" -X PUT -d "title=xx&cat=xx&text=xx"
func Update(ctx *sweetygo.Context) {
	title := ctx.Param("title")
	cat := ctx.Param("cat")
	html := ctx.Param("html")
	md := ctx.Param("md")
	if title != "" && cat != "" && html != "" && md != "" {
		err := model.UpdatePost(title, cat, html, md)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(200, "update post error", "error")
			return
		}
		ctx.JSON(201, "", "success")
		return
	}
	ctx.JSON(200, "I can't understand what u want", "fail")
}
