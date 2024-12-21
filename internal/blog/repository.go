// Handles database interactions
// Handles database interactions
package blog

import (
	"database/sql"
	"fmt"
)

// Repository struct handles database interactions
type Repository struct {
	DB *sql.DB
}

// NewRepository creates a new Repository instance
func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

// GetPosts retrieves all blog posts
func (r *Repository) GetPosts() ([]map[string]interface{}, error) {
	rows, err := r.DB.Query("SELECT id, title, content FROM posts")
	if err != nil {
		return nil, fmt.Errorf("error fetching posts: %v", err)
	}
	defer rows.Close()

	var posts []map[string]interface{}
	for rows.Next() {
		var id int
		var title, content string
		if err := rows.Scan(&id, &title, &content); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		posts = append(posts, map[string]interface{}{
			"id":      id,
			"title":   title,
			"content": content,
		})
	}

	return posts, nil
}
