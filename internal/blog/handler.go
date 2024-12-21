// HTTP handlers for blog endpoints
package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPostsHandler(c *gin.Context) {
	posts := []gin.H{
		{"id": 1, "title": "First Post", "content": "This is the first post."},
		{"id": 2, "title": "Second Post", "content": "This is the second post."},
	}
	c.JSON(http.StatusOK, posts)
}
