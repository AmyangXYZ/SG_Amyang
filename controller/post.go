package controller

import (
	"strings"

	"github.com/AmyangXYZ/SG_Amyang/model"
	"github.com/AmyangXYZ/sweetygo"
)

// Home Page Handler.
func Home(ctx *sweetygo.Context) {
	ctx.Set("title", "Home")
	posts, err := model.GetPosts("1")
	if err != nil {
		ctx.Text(500, "something error")
	}
	ctx.Set("posts", posts)
	ctx.Render(200, "home")
}

// PaginationHome returns 5 posts per page.
// for load-more button of Home Page.
func PaginationHome(ctx *sweetygo.Context) {
	if page := ctx.Param("n"); page != "" {
		posts, err := model.GetPosts(page)
		if err != nil {
			ctx.Text(500, "something error")
		}
		ctx.JSON(200, posts, "success")
	}
}

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
		b := []byte(cat)
		b[0] -= 32 // uppercase
		cat = string(b)
		ctx.Set("cat", true)
		ctx.Set("posts", posts)
		ctx.Set("title", cat)
		ctx.Render(200, "posts/cat")
	}
}

// PaginationCat returns 5 posts per page.
// for load more of Cat Page.
func PaginationCat(ctx *sweetygo.Context) {
	page := ctx.Param("n")
	cat := ctx.Param("cat")
	if page != "" && cat != "" {
		posts, err := model.GetPostsByCat(cat, page)
		if err != nil {
			ctx.Text(500, "something error")
		}
		ctx.JSON(200, posts, "success")
	}
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
	ctx.Render(200, "posts/edit")
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
			ctx.JSON(500, "create post error", "error")
			return
		}
		ctx.JSON(201, "", "success")
		return
	}
	ctx.JSON(406, "I can't understand what u want", "fail")
}

// Update Post API Handler.
//
// Usage:
// 	"/api/post" -X PUT -d "title=xx&cat=xx&text=xx"
func Update(ctx *sweetygo.Context) {
	oldTitle := ctx.Param("title")     // from url
	newTitle := ctx.Param("new-title") // from form
	cat := ctx.Param("cat")
	html := ctx.Param("html")
	md := ctx.Param("md")
	if newTitle != "" && cat != "" && html != "" && md != "" {
		err := model.UpdatePost(newTitle, cat, html, md, oldTitle)
		if err != nil {
			ctx.JSON(500, "update post error", "error")
			return
		}
		ctx.JSON(201, "", "success")
		return
	}
	ctx.JSON(406, "I can't understand what u want", "fail")
}
