package social_test

import (
	"context"
	"testing"

	social "github.com/darmonlyone/my-social"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Mock implementations of AccountRepo and PostRepo
type MockAccountRepo struct {
	mock.Mock
}

func (m *MockAccountRepo) Store(ctx context.Context, account *social.Account) error {
	args := m.Called(ctx, account)
	return args.Error(0)
}

func (m *MockAccountRepo) Find(ctx context.Context, id social.AccountID) (*social.Account, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*social.Account), args.Error(1)
}

func (m *MockAccountRepo) FindByUsername(ctx context.Context, username string) (*social.Account, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*social.Account), args.Error(1)
}

type MockPostRepo struct {
	mock.Mock
}

func (m *MockPostRepo) Store(ctx context.Context, post *social.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *MockPostRepo) Find(ctx context.Context, id social.PostID) (*social.Post, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*social.Post), args.Error(1)
}

func (m *MockPostRepo) FindAll(ctx context.Context) ([]*social.Post, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*social.Post), args.Error(1)
}

func (m *MockPostRepo) Delete(ctx context.Context, id social.PostID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockPostRepo) Update(ctx context.Context, id social.PostID, title *string, content *string) error {
	args := m.Called(ctx, id, title, content)
	return args.Error(0)
}

func TestLogin(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	username := "testuser"
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	account := &social.Account{
		ID:             uuid.NewString(),
		Username:       username,
		HashedPassword: string(hashedPassword),
	}

	mockAccountRepo.On("FindByUsername", mock.Anything, username).Return(account, nil)

	ctx := context.Background()
	result, err := svc.Login(ctx, username, password)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, account.ID, result.ID)
	mockAccountRepo.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	username := "testuser"
	password := "wrongpassword"
	account := &social.Account{
		ID:             uuid.NewString(),
		Username:       username,
		HashedPassword: "hashedpassword",
	}

	mockAccountRepo.On("FindByUsername", mock.Anything, username).Return(account, nil)

	ctx := context.Background()
	result, err := svc.Login(ctx, username, password)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, social.ErrIncorrectUsernameOrPassword, err)
	mockAccountRepo.AssertExpectations(t)
}

func TestStoreAccount(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	username := "newuser"
	password := "newpassword"
	firstname := "John"
	lastname := "Doe"

	mockAccountRepo.On("FindByUsername", mock.Anything, username).Return((*social.Account)(nil), social.ErrNotFound)

	// Mock the Store method to return no error
	mockAccountRepo.On("Store", mock.Anything, mock.MatchedBy(func(acc *social.Account) bool {
		return acc.Username == username && acc.FirstName == firstname && acc.LastName == lastname
	})).Return(nil)

	ctx := context.Background()
	err := svc.StoreAccount(ctx, username, password, firstname, lastname)

	assert.NoError(t, err)
	mockAccountRepo.AssertExpectations(t)
}


func TestStoreAccount_UserExists(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	username := "existinguser"
	password := "password"
	firstname := "Jane"
	lastname := "Doe"

	existingAccount := &social.Account{
		ID:             uuid.NewString(),
		Username:       username,
		HashedPassword: "hashedpassword",
	}

	mockAccountRepo.On("FindByUsername", mock.Anything, username).Return(existingAccount, nil)

	ctx := context.Background()
	err := svc.StoreAccount(ctx, username, password, firstname, lastname)

	assert.Error(t, err)
	assert.Equal(t, social.ErrUserAlreadyExists, err)
	mockAccountRepo.AssertExpectations(t)
}

func TestFindPost(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	postID := uuid.NewString()
	post := &social.Post{
		ID:        postID,
		CreatedBy: uuid.NewString(),
		Title:     "Post Title",
		Content:   "Post Content",
	}

	mockPostRepo.On("Find", mock.Anything, postID).Return(post, nil)

	ctx := context.Background()
	result, err := svc.FindPost(ctx, postID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, post.ID, result.ID)
	mockPostRepo.AssertExpectations(t)
}

func TestFindPost_NotFound(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)
	postID := uuid.NewString()

	mockPostRepo.On("Find", mock.Anything, postID).Return((*social.Post)(nil), social.ErrNotFound)

	ctx := context.Background()
	_, err := svc.FindPost(ctx, postID)

	assert.Error(t, err)
	assert.Equal(t, social.ErrNotFound, err)
	mockPostRepo.AssertExpectations(t)
}

func TestFindAllPosts(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	posts := []*social.Post{
		{
			ID:        uuid.NewString(),
			CreatedBy: uuid.NewString(),
			Title:     "Post 1",
			Content:   "Content 1",
		},
		{
			ID:        uuid.NewString(),
			CreatedBy: uuid.NewString(),
			Title:     "Post 2",
			Content:   "Content 2",
		},
	}

	mockPostRepo.On("FindAll", mock.Anything).Return(posts, nil)

	ctx := context.Background()
	result, err := svc.FindAllPosts(ctx)

	assert.NoError(t, err)
	assert.Len(t, result, len(posts))
	mockPostRepo.AssertExpectations(t)
}

func TestStorePost(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	userID := uuid.NewString()
	title := "New Post"
	content := "Post Content"

	mockPostRepo.On("Store", mock.Anything, mock.MatchedBy(func(post *social.Post) bool {
		return post.Title == title && post.Content == content && post.CreatedBy == userID
	})).Return(nil)

	ctx := context.WithValue(context.Background(), social.UserIDKey, userID)
	err := svc.StorePost(ctx, title, content)

	assert.NoError(t, err)
	mockPostRepo.AssertExpectations(t)
}

func TestUpdatePost(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	postID := uuid.NewString()
	userID := uuid.NewString()
	newTitle := "Updated Title"
	newContent := "Updated Content"

	post := &social.Post{
		ID:        postID,
		CreatedBy: userID,
		Title:     "Old Title",
		Content:   "Old Content",
	}

	mockPostRepo.On("Find", mock.Anything, postID).Return(post, nil)
	mockPostRepo.On("Update", mock.Anything, postID, &newTitle, &newContent).Return(nil)

	ctx := context.WithValue(context.Background(), social.UserIDKey, userID)
	err := svc.UpdatePost(ctx, postID, &newTitle, &newContent)

	assert.NoError(t, err)
	mockPostRepo.AssertExpectations(t)
}

func TestUpdatePost_NotFound(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	postID := uuid.NewString()
	userID := uuid.NewString()
	title := "test"
	content := "test"


	mockPostRepo.On("Find", mock.Anything, postID).Return((*social.Post)(nil), social.ErrNotFound)

	ctx := context.WithValue(context.Background(), social.UserIDKey, userID)
	err := svc.UpdatePost(ctx, postID, &title, &content)

	assert.Error(t, err)
	assert.Equal(t, social.ErrNotFound, err)
}

func TestUpdatePost_NoPermission(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	postID := uuid.NewString()
	userID := uuid.NewString()

	post := &social.Post{
		ID:        postID,
		CreatedBy: uuid.NewString(), // Different userID
		Title:     "Old Title",
		Content:   "Old Content",
	}

	mockPostRepo.On("Find", mock.Anything, postID).Return(post, nil)

	ctx := context.WithValue(context.Background(), social.UserIDKey, userID)
	err := svc.UpdatePost(ctx, postID, nil, nil)

	assert.Error(t, err)
	assert.Equal(t, social.ErrAuthNotHavePermission, err)
	mockPostRepo.AssertExpectations(t)
}

func TestDeletePost(t *testing.T) {
	mockAccountRepo := new(MockAccountRepo)
	mockPostRepo := new(MockPostRepo)
	svc := social.NewService(mockAccountRepo, mockPostRepo)

	postID := uuid.NewString()

	mockPostRepo.On("Delete", mock.Anything, postID).Return(nil)

	ctx := context.Background()
	err := svc.DeletePost(ctx, postID)

	assert.NoError(t, err)
	mockPostRepo.AssertExpectations(t)
}
