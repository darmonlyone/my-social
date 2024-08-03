package social

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	// Account
	Login(ctx context.Context, username string, password string) (*Account, error)
	FindAccount(ctx context.Context, id string) (*Account, error)
	StoreAccount(ctx context.Context, username string, password string, firstname string, lastname string) error

	// Post
	FindPost(ctx context.Context, id string) (*Post, error)
	FindAllPosts(ctx context.Context) ([]*Post, error)
	StorePost(ctx context.Context, title string, content string) error
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
func (s *service) Login(ctx context.Context, username string, password string) (*Account, error) {
	authAccount, err := s.accountRepo.FindByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrIncorrectUsernameOrPassword
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(authAccount.HashedPassword),
		[]byte(password),
	); err != nil {
		return nil, ErrIncorrectUsernameOrPassword
	}

	return authAccount, nil
}

func (s *service) FindAccount(ctx context.Context, id string) (*Account, error) {
	return s.accountRepo.Find(ctx, id)
}

func (s *service) StoreAccount(ctx context.Context, username string, password string, firstname string, lastname string) error {
	findAccount, err := s.accountRepo.FindByUsername(ctx, username)
	if err != nil {
		if !errors.Is(err, ErrNotFound) {
			return err
		}
	}
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
		Username:       username,
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

func (s *service) StorePost(ctx context.Context, title string, content string) error {
	userID, ok := UserIDFromContext(ctx)
	if !ok {
		return ErrBadRequest
	}
	post := &Post{
		ID:        uuid.NewString(),
		CreatedBy: userID,
		Title:     title,
		Content:   content,
	}
	return s.postRepo.Store(ctx, post)
}

func (s *service) UpdatePost(ctx context.Context, id PostID, title *string, content *string) error {
	userID, ok := UserIDFromContext(ctx)
	if !ok {
		return ErrBadRequest
	}
	post, err := s.postRepo.Find(ctx, id)
	if err != nil {
		return err
	}

	if post.CreatedBy != userID {
		return ErrAuthNotHavePermission
	}

	return s.postRepo.Update(ctx, id, title, content)
}

func (s *service) DeletePost(ctx context.Context, id string) error {
	return s.postRepo.Delete(ctx, id)
}
