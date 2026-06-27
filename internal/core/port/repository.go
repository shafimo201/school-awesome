package port

import (
	"context"

	"github.com/school-erp/project-school/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByEmail(ctx context.Context, schoolID, email string) (*domain.User, error)
	GetByID(ctx context.Context, schoolID, id string) (*domain.User, error)
	ListBySchool(ctx context.Context, schoolID string, limit, offset int) ([]*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	SoftDelete(ctx context.Context, schoolID, id, deletedBy string) error
}

type SchoolRepository interface {
	Create(ctx context.Context, school *domain.School) error
	GetByID(ctx context.Context, id string) (*domain.School, error)
	ListActive(ctx context.Context, limit, offset int) ([]*domain.School, error)
}
