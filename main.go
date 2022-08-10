package main

import (
	"fmt"
	"net/http"

	"0jinxing.github.io/gee/gee"
)

func main() {
	app := gee.New()

	app.GET("/hello", func(ctx *gee.Context) {
		msg := fmt.Sprintf("hello %s, you're at %s\n", ctx.Query("name"), ctx.Req.URL.Path)
		ctx.Status(http.StatusOK).Text(msg)
	})

	app.POST("/login", func(ctx *gee.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		ctx.Status(http.StatusOK).JSON(gee.H{"username": username, "password": password})
	})

	app.Run(":9999")
}
