package controller

import (
	"../model"
	"github.com/AmyangXYZ/sweetygo"
)

// Show Post Page Handler.
func Show(ctx *sweetygo.Context) {
	ctx.Render(200, "post/show")
}

// Get Post API Handler.
//
// Usage:
// 	"/api/post?title=Hello" 		-> the post.
// 	"/api/post?page=1"      		-> 5 posts (perpage default is 5).
// 	"/api/post?cat=sec&page=1"		-> 5 posts about Sec
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
//  "/api/post" -X POST -d "title=xx&cat=xx&text=xx"
func New(ctx *sweetygo.Context) {
	title := ctx.Param("title")
	cat := ctx.Param("cat")
	text := ctx.Param("text")
	if title != "" && cat != "" && text != "" {
		err := model.NewPost(title, cat, text)
		if err != nil {
			ctx.JSON(200, "New Post Error", "error")
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
	text := ctx.Param("text")
	if title != "" && cat != "" && text != "" {
		err := model.UpdatePost(title, cat, text)
		if err != nil {
			ctx.JSON(200, "Update Post Error", "error")
			return
		}
		ctx.JSON(201, "", "success")
		return
	}
	ctx.JSON(200, "I can't understand what u want", "fail")
}
