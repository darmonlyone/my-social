package social

import (
	"context"
	"time"
)

type PostID = string

type Post struct {
	ID        PostID    `json:"id"`
	CreatedBy AccountID `json:"created_by"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PostRepo interface {
	Store(ctx context.Context, post *Post) error
	Find(ctx context.Context, id PostID) (*Post, error)
	FindAll(ctx context.Context) ([]*Post, error)
	Delete(ctx context.Context, id PostID) error
	Update(ctx context.Context, post *Post) (*Post, error)
}
