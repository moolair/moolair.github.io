package blog

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Created string `json:"created_at,omitempty"`
}

func CreatePostHandler(c *gin.Context, repo *Repository) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// 현재 시간 추가
	post.Created = time.Now().Format(time.RFC3339)

	// 디버그 로그 추가
	log.Printf("Creating post: %+v", post)

	err := repo.CreatePost(post)
	if err != nil {
		log.Printf("Error creating post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post": post})
}

func GetPostsHandler(c *gin.Context, repo *Repository) {
	// Firestore에서 게시물 가져오기
	posts, err := repo.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 빈 배열이라도 반환하도록 수정
	if posts == nil {
		posts = []Post{}
	}

	c.JSON(http.StatusOK, posts)
}
