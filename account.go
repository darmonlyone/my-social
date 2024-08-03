package social

import (
	"context"
	"time"
)

type AccountID = string

type Account struct {
	ID             AccountID `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	FirstName      string    `json:"firstname"`
	LastName       string    `json:"lastname"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type AccountRepo interface {
	Store(ctx context.Context, account *Account) error
	Find(ctx context.Context, id AccountID) (*Account, error)
	FindByUsername(ctx context.Context, username string) (*Account, error)
}
