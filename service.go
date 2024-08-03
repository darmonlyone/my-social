package social

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	// Account
	FindAccount(ctx context.Context, id string) (*Account, error)
	StoreAccount(ctx context.Context, username string, password string, firstname string, lastname string) error

	// Post
	FindPost(ctx context.Context, id string) (*Post, error)
	FindAllPosts(ctx context.Context) ([]*Post, error)
	StorePost(ctx context.Context, createdBy AccountID, title string, content string) error
	UpdatePost(ctx context.Context, id PostID, title *string, content *string) error
	DeletePost(ctx context.Context, id string) error
}
type service struct {
	accountRepo AccountRepo
	postRepo    PostRepo
}

func NewService(accountRepo AccountRepo, postRepo PostRepo) Service {
	return &service{
		accountRepo: accountRepo,
		postRepo:    postRepo,
	}
}

// Account
func (s *service) FindAccount(ctx context.Context, id string) (*Account, error) {
	return s.accountRepo.Find(ctx, id)
}

func (s *service) StoreAccount(ctx context.Context, username string, password string, firstname string, lastname string) error {
	findAccount, _ := s.accountRepo.FindByUsername(ctx, username)
	if findAccount != nil {
		return ErrUserAlreadyExists
	}

	hashedPass, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	account := &Account{
		ID:             uuid.NewString(),
		HashedPassword: string(hashedPass),
		FirstName:      firstname,
		LastName:       lastname,
	}
	return s.accountRepo.Store(ctx, account)
}

// Post
func (s *service) FindPost(ctx context.Context, id string) (*Post, error) {
	return s.postRepo.Find(ctx, id)
}

func (s *service) FindAllPosts(ctx context.Context) ([]*Post, error) {
	return s.postRepo.FindAll(ctx)
}

func (s *service) StorePost(ctx context.Context, createdBy AccountID, title string, content string) error {
	post := &Post{
		ID:        uuid.NewString(),
		CreatedBy: createdBy,
		Title:     title,
		Content:   content,
	}
	return s.postRepo.Store(ctx, post)
}

func (s *service) UpdatePost(ctx context.Context, id PostID, title *string, content *string) error {
	return s.postRepo.Update(ctx, id, title, content)
}

func (s *service) DeletePost(ctx context.Context, id string) error {
	return s.postRepo.Delete(ctx, id)
}
