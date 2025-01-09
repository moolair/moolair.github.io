package blog

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/db"
)

type Repository struct {
	Client *db.Client
}

func NewRepository(client *db.Client) *Repository {
	return &Repository{Client: client}
}

func (r *Repository) CreatePost(post Post) error {
	ref := r.Client.NewRef("posts")

	// 새 포스트를 위한 고유 키 생성
	newRef, err := ref.Push(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to create reference: %v", err)
	}

	// 데이터 저장
	err = newRef.Set(context.Background(), post)
	if err != nil {
		return fmt.Errorf("failed to create post in database: %v", err)
	}

	return nil
}

func (r *Repository) GetPosts() ([]Post, error) {
	ref := r.Client.NewRef("posts")
	var data map[string]Post

	if err := ref.Get(context.Background(), &data); err != nil {
		return nil, fmt.Errorf("failed to fetch posts: %v", err)
	}

	// map을 slice로 변환
	posts := make([]Post, 0, len(data))
	for id, post := range data {
		post.ID = id
		posts = append(posts, post)
	}

	return posts, nil
}
