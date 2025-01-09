package main

import (
	"context"
	"log"

	"moolair.github.io/internal/blog"
	"moolair.github.io/internal/middleware"
	"moolair.github.io/pkg/db"

	"github.com/gin-gonic/gin"
	"moolair.github.io/pkg/utils"
)

func main() {
	utils.InitLogger()
	utils.LogInfo("Server started")

	// Firebase 초기화
	app, err := db.InitFirebase()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Auth: %v", err)
	}

	client, err := app.Database(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to Firestore: %v", err)
	}

	// Gin Router 설정
	repo := blog.NewRepository(client)
	r := gin.Default()

	// Static files
	r.Static("/static", "./web/static")

	// HTML templates
	r.LoadHTMLGlob("web/templates/*")

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Public Routes
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	r.GET("/posts", func(c *gin.Context) {
		blog.GetPostsHandler(c, repo)
	})

	// Admin Routes
	authMiddleware := middleware.AuthMiddleware(authClient, "a7V3odBt4STnCMXzLjWKIhFxouz2")
	admin := r.Group("/admin", authMiddleware)
	admin.POST("/post", func(c *gin.Context) {
		blog.CreatePostHandler(c, repo)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin.html", nil)
	})

	// 서버 실행
	r.Run(":8080")
}
