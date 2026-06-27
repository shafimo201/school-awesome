package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/smoha201/school-awesome/internal/core/domain"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO users (
			id, school_id, email, full_name, password_hash, role_id, status, last_login_at,
			created_at, updated_at, created_by, updated_by, deleted_at, deleted_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`,
		user.ID,
		user.SchoolID,
		user.Email,
		user.FullName,
		user.PasswordHash,
		user.RoleID,
		user.Status,
		user.LastLoginAt,
		user.CreatedAt,
		user.UpdatedAt,
		user.CreatedBy,
		user.UpdatedBy,
		user.DeletedAt,
		user.DeletedBy,
	)
	return err
}

func (r *UserRepository) GetByEmail(ctx context.Context, schoolID, email string) (*domain.User, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, school_id, email, full_name, password_hash, role_id, status, last_login_at,
			created_at, updated_at, created_by, updated_by, deleted_at, deleted_by
		FROM users
		WHERE school_id = $1 AND email = $2 AND deleted_at IS NULL
	`, schoolID, email)

	user := &domain.User{}
	if err := row.Scan(
		&user.ID,
		&user.SchoolID,
		&user.Email,
		&user.FullName,
		&user.PasswordHash,
		&user.RoleID,
		&user.Status,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
		&user.DeletedAt,
		&user.DeletedBy,
	); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, schoolID, id string) (*domain.User, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, school_id, email, full_name, password_hash, role_id, status, last_login_at,
			created_at, updated_at, created_by, updated_by, deleted_at, deleted_by
		FROM users
		WHERE school_id = $1 AND id = $2 AND deleted_at IS NULL
	`, schoolID, id)

	user := &domain.User{}
	if err := row.Scan(
		&user.ID,
		&user.SchoolID,
		&user.Email,
		&user.FullName,
		&user.PasswordHash,
		&user.RoleID,
		&user.Status,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
		&user.DeletedAt,
		&user.DeletedBy,
	); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) ListBySchool(ctx context.Context, schoolID string, limit, offset int) ([]*domain.User, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, school_id, email, full_name, password_hash, role_id, status, last_login_at,
			created_at, updated_at, created_by, updated_by, deleted_at, deleted_by
		FROM users
		WHERE school_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`, schoolID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		if err := rows.Scan(
			&user.ID,
			&user.SchoolID,
			&user.Email,
			&user.FullName,
			&user.PasswordHash,
			&user.RoleID,
			&user.Status,
			&user.LastLoginAt,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.CreatedBy,
			&user.UpdatedBy,
			&user.DeletedAt,
			&user.DeletedBy,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	_, err := r.db.Exec(ctx, `
		UPDATE users SET full_name = $1, role_id = $2, status = $3, last_login_at = $4,
			updated_at = $5, updated_by = $6, deleted_at = $7, deleted_by = $8
		WHERE school_id = $9 AND id = $10
	`,
		user.FullName,
		user.RoleID,
		user.Status,
		user.LastLoginAt,
		user.UpdatedAt,
		user.UpdatedBy,
		user.DeletedAt,
		user.DeletedBy,
		user.SchoolID,
		user.ID,
	)
	return err
}

func (r *UserRepository) SoftDelete(ctx context.Context, schoolID, id, deletedBy string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE users SET deleted_at = $1, deleted_by = $2 WHERE school_id = $3 AND id = $4
	`, time.Now().UTC(), deletedBy, schoolID, id)
	return err
}

func (r *UserRepository) UpdatePasswordByEmail(ctx context.Context, schoolID, email, passwordHash, updatedBy string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE users SET password_hash = $1, updated_at = $2, updated_by = $3 WHERE school_id = $4 AND email = $5
	`, passwordHash, time.Now().UTC(), updatedBy, schoolID, email)
	return err
}
