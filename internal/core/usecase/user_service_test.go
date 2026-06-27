package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/rs/zerolog"
	"github.com/smoha201/school-awesome/internal/core/domain"
	"github.com/smoha201/school-awesome/internal/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepo struct {
	users map[string]*domain.User
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{users: make(map[string]*domain.User)}
}

func (m *mockUserRepo) Create(_ context.Context, user *domain.User) error {
	if _, exists := m.users[user.Email]; exists {
		return errors.New("duplicate")
	}
	m.users[user.Email] = user
	return nil
}

func (m *mockUserRepo) GetByEmail(_ context.Context, schoolID, email string) (*domain.User, error) {
	user, ok := m.users[email]
	if !ok || user.SchoolID != schoolID {
		return nil, nil
	}
	return user, nil
}

func (m *mockUserRepo) GetByID(_ context.Context, schoolID, id string) (*domain.User, error) {
	for _, u := range m.users {
		if u.ID == id && u.SchoolID == schoolID {
			return u, nil
		}
	}
	return nil, nil
}

func (m *mockUserRepo) ListBySchool(_ context.Context, schoolID string, limit, offset int) ([]*domain.User, error) {
	var users []*domain.User
	for _, u := range m.users {
		if u.SchoolID == schoolID {
			users = append(users, u)
		}
	}
	return users, nil
}

func (m *mockUserRepo) Update(_ context.Context, user *domain.User) error {
	m.users[user.Email] = user
	return nil
}

func (m *mockUserRepo) SoftDelete(_ context.Context, schoolID, id, deletedBy string) error {
	for _, u := range m.users {
		if u.ID == id && u.SchoolID == schoolID {
			now := time.Now().UTC()
			u.DeletedAt = &now
			u.DeletedBy = &deletedBy
			return nil
		}
	}
	return nil
}

func TestRegisterUser_Success(t *testing.T) {
	repo := newMockUserRepo()
	hasher := auth.NewBcryptHasher(bcrypt.DefaultCost)
	logger := zerolog.Nop()
	service := NewUserService(repo, hasher, logger)

	input := RegisterUserInput{
		Username: "user1",
		FullName: "Test User",
		Password: "Password1!",
		RoleID:   "admin",
	}

	user, err := service.Register(context.Background(), "school-123", "system", input)
	assert.NoError(t, err)
	assert.Equal(t, "user1", user.Email)
	assert.Equal(t, "admin", user.RoleID)
	assert.NotEmpty(t, user.PasswordHash)
}

func TestRegisterUser_Existing(t *testing.T) {
	repo := newMockUserRepo()
	hasher := auth.NewBcryptHasher(bcrypt.DefaultCost)
	logger := zerolog.Nop()
	service := NewUserService(repo, hasher, logger)

	_, err := service.Register(context.Background(), "school-123", "system", RegisterUserInput{
		Username: "user1",
		FullName: "Test User",
		Password: "Password1!",
		RoleID:   "admin",
	})
	assert.NoError(t, err)

	_, err = service.Register(context.Background(), "school-123", "system", RegisterUserInput{
		Username: "user1",
		FullName: "Test User",
		Password: "Password1!",
		RoleID:   "admin",
	})
	assert.ErrorIs(t, err, ErrUserAlreadyExists)
}

func TestAuthenticate_Success(t *testing.T) {
	repo := newMockUserRepo()
	hasher := auth.NewBcryptHasher(bcrypt.DefaultCost)
	logger := zerolog.Nop()
	service := NewUserService(repo, hasher, logger)

	_, err := service.Register(context.Background(), "school-123", "system", RegisterUserInput{
		Username: "user1",
		FullName: "Test User",
		Password: "Password1!",
		RoleID:   "admin",
	})
	assert.NoError(t, err)

	user, err := service.Authenticate(context.Background(), "school-123", "user1", "Password1!")
	assert.NoError(t, err)
	assert.Equal(t, "user1", user.Email)
}

func TestAuthenticate_InvalidPassword(t *testing.T) {
	repo := newMockUserRepo()
	hasher := auth.NewBcryptHasher(bcrypt.DefaultCost)
	logger := zerolog.Nop()
	service := NewUserService(repo, hasher, logger)

	_, err := service.Register(context.Background(), "school-123", "system", RegisterUserInput{
		Username: "user1",
		FullName: "Test User",
		Password: "Password1!",
		RoleID:   "admin",
	})
	assert.NoError(t, err)

	_, err = service.Authenticate(context.Background(), "school-123", "user1", "WrongPassword")
	assert.ErrorIs(t, err, ErrInvalidCredentials)
}
