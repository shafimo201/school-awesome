package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/smoha201/school-awesome/internal/core/domain"
	"github.com/smoha201/school-awesome/internal/core/port"
	"github.com/smoha201/school-awesome/internal/pkg/auth"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type RegisterUserInput struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"full_name" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
	RoleID   string `json:"role_id" validate:"required"`
}

type UserService struct {
	repo         port.UserRepository
	hasher       auth.PasswordHasher
	validator    *validator.Validate
	logger       zerolog.Logger
}

func NewUserService(repo port.UserRepository, hasher auth.PasswordHasher, logger zerolog.Logger) *UserService {
	return &UserService{
		repo:      repo,
		hasher:    hasher,
		validator: validator.New(),
		logger:    logger,
	}
}

func (s *UserService) Register(ctx context.Context, schoolID, createdBy string, input RegisterUserInput) (*domain.User, error) {
	if err := s.validator.Struct(input); err != nil {
		return nil, err
	}

	input.Username = strings.TrimSpace(strings.ToLower(input.Username))
	if existing, _ := s.repo.GetByEmail(ctx, schoolID, input.Username); existing != nil {
		return nil, ErrUserAlreadyExists
	}

	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	user := &domain.User{
		AuditFields: domain.AuditFields{
			ID:        uuid.NewString(),
			SchoolID:  schoolID,
			CreatedAt: now,
			UpdatedAt: now,
			CreatedBy: createdBy,
			UpdatedBy: createdBy,
		},
		Email:        input.Username,
		FullName:     input.FullName,
		PasswordHash: passwordHash,
		RoleID:       input.RoleID,
		Status:       domain.UserStatusActive,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Authenticate(ctx context.Context, schoolID, username, password string) (*domain.User, error) {
	user, err := s.repo.GetByEmail(ctx, schoolID, strings.TrimSpace(strings.ToLower(username)))
	if err != nil {
		return nil, ErrUserNotFound
	}
	if user == nil || user.Status != domain.UserStatusActive {
		return nil, ErrInvalidCredentials
	}
	if err := s.hasher.Compare(user.PasswordHash, password); err != nil {
		return nil, ErrInvalidCredentials
	}
	return user, nil
}
