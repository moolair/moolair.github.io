// Business logic layer
// Business logic layer
package blog

import (
	"fmt"
)

// Service struct handles business logic
type Service struct {
	repo *Repository
}

// NewService creates a new Service instance
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetAllPosts retrieves all posts with additional business logic
func (s *Service) GetAllPosts() ([]map[string]interface{}, error) {
	// Fetch posts from the repository
	posts, err := s.repo.GetPosts()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts: %v", err)
	}

	// Example: Add a business logic layer (e.g., format modification)
	for _, post := range posts {
		post["summary"] = fmt.Sprintf("%.50s...", post["content"].(string)) // Add a summary field
	}

	return posts, nil
}
