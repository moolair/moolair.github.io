package main

import (
	"github.com/gin-gonic/gin"
	"moolair.github.io/internal/blog"
)

func main() {
	r := gin.Default()

	// Static files (CSS, JS, Images)
	r.Static("/static", "./web/static")

	// HTML templates
	r.LoadHTMLGlob("web/templates/*")

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/posts", blog.GetPostsHandler)

	// Start server
	r.Run(":8080")
}
